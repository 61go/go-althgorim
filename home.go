package main

import "fmt"

/**
  利用斐波契那数列解决青蛙跳台阶问题
*/

/**
问题描述：
一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个n级的台阶总共需要多少种跳法。

思路：首先考虑n等于0、1、2时的特殊情况，f(0) = 0   f(1) = 1  f(2) = 2

            其次，当n=3时，青蛙的第一跳有两种情况：跳1级台阶或者跳两级台阶

            假如跳一级，那么 剩下的两级台阶就是f(2)；假如跳两级，那么剩下的一级台阶就是f(1)，因此f(3)=f(2)+f(1)

                当n = 4时，f(4) = f(3) +f(2),以此类推...........可以联想到Fibonacci数列
---------------------
作者：imjunior
来源：CSDN
原文：https://blog.csdn.net/imsenior/article/details/53207788
版权声明：本文为博主原创文章，转载请附上博文链接！
*/

var tmp map[int]int

func main() {
	//跳1级
	fmt.Println("跳一级时：")
	fmt.Println(fiber(1))
	fmt.Println("跳二阶时：")
	fmt.Println(fiber(2))
	fmt.Println("跳三阶时：")
	fmt.Println(fiber(3))
	fmt.Println("跳4阶时：")
	fmt.Println(fiber(4))
	fmt.Println("跳10阶时：")
	fmt.Println(fiber(999))
	tmp = make(map[int]int)

}

/**
原生fiber函数
*/

func fiber(n int) int {
	v, ok := tmp[n]
	if ok {
		return v
	}
	if n <= 2 {
		tmp[n] = n
		return n
	}
	tmp[n] = fiber()
	return fiber(n-1) + fiber(n-2)
}
