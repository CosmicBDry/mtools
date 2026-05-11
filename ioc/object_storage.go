package ioc

import "fmt"

var (
	store            = NewDefaultStorage()
	DefaultNamespace = "default"
)

type IocObjectSet struct {
	Items []IocObject
}

func NewIocObjectSet() *IocObjectSet {

	return &IocObjectSet{}
}

type DefaultStorage struct {
	st map[string]*IocObjectSet
}

func NewDefaultStorage() *DefaultStorage {
	return &DefaultStorage{
		st: make(map[string]*IocObjectSet),
	}
}

//获取默认空间对象
func GetObject(name string) IocObject {

	return GetObjectWithNS(DefaultNamespace, name, DEFAULT_VERSION)

}

func RegistryObject(obj IocObject) {

	RegistryObjectWithNS(DefaultNamespace, obj)

}

//基于对象的namespace、name、version获取对象
func GetObjectWithNS(namespace, name, version string) IocObject {

	return store.Namespace(namespace).Get(name, version)

}

func RegistryObjectWithNS(namespace string, obj IocObject) {

	store.Namespace(namespace).Add(obj)
}

func InitObejects() error {
	return store.InitObejects()
}

func (ds *DefaultStorage) Namespace(namespace string) *IocObjectSet {

	if objs, ok := ds.st[namespace]; ok {
		return objs
	}

	newObjs := NewIocObjectSet()

	ds.st[namespace] = newObjs

	return newObjs

}

func (ds *DefaultStorage) InitObejects() error {
	for ns, objs := range ds.st {

		err := objs.Init()
		if err != nil {

			return fmt.Errorf("namespace: [%s] %s", ns, err.Error())
		}

	}
	return nil
}

//添加对象IocObject
func (is *IocObjectSet) Add(obj IocObject) {

	err := ValidateIocObject(obj)

	if err != nil {
		panic(err)
	}

	getObj, index := is.getWithIndex(obj.Name(), obj.Version())

	if getObj == nil {

		is.Items = append(is.Items, obj)
		return
	}

	if obj.AllowOverWrite() {

		fmt.Printf("%s object overwrite", obj.Name())

		is.setWithIndex(index, obj)
		return
	}

	fmt.Printf("%s.%s: object is exist and AllowOverWrite=false", obj.Name(), obj.Version())

}

func (is *IocObjectSet) Init() error {

	for index := range is.Items {
		obj := is.Items[index]

		err := obj.Init()
		if err != nil {
			return fmt.Errorf("%s object init failed: %s", err.Error(), obj.Name())
		}
	}
	return nil
}

func (is *IocObjectSet) Exist(name, version string) bool {

	return is.Get(name, version) != nil
}

func (is *IocObjectSet) Get(name, version string) IocObject {

	obj, _ := is.getWithIndex(name, version)

	return obj

}

//
func (is *IocObjectSet) getWithIndex(name, version string) (IocObject, int) {

	for index := range is.Items {

		obj := is.Items[index]

		if obj.Name() == name && obj.Version() == version {

			return obj, index
		}
	}
	return nil, 0

}

func (is *IocObjectSet) setWithIndex(index int, obj IocObject) {

	is.Items[index] = obj

}

func (is *IocObjectSet) ForEach(f func(IocObject)) {

	for index := range is.Items {
		obj := is.Items[index]

		f(obj)
	}

}

func (is *IocObjectSet) ObjectUids() (objuids []string) {

	for index := range is.Items {

		obj := is.Items[index]

		objuids = append(objuids, ObjectUid(obj))

	}
	return
}
