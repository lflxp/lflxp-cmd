package pkg

import (
	"fmt"
	"strings"

	"github.com/c-bata/go-prompt"
	log "github.com/sirupsen/logrus"
)

// TODO: 解析alias
var (
	Commands []prompt.Suggest
	count    map[string]string
)

func initCommands() {
	Commands = []prompt.Suggest{}
	// TODO: 解析alias无效
	// result, err := ExecCommandString("alias")
	// if err != nil {
	// 	panic(err)
	// }
	// log.Info(result)
	// for _, cmd := range strings.Split(result, "\n") {
	// 	current := strings.TrimSpace(cmd)
	// 	if current != "" {
	// 		log.Info(cmd)
	// 		tmp := strings.Split(cmd, "=")
	// 		t1 := prompt.Suggest{
	// 			Text:        tmp[0],
	// 			Description: strings.Replace(tmp[1], "'", "", -1),
	// 		}
	// 		Commands = append(Commands, t1)
	// 	}
	// }

	// TODO: 解析配置文件
	// Config.GetStringMapStringSlice()
	count = map[string]string{}
	allk := Config.AllKeys()
	log.Debug("AllKeys ", allk)
	for _, cmd := range allk {
		t1 := strings.Split(cmd, ".")
		if _, ok := count[t1[0]]; !ok {
			tmp := prompt.Suggest{
				Text:        t1[0],
				Description: fmt.Sprintf("说明:%s 快捷键alias: %s", Config.GetString(fmt.Sprintf("%s.desc", t1[0])), Config.GetString(fmt.Sprintf("%s.alias", t1[0]))),
			}
			count[t1[0]] = "1"
			Commands = append(Commands, tmp)
		}
	}

	log.Debug("AllSettings ", Config.AllSettings())
}

// TODO: 解析子配置

