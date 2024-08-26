package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/dapr-platform/common"
	"github.com/xuri/excelize/v2"
	"log"
	"strconv"
	"strings"

	"things-service/entity"
	"time"
)

type DbExcelFileUploadProcessor struct {
	Tag int64
}

func NewDbExcelFileUploadProcessor() *DbExcelFileUploadProcessor {
	return &DbExcelFileUploadProcessor{
		Tag: time.Now().Unix(),
	}
}

func (p *DbExcelFileUploadProcessor) Process(ctx context.Context, filePath string, keyPara string) error {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	for _, name := range f.GetSheetMap() {
		/*
			if strings.Index(name, "s_") < 0 && strings.Index(name, "o_") < 0 {
				log.Println("skip sheet:", name)
				continue
			}

		*/
		ret, err := common.GetDaprClient().InvokeMethod(ctx, common.DB_SERVICE_NAME, "/show/"+common.DBNAME+"/"+common.DB_SCHEMA+"/"+name, "get")
		if err != nil {
			log.Println("invoke show method error:", name, err.Error())
			continue
		}
		var columnInfo []entity.DbColumnInfo
		err = json.Unmarshal(ret, &columnInfo)
		if err != nil {
			log.Println("unmarshal show method error:", name, err.Error())
			return err
		}
		if len(columnInfo) == 0 {
			log.Println("columnInfo is empty:", name)
			continue
		}

		rows, err := f.GetRows(name)
		if err != nil {
			log.Println(err.Error())
			return err
		}

		rowMap := make(map[string]int)
		columnMap := make(map[int]entity.DbColumnInfo)
		columnDbFieldMap := make(map[string]entity.DbColumnInfo)
		key := keyPara
		if key == "" {
			key = "id"
		}

		for _, column := range columnInfo {
			columnDbFieldMap[column.ColumnName] = column

		}
		calcField := ""

		for i, row := range rows {
			if i == 0 {
				for j, col := range row {
					col = strings.TrimSpace(col)
					if strings.Index(col, "*") == 0 {
						col = col[1:]
						calcField = col
					}
					for _, column := range columnInfo {

						if col == column.ColumnName {
							rowMap[column.ColumnName] = j
							columnMap[j] = column
							break
						}
					}
				}
				continue
			}

			var jsonMap = make(map[string]interface{})
			for j, colVal := range row {
				if colVal == "" {
					continue
				}
				column, exists := columnMap[j]
				if !exists {
					continue
				}
				val, err := p.getValueByColumnInfo(column, colVal)
				if err != nil {
					log.Println(err.Error())
					continue

				}
				jsonMap[column.ColumnName] = val

				if column.ColumnName == calcField {
					id := p.CalcId(colVal)
					targetColumn, exists := columnDbFieldMap[key]
					if !exists {
						continue
					}
					val, err := p.getValueByColumnInfo(targetColumn, id)
					if err != nil {
						log.Println(err.Error())
						continue
					}
					jsonMap[key] = val
				}
				targetCol, targetValue := p.CustomProcessCol(column.ColumnName, colVal)
				if targetCol != "" {
					targetColumn, exists := columnDbFieldMap[targetCol]
					if !exists {
						continue
					}
					val, err := p.getValueByColumnInfo(targetColumn, targetValue)
					if err != nil {
						log.Println(err.Error())
						continue
					}
					jsonMap[targetCol] = val
				}

			}

			buf, _ := json.Marshal(jsonMap)
			var t map[string]any
			err = json.Unmarshal(buf, &t)
			if err != nil {
				log.Println(err.Error())
				return err
			}
			if _, exists := t[key]; !exists {
				log.Println("key is empty:", key, t)
				continue
			}

			err = common.DbUpsert(ctx, common.GetDaprClient(), t, name, key)
			if err != nil {
				log.Println(err.Error(), t)
			}
		}

	}

	return nil

	/*
		err = common.DbDeleteByOps(ctx, common.GetDaprClient(), t, []string{"tag"}, []string{"<"}, []any{p.Tag})
		if err != nil {
			log.Println("delete error:", err)
			return
		}

	*/
}

func (p *DbExcelFileUploadProcessor) getValueByColumnInfo(column entity.DbColumnInfo, col string) (val any, err error) {
	switch column.DataType {
	case "numeric":
		val, _ = strconv.ParseFloat(col, 64)
	case "bigint":
		val, _ = strconv.ParseInt(col, 10, 64)
	case "integer":
		if strings.Index(col, ".") != -1 {
			val, _ = strconv.ParseFloat(col, 64)
			val = int64(val.(float64))
		} else {
			val, _ = strconv.ParseInt(col, 10, 64)
		}

		if column.IsNullable == "YES" {
			if err != nil {
				val = nil
			}
		} else {
			if err != nil {
				val = 0
			}
		}
	case "date":
		val, _ = time.Parse("2006-01-02", col)
	case "character varying":
		val = strings.TrimSpace(col)
	case "text":
		val = strings.TrimSpace(col)
		/*
			if strings.Index(col, ".") != -1 && strings.Count(col, ".") == 1 {
				val, err = strconv.ParseFloat(col, 64)
				if err == nil {
					val = int64(val.(float64))
				} else {
					val = strings.TrimSpace(col)
				}

			} else {
				val = strings.TrimSpace(col)
			}

		*/

	case "double precision":
		val, _ = strconv.ParseFloat(col, 64)

	case "timestamp without time zone":
		val, _ = time.Parse("2006-01-02 15:04:05", col)

	default:
		err = errors.New("unknown type: " + column.DataType)
	}
	return
}

func (p *DbExcelFileUploadProcessor) CalcId(col_value string) (targetVal string) {
	h := md5.New()
	h.Write([]byte(col_value))
	re := h.Sum(nil)
	targetVal = hex.EncodeToString(re)
	return
}

func (p *DbExcelFileUploadProcessor) CustomProcessCol(db_col_name string, col_value string) (targetName string, targetVal string) {
	/*
		switch db_col_name {
		case "space_dn":
			h := md5.New()
			h.Write([]byte(col_value))
			re := h.Sum(nil)
			targetName = "space_id"
			targetVal = hex.EncodeToString(re)
		case "object_dn":
			h := md5.New()
			h.Write([]byte(col_value))
			re := h.Sum(nil)
			targetName = "object_id"
			targetVal = hex.EncodeToString(re)

		default:


		}

	*/

	return
}
