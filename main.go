package main

import (
	"log"
	"math/rand"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	app := fiber.New()

	vimTips := []string{
		"<number>G goes to the line with that number",
		"`:%s/./&/gn` counts characters in a buffer",
		"`:%s/\\i\\+/&/gn` counts characters in a buffer",
		"`:%s/^//n` counts lines in a buffer",
		"Add tokens into your quickfix list with :grep and use :cdo do run an action on all of them",
		"Mapping caps-lock to escape is a common practice for vim users",
	}

	pluginTips := []string{
		"Lazy.nvim is a modern plugin manager for Neovim.",
		"TobinPalmer/rayso.nvim allows you to take code snippets in Neovim.",
		"ThePrimeagen/harpoon is a plugin that allows you to navigate between files in Neovim.",
		"windwp/nvim-spectre is a plugin that allows you to search and replace many files at once.",
		"ibhagwan/fzf-lua is a plugin that allows you to use fzf in Neovim.",
		"nvim-telescope/telescope.nvim is a plugin that allows you to search for files in Neovim.",
	}

	allTips := append(vimTips, pluginTips...)

	app.Get("/", func(ctx *fiber.Ctx) error {
		random := rand.Intn(len(allTips))
		return ctx.SendString(allTips[random])
	})

	app.Get("/plugin", func(ctx *fiber.Ctx) error {
		random := rand.Intn(len(pluginTips))
		return ctx.SendString(pluginTips[random])
	})

	app.Get("/vim", func(ctx *fiber.Ctx) error {
		random := rand.Intn(len(vimTips))
		return ctx.SendString(vimTips[random])
	})

	port := os.Getenv("PORT")

	if os.Getenv("PORT") == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}
