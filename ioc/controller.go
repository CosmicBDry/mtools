package ioc

import "fmt"

var Controller = map[string]IocObject{}

func InitController() error {
	fmt.Println("初始化ioc.Controller--------------")
	for k, v := range Controller {

		fmt.Printf("key:%s Value:%#v\n", k, v)
		err := v.Init()
		if err != nil {
			return fmt.Errorf("%s Init Controller Error: %s", k, err.Error())
		}
	}

	return nil
}

func GetController(Name string) IocObject {

	if v, ok := Controller[Name]; ok {
		return v
	}

	panic("ioc.Getcontroller failure!")
}

func ShowController() []string {

	names := []string{}
	for k, _ := range Controller {

		names = append(names, k)

	}

	return names

}

func RegistryController(obj IocObject) {
	Controller[obj.Name()] = obj
}
