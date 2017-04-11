// Crawler project main.go
package main

import (
	"fmt"

	"github.com/henrylee2cn/pholcus/exec"
	_ "github.com/henrylee2cn/pholcus_lib"
)

func main() {

	fmt.Println("Hello World!")
	exec.DefaultRun("gui")
}
