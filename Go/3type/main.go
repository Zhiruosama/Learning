package main

import (
	"fmt"
	"strconv"
)

func main() {
	//int类型转换
	// var a1 int32 = 10
	// var a2 int64 = 20
	// // fmt.Println(int64(a1) + a2)
	// fmt.Println(a1 + int32(a2))
	// //高位向低位转换时注意不要超出范围

	// num := 30
	// fmt.Println(unsafe.Sizeof(num))
	// //表示64位计算机 int就是int64 占用8个字节

	// var a float32 = 3.12
	// fmt.Printf("值:%v--%f,类型%T", a, a, a)
	// %f表示输出float类型 %.2f表示保留两位小数

	// 科学计数法表示
	// var f float32 = 3.14e2 //3.14*10的二次方
	// fmt.Printf("值:%v--%f,类型%T", f, f, f)

	// float存在精度丢失问题
	// var f float64 = 1129.6
	// fmt.Println(f * 100)

	// GO中bool类型只有true跟false
	// 不允许将整型转换为布尔型
	// 也不存在1为true 0为false

	//strings.Split分割字符串
	// var str1 = "123-456-789"
	// arr := strings.Split(str1, "-")
	// fmt.Println(arr)

	//strings.Join合并字符串
	// var str1 = "123-456-789"
	// arr := strings.Split(str1, "-")
	// str2 := strings.Join(arr, "*")
	// fmt.Println(str2)

	//strings.HasPrefix/HasSuffix前后缀判断
	// str1 := "this is str"
	// str2 := "this"
	// flag := strings.HasPrefix(str1, str2)
	// fmt.Println(flag)

	//strings.Index子串出现的位置
	//LastIndex最后出现的位置
	// str1 := "this is str"
	// str2 := "is"
	// Index := strings.Index(str1, str2)
	// fmt.Println(Index)
	//查找不到返回-1

	//字符属于int类型
	//直接输出byte的时候输出的是其对应的ASCII码
	// var a = 'a'
	// fmt.Printf("%v-%T\n", a, a)
	// fmt.Printf("%c-%T", a, a)

	//循环输出字符串中字符
	//rune代表的是UTF-8 所以可以处理中文
	// s1 := "hello golong"
	// for i := 0; i < len(s1); i++ { //byte
	// 	fmt.Printf("%v(%c)", s1[i], s1[i])
	// }
	// s2 := "你好 golong"
	// for _, r := range s2 { //rune
	// 	fmt.Printf("%v(%c)", r, r)
	// }

	//使用Sprintf进行其他类型到string的转换
	// i := 20
	// str1 := fmt.Sprintf("%d", i)
	// fmt.Printf("%v-%T", str1, str1)

	//string转换为整型
	str := "123456"
	num, _ := strconv.ParseInt(str, 10, 64)
	fmt.Printf("%v-%T", num, num)
}
