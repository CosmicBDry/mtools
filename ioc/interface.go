package ioc

type IocObject interface {
	Name() string
	Init() error
}
