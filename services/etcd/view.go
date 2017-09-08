package etcd

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/ntfs32/etcd-proxy/utils"
)

var viewArr []string = []string{
	"v_header",
	"v_list",
	"v_info",
	"v_keys",
	"v_values",
	"v_notices",
	"v_footer",
}

//etcd ui layout
func LayoutEtcd(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	//header box
	if v, err := g.SetView(viewArr[0], 0, 0, maxX-1, 2); err != nil {
		v.BgColor = gocui.ColorBlue
		v.Overwrite =true
		if err != gocui.ErrUnknownView {
			return err
		}
		// show Title center
		fmt.Fprintln(v, utils.TextPadCenter(HEADER_INFO,maxX))
	}

	//nodes list box
	if v, err := g.SetView(viewArr[1], 0, 2, maxX/4, maxY); err != nil {
		v.Title = NODE_LIST
		v.Highlight = true
		v.Frame = true
		v.Autoscroll = true
		v.Wrap =true
		v.SelBgColor = gocui.ColorWhite
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "http:127.0.0.1:2379")
		fmt.Fprintln(v, "http:127.0.0.1:2379")
		fmt.Fprintln(v, "http:127.0.0.1:2379")
	}

	// node info box
	if v, err := g.SetView(viewArr[2], 0, maxY/4, maxX/4, maxY-4); err != nil {
		v.Title = NODE_INFO
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "12")
	}

	// key list box
	if v, err := g.SetView(viewArr[3], maxX/4, 2, maxX/4 * 3, maxY-4); err != nil {
		v.Title = KEY_LIST
		v.Highlight = true
		v.SelBgColor= gocui.ColorRed
		v.SelFgColor= gocui.ColorWhite
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "/bdh/api")
	}

	// show value box
	if v, err := g.SetView(viewArr[4], maxX/4 * 3, 2, maxX-1, maxY/3*2); err != nil {
		v.Title = VALUE_VIEW
		v.Editable = true
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "{\"name\":\"Shaddock\"}")
	}

	// small notice box
	if v, err := g.SetView(viewArr[5], maxX/4 * 3, maxY/3*2, maxX-1, maxY-4); err != nil {
		v.Title = NOTICE_INFO
		v.FgColor = gocui.ColorRed
		v.SetCursor(0,0)
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Warring...")
		fmt.Fprintln(v, "12")
	}

	//footer
	if v, err := g.SetView(viewArr[6], 0, maxY - 4, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Exit: <Ctrl+>c      Down: <Ctrl+>down       Help: <Ctrl+>h")
	}

	// register KeyMap Bind
	SetKeyBindEtcd(g)

	return nil
}

