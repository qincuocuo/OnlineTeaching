package mrds

import (
	"fmt"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
)

type RdbAlias struct {
	name       string
	driverName string
	dataSource string
	*gorm.DB
}

// DB is ...
type db struct {
	*sync.RWMutex
	dbPool map[string]*RdbAlias
}

var mdb db

// InitRDB is ...
func init() {
	pool := make(map[string]*RdbAlias)
	mdb.RWMutex = &sync.RWMutex{}
	mdb.dbPool = pool
}

// RegisterRdb is register multi-db
// the mian db is named default
// must called at pre
// params SetMaxIdleConns SetMaxOpenConns SetConnMaxLifetime
func RegisterRdb(aliasName, driverName, dataSource string, params ...int) error {
	mdb.Lock()
	defer mdb.Unlock()
	db, err := gorm.Open(driverName, dataSource)
	if err != nil {
		return err
	}
	err = db.DB().Ping()
	if err != nil {
		return err
	}
	for i, v := range params {
		switch i {
		case 0:
			db.DB().SetMaxIdleConns(v)
		case 1:
			db.DB().SetMaxOpenConns(v)
		case 2:
			db.DB().SetConnMaxLifetime(time.Second * time.Duration(int64(v)))
		}
	}
	mdb.dbPool[aliasName] = &RdbAlias{
		name:       aliasName,
		driverName: driverName,
		dataSource: dataSource,
		DB:         db,
	}
	return nil
}

// GetRdb is return DB
func GetRdb(aliasNames ...string) (*RdbAlias, error) {
	var name string
	if len(aliasNames) > 0 {
		name = aliasNames[0]
	} else {
		name = "default"
	}
	mdb.Lock()
	defer mdb.Unlock()
	al, ok := mdb.dbPool[name]
	if ok {
		return al, nil
	}
	return nil, fmt.Errorf("DataBase of alias name `%s` not found", name)
}