// 用户自定义命令
func FirstCommandFunc(in prompt.Document, args []string) []prompt.Suggest {
	if len(args) <= 1 {
		// log.Error("alias", Config.GetString(fmt.Sprintf("%s.alias", in.Text)))
		// return prompt.FilterHasPrefix(Commands, args[0], true)
		return prompt.FilterFuzzy(Commands, args[0], true)
	}

	first := args[0]
	// 过滤 -
	// 截取
	// 获取并判断
	latest := []string{}
	for _, t := range args {
		if strings.Index(t, "-") != 0 {
			latest = append(latest, t)
		}
	}
	for k, _ := range count {
		if k == first {
			second := args[len(args)-1]
			subcommands := []prompt.Suggest{}
			// 处理基础命令持续可查看功能
			if len(args) == 2 {
				cmds := fmt.Sprintf("%s.command", latest[0])
				log.Debug("1 ", cmds)
				log.Debug("1", Config.GetString(cmds))
				log.Debug("1", Config.GetStringMap(cmds))
				log.Debug("1", Config.GetStringSlice(cmds))
				log.Debug("1", Config.GetStringMapString(cmds))

				// map[strnig][]string
				if rs := Config.GetStringMapStringSlice(cmds); rs != nil {
					// 	log.Info("数据为map[string][]string", rs)
					for k, v := range rs {
						log.Debug("1 ", k, v, len(v))
						if len(v) == 1 {
							// TODO: 小bug len(values) == 1 判断是
							if v[0] != "" && v[0] != " " {
								subcommands = append(subcommands, prompt.Suggest{
									Text:        k,
									Description: v[0],
								})
							}
						} else if len(v) > 1 {
							subcommands = append(subcommands, prompt.Suggest{
								Text:        k,
								Description: GetCommon(args[0], k),
							})
						}

					}
				}

				// map[string]interface{}
				for k, v := range Config.GetStringMap(cmds) {
					switch t := v.(type) {
					case string:
						log.Debug("1 string")
						subcommands = append(subcommands, prompt.Suggest{
							Text:        k,
							Description: t,
						})
					case map[string]interface{}:
						log.Debug("1 map[string]interface{}")
						subcommands = append(subcommands, prompt.Suggest{
							Text:        k,
							Description: GetCommon(args[0], k),
						})
					case []string:
						log.Debug("1 []]string")
						for _, x := range t {
							subcommands = append(subcommands, prompt.Suggest{
								Text:        strings.Split(x, ",")[0],
								Description: strings.Split(x, ",")[1],
							})
						}
					case map[string][]string:
						log.Debug("1 map[string][]string", t)
					case map[string]string:
						log.Debug("1 map[string]string", t)
					}
				}

				// []string
				if rs := Config.GetStringSlice(cmds); len(rs) > 0 {
					// 数据为[]string
					log.Debug("7 rs []string", rs)
					for _, x := range rs {
						if strings.ContainsAny(x, ",") {
							t1 := strings.Split(x, ",")
							subcommands = append(subcommands, prompt.Suggest{
								Text:        t1[0],
								Description: t1[1],
							})
						}
					}
				}
			}
			// 处理非基础命令查看功能
			if len(args) > 2 {
				// 过滤一遍 去除空格导致数据无法获取
				filterLatest := []string{}
				for _, x := range latest {
					if x != "" && x != " " {
						filterLatest = append(filterLatest, x)
					}
				}

				// second = filterLatest[len(filterLatest)-1]

				var cmds string
				// 处理 command获取问题
				// 处理 输入字符为空的状态持续可查询问题
				if len(args) == 2 && args[1] == "" {
					cmds = fmt.Sprintf("%s.command", args[0])
				} else if len(args) == 2 && args[1] != "" {
					cmds = fmt.Sprintf("%s.command.%s", args[0], args[1])
				} else if len(args) > 2 && args[len(args)-1] == "" {
					cmds = fmt.Sprintf("%s.command.%s", args[0], strings.Join(args[1:len(args)-1], "."))
				} else if len(args) > 2 && args[len(args)-1] != "" {
					cmds = fmt.Sprintf("%s.command.%s", args[0], strings.Join(args[1:len(args)-1], "."))
				}

				log.Debug("2 cmds ", cmds)
				log.Debug("2 second ", second)
				log.Debug("2", Config.GetString(cmds))
				log.Debug("2", Config.GetStringMap(cmds))
				log.Debug("2", Config.GetStringSlice(cmds))

				// map[strnig][]string
				if rs := Config.GetStringMapStringSlice(cmds); rs != nil {
					// 	log.Info("数据为map[string][]string", rs)
					for k, v := range rs {
						log.Debug("1 ", k, v, len(v))
						if len(v) == 1 {
							// TODO: 小bug len(values) == 1 判断是
							if v[0] != "" && v[0] != " " {
								subcommands = append(subcommands, prompt.Suggest{
									Text:        k,
									Description: v[0],
								})
							}
						} else if len(v) > 1 {
							// 处理子目录说明
							// 包含【，】获取，不包含取消
							// if strings.ContainsAny(k, "-") {
							// 	subcommands = append(subcommands, prompt.Suggest{
							// 		Text:        strings.Split(k, "-")[0],
							// 		Description: strings.Split(k, "-")[1],
							// 	})
							// } else {
							// 	subcommands = append(subcommands, prompt.Suggest{
							// 		Text:        k,
							// 		Description: "子目录，非功能性命令",
							// 	})
							// }

							subcommands = append(subcommands, prompt.Suggest{
								Text:        k,
								Description: GetCommon(args[0], k),
							})
						}

					}
				}

				// 底层数据map[string]interface{}
				for k, v := range Config.GetStringMap(cmds) {
					switch t := v.(type) {
					// case string:
					// 	log.Info("2 数据为string")
					// 	subcommands = append(subcommands, prompt.Suggest{
					// 		Text:        k,
					// 		Description: t,
					// 	})
					case map[string][]string:
						log.Debug("3 数据为string")
						for k, _ := range t {
							subcommands = append(subcommands, prompt.Suggest{
								Text:        k,
								Description: GetCommon(args[0], k),
							})
						}
					case map[string]interface{}:
						log.Debug("4 map[string]interface{} ", k)
						subcommands = append(subcommands, prompt.Suggest{
							Text:        k,
							Description: GetCommon(args[0], k),
						})
					case []string:
						log.Debug("5 []string")
						for _, x := range t {
							t1 := strings.Split(x, ",")
							subcommands = append(subcommands, prompt.Suggest{
								Text:        t1[0],
								Description: t1[1],
							})
						}
					case map[string]string:
						log.Debug("6 map[string]string")
						for k, v := range t {
							subcommands = append(subcommands, prompt.Suggest{
								Text:        k,
								Description: v,
							})
						}
					}
				}
				// []string
				if rs := Config.GetStringSlice(cmds); len(rs) > 0 {
					// 数据为[]string
					log.Debug("7 rs []string", rs)
					for _, x := range rs {
						if strings.ContainsAny(x, ",") {
							t1 := strings.Split(x, ",")
							subcommands = append(subcommands, prompt.Suggest{
								Text:        t1[0],
								Description: t1[1],
							})
						}
					}
				}
			}
			// return prompt.FilterHasPrefix(subcommands, second, true)
			return prompt.FilterFuzzy(subcommands, second, true)
		}
	}

	return []prompt.Suggest{}
}

// 获取说明
func GetCommon(cmd, target string) string {
	path := fmt.Sprintf("%s.common.%s", cmd, target)
	result := Config.GetString(path)
	if result != "" {
		return result
	}
	return "子目录，非功能性命令"
}
