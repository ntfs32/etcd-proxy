package main

import (
	"log"

	"github.com/jroimartin/gocui"
	"github.com/ntfs32/etcd-proxy/services/etcd"
)


func StartMonit(){
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.Mouse = true
	g.SelFgColor = gocui.ColorBlue
	g.BgColor = gocui.ColorBlack
	g.FgColor = gocui.ColorWhite

	g.SetManagerFunc(etcd.LayoutEtcd)
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

}
