package main

import "fmt"

/**
	函数定义
		func 函数名(参数)(返回值){
		函数体
		}

	函数名：由字母、数字、下划线组成。但函数名的第一个字母不能是数字。在同一个包内，函数名也称不能重名（包的概念详见后文）。
	参数：参数由参数变量和参数变量的类型组成，多个参数之间使用,分隔。
	返回值：返回值由返回值变量和其变量类型组成，也可以只写返回值的类型，多个返回值必须用()包裹，并用,分隔。
	函数体：实现指定功能的代码块。
*/
func main() {

	////函数的调用：通过函数名()的方式调用函数
	//	sayHello()
	//	ret := intSum(10, 20)
	//	fmt.Println(ret)


	//ret1 := intSum2()
	//ret2 := intSum2(10)
	//ret3 := intSum2(10, 20)
	//ret4 := intSum2(10, 20, 30)
	//fmt.Println(ret1, ret2, ret3, ret4) //0 10 30 60


	//ret5 := intSum3(100)
	//ret6 := intSum3(100, 10)
	//ret7 := intSum3(100, 10, 20)
	//ret8 := intSum3(100, 10, 20, 30)
	//fmt.Println(ret5, ret6, ret7, ret8) //100 110 130 160




}

//求两个数之和的函数：
func intSum(x int, y int) int {
	return x + y
}

//函数的参数中如果相邻变量的类型相同，则可以省略类型，例如：
func intSum1(x, y int) int {
	return x + y
}

//无参数无返回值函数
func sayHello() {
	fmt.Println("Hello 沙河")
}

//可变参数
func intSum2(x ...int) int {
	fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

//固定参数搭配可变参数使用时，可变参数要放在固定参数的后面，示例代码如下：
func intSum3(x int, y ...int) int {
	fmt.Println(x, y)
	sum := x
	for _, v := range y {
		sum = sum + v
	}
	return sum
}

//多返回值：	Go语言中函数支持多返回值，函数如果有多个返回值时必须用()将所有返回值包裹起来。
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

//返回值命名： 函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过return关键字返回。
func calc2(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

//全局变量： 定义全局变量num
var num int64 = 10
func testGlobalVar() {
	fmt.Printf("num=%d\n", num) //函数中可以访问全局变量num
}


//局部变量： 局部变量又分为两种： 函数内定义的变量无法在该函数外使用，例如下面的示例代码main函数中无法使用testLocalVar函数中定义的变量x：
//如果局部变量和全局变量重名，优先访问局部变量。
func testLocalVar() {
	//定义一个函数局部变量x,仅在该函数内生效
	var x int64 = 100
	fmt.Printf("x=%d\n", x)
}



