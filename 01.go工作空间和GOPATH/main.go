package main

import (
	"flag"
	"fmt"
	lib1 "goCore/01.go工作空间和GOPATH/lib"
	"os"
)

func init() {
	lib1.Hello()
}

func main() {
	name := flag.String("name", "everyone", "姓名")
	age := flag.Int("age", 30, "年龄")
	flag.Usage = func() {
		_, _ = fmt.Fprintln(os.Stderr, "Usage of question")
		flag.PrintDefaults()
	}
	flag.Parse()
	fmt.Println("main", *name, age)
}
