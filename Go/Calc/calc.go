package calc

//首字母大写 公有方法
func Add(x, y int) int {
	return x + y
}

//小写表示私有 只有在当前包里能访问
func Sub(x, y int) int {
	return x - y
}
