package control

import (
	"unicode/utf8"

	"github.com/eiannone/keyboard"
)

const (
	ENTER_IP = iota
	ESC
)

func JoinGameCommandHandler(ip *string, char rune, key keyboard.Key) int {
	s_ip := *ip
	switch key {
	case keyboard.KeyEnter:
		return ENTER_IP
	case keyboard.KeyBackspace:
		s_ip = removeLastChar(s_ip)
	case keyboard.KeyBackspace2:
		s_ip = removeLastChar(s_ip)
	case keyboard.KeyEsc:
		return ESC
	default:
		if utf8.RuneCountInString(s_ip) <= 15 {
			s_ip += string(char)
		}
	}
	*ip = s_ip
	return -1
}

func removeLastChar(s string) string {
	len_s := utf8.RuneCountInString(s)
	if len_s > 0 {
		return s[:len_s-1]
	} else {
		return ""
	}
}
