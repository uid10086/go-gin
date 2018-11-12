package mysql

import (
	"github.com/go-xorm/xorm"
)
var (
	e      *MysqlEngine
	sum  int64
	readcount int64
)

func InitNewMysqlEngine(write, read []*xorm.Engine)  {
	e = new(MysqlEngine)
	e.read = read
	e.write = write
	return
}
type MysqlEngine struct {
	write, read []*xorm.Engine
}
func (this *MysqlEngine) GetWriteEngine() *xorm.Engine {
	return this.write[0]
}

func (this *MysqlEngine) GetReadEngine() *xorm.Engine {
	//c := len(this.read)
	//number := int64(3 * c)
	//sum = atomic.AddInt64(&sum,1)
	//if(sum % number == 0){
	//	return this.write[0]
	//}else {
	//	if c == 0 {
	//		return  this.write[0]
	//	}else {
	//		readcount = atomic.AddInt64(&readcount,1)
	//		return this.read[readcount % int64(c)]
	//	}
	//}
	return this.write[0]
}
func getEngine() *xorm.Engine {
	return e.GetWriteEngine()
}
func getReadEngine() *xorm.Engine {
	return e.GetReadEngine()
}

type BaseRepositoryImpl struct {
	engine *xorm.Engine
}

//获取读写实例
func (this *BaseRepositoryImpl) GetEngine() *xorm.Engine {
	return getEngine()
}

//获取只读实例
func (this *BaseRepositoryImpl) GetReadEngine() *xorm.Engine {
	return getReadEngine()
}