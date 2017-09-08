package main

import(
	"log"

	"github.com/jroimartin/gocui"
)

/**
//❤😍💕❤😍💕❤😍💕❤😍💕❤😍
 */
func SetKeyBind(g *gocui.Gui)  (err error) {
	err = g.SetKeybinding("",gocui.KeyCtrlC,gocui.ModNone,quit)
	if(err!=nil){
		return
	}

	//etcd节点切换
	err = g.SetKeybinding(NODE_LIST,gocui.KeyArrowDown,gocui.ModNone,nodeListDown)
	if(err !=nil){
		return
	}
	err = g.SetKeybinding(NODE_LIST,gocui.KeyArrowUp,gocui.ModNone,nodeListUp)
	if(err !=nil){
		return
	}

	return
}




func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}


func nodeListUp(g *gocui.Gui, v *gocui.View) error  {
	log.Println("move Up")
	return nil
}
func nodeListDown(g *gocui.Gui, v *gocui.View) error  {
	log.Println("move down")
	return  nil
}