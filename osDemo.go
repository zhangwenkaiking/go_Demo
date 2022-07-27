package main

import (
	"GO/godemo"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

//C:\Users\sky\Desktop\GoDemo\GO\osDemo.go
var wg sync.WaitGroup

func main() {
	//移动后的开头路径
	var baocunFile string
	var check int
	copyright()
	fmt.Scan(&check)
	if check != 123456 {
		return
	}
	flag := 0 //标志位
	var bar godemo.Bar

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前路径失败!")
		return
	}

	// dir := "C:\\Users\\sky\\Desktop\\测试"
	fmt.Println("当前路径是:", dir)
	fmt.Println("当前目录有如下文件夹:")
	files2, _ := ListDir(dir, "")
	// fmt.Println(files2)

	//科室数量
	var number int
	//科室代号
	var classNumber int
	className := files2

	for key, value := range className {
		fmt.Println(key, ":", value)
	}
	// //将科室切片的值复制一份给保留科室切片
	// classSlice = className
	fmt.Println("请输入移动后的‘绝对’路径:")
	fmt.Scan(&baocunFile)
	fmt.Println("请输入想保留的科室数量:")
	fmt.Scan(&number)
	number++
	//存储保存的科室的数量的切片
	classSlice := make([]string, number)
	classSlice[0] = "日志"
	for i := 1; i < number; i++ {
		fmt.Println("请输入想保留的科室的代号:")
		fmt.Scan(&classNumber)
		for classNumber > len(className) {
			fmt.Println("输入有误!!!!")
			fmt.Println("请重新输入想保留的科室的代号:")
			fmt.Scan(&classNumber)
		}
		fmt.Println("classNumber:", classNumber)
		//将要保留的科室存储到新切片中
		classSlice[i] = className[classNumber]
	}
	fmt.Println("保留的科室是:")
	for key, value := range classSlice {
		if value == "" {
			continue
		}
		fmt.Println(key, ":", value)
	}

	fmt.Println("********暂停5s********")
	for i := 1; i < 5; i++ {
		//暂停5s
		fmt.Println("********暂停", i, "s********")
		time.Sleep(time.Second * 1)
	}

	path := dir
	files, _ := WalkDir(path, "")
	//显示文件夹下的文件
	//showFile(files)
	//判断保留文件
	thisFiles, flag1 := judgeClass(classSlice, files, flag)
	bar.NewOption(0, int64(len(thisFiles)-flag1-2))
	//创建日志文件夹
	path = path + "\\日志"
	os.Mkdir(path, os.ModePerm)
	//创建日志文件
	_, newFile := godemo.CreateFile(path)
	//开始删除文件
	// wg.Add(len(thisFiles))
	wg.Add(len(thisFiles))
	//	newFile2 := "C:/Users/sky/Desktop/temp2" + "\\" + time.Now().Format("20060102150405") + ".txt"
	newFile2 := newName(thisFiles)
	go removeToFile(thisFiles, bar, newFile2, newFile, baocunFile)
	bar.Finish()
	//暂停10s
	time.Sleep(time.Second * 100)
	wg.Wait()
}

//截取文件名
func newName(thisFiles []string) []string {
	newArray := make([]string, 1024*10*10*4)
	for i := 0; i < len(thisFiles); i++ {
		i2 := strings.LastIndex(thisFiles[i], "\\")
		newArray[i] = thisFiles[i][i2+1:]
	}
	return newArray
}

//保留功能
func judgeClass(classSlice, files []string, flag int) ([]string, int) {
	//判断子串
	for i := 0; i < len(files); i++ {
		for j := 0; j < len(classSlice); j++ {
			isok := strings.Contains(files[i], classSlice[j])
			// fmt.Println("classSilce是:", classSlice)
			// fmt.Println("isok是:", isok)
			// fmt.Printf("classSlice[%v]是:%v", j, classSlice[j])
			if isok {
				files[i] = ""
				//记录保留的个数
				flag++
			}
			// if strings.Contains(files[i], "日志") {
			//  files[i] = ""
			// }
		}
	}
	return files, flag
}

//删除功能
// func removeFile(files []string, bar godemo.Bar, path string) {
// 	// for key, value := range files {
// 	//  fmt.Println(key, ":", value)
// 	//  //删除该文件
// 	//  err := os.Remove(value)
// 	//  if err != nil {
// 	//      fmt.Println(value, "删除文件失败:")
// 	//  } else {
// 	//      fmt.Println(value, "删除文件成功！")
// 	//  }
// 	// }
// 	for i := 0; i < len(files); i++ {
// 		//fmt.Println(i, ":", files[i])
// 		//删除该文件
// 		err := os.Remove(files[i])
// 		if err != nil {
// 			// fmt.Println()
// 			// fmt.Println(files[i], "删除文件失败:")
// 		} else {
// 			// fmt.Println(files[i], "删除文件成功！")
// 			writeValue := time.Now().Format("20060102150405") + ": --> " + strconv.Itoa(i) + files[i] + "\n"
// 			godemo.ToWritter(writeValue, path)
// 			godemo.ToWritter("\r\n", path)
// 			bar.Play(int64(i))
// 		}

