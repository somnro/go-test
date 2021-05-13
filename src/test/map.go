package main

import "fmt"

/**
	map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用。
	定义：map[KeyType]ValueType
	KeyType:表示键的类型。
	ValueType:表示键对应的值的类型。

	map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：
	make(map[KeyType]ValueType, [cap])



	*/

func main() {

		//scoreMap := make(map[string]int, 8)
		//scoreMap["张三"] = 90
		//scoreMap["小明"] = 100
		//fmt.Println(scoreMap)
		//fmt.Println(scoreMap["小明"])
		//fmt.Printf("type of a:%T\n", scoreMap)
		///**
		//	map[小明:100 张三:90]
		//	100
		//	type of a:map[string]int
		// */


		////map也支持在声明的时候填充元素，例如：
		//userInfo := map[string]string{
		//	"username": "沙河小王子",
		//	"password": "123456",
		//}
		//fmt.Println(userInfo) //map[password:123456 username:沙河小王子]


	////判断某个键是否存在
	////value, ok := map[key]
	//scoreMap := make(map[string]int)
	//scoreMap["张三"] = 90
	//scoreMap["小明"] = 100
	//// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	//v, ok := scoreMap["张三"]
	//if ok {
	//	fmt.Println(v)//90
	//} else {
	//	fmt.Println("查无此人")
	//}


	////map的遍历
	//	scoreMap := make(map[string]int)
	//	scoreMap["张三"] = 90
	//	scoreMap["小明"] = 100
	//	scoreMap["娜扎"] = 60
	//	for k, v := range scoreMap {
	//		fmt.Println(k, v)//遍历
	//		fmt.Println(k)//只遍历key
	//	}


	////使用delete()函数删除键值对
	///**
	//	delete(map, key)
	//	map:表示要删除键值对的map
	//	key:表示要删除的键值对的键
	// */
	//	scoreMap := make(map[string]int)
	//	scoreMap["张三"] = 90
	//	scoreMap["小明"] = 100
	//	scoreMap["娜扎"] = 60
	//	delete(scoreMap, "小明")//将小明:100从map中删除
	//	for k,v := range scoreMap{
	//		fmt.Println(k, v)
	//	}


	////按照指定顺序遍历map
	//rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	//
	//var scoreMap = make(map[string]int, 200)
	//
	//for i := 0; i < 100; i++ {
	//	key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
	//	value := rand.Intn(100)          //生成0~99的随机整数
	//	scoreMap[key] = value
	//}
	////取出map中的所有key存入切片keys
	//var keys = make([]string, 0, 200)
	//for key := range scoreMap {
	//	keys = append(keys, key)
	//}
	////对切片进行排序
	//sort.Strings(keys)
	////按照排序后的key遍历map
	//for _, key := range keys {
	//	fmt.Println(key, scoreMap[key])
	//}


	//值为切片类型的map
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}
