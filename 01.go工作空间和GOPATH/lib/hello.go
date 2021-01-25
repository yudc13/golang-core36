package lib1

import (
	"fmt"
	"goCore/01.go工作空间和GOPATH/lib/internal"
)

func Hello()  {
	internal.TestInternal()
	fmt.Println("lib1 Hello")
}
