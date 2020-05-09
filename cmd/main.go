package main

import "github.com/lflxp/lflxp-cmd/pkg"

func main() {
	pkg.Exec()
}

// func main() {
//     cli := ssh.Ssh("1.1.1.1", "root", "xxx", 22)
//     output, err := cli.Run("df -h")
//     fmt.Printf("%v\n%v", output, err)
// }
