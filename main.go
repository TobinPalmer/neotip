package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/heroku/x/hmetrics/onload"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
)

//type DotfyleRequest interface {
//	Result() Result
//}
//
//type Result interface {
//	Data() Data
//}
//
//type Data interface {
//	Meta() Meta
//	Data() []DotfyleData
//}
//
//type Meta interface {
//	Total() int
//	LastPage() int
//	CurrentPage() int
//	PerPage() int
//	Prev() interface{}
//	Next() int
//}
//
//type DotfyleData interface {
//	Id() int
//	Owner() string
//	Name() string
//	Type() string
//	Source() string
//	Category() string
//	Link() string
//	Description() string
//	ShortDescription() string
//	CreatedAt() string
//	LastSyncedAt() string
//	Stars() int
//	AddedLastWeek() int
//	DotfyleShieldAddedAt() string
//	Media() []Media
//	ConfigCount() int
//}
//
//type Media interface {
//	Id() int
//	Url() string
//	Type() string
//	Thumbnail() bool
//	NeovimPluginId() int
//}

type Response struct {
	Result Result `json:"result"`
}

type Result struct {
	Data Data `json:"data"`
}

type Data struct {
	Data []DotfyleRequest `json:"data"`
}

type DotfyleRequest struct {
	Id                   int    `json:"id"`
	Owner                string `json:"owner"`
	Name                 string `json:"name"`
	Type                 string `json:"type"`
	Source               string `json:"source"`
	Category             string `json:"category"`
	Link                 string `json:"link"`
	Description          string `json:"description"`
	ShortDescription     string `json:"shortDescription"`
	CreatedAt            string `json:"createdAt"`
	LastSyncedAt         string `json:"lastSyncedAt"`
	Stars                int    `json:"stars"`
	AddedLastWeek        int    `json:"addedLastWeek"`
	DotfyleShieldAddedAt string `json:"dotfyleShieldAddedAt"`
	Media                []Media
	ConfigCount          int `json:"configCount"`
}

type Media struct {
	Id             int    `json:"id"`
	Url            string `json:"url"`
	Type           string `json:"type"`
	Thumbnail      bool   `json:"thumbnail"`
	NeovimPluginId int    `json:"neovimPluginId"`
}

func requestDotfyle() ([]DotfyleRequest, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	get, err := client.Get("https://dotfyle.com/trpc/searchPluginsWithMedia?batch=1&input=%7B%220%22%3A%7B%22category%22%3A%22colorscheme%22%2C%22sorting%22%3A%22trending%22%2C%22page%22%3A1%2C%22take%22%3A25%7D%7D")
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(get.Body)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(get.Body)

	var response []Response

	err = json.NewDecoder(get.Body).Decode(&response)
	if err != nil {
		log.Printf("Error unmarshalling: %v", err)
		return nil, err
	}

	var dotfyleRequest []DotfyleRequest
	for _, data := range response {
		for _, d := range data.Result.Data.Data {
			dotfyleRequest = append(dotfyleRequest, d)
		}
	}

	return dotfyleRequest, nil
}

func concatSlices(slices ...[]string) []string {
	var result []string
	for _, slice := range slices {
		for _, elem := range slice {
			result = append(result, elem)
		}
	}
	return result
}

