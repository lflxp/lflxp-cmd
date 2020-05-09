package pkg

import (
	"fmt"
	"os"
	"strings"
)

// 实时左标显示
var LivePrefixState struct {
	LivePrefix string
	IsEnable   bool
}

func changeLivePrefix() (string, bool) {
	return LivePrefixState.LivePrefix, LivePrefixState.IsEnable
}

// 执行器入口
func executor(in string) {
	if in == "" {
		LivePrefixState.IsEnable = false
		LivePrefixState.LivePrefix = in
		return
	} else {
		// 判断输入in是否只含有一个字符
		// 判断Config中是否有alias
		// 没有则正常执行
		cmdline := in
		tmp := strings.Split(in, " ")
		if len(tmp) == 1 && Config.GetString(fmt.Sprintf("%s.alias", in)) != "" {
			cmdline = Config.GetString(fmt.Sprintf("%s.alias", in))
		}
		thisisit, status := ParseExecutors(cmdline)
		if status {
			thisisit()
		}
	}
	LivePrefixState.LivePrefix = "➜ " + GetCurrPath(in) + "> "
	LivePrefixState.IsEnable = true
}

func GetCurrPath(in string) string {
	var rs string
	t := strings.Split(in, " ")
	if t[0] == "cd" {
		if len(t) == 1 {
			err := Chdir("~")
			if err != nil {
				return err.Error()
			}
		} else if len(t) > 1 {
			err := Chdir(t[1])
			if err != nil {
				return err.Error()
			}
		}
		if len(t) == 0 {
			rs = ""
		} else if len(t) == 1 {
			rs = t[0]
		} else if len(t) > 1 {
			rs = strings.Join(t[1:], " ")
		}
	} else {
		rs = in
	}
	// dir, _ := os.Executable()
	// exPath := filepath.Dir(dir)
	// if exPath == string(os.PathSeparator) {
	// 	return string(os.PathSeparator)
	// }
	// tmp := strings.Split(exPath, string(os.PathSeparator))
	// return tmp[len(tmp)-1]

	return rs
}

// Chdir 将程序工作路径修改成程序所在位置
func Chdir(in string) error {
	// fmt.Println(in)
	// dir, err := filepath.Abs(filepath.Dir(in))
	// if err != nil {
	// 	return err
	// }

	// fmt.Println(dir)
	err := os.Chdir(in)
	return err
}
