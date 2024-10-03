package main
import (
	"fmt"
	"math"
)

type Sixflower struct {       //定义六花结构体
	 i float64        //i是实部
	 j float64        //j是虚部
}

// 复数加法方法
func (c Sixflower) Add(other Sixflower) Sixflower {
	return Sixflower{
		i: c.i + other.i,
		j: c.j + other.j,
	}
}

// 复数减法方法
func (c Sixflower) Subtract(other Sixflower) Sixflower {
	return Sixflower{
		i: c.i - other.i,
		j: c.j - other.j,
	}
}

// 复数乘法方法
func (c Sixflower) Multiply(other Sixflower) Sixflower {
	return Sixflower{
		i: c.i*other.i - c.j*other.j,
		j: c.i*other.j + c.j*other.i,
	}
}

// 复数除法方法
func (c Sixflower) Divide(other Sixflower) Sixflower {
	if other.i == 0 && other.j == 0 {
		// 防止除以零的情况
		panic("分母不为零")
	}
	divisor := other.i*other.i + other.j*other.j
	iPart := (c.i*other.i + c.j*other.j) / divisor
	jPart := (c.j*other.i - c.i*other.j) / divisor
	return Sixflower{
		i: iPart,
		j: jPart,
	}
}

// 复数求模长方法
func (c Sixflower) Modulus() float64 {
	return math.Sqrt(c.i*c.i + c.j*c.j)
}

// 复数返回String 方法
func (c Sixflower) String() string {
	return fmt.Sprintf("(%g + %gi)", c.i, c.j)
}

func main() {

	c1 := Sixflower{i: 1, j: 2}        //初始化六花 进行变量赋值
	c2 := Sixflower{i: 3, j: 4}

	fmt.Println("c1 + c2 =", c1.Add(c2).String())
	fmt.Println("c1 - c2 =", c1.Subtract(c2).String())          //加减乘除求模长并返回字符串
	fmt.Println("c1 * c2 =", c1.Multiply(c2).String())
	fmt.Println("c1 / c2 =", c1.Divide(c2).String())
	fmt.Println("c1 modulus =", c1.Modulus())
}       //六花结构体嘎嘎好用 谁懂