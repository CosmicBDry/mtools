package ioc

import "fmt"

func ValidateIocObject(obj IocObject) error {

	if obj.Name() == "" {

		return fmt.Errorf("%T obj.Name() return empty,object name reqiured", obj)

	}

	if obj.Version() == "" {
		return fmt.Errorf("%T obj.Version() return empty,object version reqiured", obj)
	}

	return nil

}
