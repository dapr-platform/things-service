package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dapr-platform/common"
	"github.com/mozillazg/go-pinyin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cast"
	"strings"

	"things-service/entity"
	"things-service/model"
	"time"
)

var businessGaugeMap = make(map[string]*prometheus.GaugeVec)

func init() {
	go registerBusinessGauge()
}

func SetDeviceGaugeValue(productName, propertyIdentifier, deviceIdentifier string, value any) {
	go setDeviceGaugeValue(productName, propertyIdentifier, deviceIdentifier, value)
}
func setDeviceGaugeValue(productName, propertyIdentifier, deviceIdentifier string, value any) {
	defer func() {
		if r := recover(); r != nil {
			common.Logger.Error("set device gauge value error", r)
		}
	}()
	productName = getProductNamePinyin(productName)
	propertyIdentifier = fixPropertyIdentifier(propertyIdentifier)
	gauge, exist := businessGaugeMap[productName+"_"+propertyIdentifier]
	if exist {
		gauge.WithLabelValues(deviceIdentifier).Set(cast.ToFloat64(value))
	}
}
func fixPropertyIdentifier(identifier string) string {
	return strings.ReplaceAll(strings.ReplaceAll(identifier, "(", ""), ")", "")
}

func registerBusinessGauge() {
	for {
		products, err := common.DbQuery[model.Product](context.Background(), common.GetDaprClient(), model.ProductTableInfo.Name, "")
		if err != nil {
			common.Logger.Error("get product error", err)
			time.Sleep(time.Second * 60)
			continue
		}
		for _, product := range products {
			err := registerOneProductBusinessGauge(context.Background(), &product)
			if err != nil {
				common.Logger.Error("register business gauge error", err)
				time.Sleep(time.Second * 60)
				continue
			}
		}
		time.Sleep(time.Second * 60)
	}

}

func getProductNamePinyin(productName string) string {
	pinyins := pinyin.LazyConvert(productName, nil)
	result := strings.Trim(fmt.Sprint(pinyins), "[]")
	result = strings.Replace(result, " , ", "", -1)
	result = strings.Replace(result, " ", "", -1)
	if result == "" { //英文的
		result = productName
	}
	return result
}

func registerOneProductBusinessGauge(ctx context.Context, product *model.Product) (err error) {
	pmodel := &entity.ProductModel{}
	err = json.Unmarshal([]byte(product.JSONData), pmodel)
	if err != nil {
		common.Logger.Error("json unmarshal error", err)
		return
	}
	for _, deviceModel := range pmodel.DeviceModels {
		for _, property := range deviceModel.Properties {
			gaugeKey := getProductNamePinyin(product.Name) + "_" + fixPropertyIdentifier(property.Identifier)

			_, exist := businessGaugeMap[gaugeKey]
			if !exist {
				gauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{
					Namespace: "iot",
					Subsystem: "things_service",
					Name:      gaugeKey,
					Help:      "",
				}, []string{"device_identifier"})
				prometheus.MustRegister(gauge)
				businessGaugeMap[gaugeKey] = gauge
				common.Logger.Debugf("register gauge %s", gaugeKey)
			}
		}
	}
	return
}
