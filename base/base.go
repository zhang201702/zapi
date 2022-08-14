package base

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type Dao interface {
	Save(data g.Map) (result interface{}, err error)
	Insert(data g.Map) (result interface{}, err error)
	Update(data g.Map, id string) (result interface{}, err error)
	Delete(id string) (result interface{}, err error)
	DeleteAll(data g.Map) (result interface{}, err error)
	List(data g.Map) (list g.List, err error)
	Fetch(id string) (result g.Map, err error)
	ListSql(data g.Map) (list g.List, err error)
}

type Service interface {
	Save(data g.Map) (result interface{}, err error)
	Insert(data g.Map) (result interface{}, err error)
	Update(data g.Map, id string) (result interface{}, err error)
	Delete(id string) (result interface{}, err error)
	DeleteAll(data g.Map) (result interface{}, err error)
	List(data g.Map) (list g.List, err error)
	Fetch(id string) (result g.Map, err error)
}

type Api interface {
	Fetch(r *ghttp.Request)
	List(r *ghttp.Request)
	Save(r *ghttp.Request)
	Delete(r *ghttp.Request)
}
