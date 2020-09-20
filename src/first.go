package main

import (
	"fmt"
)

func main() {

	var userVec  []float32

	fmt.Println("hello go ",userVec)

	const age  = 10
	fmt.Println(age)


	var a int = 21
	var b int = 10
	fmt.Println(a + b)

	// for init; condition; post { }
	for i := 0; i <= 10; i++  {
		fmt.Println(i)
	}

    // for condition { }
	i := 0
	for i <= 10 {
		i++
		fmt.Println(i)
	}

    // for { }
	//var aa= 0
	//for {
	//	aa++ // 无限循环下去
	//	fmt.Println(aa)
	//}


	funTest("yangsc","12")

	a,b = swap(1,2);
	fmt.Println(a , b)



	var balance [10] float32
	for i := 0; i < len(balance);i ++ {
		balance[i] = float32(i)
	}
	fmt.Println(balance)

	var balance2 = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance2)


	var aaa int = 10

	fmt.Printf("变量的地址: %x\n", &aaa  )


	var stu Student = Student{12,"yangsc"}
	fmt.Println(stu.age,stu.name)




	numbers := []int{0,1,2,3,4,5,6,7,8}
	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])


	numbers1 := make([]int,0,5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)


	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)

}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

func swap(x, y int) (int, int)  {
	return  y,x
}

func funTest (name , age string)  {
	fmt.Println(name,age)
}

type Student struct {
	age int
	name string
}



func varTest(){

	// 第一种，指定变量类型，如果没有初始化，则变量默认为零值。
	var age int
	age = 10
	fmt.Println(age)

	// 第二种，根据值自行判定变量类型。
	var name  = "yangsanchao"
	fmt.Println(name)

	// 第三种，省略 var, 注意 := 左侧如果没有声明新的变量，就产生编译错误，格式：
	sex := "girl"
	fmt.Println(sex)

}