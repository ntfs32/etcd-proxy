package main

import (
	"github.com/jroimartin/gocui"
	"fmt"
	"log"
)

var (
	HEADER_INFO = "        Etcd Config Manager"
	NODE_LIST string = "Etcd Cluster Nodes"
	NODE_INFO string = "Etcd Node Info"
	KEY_LIST string = "Keys List Check"
	VALUE_VIEW string = "View Key Value"
	NOTICE_INFO string = "Notice Info"
	HELP_FOOTER string = "Help Infos"
)

func StartMonit(){
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()
	// Highlight active view.
	g.Highlight = true
	g.Mouse = true
	g.SelFgColor = gocui.ColorBlue
	g.BgColor = gocui.ColorBlack
	g.FgColor = gocui.ColorWhite

	g.SetManagerFunc(layout)
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

}


func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	//header box
	if v, err := g.SetView(HEADER_INFO, 0, 0, maxX-1, 2); err != nil {
		v.BgColor = gocui.ColorBlue
		v.Overwrite =true
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, HEADER_INFO)
	}

	//nodes list box
	if v, err := g.SetView(NODE_LIST, 0, 2, maxX/4, maxY); err != nil {
		v.Title = NODE_LIST
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, ".1..")
	}

	// node info box
	if v, err := g.SetView(NODE_INFO, 0, maxY/4, maxX/4, maxY-4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, NODE_INFO)
	}

	// key list box
	if v, err := g.SetView(KEY_LIST, maxX/4, 2, maxX/4 * 3, maxY-4); err != nil {
		v.Title = KEY_LIST
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, NODE_INFO)
	}

	// show value box
	if v, err := g.SetView(VALUE_VIEW, maxX/4 * 3, 2, maxX-1, maxY/3*2); err != nil {
		v.Title = VALUE_VIEW
		v.FgColor = gocui.ColorCyan
		v.SelBgColor = gocui.ColorBlue
		v.SelFgColor = gocui.ColorBlack
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, VALUE_VIEW)
	}

	// small notice box
	if v, err := g.SetView(NOTICE_INFO, maxX/4 * 3, maxY/3*2, maxX-1, maxY-4); err != nil {
		v.Title = NOTICE_INFO
		g.Cursor = true
		g.Cursor = true
		v.SetCursor(0,0)
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, NODE_INFO)
		fmt.Fprintln(v, "12")
	}

	//footer
	if v, err := g.SetView(HELP_FOOTER, 0, maxY - 4, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Exit: <Ctrl+>c      Down: <Ctrl+>down       Help: <Ctrl+>h")
	}

	// register KeyMap Bind
	SetKeyBind(g)

	return nil
}
