package ioc

import (
	"fmt"
)

type IocObject interface {
	Name() string
	Init() error
	Version() string
	Priority() int
	Destroy()
	AllowOverWrite() bool
}

const (
	DEFAULT_VERSION = "v1"
)

func ObjectUid(obj IocObject) string {

	return fmt.Sprintf("%s.%s", obj.Name(), obj.Version())
}

type IocObjectImpl struct {
}

func (i IocObjectImpl) Name() string {
	return ""
}

func (i IocObjectImpl) Init() error {
	return nil
}

func (i IocObjectImpl) Version() string {
	return DEFAULT_VERSION
}

func (i IocObjectImpl) Priority() int {
	return 0
}

func (i IocObjectImpl) AllowOverWrite() bool {
	return false
}

func (i IocObjectImpl) Destroy() {
}
