// https://www.runoob.com/go/go-for-loop.html
/*
Go 语言的 For 循环有 3 种形式，只有其中的一种使用分号。

和 C 语言的 for 一样：

for init; condition; post { }
和 C 的 while 一样：

for condition { }
和 C 的 for(;;) 一样：

for { }

for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：

for key, value := range oldMap {
    newMap[key] = value
}
*/
package main

import "fmt"

func main() {

	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}

	fmt.Println(sum)
}
