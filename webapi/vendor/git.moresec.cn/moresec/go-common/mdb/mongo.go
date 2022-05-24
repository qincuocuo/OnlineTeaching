package mdb

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"
)

// mongo 实现. api文档:https://godoc.org/github.com/globalsign/mgo
// 该库封装函数默认一个库.
type MongoSession struct {
	session *mgo.Session
	dbName  string
}

func NewMongoSession() *MongoSession {
	return &MongoSession{}
}

func (ms *MongoSession) Connect(opt interface{}) error {
	url, ok := opt.(string)
	if !ok {
		return errors.New("mongo connection param error")
	}

	dialInfo, err := mgo.ParseURL(url)
	ms.dbName = dialInfo.Database

	s, err := mgo.Dial(url)
	if err != nil {
		return err
	}

	ms.session = s
	ms.session.SetMode(mgo.Monotonic, true)

	return nil
}

func (ms *MongoSession) SetPoolLimit(limit int) {
	ms.session.SetPoolLimit(limit)
}

// 获取原始操作 session.
// 类型:  mgo.Session 类型
// RT: 连接使用完成，使用方必须调用 Close() 释放资源.
// example:
// session, _ = GetOrg().(*mgo.Session)
// defer session.Close()
// session.Find(....)...
func (ms *MongoSession) GetOrg() interface{} {
	return ms.session.Clone()
}

func (ms *MongoSession) clone() *mgo.Session {
	return ms.session.Clone()
}

// 实际操作.
func (ms *MongoSession) FindOne(name string, query, result interface{}) (exist bool, err error) {
	s := ms.clone()
	defer s.Close()

	err = s.DB(ms.dbName).C(name).Find(query).One(result)
	exist = true

	if err == mgo.ErrNotFound {
		exist = false
	}
	return exist, err
}

func (ms *MongoSession) FindById(name string, id interface{}, result interface{}) (exist bool, err error) {
	s := ms.clone()
	defer s.Close()

	err = s.DB(ms.dbName).C(name).FindId(id).One(result)
	exist = true

	if err == mgo.ErrNotFound {
		exist = false
	}
	return exist, err
}

func (ms *MongoSession) FindAllWithSelect(name string, query, selection, result interface{}) error {
	s := ms.clone()
	defer s.Close()

	return s.DB(ms.dbName).C(name).Find(query).Select(selection).All(result)
}

// 查询固定字段.
func (ms *MongoSession) FindOneWithSelect(name string, query, selection, result interface{}) (exist bool, err error) {
	s := ms.clone()
	defer s.Close()

	err = s.DB(ms.dbName).C(name).Find(query).Select(selection).One(result)
	exist = true

	if err == mgo.ErrNotFound {
		exist = false
	}
	return exist, err
}

// selector, result, limit
// result 必须是slice.
func (ms *MongoSession) Find(name string, query, result interface{}, limit int) error {
	if limit <= 0 {
		return ErrorLimit
	}

	s := ms.clone()
	defer s.Close()

	err := s.DB(ms.dbName).C(name).Find(query).Limit(limit).All(result)
	return err
}

func (ms *MongoSession) FindByLimitAndSkip(name string, query, result interface{}, limit, skip int) error {
	if limit <= 0 || skip < 0 {
		return ErrorLimit
	}

	s := ms.clone()
	defer s.Close()

	err := s.DB(ms.dbName).C(name).Find(query).Skip(skip).Limit(limit).All(result)
	return err
}

func (ms *MongoSession) FindAll(name string, query, result interface{}) error {
	s := ms.clone()
	defer s.Close()

	return s.DB(ms.dbName).C(name).Find(query).All(result)
}

func (ms *MongoSession) FindMin(name string, query interface{}, field string, result interface{}) error {
	s := ms.clone()
	defer s.Close()

	return s.DB(ms.dbName).C(name).Find(query).Select(bson.M{field: 1}).Sort(field).One(result)
}

func (ms *MongoSession) FindMax(name string, query interface{}, field string, result interface{}) error {
	s := ms.clone()
	defer s.Close()

	return s.DB(ms.dbName).C(name).Find(query).Select(bson.M{field: 1}).Sort(fmt.Sprintf("-%s", field)).One(result)
}

func (ms *MongoSession) RemoveById(name string, id interface{}) error {
	s := ms.clone()
	defer s.Close()

	err := s.DB(ms.dbName).C(name).RemoveId(id)
	return err
}

func (ms *MongoSession) RemoveAll(name string, query interface{}) error {
	s := ms.clone()
	defer s.Close()

	_, err := s.DB(ms.dbName).C(name).RemoveAll(query)
	return err
}

