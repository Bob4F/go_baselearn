package unit
/*
  Go 指针学习
  变量是使用方便的占位符，是应用计算机内存地址
  指针 = 值的内存地址
  如何获取指针 ：&var_name / 变量名前面加上&
 */
import "fmt"


func pointer(){
	var variable int = 20
	var pointer  *int
	pointer = &variable
	fmt.Printf("variable 变量地址:%x\n",&variable)
	fmt.Printf("pointer 存储了指针地址：%x\n",pointer)
	fmt.Printf("获取指针 pointer 对应的值%d\n",*pointer)
}

func Pointer(){
	fmt.Println("Start Pointer")
	pointer()
}


