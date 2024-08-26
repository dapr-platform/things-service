package service

import (
	"github.com/Knetic/govaluate"
	"strings"
	"testing"
	"time"
)

func TestString(t *testing.T) {
	s := "abc"
	t.Log(len(strings.Split(s, " ")))
}
func TestCalc(t *testing.T) {
	expr := "value / 60"
	expression, err := govaluate.NewEvaluableExpression(expr)
	if err != nil {
		t.Error(err)
		return
	}
	parameters := make(map[string]interface{}, 8)

	parameters["value"] = float64(500)
	parameters["interval"] = 60

	result, err := expression.Evaluate(parameters)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestDate(t *testing.T) {
	ts := time.Now().Truncate(time.Second)
	t.Log(ts.Location().String())
	t.Log(ts.Format("2006-01-02 15:04:05"))
	qts := time.Date(ts.Year(), ts.Month(), ts.Day(), 0, 0, 0, 0, time.Local)
	t.Log(qts.Format("2006-01-02 15:04:05"))
	qts = time.Date(ts.Year(), ts.Month(), 1, 0, 0, 0, 0, time.Local)
	t.Log(qts.Format("2006-01-02 15:04:05"))
	qts = time.Date(ts.Year(), 1, 1, 0, 0, 0, 0, time.Local)
	t.Log(qts.Format("2006-01-02 15:04:05"))

}
