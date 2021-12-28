package comp

//计算用的参数
type CompEvent struct {
	//对象
	ObjList []Iobject
	//变量
	Param []int64
}

//获取变量值；
//使用a,b,c,d,e,f,g来定义变量，所以a对应的就是数组索引0，依次类推
func (et *CompEvent) GetVal(cn *CompNode) (result int64) {
	if cn.Var == "" {
		return cn.Val
	} else {
		i := int(cn.Var[0] - Coust_a) //拿到变量索引
		if len(et.Param) > i {
			result = et.Param[i]
		}
	}
	return
}

//获取对象的属性值
func (et *CompEvent) GetObjAttr(cn *CompNode) (result int64) {
	for i := range et.ObjList {
		if et.ObjList[i].GetObjName() == cn.Obj {
			return et.ObjList[i].GetAttr(cn.Var)
		}
	}
	return 0
}