func (ms *MongoSession) Insert(name string, docs ...interface{}) error {
	s := ms.clone()
	defer s.Close()

	return s.DB(ms.dbName).C(name).Insert(docs...)
}

// Update 更新函数只会更新给出的字段，而不会替换整个文档.
func (ms *MongoSession) Update(name string, query interface{}, update interface{}, multi bool) error {
	s := ms.clone()
	defer s.Close()

	value := make(bson.M)
	value["$set"] = update
	if multi {
		_, err := s.DB(ms.dbName).C(name).UpdateAll(query, value)
		return err
	}
	return s.DB(ms.dbName).C(name).Update(query, value)
}

// UpdateManual 更新函数只会更新给出的字段，会替换整个文档.
func (ms *MongoSession) UpdateManual(name string, query interface{}, update interface{}, multi bool) error {
	s := ms.clone()
	defer s.Close()

	if multi {
		_, err := s.DB(ms.dbName).C(name).UpdateAll(query, update)
		return err
	}
	return s.DB(ms.dbName).C(name).Update(query, update)
}

// UpdateById 根据主键id更新函数只会更新给出的字段，而不会替换整个文档.
func (ms *MongoSession) UpdateById(name string, id interface{}, update interface{}) error {
	s := ms.clone()
	defer s.Close()

	upset := make(bson.M)
	upset["$set"] = update
	return s.DB(ms.dbName).C(name).UpdateId(id, upset)
}

// UpdateByIdManual 根据主键id更新函数只会更新给出的字段，会替换整个文档.
func (ms *MongoSession) UpdateByIdManual(name string, id interface{}, update interface{}) error {
	s := ms.clone()
	defer s.Close()

	return s.DB(ms.dbName).C(name).UpdateId(id, update)
}

func (ms *MongoSession) Close() {
	ms.session.Close()
}

// FindWithSelect 查询固定字段(返回多条结果), limit=0时查询所有
func (ms *MongoSession) FindWithSelect(name string, query, selection, result interface{}, limit int) error {
	s := ms.clone()
	defer s.Close()

	if limit == 0 {
		return s.DB(ms.dbName).C(name).Find(query).Select(selection).All(result)
	} else {
		return s.DB(ms.dbName).C(name).Find(query).Limit(limit).Select(selection).All(result)
	}
}

// FindCount 查询统计
func (ms *MongoSession) FindCount(name string, query interface{}) (int, error) {
	s := ms.clone()
	defer s.Close()

	return s.DB(ms.dbName).C(name).Find(query).Count()
}

// FindSort 查询排序
func (ms *MongoSession) FindSort(name string, query interface{}, sorter string, result interface{}, limit int) error {
	s := ms.clone()
	defer s.Close()

	if limit == 0 {
		return s.DB(ms.dbName).C(name).Find(query).Sort(sorter).All(result)
	} else {
		return s.DB(ms.dbName).C(name).Find(query).Sort(sorter).Limit(limit).All(result)
	}
}

func (ms *MongoSession) FindSortByLimitAndSkip(name string, query interface{}, result interface{}, limit, skip int, sorter ...string) error {
	if limit <= 0 || skip < 0 {
		return ErrorLimit
	}

	s := ms.clone()
	defer s.Close()
	if limit == 0 {
		return s.DB(ms.dbName).C(name).Find(query).Sort(sorter...).All(result)
	} else {
		return s.DB(ms.dbName).C(name).Find(query).Sort(sorter...).Skip(skip).Limit(limit).All(result)
	}
}

// 聚合管道
func (ms *MongoSession) FindWithAggregation(name string, pipeline interface{}, result interface{}) error {
	s := ms.clone()
	defer s.Close()

	return s.DB(ms.dbName).C(name).Pipe(pipeline).All(result)
}

func (ms *MongoSession) Upsert(name string, selector interface{}, upset interface{}) (*mgo.ChangeInfo, error) {
	s := ms.clone()
	defer s.Close()

	ci, err := s.DB(ms.dbName).C(name).Upsert(selector, upset)
	return ci, err
}

func (ms *MongoSession) UpsertId(name string, id string, upset interface{}) (*mgo.ChangeInfo, error) {
	s := ms.clone()
	defer s.Close()

	ci, err := s.DB(ms.dbName).C(name).UpsertId(id, upset)
	return ci, err
}

func (ms *MongoSession) GetCollectionNames() ([]string, error) {
	s := ms.clone()
	defer s.Close()

	collections, err := s.DB(ms.dbName).CollectionNames()
	return collections, err
}
