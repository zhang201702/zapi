package zapi

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/zhang201702/zapi/base"
	"github.com/zhang201702/zapi/utils"
	"github.com/zhang201702/zhang"
	"reflect"
)

var apis = make(map[string]base.Api)
var ops = map[string]string{"list": "List", "fetch": "Fetch", "save": "Save", "delete": "Delete"}

func RegisterSimpleApi(group *ghttp.RouterGroup) {
	group.POST("/{_name}/{_op}", func(r *ghttp.Request) {
		name := r.GetString("_name")
		op := r.GetString("_op")
		a, ok := apis[name]
		if !ok {
			a = utils.GetApi(name, "")
			apis[name] = a
		}
		inputs := []reflect.Value{reflect.ValueOf(r)}
		reflect.ValueOf(a).MethodByName(ops[op]).Call(inputs)
	})
}

func SimpleApiRun() {
	s := zhang.Default()
	s.Group("/api", func(group *ghttp.RouterGroup) {
		RegisterSimpleApi(group)
		group.Middleware()
	})
	s.Run()
}
