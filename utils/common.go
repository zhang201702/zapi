package utils

import (
	"github.com/zhang201702/zapi/api"
	"github.com/zhang201702/zapi/base"
	"github.com/zhang201702/zapi/dao"
	"github.com/zhang201702/zapi/service"
	"github.com/zhang201702/zhang/z"
)

var GetService = func(name string) base.Service {
	return &service.BaseService{Name: name, Dao: GetDao(name)}
}

var GetDao = func(name string) base.Dao {
	return &dao.BaseDao{Name: name, DB: z.DB()}
}

var GetApi = func(name, path string) base.Api {
	return &api.BaseApi{Name: name, Path: path, Service: GetService(name)}
}
