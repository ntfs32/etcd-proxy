package etcd

import (
	"github.com/jroimartin/gocui"
	"fmt"
	util "github.com/ntfs32/etcd-proxy/utils"

	"github.com/willf/pad"
)

// change top with node list view
func changeEtcdNode(g *gocui.Gui, v *gocui.View) error  {
	vName := viewArr[1]
	_, err := g.View(viewArr[1])
	if err != nil {
		return err
	}
	if _, err := g.SetCurrentView(vName); err != nil {
		return err
	}
	return  nil
}




func actionGlobalQuit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func nodeListEnter(g *gocui.Gui, v *gocui.View) error  {
	fmt.Fprintln(v,v.Title)
	return nil
}

func nodeListMouseLeft(g *gocui.Gui, v *gocui.View) error  {
	return nil
}
func nodeListDown(g *gocui.Gui, v *gocui.View) error  {
	actionGlobalQuit(g,v)
	v.Highlight = true
	return  nil
}

//set footer help message
func changeStatusContext(g *gocui.Gui,vName string, c string) error {
	lMaxX, _ := g.Size()
	v, err := g.View(vName)
	if err != nil {
		return err
	}

	v.Clear()

	i := lMaxX + 4
	b := ""

	switch c {
	case "D":
		i = 150 + i
		b = b + util.FrameText("↑") + " Up   "
		b = b + util.FrameText("↓") + " Down   "
		b = b + util.FrameText("D") + " Delete   "
		b = b + util.FrameText("L") + " Show Logs   "
		b = b + util.FrameText("CTRL+N") + " Namespace   "
	case "SE":
		i = i + 100
		b = b + util.FrameText("↑") + " Up   "
		b = b + util.FrameText("↓") + " Down   "
		b = b + util.FrameText("Enter") + " Select   "
	case "SL":
		i = i + 100
		b = b + util.FrameText("↑") + " Up   "
		b = b + util.FrameText("↓") + " Down   "
		b = b + util.FrameText("L") + " Hide Logs   "
	}
	b = b + util.FrameText("CTRL+C") + " Exit"

	fmt.Fprintln(v, pad.Left(b, i, " "))

	return nil
}
