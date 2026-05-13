package ioc

import (
	"fmt"

	"github.com/emicklei/go-restful/v3"
	//"github.com/gin-gonic/gin"
)

// type ApiHandlerController interface {
// 	IocObject
// 	Registry(*gin.RouterGroup)
// }

var (
	ApiNamespace = "apis"
)

type ApiGoRestfulInterface interface {
	IocObject
	Registry(*restful.WebService)
}

func RegistryApi(obj IocObject) {

	RegistryObjectWithNS(ApiNamespace, obj)
}

func GetApi(name string) IocObject {

	return GetApiWithVersion(name, DEFAULT_VERSION)

}

func GetApiWithVersion(name, version string) IocObject {

	return GetObjectWithNS(ApiNamespace, name, version)

}

func ShowApis() []string {
	return store.Namespace(ApiNamespace).ObjectUids()
}

func LoadGoRestfulRouterApi(path_prefix string, rc *restful.Container) {

	store.Namespace(ApiNamespace).ForEach(func(obj IocObject) {

		//断言为ApiGoRestfulInterface接口类型
		api, ok := obj.(ApiGoRestfulInterface)

		if !ok {
			return
		}

		ws := new(restful.WebService)

		//Consumes限制客户请求格式类型，Produces限制服务端响应格式类型
		ws.Path(fmt.Sprintf("%s %s", api.Version(), api.Name())).Consumes(restful.MIME_JSON, restful.MIME_XML).Produces(restful.MIME_JSON, restful.MIME_XML)

		api.Registry(ws)

		rc.Add(ws)

	})

}
