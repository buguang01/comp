package comp

import (
	"fmt"
)

var (
	//需要一个全局控制的函数
	FuncLi map[string]CompMark
)

const (
	//常量：最小的变量索引
	Coust_a byte = 97
)

func init() {
	FuncLi = make(map[string]CompMark)
}

func RegisterComp(name string, f CompMark) {
	FuncLi[name] = f
}

/*
输入字符串；输出运算对象
现在支持+-*／还有自定义的函数
"(a+b*c-4)/rand(6+21,100)"
不支持负号开头
*/
func NewCompNode(exp string) *CompNode {
	arr := compStringToArray(exp)
	result, _ := createComp(arr)
	return result
}

type CompMark func(node *CompNode, arr []int64) int64

//计算对象
type CompNode struct {
	Var  string    //值
	Val  int64     //值数据
	Num1 *CompNode //数据1
	Mark string    //运算符号
	Num2 *CompNode //数据2
}

//循环计算
func (cn *CompNode) CompVal(param []int64) int64 {

	if cn.Mark != "" {
		return cn.markComp(param)
	} else {
		return cn.getVal(param)
	}
	// } else if cn.Mark == "+" || cn.Mark == "-" || cn.Mark == "*" || cn.Mark == "/" {
	// 	return "(" + cn.Num1.String() + " " + cn.Mark + " " + cn.Num2.String() + ")"
	// } else {
	// 	return cn.Mark + "(" + cn.Num1.String() + "," + cn.Num2.String() + ")"
	// }
}

func (cn *CompNode) String() string {
	if cn.Var != "" {
		return cn.Var
	} else if cn.Mark == "+" || cn.Mark == "-" || cn.Mark == "*" || cn.Mark == "/" || cn.Mark == "%" {
		return "(" + cn.Num1.String() + " " + cn.Mark + " " + cn.Num2.String() + ")"
	} else if cn.Mark != "" {
		return cn.Mark + "(" + cn.Num1.String() + "," + cn.Num2.String() + ")"
	} else {
		return fmt.Sprint(cn.Val)
	}
}

func (cn *CompNode) markComp(arr []int64) (result int64) {
	switch cn.Mark {
	case "+":
		result = cn.Num1.CompVal(arr) + cn.Num2.CompVal(arr)
	case "-":
		result = cn.Num1.CompVal(arr) - cn.Num2.CompVal(arr)
	case "*":
		result = cn.Num1.CompVal(arr) * cn.Num2.CompVal(arr)
	case "/":
		result = cn.Num1.CompVal(arr) / cn.Num2.CompVal(arr)
	case "%":
		result = cn.Num1.CompVal(arr) % cn.Num2.CompVal(arr)
	default:
		if f, ok := FuncLi[cn.Mark]; ok {
			//找到函数
			result = f(cn, arr)
		} else {
			fmt.Println(" CompNode Error. Not exist Func ", cn.Mark, ". on CompNode:", cn)
		}
	}
	return
}

//拿到值
func (cn *CompNode) getVal(arr []int64) (result int64) {
	//使用a,b,c,d,e,f,g来定义变量，所以a对应的就是数组索引0，依次类推
	if cn.Var != "" {
		//是变量
		i := int(cn.Var[0] - Coust_a) //拿到变量索引
		if len(arr) > i {
			result = arr[i]
		}
	} else {
		result = cn.Val
	}
	return
}

//运算数据
func compStringToArray(exp string) (result []string) {
	arr := []rune(exp)
	tmp := NewStringBuilder()
	result = make([]string, 0, len(arr))
	for len(arr) > 0 {
		switch arr[0] {
		case '-', '+', '*', '/', '%', '(', ')', ',':
			if !tmp.IsEmpty() {
				result = append(result, tmp.ToString())
				tmp.Clear()
			}
			result = append(result, string(arr[0]))
		default:
			tmp.AppendRune(arr[0])
		}
		arr = arr[1:]
	}
	if !tmp.IsEmpty() {
		result = append(result, tmp.ToString())
	}
	// fmt.Println(result)
	return
}

//算符号优先
func checkPriority(mark string) int {
	priority := 10
	switch mark {
	case "+", "-":
		priority = 0
	case "*", "/", "%":
		priority = 1

	}
	return priority
}

func newComp(v string) *CompNode {
	result := new(CompNode)
	if t, ok := NewString(v).ToInt64(); ok == nil {
		result.Val = t
	} else {
		result.Var = v
	}
	return result
}

//生成运算对象
func createComp(arr []string) (result *CompNode, resarr []string) {
	numsk := make(Stack, 0, 20)
	marksk := make(StackMark, 0, 20)
arrfor:
	for len(arr) > 0 {
		data := arr[0]
		arr = arr[1:]
		if _, ok := FuncLi[data]; ok {
			//找到了函数
			//函数部分,双参数的部分
			marksk.Push(data)
			result, resarr = createComp(arr[1:]) //函数都会带括号
			numsk.Push(result)
			arr = resarr
			result, resarr = createComp(arr)
			numsk.Push(result)
			arr = resarr
		} else {
			switch data {
			case "(":
				//要嵌套调用
				result, resarr = createComp(arr)
				numsk.Push(result)
				arr = resarr
			case ")", ",":
				//结束当前调用
				break arrfor
			case "-", "+", "*", "/", "%":
				//符号
				if marksk.Len() == 0 {
					marksk.Push(data)
				} else {
					curpt := checkPriority(data)
					for numsk.Len() >= 2 {
						propt := checkPriority(marksk.Get())
						//优先级
						if curpt <= propt {
							num2, num1 := numsk.Pop(), numsk.Pop()
							mark := marksk.Pop()
							cn := new(CompNode)
							cn.Num1 = num1
							cn.Num2 = num2
							cn.Mark = mark
							numsk.Push(cn)
						} else {
							break
						}
					}
					marksk.Push(data)
				}
			default:
				//数值
				cn := newComp(data)
				numsk.Push(cn)
			}
		}

	}
	for numsk.Len() >= 2 {
		num2, num1 := numsk.Pop(), numsk.Pop()
		mark := marksk.Pop()
		cn := new(CompNode)
		cn.Num1 = num1
		cn.Num2 = num2
		cn.Mark = mark

		numsk.Push(cn)
	}
	cn := numsk.Get()
	return cn, arr
}

type Stack []*CompNode

func (sk Stack) Len() int {
	return len(sk)
}

func (sk *Stack) Push(v *CompNode) {
	*sk = append(*sk, v)
}

func (sk *Stack) Pop() (result *CompNode) {
	result = (*sk)[sk.Len()-1]
	*sk = (*sk)[:sk.Len()-1]
	return
}

func (sk Stack) IsEmpty() bool {
	return len(sk) == 0
}

func (sk Stack) Get() (result *CompNode) {
	return sk[sk.Len()-1]
}

type StackMark []string

func (sk StackMark) Len() int {
	return len(sk)
}

func (sk *StackMark) Push(v string) {
	*sk = append(*sk, v)
}

func (sk *StackMark) Pop() (result string) {
	result = (*sk)[sk.Len()-1]
	*sk = (*sk)[:sk.Len()-1]
	return
}

func (sk StackMark) IsEmpty() bool {
	return len(sk) == 0
}

func (sk StackMark) Get() (result string) {
	return sk[sk.Len()-1]
}
