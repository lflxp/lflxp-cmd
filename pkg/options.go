package pkg

import (
	"fmt"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/prometheus/common/log"
)

// long == --
func OptionsCompleters(args []string) []prompt.Suggest {
	var options []prompt.Suggest
	if len(args) == 2 {

		// return prompt.FilterHasPrefix(optionHelp, "--", false)
		first := args[0]
		for k, _ := range count {
			if k == first {
				latest := args[len(args)-1]
				options = []prompt.Suggest{}
				for _, x := range Config.GetStringSlice(fmt.Sprintf("%s.flags.default", k)) {
					t1 := strings.Split(x, ",")
					tmp := prompt.Suggest{
						Text:        t1[0],
						Description: t1[1],
					}
					options = append(options, tmp)
				}
				return prompt.FilterHasPrefix(options, latest, true)
			}
		}
	}
	if len(args) > 2 {
		first := args[0]
		for k, _ := range count {
			if k == first {
				latest := []string{}
				for _, t := range args {
					if strings.Index(t, "-") != 0 {
						latest = append(latest, t)
					}
				}
				log.Debug("latest ", latest[len(latest)-1])
				options = []prompt.Suggest{}
				for _, x := range Config.GetStringSlice(fmt.Sprintf("%s.flags.%s", k, latest[len(latest)-1])) {
					t1 := strings.Split(x, ",")
					tmp := prompt.Suggest{
						Text:        t1[0],
						Description: t1[1],
					}
					options = append(options, tmp)
				}
				return prompt.FilterHasPrefix(options, args[len(args)-1], true)
			}
		}
	}

	// return prompt.FilterContains(options, strings.TrimLeft(args[len(args)-1], "-"), true)
	return options
}
