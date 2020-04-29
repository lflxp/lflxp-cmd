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
	)
	p.Run()
}
