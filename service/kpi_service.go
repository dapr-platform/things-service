package service

import (
	"context"
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/dapr-platform/common"
	"github.com/go-co-op/gocron"
	"github.com/spf13/cast"
	"sort"
	"strconv"

	"things-service/entity"
	"things-service/model"
	"time"
)

type ProductKpiProcessor struct {
	ProductName   string
	Schedulers    map[int]*gocron.Scheduler
	KpiProcessors map[string]*KpiCalcProcessor
}

type KpiCalcProcessor struct {
	KpiInfo    model.Kpi_info
	Evaluable  *govaluate.EvaluableExpression
	CalcScript string
}

var ProductKpiProcessors = make(map[string]*ProductKpiProcessor, 0)

var KpiCalcProcessors = make(map[string]*KpiCalcProcessor, 0)

func init() {
	go refreshKpiInfoCalc()
}

func refreshKpiInfoCalc() {
	time.Sleep(time.Second * 5)
	for {
		kpis, err := common.DbQuery[model.Kpi_info](context.Background(), common.GetDaprClient(), model.Kpi_infoTableInfo.Name, "")
		if err != nil {
			common.Logger.Error("get all kpis error " + err.Error())
			time.Sleep(time.Minute)
			continue
		}
		for _, kpi := range kpis {
			if _, exist := ProductKpiProcessors[kpi.ProductName]; !exist {
				ProductKpiProcessors[kpi.ProductName] = &ProductKpiProcessor{
					ProductName:   kpi.ProductName,
					KpiProcessors: make(map[string]*KpiCalcProcessor, 0),
					Schedulers:    make(map[int]*gocron.Scheduler, 0),
				}
			}
			productKpiProcessor := ProductKpiProcessors[kpi.ProductName]
			if kpi.Type == 0 { //原始指标需要计算
				if _, exist := productKpiProcessor.Schedulers[int(kpi.Interval)]; !exist {
					productKpiProcessor.Schedulers[int(kpi.Interval)] = gocron.NewScheduler(time.UTC)
					_, err = productKpiProcessor.Schedulers[int(kpi.Interval)].Every(int(kpi.Interval)).Minute().Do(productKpiProcessor.calcAllKpis)
					if err != nil {
						common.Logger.Errorf("kpiId=%s %s start schedule interval %d minutes error %s", kpi.ID, kpi.Name, kpi.Interval, err.Error())
						continue
					}
					productKpiProcessor.Schedulers[int(kpi.Interval)].StartAsync()
				}
			}

			if _, exist := productKpiProcessor.KpiProcessors[kpi.ID]; !exist {
				productKpiProcessor.KpiProcessors[kpi.ID] = &KpiCalcProcessor{
					KpiInfo: kpi,
				}
			}
			kpiCalcProcessor := productKpiProcessor.KpiProcessors[kpi.ID]
			kpiCalcProcessor.KpiInfo = kpi

			if kpi.CalcScript == "" {
				common.Logger.Error("kpiId=" + kpi.ID + " " + kpi.Name + " calcScript is empty")
				continue
			}
			if kpiCalcProcessor.Evaluable == nil || kpi.CalcScript != kpiCalcProcessor.CalcScript {
				common.Logger.Debug("kpiId=" + kpi.ID + " " + kpi.Name + " calcScript new or changed " + kpi.CalcScript)
				expression, err := govaluate.NewEvaluableExpression(kpi.CalcScript)
				if err != nil {
					common.Logger.Error("kpiId=" + kpi.ID + " " + kpi.Name + " calcScript error " + err.Error())
					continue
				}
				kpiCalcProcessor.Evaluable = expression
				kpiCalcProcessor.CalcScript = kpi.CalcScript
			}

		}

		time.Sleep(time.Minute)
	}
}

