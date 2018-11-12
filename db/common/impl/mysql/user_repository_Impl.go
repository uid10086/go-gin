package mysql

import (
	"lcq/db/mysql"
	"errors"
)


type UserRepositoryImpl struct {
	BaseRepositoryImpl
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return new(UserRepositoryImpl)
}

func (this *UserRepositoryImpl) Add(i *mysql.User) (err error) {
	engine := this.GetEngine()
	_, err = engine.InsertOne(i)
	return
}

func (this *UserRepositoryImpl) Get(query map[string]interface{}) (o *mysql.User, err error) {
	engine := this.GetReadEngine()
	o = new(mysql.User)
	has, err := engine.Where(query).Get(o)
	if !has {
		return nil, errors.New("没有查询到用户")
	}
	return
}

func (this *UserRepositoryImpl) EditById(id int, fields []string, i *mysql.User) (err error) {
	engine := this.GetEngine()
	_, err = engine.Id(id).Cols(fields...).Update(i)
	if err != nil {
		return
	}
	return nil
}

func (this *UserRepositoryImpl) DeleteById(id int) (err error) {
	engine := this.GetEngine()
	o := new(mysql.User)
	_, err = engine.ID(id).Delete(o)
	if err != nil {
		return
	}
	return nil
}

