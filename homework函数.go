package main

import "fmt"
import "math"


type Sixflower struct {  //定义六花结构体
    i float64    //复数的实部
    j float64    //复数的虚部
}

func Add(a, b Sixflower) Sixflower {        //定义复数加法函数
    return Sixflower{i: a.i + b.i, j: a.j + b.j}
}

func Sub(a, b Sixflower) Sixflower {      //定义复数减法函数
    return Sixflower{i: a.i - b.i, j: a.j - b.j}
}

func Mul(a, b Sixflower) Sixflower {    //定义复数的乘法函数
    return Sixflower{i: (a.i*b.i - a.j*b.j), j: (a.i*b.j + b.i*a.j)}
}

func Div(a, b Sixflower) Sixflower {  //定义复数的除法函数
	if b.i == 0 && b.j == 0 {
		// 防止除以零的情况
		panic("分母不为零")
	}
    return Sixflower{i: (a.i*b.i + a.j*b.j) / (b.i*b.i + b.j*b.j), j: (a.j*b.i - a.i*b.j) / (b.i*b.i + b.j*b.j)}
}
func Mod(c Sixflower)float64 {
	return math.Sqrt(c.i*c.i + c.j*c.j)
}
func ToString (a Sixflower)string {    //定义转化字符串函数
	return fmt.Sprintf("Sixflower{i :%v,j :%v}",a.i,a.j)

}

func main() {
    flower1 := Sixflower{i: 1, j: 1}//初始化声名两个变量
    flower2 := Sixflower{i: 1, j: 1}

    fmt.Println("Add:", Add(flower1, flower2))
    fmt.Println("Subtract:", Sub(flower1, flower2))
    fmt.Println("Multiply:", Mul(flower1, flower2))       //加减乘除 字符串的调用
    fmt.Println("Divide:", Div(flower1, flower2))
	fmt.Println(ToString(flower1))
	fmt.Println(ToString(flower2))    //综上所述 六花结构体嘎嘎好用 大家快用起来
}