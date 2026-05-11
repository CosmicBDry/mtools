package ioc

var (
	ControllerNamespace = "controllers"
)

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
