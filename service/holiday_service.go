package service

import (
	"context"
	"github.com/dapr-platform/common"
	"io"
	"net/http"

	"things-service/model"
	"time"
)

func init() {
	/*
		time.Sleep(time.Second * 10)

		go func() {
			defer func() { // 必须要先声明defer，否则不能捕获到panic异常
				fmt.Println("d")
				if err := recover(); err != nil {
					fmt.Println(err) // 这里的err其实就是panic传入的内容
				}
				fmt.Println("e")
			}()
			refreshData()
		}()*/
}

func refreshData() {
	common.Logger.Debug("refresh holiday data")
	yearStr := time.Now().Format("2006")
	refreshOneYear(yearStr)
	yearStr = time.Now().AddDate(1, 0, 0).Format("2006")
	refreshOneYear(yearStr)
}
func refreshOneYear(yearStr string) {
	url := "https://ghproxy.com/https://raw.githubusercontent.com/NateScarlet/holiday-cn/master/" + yearStr + ".json"
	resp, err := http.Get(url)
	if err != nil {
		common.Logger.Error("get holiday data error,from url="+url, err)
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		common.Logger.Error("read holiday data error,from url="+url, err)
		return
	}
	if len(data) == 0 {
		common.Logger.Error("read holiday data error,from url="+url, err)
		return
	}
	holiday := model.Holiday_json{
		ID:       yearStr,
		JSONData: string(data),
	}
	err = common.DbUpsert[model.Holiday_json](context.Background(), common.GetDaprClient(), holiday, model.Holiday_jsonTableInfo.Name, model.Holiday_json_FIELD_NAME_id)
	if err != nil {
		common.Logger.Error("refresh holiday data error,from url="+url, err)
		return
	}
	return
}
