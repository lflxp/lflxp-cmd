package pkg

import (
	prompt "github.com/c-bata/go-prompt"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func Exec() {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix("➜ "),
		prompt.OptionLivePrefix(changeLivePrefix),
		prompt.OptionTitle("pkg"),
		prompt.OptionPrefixTextColor(prompt.Red),
		prompt.OptionInputTextColor(prompt.LightGray),
	)
	p.Run()
}
