package main

import "fmt"

var (
	str01,str02 string
)


func get_sum(pa01,pa02 int) int {
	return pa01 + pa02
}

func main(){
	for true {
		var func_selection string

		fmt.Print("please input : ")
		fmt.Scanln(&func_selection)

		fmt.Println("your input is --> ",func_selection,"!")

		if func_selection == "exit" {
			break
		}else if func_selection == "sum" { //golang调用函数求和
			var num01,num02,sum_num int

			fmt.Print("please input num 01: ")
			fmt.Scanln(&num01)
			fmt.Print("please input num 02: ")
			fmt.Scanln(&num02)
			sum_num = get_sum(num01,num02)
			fmt.Println("sum result is : ",num01,"+",num02,"=",sum_num)
		}else if func_selection == "array"{ // 数组使用
			var demo_array [6] int
			for i:=1;i<7;i++{
				demo_array[i-1] = i
			}
			fmt.Println(demo_array)
			for j:=6;j>0;j--{
				demo_array[6-j] = demo_array[j-1]
			}
			fmt.Println(demo_array)
		}

	}

}
