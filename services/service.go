package services

import (
	"lcq/db/common/impl/mysql"
	mysql2 "lcq/db/mysql"
)
type Service struct {
	user * mysql.UserRepositoryImpl
}

func NewService() * Service  {
	s := new(Service)
	s.user = mysql.NewUserRepositoryImpl()
	return s
}

func (this * Service) Getuser (m map[string]interface{}) (*mysql2.User,error){

	o,err := this.user.Get(m)
	if err != nil {
		return nil,err
	}
	return o,nil
}
