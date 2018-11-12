package services

import (
	"sync"
)
var (
	only *ServiceFactory
	lock sync.Once
)
func ServicesInit()  {
	lock.Do(func() {
		only = NewServiceFactory()
	})
}

type ServiceFactory struct {
	Service        *Service

}

func NewServiceFactory() *ServiceFactory {
	o := new(ServiceFactory)
	o.Service = NewService()

	return o
}

func Only() *ServiceFactory {
	return only
}