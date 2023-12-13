package core

import (
	"runtime"
	"shares/internal/config"

	"github.com/xxjwxc/public/errors"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/mysqldb"
)

// Dao ...
var Dao DaoCore

// DaoCore core dao
type DaoCore struct {
	dbr *mysqldb.MySqlDB // 读库
	dbw *mysqldb.MySqlDB // 写库
}

func init() {
	Dao.InitDao()
}

// GetDB 获取读数据库
func (dao *DaoCore) GetDB() *mysqldb.MySqlDB {
	if dao.dbr == nil {
		mylog.Error(errors.New("dbr is nil. "))
	}
	return dao.dbr
}

// GetDBr 获取读数据库
func (dao *DaoCore) GetDBr() *mysqldb.MySqlDB {
	if dao.dbr == nil {
		mylog.Error(errors.New("dbr is nil. "))
	}
	return dao.dbr
}

// GetDBw 获取写数据库
func (dao *DaoCore) GetDBw() *mysqldb.MySqlDB {
	if dao.dbw == nil {
		mylog.Error(errors.New("dbw is nil. "))
	}
	return dao.dbw
}

// InitDao 初始化dao
func (dao *DaoCore) InitDao() {
	runtime.SetFinalizer(dao, dao.Destroy) //析构时触发

	dao.dbr = mysqldb.OnInitDBOrm(config.GetMysqlConStr(), 5, 20, true)
	dao.dbw = mysqldb.OnInitDBOrm(config.GetMysqlConStr(), 5, 20, true)
}

// Destroy 释放
func (dao *DaoCore) Destroy() {
	if dao.dbr != nil {
		dao.dbr.OnDestoryDB()
	}

	if dao.dbw != nil {
		dao.dbw.OnDestoryDB()
	}
}
