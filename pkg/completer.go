package pkg

import (
	"strings"

	"github.com/c-bata/go-prompt"
)

// 功能筛选信息
func completer(in prompt.Document) []prompt.Suggest {
	// 常规过滤
	// 如果输入值为空 返回空字符串
	if in.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}

	// 	获取所有输入字符串并以空格分割
	args := strings.Split(in.TextBeforeCursor(), " ")
	// 获取当前输入字符
	current := in.GetWordBeforeCursor()
	// log.Debug(current)

	// If PIPE is in text before the cursor, returns empty suggestions.
	for i := range args {
		if args[i] == "|" {
			return []prompt.Suggest{}
		}
	}

	// If word before the cursor starts with "-", returns CLI flag options.
	if strings.HasPrefix(current, "-") {
		return OptionsCompleters(args)
	}

	// 功能列表 排除包含“-”的字符，遇到-则返回空交由下面函数处理
	// if suggests, found := completers.GlobalOptionFunc(in); found {
	// 	return suggests
	// }

	// 输入即取消提示
	// 非常规过滤

	// s := []prompt.Suggest{
	// 	{Text: "users", Description: "Store the username and age"},
	// 	{Text: "articles", Description: "Store the article text posted by user"},
	// 	{Text: "comments", Description: "Store the text commented to articles"},
	// }

	// return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
	return FirstCommandFunc(in, args)
}
