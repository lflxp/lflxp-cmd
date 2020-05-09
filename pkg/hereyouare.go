package pkg

import (
	prompt "github.com/c-bata/go-prompt"
)

func Exec() {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(">>> "),
		prompt.OptionLivePrefix(changeLivePrefix),
		prompt.OptionTitle("pkg"),
		prompt.OptionPrefixTextColor(prompt.Green),
		prompt.OptionInputTextColor(prompt.Blue),
	)
	p.Run()
}
