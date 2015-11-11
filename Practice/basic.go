package main

import (
	"fmt"
	"runtime"
)

//自定义类型
type (
	IZ  int
	FZ  float64
	STR string
)

//private 首字母小写
var gVal  int = 100 //全局变量声明后，未使用，也可以编译通过

//public 首字母大写
var GVal2 int = 200 



// init 函数  ====================
/*
  变量除了可以在全局声明中初始化，也可以在 init 函数中初始化。
  这是一类非常特殊的函数，它不能够被人为调用，而是在每个包完成初始化后自动执行，并且执行优先级比 main 函数高。
  每一个源文件都可以包含且只包含一个 init 函数。初始化总是以单线程执行，并且按照包的依赖关系顺序执行。
  用途1：在开始执行程序之前对数据进行检验或修复，以保证程序状态的正确性。
*/

/*
用途2：
  init 函数也经常被用在当一个程序开始之前调用后台执行的 goroutine，如下面这个例子当中的 backend()：

	func init() {
	// setup preparations
	go backend()
	}

*/

func init() {
	fmt.Print("enter init() \n")
	fmt.Printf("go version ==> %s \n", runtime.Version()) //类似 POSIX 的printf
/*
	var a IZ = 100
	var b FZ = 3.14
	var c STR = "justString"



//printf与printfln的区别
	fmt.Print(a, b, c) //Print() 不会自动添加 回车&空格


	//如果2个类型 都是 数值类型，写在一起，会自动添加了空格
	fmt.Print(a,a,b,a,b,b,a,a)

	//分开写 没有添加空格
	fmt.Print(a)
	fmt.Print(b)
	
	//fmt.Println(234)
	fmt.Println("==分隔符=====")

	//Println()自动添加 空格
	fmt.Println(a, b, c)

	//Println()自动添加 换行符
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Print(  "a=", a, ",b=", b, ",c=", c )//与字符一起输出
	fmt.Println("a=", a, ",b=", b, ",c=", c )//与字符一起输出
	fmt.Printf("a=%d, b=%d, c=%s \n",a,b,c)//与C语言类似的格式化输出
*/
	fmt.Print("leave init() \n")
	
	
}

func main() {
	
//声明 + 赋值 =============================
	var val int32 //默认初始值 为0
	var val2 int32 = 223 //显式声明 赋值
	fmt.Println("val=",val,", val2=",val2)//局部变量。 如果未使用，则编译报错，所以都加了Println

   

//全局变量================
	//fmt.Print(G_val) //全局变量声明后，未使用。g 也可以编译通过
	
	var gVal int32 =101 //函数体中声明的局部变量（变量名名 与全局变量 相同），此函数体使用的是局部变量 
	GVal2 = 300
	
	fmt.Printf("gVal=%d, GVal2=%d \n",gVal,GVal2)


//显式声明,(隐式)声明 ,并行赋值 ===========

	var a, b, c int //显式声明 初始值 为0
	a, b, c = 5, 7, 9 //并行赋值
	fmt.Printf("a=%d, b=%d, c=%d \n",a,b,c)

	

	a2 := 123          //(隐式)声明 + 赋值, 自动根据值 来匹配类型
	a3, b3, c3 := 5, 7, "abc" //(隐式)声明 + 并行赋值
	fmt.Printf("a2=%d \n", a2)
	fmt.Printf("a3=%d, b3=%d, c3=%s \n",a3,b3,c3)


//转换类型 必须显式====================
	var a1 IZ = 5
	b1 := int(a1)
	c1 := IZ(b1)
	fmt.Printf("a1=%d , b1=%d, c1=%d \n",a1,b1,c1)



//函数调用，函数多返回值===============
	ret1, ret2 := funcA(  int32(a), val2 ) //函数返回多个值
	fmt.Println(ret1, ret2)



//两个变量的值，则可以简单地使用 a, b = b, a ==============
	a,b = b,a

//空白标识符 `_` 也被用于抛弃值，
	_, b = 5, 7 //如值 `5` 被抛弃。

	

}

func funcA(a int32, b int32) (ret1 int32, ret2 int32) {
	//fmt.Println("enter funA ", a,b)
	//fmt.Println(a+b)
	return a + 1, b + 1
}
