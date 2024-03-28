package db

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"sync"
	viper1 "test/t_viper"
	"time"
)

type MyDB struct {
	*gorm.DB
}
type Db struct {
	Ctx   context.Context
	Write *MyDB
	Read  *MyDB
}

var (
	write *MyDB
	read  *MyDB
	once  sync.Once
)

type ctxTransactionKey struct {
}

func New(ctx context.Context) *Db {
	once.Do(func() {
		conf := viper1.DBViper()
		write = &MyDB{DB: dial(conf["write"])}
		read = &MyDB{DB: dial(conf["read"])}
	})
	// 根据实例,修改上下文
	var writeC *gorm.DB
	var readC *gorm.DB
	// 如果是在事务中, 则读写库全部使用事务的db
	if tx := ctx.Value(ctxTransactionKey{}); tx != nil {
		writeC = tx.(*gorm.DB)
		readC = tx.(*gorm.DB)
	} else {
		writeC = write.WithContext(ctx)
		readC = read.WithContext(ctx)
	}

	return &Db{
		Ctx:   ctx,
		Write: &MyDB{DB: writeC},
		Read:  &MyDB{DB: readC},
	}
}

func dial(conf viper1.DataBaseConf) *gorm.DB {
	_db, err := gorm.Open(mysql.Open(conf.Source), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "zby_", // table name prefix, table for `User` would be `t_users`
			SingularTable: true,   // use singular table name, table for `User` would be `user` with this option enabled
		},
	})
	if err != nil {
		panic(err)
	}
	err = polling(_db)
	if err != nil {
		panic(err)
	}

	return _db
}

// 设置程序池;
func polling(_db *gorm.DB) error {
	sqlDb, err := _db.DB()
	if err != nil {
		return err
	}
	// TODO: 这些参数要迁移到配置文件中
	// 设置连接池;
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDb.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDb.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间
	sqlDb.SetConnMaxLifetime(time.Hour)
	return nil
}
