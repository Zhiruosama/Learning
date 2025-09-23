package main

import (
	"fmt"
	"strings"
)

func main() {
	//make创建Map
	// var userinfo = make(map[string]string)
	// userinfo["username"] = "老子"
	// userinfo["age"] = "20"
	// userinfo["sex"] = "男"

	// var userinfo = map[string]string{
	// 	"username": "张三",
	// 	"age":      "20",
	// 	"sex":      "男",
	// }

	// userinfo := map[string]string{
	// 	"username": "张三",
	// 	"age":      "20",
	// 	"sex":      "男",
	// }
	// fmt.Println(userinfo)

	//用for range进行内容遍历
	// userinfo := map[string]string{
	// 	"username": "张三",
	// 	"age":      "20",
	// 	"sex":      "男",
	// }
	// for k, v := range userinfo {
	// 	fmt.Printf("key:%v value:%v\n", k, v)
	// }

	//判断map中是否有某个内容
	// v1, ok1 := userinfo["age"]
	// v2, ok2 := userinfo["inte"]
	// fmt.Println(v1, ok1)
	// fmt.Println(v2, ok2)

	//数据删除
	// delete(userinfo, "sex")
	// fmt.Println(userinfo)

	//切片综合map
	// var userinfo = make([]map[string]string, 3, 3)
	// if userinfo[0] == nil {
	// 	userinfo[0] = make(map[string]string)
	// 	userinfo[0]["username"] = "张三"
	// 	userinfo[0]["age"] = "20"
	// }
	// if userinfo[1] == nil {
	// 	userinfo[1] = make(map[string]string)
	// 	userinfo[1]["username"] = "李四"
	// 	userinfo[1]["age"] = "21"
	// }
	// if userinfo[2] == nil {
	// 	userinfo[2] = make(map[string]string)
	// 	userinfo[2]["username"] = "王五"
	// 	userinfo[2]["age"] = "22"
	// }
	// for _, v := range userinfo {
	// 	fmt.Println(v)
	// 	for k, d := range v {
	// 		fmt.Printf("%v-%v\n", k, d)
	// 	}
	// }
	// var userinfo = make(map[string][]string)
	// userinfo["username"] = []string{"那个人"}
	// userinfo["hobby"] = []string{"吃饭", "睡觉"}
	// fmt.Println(userinfo)

	//map类型也是引用数据类型

	var str = "how do you do"
	var strSlice = strings.Split(str, " ")
	fmt.Println(strSlice)
	var strMap = make(map[string]int)
	for _, v := range strSlice {
		strMap[v]++
	}
	fmt.Println(strMap)
}
