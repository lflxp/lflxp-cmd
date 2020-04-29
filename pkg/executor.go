package pkg

import (
	"fmt"
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
	LivePrefixState.LivePrefix = in + "> "
	LivePrefixState.IsEnable = true
}
