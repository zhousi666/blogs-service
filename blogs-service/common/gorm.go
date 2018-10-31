package common

import (
	"database/sql"
	"github.com/jinzhu/gorm"

	// 引入数据库驱动注册及初始化,不添加会报错
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB gorm数据库实例
var DB *gorm.DB

// GormDB 封装后的gorm数据库实例
type GormDB struct {
	*gorm.DB
	gdbDone bool
}

// InitDB 初始化数据库
func InitDB(mysqlurl string, mysqlidle, mysqlmaxopen int, debug bool) {
	Logger.Info("db info: ", mysqlurl, mysqlidle, mysqlmaxopen)

	idb, err := gorm.Open("mysql", mysqlurl)
	if err != nil {
		panic(err)
	}

	idb.DB().SetMaxIdleConns(mysqlidle)
	idb.DB().SetMaxOpenConns(mysqlmaxopen)
	idb.LogMode(debug)

	DB = idb
}

// DBClose 关闭数据库
func DBClose() {
	DB.Close()
}

// DBBegin 打开一个transaction
func DBBegin() *GormDB {
	txn := DB.Begin()
	if txn.Error != nil {
		panic(txn.Error)
	}
	return &GormDB{txn, false}
}

// DBCommit 提交并关闭transaction(用于插入)
func (c *GormDB) DBCommit() {
	if c.gdbDone {
		return
	}
	tx := c.Commit()
	c.gdbDone = true
	if err := tx.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
}

// DBRollback 回滚并关闭transaction(一般用于查询)
func (c *GormDB) DBRollback() {
	if c.gdbDone {
		return
	}
	tx := c.Rollback()
	c.gdbDone = true
	if err := tx.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
}
