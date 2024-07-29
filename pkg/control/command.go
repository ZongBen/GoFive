package control

import (
	"github.com/eiannone/keyboard"
)

func ExecuteCommand[Tin any, Tout any](t Tin, commandHandler func(Tin, rune, keyboard.Key) Tout) Tout {
	char, key, _ := keyboard.GetSingleKey()
	return commandHandler(t, char, key)
}

func ReadCommand() (rune, keyboard.Key) {
	char, key, _ := keyboard.GetSingleKey()
	return char, key
}
