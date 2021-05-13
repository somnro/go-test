package main

/**
		数组定义：
		var 数组变量名 [元素数量]T
		var a [3]int
	*/
func main() {

	/**
	  方法一
	 */
	//var testArray [3]int                        //数组会初始化为int类型的零值
	//var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
	//var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
	//fmt.Println(testArray)                      //[0 0 0]
	//fmt.Println(numArray)                       //[1 2 0]
	//fmt.Println(cityArray)                      //[北京 上海 深圳]


	/**
	  方法2
	*/
	//var testArray [3]int
	//var numArray = [...]int{1, 2}
	//var cityArray = [...]string{"北京", "上海", "深圳"}
	//fmt.Println(testArray)                          //[0 0 0]
	//fmt.Println(numArray)                           //[1 2]
	//fmt.Printf("type of numArray:%T\n", numArray)   //type of numArray:[2]int
	//fmt.Println(cityArray)                          //[北京 上海 深圳]
	//fmt.Printf("type of cityArray:%T\n", cityArray) //type of cityArray:[3]string


	/**
	  方法3
	*/
	//a := [...]int{1: 1, 3: 5}
	//fmt.Println(a)                  // [0 1 0 5]
	//fmt.Printf("type of a:%T\n", a) //type of a:[4]int


	//var a = [...]string{"北京", "上海", "深圳"}
	//// 方法1：for循环遍历
	//for i := 0; i < len(a); i++ {
	//	fmt.Println(a[i])
	//}
	//
	//// 方法2：for range遍历
	//for index, value := range a {
	//	fmt.Println(index, value)
	//}

	/**
		多维数组
	 */
	//a := [3][2]string{
	//	{"北京", "上海"},
	//	{"广州", "深圳"},
	//	{"成都", "重庆"},
	//}
	//fmt.Println(a) //[[北京 上海] [广州 深圳] [成都 重庆]]
	//fmt.Println(a[2][1]) //支持索引取值:重庆


	/**
		二维数组遍历
	 */
	//a := [3][2]string{
	//	{"北京", "上海"},
	//	{"广州", "深圳"},
	//	{"成都", "重庆"},
	//}
	//for _, v1 := range a {
	//	for _, v2 := range v1 {
	//		fmt.Printf("%s\t", v2)
	//	}
	//	fmt.Println()
	//}

}
