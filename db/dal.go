package db

import (
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	"fmt"
	"lcq/db/mysql"
	impl "lcq/db/common/impl/mysql"
	_ "github.com/go-sql-driver/mysql"
)




type DbConfig struct {
	Types              string
}


func mysqlsync(e *xorm.Engine) error{
	e.SetMapper(core.SameMapper{})
	//同步数据库

	syncArray := []interface{}{
		new(mysql.User),
	}
	err := e.Sync2(syncArray...)
	return err
}
func MysqlInit() error {
	w := make([]*xorm.Engine, 0)
	r := make([]*xorm.Engine, 0)
	//for _, v := range config.Write {

	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=%s", "lcq", "lcq123",
		"120.78.222.255", 3306, "lcq", "Asia%2FShanghai")
		e, err := xorm.NewEngine("mysql", dbUrl)
		if err != nil {
			panic(err)
		}
		err = e.Ping()
		//e.SetLogger(new(Logger))
		if err != nil {
			panic(err)
			fmt.Println(err)
		}
		//log.Debug("mysql write instance ", dbUrl, " is connected")

		//e.SetLogLevel(core.LOG_DEBUG)
		//e.ShowSQL(v.SqlLog)
		//设置字段映射规则
		e.SetMapper(core.SameMapper{})
		e.SetMaxIdleConns(100)
		//if v.Sync == true {
			//log.Debug("mysql write sync")
			err = mysqlsync(e)
	if err != nil {
		panic(err)
		fmt.Println(err)
	}

		//}
		//分页缓存
		//cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
		//a := new(mysql.Stream)
		//e.MapCacher(&a, cacher)
		w = append(w, e)
//	}

    impl.InitNewMysqlEngine(w,r)
	return nil
}
