package main

import "fmt"

var (
	str01, str02 string
)

func get_sum(pa01, pa02 int) int {
	/**
	两个数求和
	 */
	return pa01 + pa02
}

type mUser struct {
	/**
	定义结构体
	 */
	username string
	password string
}

func parsemUser(po mUser) {
	/**
	结构体作为函数参数
	 */
	fmt.Println("begin parse user...")
	fmt.Println(po.username, po.password)
}

func main() {
	for true {
		var func_selection string

		fmt.Print("please select function : ")
		fmt.Scanln(&func_selection)

		fmt.Println("your input is --> ", func_selection, "!")

		if func_selection == "exit" || func_selection == "q" {
			break
		} else if func_selection == "sum" { //golang调用函数求和
			var num01, num02, sum_num int

			fmt.Print("please input num 01: ")
			fmt.Scanln(&num01)
			fmt.Print("please input num 02: ")
			fmt.Scanln(&num02)
			sum_num = get_sum(num01, num02)
			fmt.Println("sum result is : ", num01, "+", num02, "=", sum_num)
		} else if func_selection == "array" { // 数组使用
			var demo_array [6] int
			for i := 1; i < 7; i++ {
				demo_array[i-1] = i
			}
			fmt.Println(demo_array)
			for j := 6; j > 0; j-- {
				demo_array[6-j] = demo_array[j-1]
			}
			fmt.Println(demo_array)
		} else if func_selection == "point" { //指针使用
			var num int = 30
			var addr *int

			addr = &num

			fmt.Println(" num=", num, "\r\n &num=", &num, "\r\n addr=", addr, "\r\n &addr=", &addr)
		} else if func_selection == "struct" { //结构体使用
			xm := mUser{"xiaoming", "123456"}
			parsemUser(xm)
		}

	}

}
