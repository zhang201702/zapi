package service

import (
	"github.com/gogf/gf/frame/g"
	"github.com/zhang201702/zapi/base"
)

type BaseService struct {
	Name string
	Dao  base.Dao
}

func (service *BaseService) Save(data g.Map) (result interface{}, err error) {
	result, err = service.Dao.Save(data)
	if err == nil {
		result = data
	}
	return result, err
}
func (service *BaseService) Insert(data g.Map) (result interface{}, err error) {
	return service.Dao.Insert(data)
}
func (service *BaseService) Update(data g.Map, id string) (result interface{}, err error) {
	return service.Dao.Update(data, id)
}
func (service *BaseService) Delete(id string) (result interface{}, err error) {
	return service.Dao.Delete(id)
}
func (service *BaseService) DeleteAll(data g.Map) (result interface{}, err error) {
	return service.Dao.DeleteAll(data)
}
func (service *BaseService) List(data g.Map) (list g.List, err error) {
	return service.Dao.List(data)
}
func (service *BaseService) Fetch(id string) (result g.Map, err error) {
	return service.Dao.Fetch(id)
}
