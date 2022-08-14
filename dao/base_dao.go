package dao

import (
	"errors"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/google/uuid"
	"github.com/zhang201702/zhang/zlog"
)

type BaseDao struct {
	DB   gdb.DB
	Name string
}

func (dao *BaseDao) Save(data g.Map) (result interface{}, err error) {
	if id, ok := data["id"]; ok {
		return dao.Update(data, gconv.String(id))
	} else {
		return dao.Insert(data)
	}
}

func (dao *BaseDao) Insert(data g.Map) (result interface{}, err error) {

	table := dao.DB.Table(dao.Name)
	if gconv.String(data["id"]) == "" {
		data["id"] = uuid.New().String()
	}
	result, err = table.Data(data).Insert()
	if err != nil {
		zlog.LogError(err, dao.Name, "添加异常", data, result)
	}
	return result, err
}

func (dao *BaseDao) Update(data g.Map, id string) (result interface{}, err error) {
	table := dao.DB.Table(dao.Name)
	result, err = table.Data(data).Where("id=?", id).Update()
	if err != nil {
		zlog.LogError(err, dao.Name, "修改异常", data, result)
	}
	return result, err
}

func (dao *BaseDao) Delete(id string) (result interface{}, err error) {
	table := dao.DB.Table(dao.Name)
	result, err = table.Where("id=?", id).Delete()
	if err != nil {
		zlog.LogError(err, dao.Name, "删除异常", id, result)
	}
	return result, err
}

func (dao *BaseDao) DeleteAll(data g.Map) (result interface{}, err error) {
	table := dao.DB.Table(dao.Name)
	result, err = table.Where(data).Delete()
	if err != nil {
		zlog.LogError(err, dao.Name, "删除所有异常", data, result)
	}
	return result, err
}

func (dao *BaseDao) List(data g.Map) (list g.List, err error) {

	table := dao.DB.Table(dao.Name)

	if orderBy, ok := data["_orderby"]; ok {
		table.OrderBy(gconv.String(orderBy))
		delete(data, "_orderby")
	}
	r, err := table.Where(data).All()
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return g.List{}, nil
		}
		zlog.LogError(err, dao.Name, "查询异常", data, r)
	} else {
		list = r.List()
	}
	return list, err
}

func (dao *BaseDao) Fetch(id string) (result g.Map, err error) {
	table := dao.DB.Table(dao.Name)
	r, err := table.Where("id=?", id).One()
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		zlog.LogError(err, dao.Name, "查询异常", r)
		result = nil
	} else {
		result = r.Map()
	}
	return result, err
}

func (dao *BaseDao) ListSql(data g.Map) (list g.List, err error) {
	if sql, ok := data["_sql"]; ok {
		args := gconv.Interfaces(data["_args"])
		r, err := dao.DB.GetAll(gconv.String(sql), args...)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				return g.List{}, nil
			}
			zlog.LogError(err, dao.Name, "查询异常.ListSql", data, r)
		} else {
			list = r.List()
		}
		return list, err
	}
	return nil, errors.New("未知！")
}
