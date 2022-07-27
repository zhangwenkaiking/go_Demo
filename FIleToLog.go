package godemo

import (
	"fmt"
	"log"
	"os"
	"time"
)

func CreateFile(path string) (error, string) {

	// dir, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println("获取当前路径失败!")
	// 	return err
	// }
	//给文件名加上时间戳
	path = path + "\\" + time.Now().Format("20060102150405") + ".txt"
	fmt.Println("path是:", path)
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("创建失败！！！！")
		return err, ""
	}
	fmt.Println("日志文件成功创建:", path)
	defer file.Close()
	return nil, path
}

//写入方法
func ToWritter(param string, path string) {
	f1, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("open file error :", err)
		return
	}
	_, err2 := f1.WriteString(param) //"writeString:" +
	if err2 != nil {
		log.Println("open file error :", err)
		return
	}
}
