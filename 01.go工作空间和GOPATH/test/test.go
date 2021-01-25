package test

import (
	"fmt"
	lib1 "goCore/01.go工作空间和GOPATH/lib"
	//"goCore/01.go工作空间和GOPATH/lib/internal" // 使用internal的包 编译不通过
)

func Test()  {
	//internal.TestInternal()
	lib1.Hello()
	fmt.Println("test")
}
