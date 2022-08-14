package api

import (
	"encoding/json"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/zhang201702/zapi/base"
	"github.com/zhang201702/zhang/zweb"
	"io/ioutil"
	"strings"
)

type BaseApi struct {
	zweb.ApiBase
	Service base.Service
	Name    string
	Path    string
}

//
//func (api *BaseApi) Register(s *zhang.ServerGF) {
//	s.BindHandler("/"+api.Path+"/list", api.List)
//	s.BindHandler("/"+api.Path+"/fetch", api.Fetch)
//	//s.BindHandler("/"+api.Name, api.fetch)
//	s.BindHandler("/"+api.Path+"/save", api.Save)
//	s.BindHandler("/"+api.Path+"/delete", api.Delete)
//}
func (api *BaseApi) Fetch(r *ghttp.Request) {
	id := r.GetString("id")
	if id == "" {
		api.ErrorResult("fetch,非法参数", nil, r)
		return
	}
	result, err := api.Service.Fetch(id)
	api.Result(result, err, r)
}
func (api *BaseApi) List(r *ghttp.Request) {

	var data g.Map = GetJsonMap(r)
	result, err := api.Service.List(data)
	api.Result(result, err, r)
}
func (api *BaseApi) Save(r *ghttp.Request) {
	var data = GetJsonMap(r)
	result, err := api.Service.Save(data)
	api.Result(result, err, r)
}
func (api *BaseApi) Delete(r *ghttp.Request) {
	id := r.GetString("id")
	if id == "" {
		api.ErrorResult("delete,非法参数", nil, r)
		return
	}
	result, err := api.Service.Delete(id)
	api.Result(result, err, r)
}

func GetJsonMap(r *ghttp.Request) g.Map {
	contentType := r.Request.Header.Get("Content-type")
	if strings.Index(contentType, "application/json") > -1 {
		body, _ := ioutil.ReadAll(r.Body)
		var queryMap = g.Map{}
		err := json.Unmarshal(body, &queryMap)
		if err == nil {
			return queryMap
		}
	}
	data := r.GetMap()
	delete(data, "_name")
	delete(data, "_op")
	return data
}