func main() {
	app := fiber.New()

	/*
		https://neovim.io/doc/user/tips.html
		https://vim.fandom.com/wiki/Best_Vim_Tips
	*/
	vimTips := []string{
		"<number>G goes to the line with that number",
		"`:%s/./&/gn` counts characters in a buffer",
		"`:%s/\\i\\+/&/gn` counts characters in a buffer",
		"`:%s/^//n` counts lines in a buffer",
		"Add tokens into your quickfix list with :grep and use :cdo do run an action on all of them",
		"Mapping caps-lock to escape is a common practice for vim users",
		"Use \"_ to yank into the black hole register ",
		"Use `:help g` to learn about the powerful uses of the g command",
		"If text is wrapping use gk and gj to move up and down",
		"Use CTRL-A and CTRL-X to increment and decrement numbers",
		"You can use a vimscript function to replace text `%s/replace/\\=1+1`",
		"Switch registers with \"<reg> to copy/paste from different registers",
		"q: opens the recent command history",
		"Use :%! nl -ba to add line numbers to a file",
		"Use CTRL-O and CTRL-I to jump back and forth between locations",
	}

	/*
		https://github.com/rockerBOO/awesome-neovim
	*/
	pluginTips := []string{
		"folke/lazy.nvim is a modern plugin manager for Neovim.",
		"TobinPalmer/rayso.nvim allows you to take code snippets in Neovim.",
		"ThePrimeagen/harpoon is a plugin that allows you to navigate between files in Neovim.",
		"windwp/nvim-spectre is a plugin that allows you to search and replace many files at once.",
		"ibhagwan/fzf-lua is a plugin that allows you to use fzf in Neovim.",
		"nvim-telescope/telescope.nvim is a plugin that allows you to search for files in Neovim.",
		"echasnovski/mini.nvim is a collection of 35+ independent lua plugins for Neovim.",
		"pluffie/neoproj is a small yet powerful project (and session) manager.",
		"ahmedkhalf/project.nvim is an all in one Neovim plugin that provides superior project management.",
		"cljoly/telescope-repo.nvim is a Telescope picker that lets you jump to any repository (git or other) on the file system.",
		"ziontee113/color-picker.nvim is a color picking plugin that lets users choose & modify RGB/HSL/HEX colors inside Neovim.",
	}

	/*
		https://github.com/rockerBOO/awesome-neovim
	*/
	colorSchemes := []string{
		"sontungexpt/witch - The primary stinvim distro colorscheme includes the default feature of dimming inactive windows, along with various other customization options for users.",
		"shaeinst/roshnivim-cs - Colorscheme written in Lua, specially made for roshnivim with Tree-sitter support.",
		"rafamadriz/neon - Customizable colorscheme with excellent italic and bold support, dark and light variants. Made to work and look good with Tree-sitter.",
		"tomasiser/vim-code-dark - A dark color scheme heavily inspired by the look of the Dark+ scheme of Visual Studio Code.",
		"Mofiqul/vscode.nvim - A Lua port of vim-code-dark colorscheme with vscode light and dark theme.",
		"askfiy/visual_studio_code - A Neovim theme that highly restores vscode, so that your friends will no longer be surprised that you use Neovim, because they will think you are using vscode.",
		"marko-cerovac/material.nvim - Material.nvim is a highly configurable colorscheme written in Lua and based on the material palette.",
		"bluz71/vim-nightfly-colors - A dark midnight colorscheme with modern Neovim support including Tree-sitter.",
		"bluz71/vim-moonfly-colors - A dark charcoal colorscheme with modern Neovim support including Tree-sitter.",
		"ChristianChiarulli/nvcode-color-schemes.vim - Nvcode, onedark, nord colorschemes with Tree-sitter support.",
		"folke/tokyonight.nvim - A clean, dark and light Neovim theme written in Lua, with support for LSP, Tree-sitter and lots of plugins.",
		"crispybaccoon/evergarden - A comfy Neovim colorscheme for cozy morning coding.",
		"sainnhe/sonokai - High Contrast & Vivid Color Scheme based on Monokai Pro.",
		"nyoom-engineering/oxocarbon.nvim - A dark and light Neovim theme written in fennel, inspired by IBM Carbon.",
		"kyazdani42/blue-moon - A dark color scheme derived from palenight and carbonight.",
		"mhartington/oceanic-next - Oceanic Next theme.",
		"nvimdev/zephyr-nvim - A dark colorscheme with Tree-sitter support.",
		"rockerBOO/boo-colorscheme-nvim - A colorscheme with handcrafted support for LSP, Tree-sitter.",
		"jim-at-jibba/ariake-vim-colors - A port of the great Atom theme. Dark and light with Tree-sitter support.",
		"Th3Whit3Wolf/onebuddy - Light and dark atom one theme.",
		"ishan9299/modus-theme-vim - This is a color scheme developed by Protesilaos Stavrou for emacs.",
		"sainnhe/edge - Clean & Elegant Color Scheme inspired by Atom One and Material.",
		"theniceboy/nvim-deus - Vim-deus with Tree-sitter support.",
		"bkegley/gloombuddy - Gloom inspired theme.",
		"Th3Whit3Wolf/one-nvim - An Atom One inspired dark and light colorscheme.",
		"PHSix/nvim-hybrid - A Neovim colorscheme write in Lua.",
		"Th3Whit3Wolf/space-nvim - A spacemacs inspired dark and light colorscheme.",
		"yonlu/omni.vim - Omni color scheme for Vim.",
		"ray-x/aurora - A 24-bit dark theme with Tree-sitter and LSP support.",
		"ray-x/starry.nvim - A collection of modern Neovim colorschemes: material, moonlight, dracula (blood), monokai, mariana, emerald, earlysummer, middlenight_blue, darksolar.",
		"tanvirtin/monokai.nvim - Monokai theme written in Lua.",
		"ofirgall/ofirkai.nvim - Monokai theme that aims to feel like Sublime Text.",
		"savq/melange-nvim - Warm colorscheme written in Lua with support for various terminal emulators.",
		"RRethy/nvim-base16 - Neovim plugin for building base16 colorschemes. Includes support for Treesitter and LSP highlight groups.",
		"fenetikm/falcon - A colour scheme for terminals, Vim and friends.",
		"andersevenrud/nordic.nvim - A nord-esque colorscheme.",
		"AlexvZyl/nordic.nvim - Nord for Neovim, but warmer and darker. Supports a variety of plugins and other platforms.",
		"shaunsingh/nord.nvim - Neovim theme based off of the Nord Color Palette.",
		"Tsuzat/NeoSolarized.nvim - NeoSolarized colorscheme with full transparency.",
		"svrana/neosolarized.nvim - Dark solarized colorscheme using colorbuddy for easy customization.",
		"ishan9299/nvim-solarized-lua - Solarized colorscheme in Lua (Neovim >= 0.5).",
		"shaunsingh/moonlight.nvim - Port of VSCode's Moonlight colorscheme, written in Lua with built-in support for native LSP, Tree-sitter and many more plugins.",
		"navarasu/onedark.nvim - A One Dark Theme (Neovim >= 0.5) written in Lua based on Atom's One Dark Theme.",
		"lourenci/github-colors - GitHub colors leveraging Tree-sitter to get 100% accuracy.",
		"sainnhe/gruvbox-material - Gruvbox modification with softer contrast and Tree-sitter support.",
		"sainnhe/everforest - A green based colorscheme designed to be warm, soft and easy on the eyes.",
		"neanias/everforest-nvim - A Lua port of the Everforest colour scheme.",
		"NTBBloodbath/doom-one.nvim - Lua port of doom-emacs' doom-one.",
		"dracula/vim - Famous beautiful dark powered theme.",
		"Mofiqul/dracula.nvim - Dracula colorscheme for neovim written in Lua.",
		"yashguptaz/calvera-dark.nvim - A port of VSCode Calvara Dark Theme to Neovim with Tree-sitter and many other plugins support.",
		"nxvu699134/vn-night.nvim - A dark Neovim colorscheme written in Lua. Support built-in LSP and Tree-sitter.",
		"adisen99/codeschool.nvim - Codeschool colorscheme written in Lua with Tree-sitter and built-in lsp support.",
		"projekt0n/github-nvim-theme - A GitHub theme, kitty, alacritty written in Lua. Support built-in LSP and Tree-sitter.",
		"kdheepak/monochrome.nvim - A 16 bit monochrome colorscheme that uses hsluv for perceptually distinct gray colors, with support for Tree-sitter and other commonly used plugins.",
		"rose-pine/neovim - All natural pine, faux fur and a bit of soho vibes for the classy minimalist.",
		"mcchrish/zenbones.nvim - A collection of Vim/Neovim colorschemes designed to highlight code using contrasts and font variations.",
		"catppuccin/nvim - Warm mid-tone dark theme to show off your vibrant self! with support for native LSP, Tree-sitter, and more 🍨!",
		"FrenzyExists/aquarium-vim - A dark, yet vibrant colorscheme.",
		"EdenEast/nightfox.nvim - A soft dark, fully customizable Neovim theme, with support for lsp, treesitter and a variety of plugins.",
		"kvrohit/substrata.nvim - A cold, dark color scheme written in Lua ported from arzg/vim-substrata theme.",
		"ldelossa/vimdark - A minimal Vim theme for night time. Loosely based on vim-monotonic and chrome's dark reader extension. A light theme is included as well for the day time.",
		"Everblush/everblush.nvim - A dark, vibrant and beautiful colorscheme written in Lua.",
		"adisen99/apprentice.nvim - Colorscheme written in Lua based on the Apprentice color pattete with Tree-sitter and built-in lsp support.",
		"olimorris/onedarkpro.nvim - Atom's iconic One Dark theme. Cacheable, fully customisable, Tree-sitter and LSP semantic token support. Comes with light and dark variants.",
		"rmehri01/onenord.nvim - A Neovim theme that combines the Nord and Atom One Dark color palettes for a more vibrant programming experience.",
		"RishabhRD/gruvy - Gruvbuddy without colorbuddy using Lush.",
		"echasnovski/mini.nvim#colorschemes - Color schemes included in mini.nvim plugin. All of them prioritize high contrast ratio for reading text and computing palettes in perceptually uniform color spaces.",
		"luisiacc/gruvbox-baby - A modern gruvbox theme with full treesitter support.",
		"titanzero/zephyrium - A zephyr-esque theme, written in Lua, with TreeSitter support.",
		"rebelot/kanagawa.nvim - Neovim dark colorscheme inspired by the colors of the famous painting by Katsushika Hokusai.",
		"tiagovla/tokyodark.nvim - A clean dark theme written in Lua (Neovim >= 0.5) and above.",
		"cpea2506/one_monokai.nvim - One Monokai theme written in Lua.",
		"phha/zenburn.nvim - A low-contrast dark colorscheme with support for various plugins.",
		"kvrohit/rasmus.nvim - A dark color scheme written in Lua ported from rsms/sublime-theme theme.",
		"chrsm/paramount-ng.nvim - A dark color scheme written using Lush. Treesitter supported.",
		"kaiuri/nvim-juliana - Port of Sublime's Mariana Theme to Neovim for short attention span developers with Tree-sitter support.",
		"lmburns/kimbox - A colorscheme with a dark background, and vibrant foreground that is centered around the color brown. A modification of Kimbie Dark.",
		"rockyzhang24/arctic.nvim - A Neovim colorscheme ported from VSCode Dark+ theme with the strict and precise color picking for both the editor and UI.",
		"ramojus/mellifluous.nvim - Pleasant and productive colorscheme.",
		"Yazeed1s/minimal.nvim - Two tree-sitter supported colorschemes that are inspired by base16-tomorrow-night and monokai-pro.",
		"lewpoly/sherbet.nvim - A soothing colorscheme with support for popular plugins and tree-sitter.",
		"Mofiqul/adwaita.nvim - Colorscheme based on GNOME Adwaita syntax with support for popular plugins.",
		"olivercederborg/poimandres.nvim - Neovim port of poimandres VSCode theme with Tree-sitter support, written in Lua.",
		"kvrohit/mellow.nvim - A soothing dark color scheme with tree-sitter support.",
		"gbprod/nord.nvim - An arctic, north-bluish clean and elegant Neovim theme, based on Nord Palette.",
		"Yazeed1s/oh-lucy.nvim - Two tree-sitter supported colorschemes, inspired by oh-lucy in vscode.",
		"embark-theme/vim - A deep inky purple theme leveraging bright colors.",
		"nyngwang/nvimgelion - Neon Genesis Evangelion but for Vimmers.",
		"maxmx03/FluoroMachine.nvim - Synthwave x Fluoromachine port.",
		"dasupradyumna/midnight.nvim - A modern black Neovim theme with comfortable color contrast for a pleasant visual experience, with LSP and Tree-sitter support.",
		"sonjiku/yawnc.nvim - Theming using pywal, with a Base16 twist.",
		"sekke276/dark_flat.nvim - A Neovim colorscheme written in Lua ported from Dark Flat iTerm2 theme, with LSP and Tree-sitter support.",
		"zootedb0t/citruszest.nvim - A colorscheme that features a combination of bright and juicy colors reminiscent of various citrus fruits, with LSP and Tree-sitter support.",
		"2nthony/vitesse.nvim - Vitesse theme Lua port.",
		"xero/miasma.nvim - A dark pastel color scheme inspired by the woods. Built using lush and supports Tree-sitter, diagnostics, CMP, Git-Signs, Telescope, Which-key, Lazy, and more.",
		"Verf/deepwhite.nvim - A light colorscheme inspired by flatwhite-syntax and elegant-emacs.",
		"judaew/ronny.nvim - A dark colorscheme, which mostly was inspired by the Monokai originally created by Wimem Hazenberg.",
		"ribru17/bamboo.nvim - A warm green theme.",
		"cryptomilk/nightcity.nvim - A dark colorscheme inspired by Inkpot, Jellybeans, Gruvbox and Tokyonight with LSP support.",
		"polirritmico/monokai-nightasty.nvim - A dark/light theme based on the Monokai color palette written in Lua, support for LSP, Tree-sitter and lots of plugins.",
		"oxfist/night-owl.nvim - A Night Owl colorscheme port from VSCode with support for Tree-sitter and semantic tokens.",
		"text-to-colorscheme - Dynamically generated colorschemes generated on the fly with a text prompt using ChatGPT.",
		"miikanissi/modus-themes.nvim - Accessible theme, conforming with the highest standard for color contrast (WCAG AAA).",
		"alexmozaidze/palenight.nvim - Palenight colorscheme supporting Tree-sitter, LSP (including semantic tokens) and lots of plugins.",
		"scottmckendry/cyberdream.nvim - A high-contrast, futuristic & vibrant coloursheme.",
		"HoNamDuong/hybrid.nvim - A dark theme written in Lua.",
	}

	allTips := concatSlices(vimTips, pluginTips, colorSchemes)

	app.Get("/", func(ctx *fiber.Ctx) error {
		queries := ctx.Queries()
		if len(queries) == 0 {
			return ctx.SendString(allTips[rand.Intn(len(allTips))])
		}

		var tips []string
		for query := range queries {
			fmt.Println("QUERY", query)
			switch query {
			case "vim":
				tips = concatSlices(tips, vimTips)
			case "plugin":
				tips = concatSlices(tips, pluginTips)
			case "colorscheme":
				tips = concatSlices(tips, colorSchemes)
			case "all":
				tips = allTips
			default:
				tips = allTips
			}
		}

		return ctx.SendString(tips[rand.Intn(len(tips))])
	})

	app.Get("/v2", func(ctx *fiber.Ctx) error {
		get, err := requestDotfyle()
		if err != nil {
			log.Fatal("Error requesting", err)
		}

		// pick a random tip
		return ctx.SendString(get[rand.Intn(len(get))].Name)
	})

	port := os.Getenv("PORT")

	if os.Getenv("PORT") == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}
