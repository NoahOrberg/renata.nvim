package main

import (
	"github.com/NoahOrberg/renata.nvim/command"
	"github.com/neovim/go-client/nvim/plugin"
)

func main() {
	g := &command.Renata{}
	plugin.Main(func(p *plugin.Plugin) error {
		p.HandleCommand(&plugin.CommandOptions{Name: "Http", NArgs: "*"}, g.RenataHttp)
		return nil
	})
}
