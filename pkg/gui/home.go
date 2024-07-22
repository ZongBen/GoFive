package gui

import (
	"net"
	"strings"

	"github.com/ZongBen/GoFive/pkg/menu"
	"github.com/ZongBen/tanvas"
)

var (
	_menuCanvas      tanvas.Canvas
	_menuSection     tanvas.Section
	_joinGameSection tanvas.Section
	_hostGameSection tanvas.Section
)

var title string

func init() {
	menuCanvas := tanvas.CreateCanvas(50, 5, 1)
	menuCanvas.SetOffset(0, 0)
	_menuCanvas = &menuCanvas

	menuSection := menuCanvas.CreateSection(7, 0, 35, 5, 0)
	_menuSection = &menuSection

	joinGameSection := menuCanvas.CreateSection(2, 0, 35, 5, 0)
	_joinGameSection = &joinGameSection

	hostGameSection := menuCanvas.CreateSection(7, 0, 35, 5, 0)
	_hostGameSection = &hostGameSection

	title = renderTitle()
}

func RenderHome(homeMenu menu.HomeMenu) string {
	_menuCanvas.Clear()
	result := title + renderHomeMenu(homeMenu)
	return result
}

func RenderOnline(onlineMenu menu.OnlineMenu) string {
	_menuCanvas.Clear()
	result := title + renderOnlineMenu(onlineMenu)
	return result
}

func RenderJoinGame(ip string) string {
	_menuCanvas.Clear()
	result := title + renderJoinGame(ip)
	return result
}

func RenderHostGame(dot int) string {
	_menuCanvas.Clear()
	result := title + renderHostGame(dot)
	return result
}

func renderTitle() string {
	titleCanvas := tanvas.CreateCanvas(35, 7, 1)
	titleSection := titleCanvas.CreateSection(0, 0, 35, 7, 0)
	_titleSection := &titleSection

	title :=
		`
   _____       _____  _           
  / ____|     |  ___|(_)          
 | |  __  ___ | |__   ___   _____ 
 | | |_ |/ _ \|  __| | \ \ / / _ \
 | |__| | (_) | |    | |\ V /  __/
  \_____|\___/|_|    |_| \_/ \___|
  `
	title = strings.Trim(title, "\n")
	lines := strings.Split(title, "\n")
	for i, line := range lines {
		_titleSection.SetRow(0, i, line)
	}
	return titleCanvas.Project()
}

func renderHomeMenu(m menu.HomeMenu) string {
	menu :=
		`
  1. Local  Game
  2. Online Game
  3. Exit       `
	lines := strings.Split(menu, "\n")
	for i, line := range lines {
		if i == m.GetMenuSelect()+1 {
			line += " <=="
		} else {
			line += "    "
		}
		_menuSection.SetRow(0, i, line)
	}
	return _menuCanvas.Project()
}

func renderOnlineMenu(m menu.OnlineMenu) string {
	menu :=
		`
  1. Join Game
  2. Host Game
  3. Back     `
	lines := strings.Split(menu, "\n")
	for i, line := range lines {
		if i == m.GetMenuSelect()+1 {
			line += " <=="
		} else {
			line += "    "
		}
		_menuSection.SetRow(0, i, line)
	}
	return _menuCanvas.Project()
}

func renderJoinGame(ip string) string {
	menu := "Enter the IP address of the host"
	input := "IP: " + ip
	_joinGameSection.SetRow(0, 1, menu)
	_joinGameSection.SetRow(0, 2, input)
	return _menuCanvas.Project()
}

func renderHostGame(dot int) string {
	s_dot := ""
	for i := 0; i < dot; i++ {
		s_dot += "."
	}
	ipAddrs, _ := getLocalIPs()
	var ip string
	for _, ipAddr := range ipAddrs {
		if strings.HasPrefix(ipAddr.String(), "192.168.") {
			ip = ipAddr.String()
			break
		}
	}
	menu := "Waiting for connection" + s_dot
	_hostGameSection.SetRow(0, 1, "Your IP: "+ip)
	_hostGameSection.SetRow(0, 2, menu)
	return _menuCanvas.Project()
}

func getLocalIPs() ([]net.IP, error) {
	var ips []net.IP
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, addr := range addresses {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP)
			}
		}
	}
	return ips, nil
}
