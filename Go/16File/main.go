package main

import (
	"fmt"
	"io"
	"os"
)

// func test1() {
// 	file, _ := os.Open("D:/Github/Learning/Go/16File/main.go")
// 	defer file.Close()
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }
// 	// fmt.Println(file) //返回的file是一个地址
// 	// &{0xc0000a06c8}

// 	//读取文件内的内容
// 	var strSlice []byte
// 	var tempSlice = make([]byte, 128)
// 	for {
// 		n, err := file.Read(tempSlice)
// 		if err == io.EOF {
// 			fmt.Println("读取完毕")
// 			break
// 		}
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		// fmt.Println(n)
// 		// fmt.Println(string(timeSlice))
// 		strSlice = append(strSlice, tempSlice[:n]...)
// 	}
// 	fmt.Println(string(strSlice))
// }

// func test2() {
// 	//bufio读取文件
// 	file, err := os.Open("./main.go")
// 	defer file.Close()

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	var fileStr string
// 	reader := bufio.NewReader(file)
// 	for {
// 		str, err := reader.ReadString('\n') //表示一次读取一行
// 		if err == io.EOF {
// 			fileStr += str
// 			break
// 		}
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		fileStr += str
// 	}
// 	fmt.Println(fileStr)
// }

// ioutil读文件
// func test3() {
// 	byteStr, err := ioutil.ReadFile("./main.go")

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(string(byteStr))
// }

// 文件写入
// func test4() {
// 	file, err := os.OpenFile("D:/Github/Learning/Go/16File/test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
// 	defer file.Close()

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	//写入文件
// 	// for i := 0; i < 10; i++ {
// 	// 	file.WriteString("写入文件\n")
// 	// }
// 	var str = "直接写入\n"
// 	file.Write([]byte(str))

// 	writer := bufio.NewWriter(file)
// 	writer.WriteString("你好golang\n") //这部分会先被写入缓存
// 	writer.Flush()                   //通过Flush可以将缓存中的数据写入文件
// }

// 文件复制
func CopyFile(srcFileName string, dstFileName string) (err error) {
	sFile, err1 := os.Open(srcFileName)
	dFile, err2 := os.OpenFile(dstFileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	defer sFile.Close()
	defer dFile.Close()
	var tempSlice = make([]byte, 128)
	for {
		//读取数据
		n1, err := sFile.Read(tempSlice)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		//写入数据
		if _, err := dFile.Write(tempSlice[:n1]); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	srcFile := "D:/Mad/END/yuyin/sobitterrrrr.mp4"
	dstFile := "D:/Mad/END/yuyin/copy.mp4"
	err := CopyFile(srcFile, dstFile)
	if err == nil {
		fmt.Println("拷贝成功")
	} else {
		fmt.Println("拷贝失败")
	}

}
