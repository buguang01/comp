package comp

//对象实现接口
type Iobject interface {
	//获取自己的名字
	GetObjName() string
	//获取属性
	GetAttr(attrname string) int64
}

type IAttr interface {
	//获取属性
	GetAttr(attrname string) int64
}

func NewObjectAttr(name string, obj IAttr) Iobject {
	result := new(ObjectAttr)
	result.name = name
	result.IAttr = obj
	return result
}

type ObjectAttr struct {
	name string
	IAttr
}

//获取自己的名字
func (obj *ObjectAttr) GetObjName() string {
	return obj.name
}
