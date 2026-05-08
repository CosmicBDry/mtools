package ioc

type IocObject interface {
	Name() string
	Init() error
}

type IocObjImpl struct {
}

func (i IocObjImpl) Name() string {
	return ""
}

func (i IocObjImpl) Init() error {
	return nil
}
