package main

import (
    "fmt"
    "runtime"
)


//自定义类型
	type (
   IZ int
   FZ float64
   STR string
)



//特殊的函数，每个含有该函数的包都会首先执行这个函数
func init(){
	fmt.Print("enter init() \n")
	fmt.Printf("go version ==> %s \n", runtime.Version())//类似 POSIX 的printf
	
	var a IZ = 100
	var b FZ = 3.14
	var c STR = "justString"

	fmt.Print(a,b,c)//Print() 不会自动添加 回车&空格
	
/*	
	//如果2个类型 都是 数值类型，写在一起，会自动添加了空格
	fmt.Print(a,a,b,a,b,b,a,a)
	
	//分开写 没有添加空格
	fmt.Print(a)  
	fmt.Print(b)
*/	
  //fmt.Println(234)
	fmt.Println("==分隔符=====")
	
	
	//Println()自动添加 空格
	fmt.Println(a,b,c)
	
	//Println()自动添加 换行符
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	
	fmt.Print("leave init() \n")
}

func main() {
	
  var val int32  //初始值 为0
	fmt.Println(val)
	
	a := 123;//自动根据值 来匹配类型
	var b int32 = 223;//显式 赋值
	fmt.Println(a,b)

//必须显式 转换类型	
	var a1 IZ = 5
	b1 := int(a1)
	c1 := IZ(b1)
  fmt.Println(a1, b1,c1)

	fmt.Println( funcA( int32(a), b) )// a 由int显式转换成int32, 各个
}




func funcA(a int32, b int32)(ret1 int32, ret2 int32){
	//fmt.Println("enter funA ", a,b)
	//fmt.Println(a+b)
	return a+1,b+1
}
