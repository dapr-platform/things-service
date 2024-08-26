package parsescript

import (
	"github.com/pkg/errors"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	"sync"
)

type GolangScriptProcessor struct {
	Content                         string
	Interp                          *interp.Interpreter
	PropertyTranslateFunc           func(map[string]any) map[string]any
	PropertySetTranslateFunc        func(map[string]any, string, any) map[string]any
	PropertySetTranslateDesiredFunc func(map[string]any, string, any) map[string]any
	CheckAlertFunc                  func(map[string]any) []map[string]any
}

var golangScriptProcessorMap map[string]*GolangScriptProcessor
var golangScriptProcessorMapLock sync.Mutex

func init() {
	golangScriptProcessorMap = make(map[string]*GolangScriptProcessor, 0)
}
func GetGolangScriptProcessor(content string) (processor *GolangScriptProcessor, err error) {
	golangScriptProcessorMapLock.Lock()
	defer golangScriptProcessorMapLock.Unlock()
	if processor, ok := golangScriptProcessorMap[content]; ok {
		return processor, nil
	} else {
		processor := &GolangScriptProcessor{
			Content: content,
		}
		err = processor.Init()
		if err != nil {
			err = errors.Wrap(err, "init golangScriptProcessor error")
			return nil, err
		}
		golangScriptProcessorMap[content] = processor
		return processor, nil
	}

}

func (p *GolangScriptProcessor) Init() (err error) {
	p.Interp = interp.New(interp.Options{GoPath: "./_gopath/"})
	if err = p.Interp.Use(stdlib.Symbols); err != nil {
		err = errors.Wrap(err, "use stdlib symbols error")
		return
	}
	err = p.Interp.Use(Symbols)
	if err != nil {
		err = errors.Wrap(err, "use symbols error")
		return
	}
	_, err = p.Interp.Eval(p.Content)
	if err != nil {
		err = errors.Wrap(err, "eval content error")
		return
	}
	v, err := p.Interp.Eval("parse.TranslateProperty")
	if err == nil {
		p.PropertyTranslateFunc = v.Interface().(func(map[string]any) map[string]any)
	}

	vv, err := p.Interp.Eval("parse.TranslatePropertySet")
	if err == nil {
		p.PropertySetTranslateFunc = vv.Interface().(func(map[string]any, string, any) map[string]any)

	}

	vvv, err := p.Interp.Eval("parse.TranslatePropertySetDesired")
	if err == nil {
		p.PropertySetTranslateDesiredFunc = vvv.Interface().(func(map[string]any, string, any) map[string]any)

	}
	vvvv, err := p.Interp.Eval("parse.CheckAlert")
	if err == nil {
		p.CheckAlertFunc = vvvv.Interface().(func(map[string]any) []map[string]any)
	}

	return
}

func (p *GolangScriptProcessor) ProcessTranslatePropertyValue(data map[string]any) (result map[string]any) {
	if p.PropertyTranslateFunc == nil {
		return data
	}
	result = p.PropertyTranslateFunc(data)
	return
}
func (p *GolangScriptProcessor) ProcessTranslatePropertySet(data map[string]any, propName string, val any) (result map[string]any) {
	if p.PropertySetTranslateFunc == nil {
		result = make(map[string]any)
		result[propName] = val
		return
	}
	result = p.PropertySetTranslateFunc(data, propName, val)
	return
}

func (p *GolangScriptProcessor) ProcessTranslatePropertySetDesired(data map[string]any, propName string, val any) (result map[string]any) {
	if p.PropertySetTranslateDesiredFunc == nil {
		result = make(map[string]any)
		result[propName] = val
		return
	}
	result = p.PropertySetTranslateDesiredFunc(data, propName, val)
	return
}

func (p *GolangScriptProcessor) CheckAlert(data map[string]any) (result []map[string]any) {
	if p.CheckAlertFunc == nil {
		return make([]map[string]any, 0)
	}
	result = p.CheckAlertFunc(data)
	return
}
