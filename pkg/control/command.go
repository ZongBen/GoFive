package control

import (
	"github.com/eiannone/keyboard"
)

func ExecuteCommand[Tin any, Tout any](t Tin, commandHandler func(Tin, rune) Tout) Tout {
	key := readKey()
	return commandHandler(t, key)
}

func readKey() rune {
	key, _, _ := keyboard.GetSingleKey()
	return key
}
