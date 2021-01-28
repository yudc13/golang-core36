package internal

import "fmt"

// 使用internal包（模块级私有）仅仅能被当前模块包引用
// 只能被直接父包及其子包中的代码引用

func TestInternal() {
	fmt.Println("internal")
}
