package main

/**
		切片
	切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。
	切片是一个引用类型，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合。

		var name []T
	name:表示变量名
	T:表示切片中的元素类型

	*/

func main() {

	//// 声明切片类型
	//var a []string              //声明一个字符串切片
	//var b = []int{}             //声明一个整型切片并初始化
	//var c = []bool{false, true} //声明一个布尔切片并初始化
	////var d = []bool{false, true} //声明一个布尔切片并初始化
	//fmt.Println(a)              //[]
	//fmt.Println(b)              //[]
	//fmt.Println(c)              //[false true]
	//fmt.Println(a == nil)       //true
	//fmt.Println(b == nil)       //false
	//fmt.Println(c == nil)       //false
	//// fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较


	////切片的底层就是一个数组，所以我们可以基于数组通过切片表达式得到切片。
	////切片表达式中的low和high表示一个索引范围（左包含，右不包含），也就是下面代码中从数组a中选出1<=索引值<4的元素组成切片s，得到的切片长度=high-low，容量等于得到的切片的底层数组的容量。
	//a := [5]int{1, 2, 3, 4, 5}
	//s := a[1:3]  // s := a[low:high]
	//fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	////输出：s:[2 3] len(s):2 cap(s):4


	//为了方便起见，可以省略切片表达式中的任何索引。省略了low则默认为0；省略了high则默认为切片操作数的长度:
	//a[2:]  // 等同于 a[2:len(a)]
	//a[:3]  // 等同于 a[0:3]
	//a[:]   // 等同于 a[0:len(a)]



	//a := [5]int{1, 2, 3, 4, 5}
	//s := a[1:3]  // s := a[low:high]
	//fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	//s2 := s[3:4]  // 索引的上限是cap(s)而不是len(s)
	//fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))
	////输出:
	////	s:[2 3] len(s):2 cap(s):4
	////	s2:[5] len(s2):1 cap(s2):1



	//a := [5]int{1, 2, 3, 4, 5}
	//t := a[1:3:5]
	//fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t))
	////输出：t:[2 3] len(t):2 cap(t):4


	/**
		使用make()函数构造切片:
		make([]T, size, cap)
		T:切片的元素类型
		size:切片中元素的数量
		cap:切片的容量
	 */
		//a := make([]int, 2, 10)
		//fmt.Println(a)      //[0 0]
		//fmt.Println(len(a)) //2
		//fmt.Println(cap(a)) //10

	//切片的本质 https://www.liwenzhou.com/posts/Go/06_slice/

	//判断切片是否为空
	//要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断。


	//切片的赋值拷贝
	//	s1 := make([]int, 3) //[0 0 0]
	//	s2 := s1             //将s1直接赋值给s2，s1和s2共用一个底层数组
	//	s2[0] = 100
	//	fmt.Println(s1) //[100 0 0]
	//	fmt.Println(s2) //[100 0 0]


	//切片遍历:切片的遍历方式和数组是一致的，支持索引遍历和for range遍历。
	//	s := []int{1, 3, 5}
	//
	//	for i := 0; i < len(s); i++ {
	//		fmt.Println(i, s[i])
	//	}
	//
	//	for index, value := range s {
	//		fmt.Println(index, value)
	//	}


	//append()方法为切片添加元素
	//	var s []int
	//	s = append(s, 1)        // [1]
	//	s = append(s, 2, 3, 4)  // [1 2 3 4]
	//	s2 := []int{5, 6, 7}
	//	s = append(s, s2...)    // [1 2 3 4 5 6 7]

	//注意：通过var声明的零值切片可以在append()函数直接使用，无需初始化。
	//var s []int
	//s = append(s, 1, 2, 3)


	////append()添加元素和切片扩容
	//var numSlice []int
	//for i := 0; i < 10; i++ {
	//	numSlice = append(numSlice, i)
	//	fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	//}
	///**
	//[0]  len:1  cap:1  ptr:0xc0000a8000
	//[0 1]  len:2  cap:2  ptr:0xc0000a8040
	//[0 1 2]  len:3  cap:4  ptr:0xc0000b2020
	//[0 1 2 3]  len:4  cap:4  ptr:0xc0000b2020
	//[0 1 2 3 4]  len:5  cap:8  ptr:0xc0000b6000
	//[0 1 2 3 4 5]  len:6  cap:8  ptr:0xc0000b6000
	//[0 1 2 3 4 5 6]  len:7  cap:8  ptr:0xc0000b6000
	//[0 1 2 3 4 5 6 7]  len:8  cap:8  ptr:0xc0000b6000
	//[0 1 2 3 4 5 6 7 8]  len:9  cap:16  ptr:0xc0000b8000
	//[0 1 2 3 4 5 6 7 8 9]  len:10  cap:16  ptr:0xc0000b8000
	//
	//从上面的结果可以看出：
	//	append()函数将元素追加到切片的最后并返回该切片。
	//	切片numSlice的容量按照1，2，4，8，16这样的规则自动进行扩容，每次扩容后都是扩容前的2倍。
	//*/


	////append()函数还支持一次性追加多个元素。 例如：
	//var citySlice []string
	//// 追加一个元素
	//citySlice = append(citySlice, "北京")
	//// 追加多个元素
	//citySlice = append(citySlice, "上海", "广州", "深圳")
	//// 追加切片
	//a := []string{"成都", "重庆"}
	//citySlice = append(citySlice, a...)
	//fmt.Println(citySlice) //[北京 上海 广州 深圳 成都 重庆]


	/**
	切片的扩容策略
	可以通过查看$GOROOT/src/runtime/slice.go源码，其中扩容相关代码如下：

	首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）。
	否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap），
	否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的1/4，即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
	如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）。
	*/

	////使用copy()函数复制切片
	//	a := []int{1, 2, 3, 4, 5}
	//	b := a
	//	fmt.Println(a) //[1 2 3 4 5]
	//	fmt.Println(b) //[1 2 3 4 5]
	//	b[0] = 1000
	//	fmt.Println(a) //[1000 2 3 4 5]
	//	fmt.Println(b) //[1000 2 3 4 5]
	//由于切片是引用类型，所以a和b其实都指向了同一块内存地址。修改b的同时a的值也会发生变化。


	////Go语言内建的copy()函数可以迅速地将一个切片的数据复制到另外一个切片空间中，copy()函数的使用格式如下：
	///**
	//copy(destSlice, srcSlice []T)
	//srcSlice: 数据来源切片
	//destSlice: 目标切片
	//*/
	//	// copy()复制切片
	//	a := []int{1, 2, 3, 4, 5}
	//	c := make([]int, 5, 5)
	//	copy(c, a)     //使用copy()函数将切片a中的元素复制到切片c
	//	fmt.Println(a) //[1 2 3 4 5]
	//	fmt.Println(c) //[1 2 3 4 5]
	//	c[0] = 1000
	//	fmt.Println(a) //[1 2 3 4 5]
	//	fmt.Println(c) //[1000 2 3 4 5]


	////从切片中删除元素
	//	a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	//	// 要删除索引为2的元素
	//	a = append(a[:2], a[3:]...)
	//	fmt.Println(a) //[30 31 33 34 35 36 37]
	////总结一下就是：要从切片a中删除索引为index的元素，操作方法是a = append(a[:index], a[index+1:]...)
}
