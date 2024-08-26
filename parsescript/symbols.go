package parsescript

import "reflect"

//go:generate go install github.com/traefik/yaegi/cmd/yaegi@v0.15.1
//go:generate yaegi extract github.com/spf13/cast
var Symbols = map[string]map[string]reflect.Value{}
