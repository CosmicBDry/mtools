package ioc

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ApiHandlerController interface {
	IocObject
	Registry(*gin.RouterGroup)
}

var ApiController = map[string]ApiHandlerController{}

func InitApiController(rootPath string, r *gin.Engine) error {
	fmt.Println("初始化ioc.ApiController--------------")
	for k, v := range ApiController {
		err := v.Init()

		if err != nil {
			return fmt.Errorf("%s InitApiController Error: %s", k, err.Error())
		}

		router := r.Group(rootPath + "/" + v.Name())
		fmt.Printf("key:%s Value:%#v Path:%s\n", k, v, rootPath+"/"+v.Name())
		v.Registry(router)
	}
	return nil

}

func GetApiHandlerController(NAME string) ApiHandlerController {

	if v, ok := ApiController[NAME]; ok {
		return v
	}

	panic("GET ApiHandlerController failure...!")

}

func ShowApiHandlerControllers() []string {
	names := []string{}

	for k, _ := range ApiController {

		names = append(names, k)

	}
	return names

}

func RegistryApiHandlerController(obj ApiHandlerController) {

	ApiController[obj.Name()] = obj
}
