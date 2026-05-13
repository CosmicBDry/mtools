package ioc

import (
	grpc "google.golang.org/grpc"
)

var (
	ControllerNamespace = "controllers"
)

type GrpcInterface interface {
	IocObject
	Registry(g *grpc.Server)
}

func RegistryController(obj IocObject) {

	RegistryObjectWithNS(ControllerNamespace, obj)

}

func GetController(name string) IocObject {
	return GetControllerWithVersion(name, DEFAULT_VERSION)

}

func GetControllerWithVersion(name, version string) IocObject {

	return GetObjectWithNS(ControllerNamespace, name, version)
}

func ShowControllers() []string {

	return store.st[ControllerNamespace].ObjectUids()

}

func LoadGrpcServerController(gserver *grpc.Server) {

	objs := store.Namespace(ControllerNamespace)

	objs.ForEach(func(obj IocObject) {

		grpc_obj, ok := obj.(GrpcInterface)

		if !ok {
			return
		}

		grpc_obj.Registry(gserver)

	})

}