func (p *ProductKpiProcessor) calcAllKpis() {
	ts := time.Now().Truncate(time.Second)
	common.Logger.Debugf("calc kpi %s start", p.ProductName)
	deviceIdentifiers, err := GetEnabledDeviceIdentifiersByProductName(context.Background(), p.ProductName)
	if err != nil {
		common.Logger.Errorf("calc kpi %s error %s", p.ProductName, err.Error())
		return
	}
	for _, deviceIdentifier := range deviceIdentifiers {
		deviceMirror, err := GetDeviceMirror(context.Background(), deviceIdentifier)
		if err != nil {
			common.Logger.Errorf("calc kpi %s get device mirror error %s", p.ProductName, err.Error())
			continue
		}
		if deviceMirror == nil {
			common.Logger.Errorf("calc deviceIdentifier=%s ProductName=%s deviceMirror is nil", deviceIdentifier, p.ProductName)
			continue
		}
		kpiDatas := make([]entity.CurrentKpiData, 0)
		for _, v := range p.KpiProcessors {
			if v.Evaluable == nil {
				continue
			}
			params := make(map[string]interface{})
			params["interval"] = v.KpiInfo.Interval
			for k, v := range deviceMirror.State.Reported {
				params[k] = v
			}
			if v.KpiInfo.Type == 0 {

				val, err := v.Evaluable.Evaluate(params)
				if err != nil {
					common.Logger.Errorf("calc deviceIdentifier=%s kpiId=%s kpiName=%s evaluate error %s", deviceIdentifier, v.KpiInfo.ID, v.KpiInfo.Name, err.Error())

					continue
				}
				err = saveKpiData(context.Background(), common.GetMD5Hash(deviceIdentifier), v.KpiInfo.ID, ts, val)
				if err != nil {
					common.Logger.Errorf("calc deviceIdentifier=%s kpiId=%s kpiName=%s save data error %s", deviceIdentifier, v.KpiInfo.ID, v.KpiInfo.Name, err.Error())
				}
			} else {
				qts := ts
				qtbl := "f_kpi_metrics_"
				switch v.KpiInfo.SummaryType {
				case "daily":
					qts = time.Date(ts.Year(), ts.Month(), ts.Day(), 0, 0, 0, 0, time.Local)
				case "monthly":
					qts = time.Date(ts.Year(), ts.Month(), 1, 0, 0, 0, 0, time.Local)
				case "yearly":
					qts = time.Date(ts.Year(), 1, 1, 0, 0, 0, 0, time.Local)
				default:
					common.Logger.Errorf("%s kpiId=%s %s summaryType error %s", p.ProductName, v.KpiInfo.ID, v.KpiInfo.Name, v.KpiInfo.SummaryType)
					continue
				}
				qtbl += v.KpiInfo.SummaryType
				qstr := "_select=" + v.KpiInfo.ValueType + "_value&bucket=" + common.LocalTime(qts).DbString() + "&kpi_id=" + cast.ToString(v.KpiInfo.OrgID) + "&device_id=" + common.GetMD5Hash(deviceIdentifier)
				common.Logger.Debugf("calc kpi %s %s %s", p.ProductName, v.KpiInfo.Name, qstr)
				data, err := common.DbGetOne[map[string]interface{}](context.Background(), common.GetDaprClient(), qtbl, qstr)
				if err != nil {
					common.Logger.Errorf("calc kpi %s %s %s error %s", p.ProductName, v.KpiInfo.Name, qstr, err.Error())
					continue
				}
				if data != nil {
					dbMap := *data
					params["value"] = cast.ToFloat64(dbMap[v.KpiInfo.ValueType+"_value"])
					//common.Logger.Debugf("params %v", params)
					val, err := v.Evaluable.Evaluate(params)
					if err != nil {
						common.Logger.Errorf("calc deviceIdentifier=%s kpiId=%s kpiName=%s evaluate error value=%v %s", deviceIdentifier, v.KpiInfo.ID, v.KpiInfo.Name, dbMap[v.KpiInfo.ValueType+"_value"], err.Error())
						continue
					}
					val, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", val), 64)
					kpiDatas = append(kpiDatas, entity.CurrentKpiData{
						Name:  v.KpiInfo.Label,
						Id:    v.KpiInfo.ID,
						Value: val,
						Unit:  v.KpiInfo.Unit,
						Ts:    common.LocalTime(qts),
					})
				}
			}

		}
		less := func(i, j int) bool {
			return cast.ToInt(kpiDatas[i].Id) < cast.ToInt(kpiDatas[j].Id)
		}
		sort.Slice(kpiDatas, less)
		err = DeviceMirroerMergeKpiData(context.Background(), deviceIdentifier, kpiDatas)
		if err != nil {
			common.Logger.Errorf("calc deviceIdentifier=%s kpiDatas error %s", deviceIdentifier, err.Error())
		}

	}
	return
}

func (p *KpiCalcProcessor) calcKpi() {
	ts := time.Now().Truncate(time.Second)
	deviceIdentifiers, err := GetEnabledDeviceIdentifiersByProductName(context.Background(), p.KpiInfo.ProductName)
	if err != nil {
		common.Logger.Errorf("calc kpiId=%s %s error %s", p.KpiInfo.ID, p.KpiInfo.Name, err.Error())
		return
	}
	for _, deviceIdentifier := range deviceIdentifiers {
		deviceMirror, err := GetDeviceMirror(context.Background(), deviceIdentifier)
		if err != nil {
			common.Logger.Errorf("calc kpiId=%s %s error %s", p.KpiInfo.ID, p.KpiInfo.Name, err.Error())
			continue
		}
		if deviceMirror == nil {
			common.Logger.Errorf("calc deviceIdentifier=%s kpiId=%s kpiName=%s deviceMirror is nil", deviceIdentifier, p.KpiInfo.ID, p.KpiInfo.Name)
			continue
		}
		params := make(map[string]interface{})
		params["interval"] = p.KpiInfo.Interval
		for k, v := range deviceMirror.State.Reported {
			params[k] = v
		}
		val, err := p.Evaluable.Evaluate(params)
		if err != nil {
			common.Logger.Errorf("calc deviceIdentifier=%s kpiId=%s kpiName=%s evaluate error %s", deviceIdentifier, p.KpiInfo.ID, p.KpiInfo.Name, err.Error())

			continue
		}
		err = saveKpiData(context.Background(), common.GetMD5Hash(deviceIdentifier), p.KpiInfo.ID, ts, val)
		if err != nil {
			common.Logger.Errorf("calc deviceIdentifier=%s kpiId=%s kpiName=%s save data error %s", deviceIdentifier, p.KpiInfo.ID, p.KpiInfo.Name, err.Error())
		}
	}
	return
}

func saveKpiData(ctx context.Context, deviceId string, kpiId string, ts time.Time, value interface{}) (err error) {
	data := model.Kpi_metrics_5m{
		ID:       common.GetMD5Hash(deviceId + "_" + kpiId + "_" + ts.Format("2006-01-02 15:04:05")),
		Ts:       common.LocalTime(ts),
		DeviceID: deviceId,
		KpiID:    kpiId,
		Value:    cast.ToFloat64(value),
	}
	err = common.DbUpsert[model.Kpi_metrics_5m](ctx, common.GetDaprClient(), data, model.Kpi_metrics_5mTableInfo.Name, model.Kpi_metrics_5m_FIELD_NAME_kpi_id)
	if err != nil {
		common.Logger.Errorf("save kpi data error %s", err.Error())
		return
	}
	return
}
