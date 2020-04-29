package pkg

import (
	"os"
	"os/exec"
)

/** 解析执行命令函数
@param in   // command from
@result func() // function
@result bool // 状态 是否执行
*/
func ParseExecutors(in string) (func(), bool) {
	var result func()
	status := false
	// if in == "dashboard show" {
	// 	result = func() {
	// 		dashboard.Run()
	// 	}
	// 	status = true
	// } else if in == "dashboard helloworld" {
	// 	result = func() {
	// 		helloworld.Run()
	// 	}
	// 	status = true
	// } else if in == "kubectl" {
	// 	result = func() {
	// 		kubectl.ManualInit()
	// 	}
	// 	status = true
	// }

	result = func() {
		cmd := exec.Command("bash", "-c", in)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
	status = true
	return result, status
}
