package etcd

import(

	"github.com/jroimartin/gocui"
)

type keyBind struct {
	viewname string
	key gocui.Key
	handler  func(*gocui.Gui, *gocui.View) error
}

// bind keys list
var keys = []keyBind{
	{
		viewname: "",
		key:      gocui.KeyCtrlC,
		handler:actionGlobalQuit,
	},
	{
		viewname: "",
		key:      gocui.KeyTab,
		handler:  changeEtcdNode,
	},
	{
		viewname: viewArr[1],
		key:      gocui.MouseLeft,
		handler:  nodeListMouseLeft,
	},
	{
		viewname: viewArr[1],
		key:      gocui.KeyEnter,
		handler:  nodeListEnter,
	},
	{
		viewname: viewArr[1],
		key:      gocui.KeyArrowDown,
		handler:  nodeListDown,
	},

	{
		viewname: viewArr[5],
		key:      gocui.KeyF3,
		handler:  nodeListDown,
	},
}


// bind keys
func SetKeyBindEtcd(g *gocui.Gui) error {
	for _, key := range keys {
		if err := g.SetKeybinding(key.viewname, key.key, gocui.ModNone, key.handler); err != nil {
			return err
		}
	}
	return nil
}
