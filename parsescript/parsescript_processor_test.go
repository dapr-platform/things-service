package parsescript

import "testing"

func TestInit(t *testing.T) {
	content := `
package parse
import "github.com/spf13/cast"
//读转换
func TranslateProperty(data map[string]any) (value map[string]any){
    value=make(map[string]any)
    for k, v:=range data{
        
	    value[k]=v
    }
    value["新风阀开关"] = data["新风阀开状态"]
    value["排风阀开关"] = data["排风阀开状态"]
    return
}

//写转换

//写转换desired
func TranslatePropertySetDesired(data map[string]any,propName string,val any) (value map[string]any){
value=make(map[string]any)
switch propName{
case "新风阀开关":
    value["新风阀开关"]=val
   
case "排风阀开关":
   
    value["排风阀开关"]=val

default:
value[propName]=val
}
return
}

func CheckAlert(data map[string]any) (alerts []map[string]any){
    alerts = make( []map[string]any,0)
    if cast.ToInt(data["配电1-故障告警"])==1{
        event := make(map[string]any,0)
        event["alert_property"] = "配电1-故障告警"
        event["alert_value"] = 1
        alerts = append(alerts,event)
    }
    return

}
`
	processor, err := GetGolangScriptProcessor(content)
	if err != nil {
		t.Error(err)
		return
	}
	data := make(map[string]any)
	data["age"] = "100"
	val := processor.ProcessTranslatePropertyValue(data)
	t.Log(val)
	val = processor.ProcessTranslatePropertySet(data, "新风阀停止位置", "100")
	t.Log(val)
	alert := processor.CheckAlert(data)
	t.Log(alert)

}