// 	}
// 	fmt.Println("      *****删除完成!!!!*****")
// 	wg.Done()
// }
//移动功能
func removeToFile(files []string, bar godemo.Bar, newpath []string, path string, baocunFile string) {
	// for key, value := range files {
	//  fmt.Println(key, ":", value)
	//  //删除该文件
	//  err := os.Remove(value)
	//  if err != nil {
	//      fmt.Println(value, "删除文件失败:")
	//  } else {
	//      fmt.Println(value, "删除文件成功！")
	//  }
	// }
	for i := 0; i < len(files); i++ {
		//fmt.Println(i, ":", files[i])
		//删除该文件
		err := os.Rename(files[i], baocunFile+"\\"+"("+strconv.Itoa(i)+")"+newpath[i])
		if err != nil {
			// fmt.Println()
			// fmt.Println(files[i], "删除文件失败:")
		} else {
			// fmt.Println(files[i], "删除文件成功！")
			writeValue := time.Now().Format("20060102150405") + ": --> " + "(" + strconv.Itoa(i) + ")" + files[i] + "\n"
			godemo.ToWritter(writeValue, path)
			godemo.ToWritter("\r\n", path)
			bar.Play(int64(i))
		}

	}
	fmt.Println(" *****移动完成!!!!*****")
	fmt.Println(" *****移动到了:", baocunFile)
	wg.Done()
}

//显示文件夹下的全部文件
func showFile(files []string) {
	for key, value := range files {
		fmt.Println(key, ":", value)
	}
}
func WalkDir(dir, suffix string) (files []string, err error) {
	files = []string{}
	err = filepath.Walk(dir, func(fname string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			//忽略目录
			return nil
		}

		if len(suffix) == 0 || strings.HasSuffix(strings.ToLower(fi.Name()), suffix) {
			//文件后缀匹配
			files = append(files, fname)
		}

		return nil
	})

	return files, err
}

/* 获取指定路径下的所有文件，只搜索当前路径，不进入下一级目录，可匹配后缀过滤（suffix为空则不过滤）*/
func ListDir(dir, suffix string) (files []string, err error) {
	files = []string{}

	_dir, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	suffix = strings.ToLower(suffix) //匹配后缀

	for _, _file := range _dir {
		if _file.IsDir() {
			files = append(files, _file.Name())
		}
		if len(suffix) == 0 || strings.HasSuffix(strings.ToLower(_file.Name()), suffix) {
			//文件后缀匹配
			//   files = append(files, path.Join(dir, _file.Name()))
			continue //忽略
		}
	}

	return files, nil
}

//版权声明
func copyright() {
	// str := "*"
	// fmt.Print("copyright@版权声明")
	// for i := 0; i < 15; i++ {
	//  fmt.Print(str)
	// }
	// for j := 0; j < 20; j++ {
	//  for i := 0; i < 105; i++ {
	//      if j >= 0 && j < 3 {
	//          fmt.Printf(str)
	//      } else if j > 20-4 && j < 20 {
	//          fmt.Printf(str)
	//      } else if j >= 3 && j <= 20-4 && i >= 3 && i <= 105-3 {
	//          fmt.Printf(" ")
	//      } else if i >= 0 && i < 2 {
	//          fmt.Printf(str)
	//      } else if i >= 105-3 && i <= 105 {
	//          fmt.Printf(str)
	//      } else if j == 4 {
	//          fmt.Printf("**                                        copyright@版权声明                                         **")
	//      }
	//  }
	//  fmt.Println()
	// }
	fmt.Println("******************************************************************************************************")
	fmt.Println("**                                        copyright@版权声明                                        **")
	fmt.Println("**         1.功能:删除当前文件夹下的文件,可以选择保留特定文件夹下的文件。                           **")
	fmt.Println("**         2.注意:如果文件不在文件夹内，将无法受到保护，会被删除，删除的文件不会出现在回收站!!!     **")
	fmt.Println("**             请谨慎操作!!!                                                                        **")
	fmt.Println("**         3.本程序为自用程序,请勿移作他用,否则后果自负!!!                                          **")
	fmt.Println("**         4.删除文件数量大于300w个时可能会不成功，请谨慎删除！                                     **")
	fmt.Println("**         5.功能可定制 联系QQ:313426871                                                            **")
	fmt.Println("**                                                                                                  **")
	fmt.Println("**                                                                                                  **")
	fmt.Println("**                                   最终解释权归程序开发者所有                                     **")
	fmt.Println("**                                                                                                  **")
	fmt.Println("**                                        version0.1                                                **")
	fmt.Println("**                                                                                                  **")
	fmt.Println("**                                        2022/01/07                                                **")
	fmt.Println("**                                                                                                  **")
	fmt.Println("******************************************************************************************************")
}
