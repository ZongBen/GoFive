package scenes

import (
	"github.com/ZongBen/tanvas/pkg/utils"
)

func RenderHome() string {
	c := utils.CreateCanvas(40, 7, 1)
	titleSection := c.CreateSection(0, 0, 40, 7, 0)

	titleSection.SetRow(6, 0, "   _____       ______ _           ")
	titleSection.SetRow(6, 1, "  / ____|     |  ____(_)          ")
	titleSection.SetRow(6, 2, " | |  __  ___ | |__   ___   _____ ")
	titleSection.SetRow(6, 3, " | | |_ |/ _ \\|  __| | \\ \\ / / _ \\")
	titleSection.SetRow(6, 4, " | |__| | (_) | |    | |\\ V /  __/")
	titleSection.SetRow(6, 5, "  \\_____|\\___/|_|    |_| \\_/ \\___|")

	return c.Render()
}
