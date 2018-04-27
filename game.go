package main 

import(
	"fmt"
	"log"
	"github.com/jroimartin/gocui"
)


func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	// Set GUI managers and key bindings
	setupGrid(g)

	g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit)
	g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, keyUp)
	g.SetKeybinding("", gocui.KeyArrowLeft, gocui.ModNone, keyLeft)
	g.SetKeybinding("", gocui.KeyArrowRight, gocui.ModNone, keyRight)
	g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, keyDown)

	g.SetCurrentView("MidMid")
	g.CurrentView().Highlight = true
	g.CurrentView().SelBgColor = gocui.ColorGreen
	g.CurrentView().SelFgColor = gocui.ColorBlack

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func keyLeft(g *gocui.Gui, v *gocui.View) error {
	switch(v.Name()) {
		case "MidMid":
			v.Highlight = false
			g.SetCurrentView("MidLeft")
			g.CurrentView().Highlight = true
			g.CurrentView().SelBgColor = gocui.ColorGreen
			g.CurrentView().SelFgColor = gocui.ColorBlack
		case "MidRight":
			v.Highlight = false
			g.SetCurrentView("MidMid")
			g.CurrentView().Highlight = true
			g.CurrentView().SelBgColor = gocui.ColorGreen
			g.CurrentView().SelFgColor = gocui.ColorBlack
		case "TopMid":
			v.Highlight = false
			g.SetCurrentView("TopLeft")
			g.CurrentView().Highlight = true
			g.CurrentView().SelBgColor = gocui.ColorGreen
			g.CurrentView().SelFgColor = gocui.ColorBlack
		case "TopRight":
			v.Highlight = false
			g.SetCurrentView("TopMid")
			g.CurrentView().Highlight = true
			g.CurrentView().SelBgColor = gocui.ColorGreen
			g.CurrentView().SelFgColor = gocui.ColorBlack
		case "BottomMid":
			v.Highlight = false
			g.SetCurrentView("BottomLeft")
			g.CurrentView().Highlight = true
			g.CurrentView().SelBgColor = gocui.ColorGreen
			g.CurrentView().SelFgColor = gocui.ColorBlack
		case "BottomRight":
			v.Highlight = false
			g.SetCurrentView("BottomMid")
			g.CurrentView().Highlight = true
			g.CurrentView().SelBgColor = gocui.ColorGreen
			g.CurrentView().SelFgColor = gocui.ColorBlack
		default:	
	}

	return nil
}

func keyUp(g *gocui.Gui, v *gocui.View) error {

	switch(v.Name()) {
		case "MidLeft":
			v.Highlight = false
			g.SetCurrentView("TopLeft")
			g.CurrentView().Highlight = true
			g.CurrentView().SelBgColor = gocui.ColorGreen
			g.CurrentView().SelFgColor = gocui.ColorBlack
		case "MidMid":
			v.Highlight = false
			g.SetCurrentView("TopMid")
			g.CurrentView().Highlight = true
			g.CurrentView().SelBgColor = gocui.ColorGreen
			g.CurrentView().SelFgColor = gocui.ColorBlack
		case "MidRight":
			v.Highlight = false
			g.SetCurrentView("TopRight")
			g.CurrentView().Highlight = true
			g.CurrentView().SelBgColor = gocui.ColorGreen
			g.CurrentView().SelFgColor = gocui.ColorBlack
		case "BottomLeft":
			v.Highlight = false
			g.SetCurrentView("MidLeft")
			g.CurrentView().Highlight = true
			g.CurrentView().SelBgColor = gocui.ColorGreen
			g.CurrentView().SelFgColor = gocui.ColorBlack
		case "BottomMid":
			v.Highlight = false
			g.SetCurrentView("MidMid")
			g.CurrentView().Highlight = true
			g.CurrentView().SelBgColor = gocui.ColorGreen
			g.CurrentView().SelFgColor = gocui.ColorBlack
		case "BottomRight":
			v.Highlight = false
			g.SetCurrentView("MidRight")
			g.CurrentView().Highlight = true
			g.CurrentView().SelBgColor = gocui.ColorGreen
			g.CurrentView().SelFgColor = gocui.ColorBlack
		default:

	}

	return nil
}

func keyRight(g *gocui.Gui, v *gocui.View) error {
	switch(v.Name()) {
		case "MidLeft":
			v.Highlight = false
			g.SetCurrentView("MidMid")
			g.CurrentView().Highlight = true
			g.CurrentView().SelBgColor = gocui.ColorGreen
			g.CurrentView().SelFgColor = gocui.ColorBlack
		case "MidMid":
			v.Highlight = false
			g.SetCurrentView("MidRight")
			g.CurrentView().Highlight = true
			g.CurrentView().SelBgColor = gocui.ColorGreen
			g.CurrentView().SelFgColor = gocui.ColorBlack
		case "TopLeft":
			v.Highlight = false
			g.SetCurrentView("TopMid")
			g.CurrentView().Highlight = true
			g.CurrentView().SelBgColor = gocui.ColorGreen
			g.CurrentView().SelFgColor = gocui.ColorBlack
		case "TopMid":
			v.Highlight = false
			g.SetCurrentView("TopRight")
			g.CurrentView().Highlight = true
			g.CurrentView().SelBgColor = gocui.ColorGreen
			g.CurrentView().SelFgColor = gocui.ColorBlack
		case "BottomLeft":
			v.Highlight = false
			g.SetCurrentView("BottomMid")
			g.CurrentView().Highlight = true
			g.CurrentView().SelBgColor = gocui.ColorGreen
			g.CurrentView().SelFgColor = gocui.ColorBlack
		case "BottomMid":
			v.Highlight = false
			g.SetCurrentView("BottomRight")
			g.CurrentView().Highlight = true
			g.CurrentView().SelBgColor = gocui.ColorGreen
			g.CurrentView().SelFgColor = gocui.ColorBlack
		default:

	}

	return nil
}

func keyDown(g *gocui.Gui, v *gocui.View) error {
	switch(v.Name()) {
		case "MidLeft":
			v.Highlight = false
			g.SetCurrentView("BottomLeft")
			g.CurrentView().Highlight = true
			g.CurrentView().SelBgColor = gocui.ColorGreen
			g.CurrentView().SelFgColor = gocui.ColorBlack
		case "MidMid":
			v.Highlight = false
			g.SetCurrentView("BottomMid")
			g.CurrentView().Highlight = true
			g.CurrentView().SelBgColor = gocui.ColorGreen
			g.CurrentView().SelFgColor = gocui.ColorBlack
		case "MidRight":
			v.Highlight = false
			g.SetCurrentView("BottomRight")
			g.CurrentView().Highlight = true
			g.CurrentView().SelBgColor = gocui.ColorGreen
			g.CurrentView().SelFgColor = gocui.ColorBlack
		case "TopLeft":
			v.Highlight = false
			g.SetCurrentView("MidLeft")
			g.CurrentView().Highlight = true
			g.CurrentView().SelBgColor = gocui.ColorGreen
			g.CurrentView().SelFgColor = gocui.ColorBlack
		case "TopMid":
			v.Highlight = false
			g.SetCurrentView("MidMid")
			g.CurrentView().Highlight = true
			g.CurrentView().SelBgColor = gocui.ColorGreen
			g.CurrentView().SelFgColor = gocui.ColorBlack
		case "TopRight":
			v.Highlight = false
			g.SetCurrentView("MidRight")
			g.CurrentView().Highlight = true
			g.CurrentView().SelBgColor = gocui.ColorGreen
			g.CurrentView().SelFgColor = gocui.ColorBlack
		default:

	}

	return nil

}


func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func setupGrid(g *gocui.Gui) {
	mX, mY := g.Size()
	maxX := float64(mX)
	maxY := float64(mY)

	x1 := int(maxX/2-0.15*maxX)
	x2 := int(maxX/2-0.05*maxX)
	x3 := int(maxX/2+0.05*maxX)
	z1 := int(maxX/2-0.05*maxX)
	z2 := int(maxX/2+0.05*maxX)
	z3 := int(maxX/2+0.15*maxX)
	yt := int(maxY/2-0.15*maxY)
	ym := int(maxY/2-0.05*maxY)
	yb := int(maxY/2+0.05*maxY)
	jt := int(maxY/2-0.05*maxY)
	jm := int(maxY/2+0.05*maxY)
	jb := int(maxY/2+0.15*maxY)

	if v, err := g.SetView("TopLeft", x1, yt, z1, jt); err != nil {
		fmt.Fprintln(v, "-");
	}
	if v, err := g.SetView("TopMid", x2, yt, z2, jt); err != nil {
		fmt.Fprintln(v, "-");
	}
	if v, err := g.SetView("TopRight", x3, yt, z3, jt); err != nil {
		fmt.Fprintln(v, "-");
	}
	if v, err := g.SetView("MidLeft", x1, ym, z1, jm); err != nil {
		fmt.Fprintln(v, "-");
	}
	if v, err := g.SetView("MidMid", x2, ym, z2, jm); err != nil {
		fmt.Fprintln(v, "-");
	}
	if v, err := g.SetView("MidRight", x3, ym, z3, jm); err != nil {
		fmt.Fprintln(v, "-");
	}
	if v, err := g.SetView("BottomLeft", x1, yb, z1, jb); err != nil {
		fmt.Fprintln(v, "-");
	}
	if v, err := g.SetView("BottomMid", x2, yb, z2, jb); err != nil {
		fmt.Fprintln(v, "-");
	}
	if v, err := g.SetView("BottomRight", x3, yb, z3, jb); err != nil {
		fmt.Fprintln(v, "-");
	}


	grid := make([][]int, 3)
	for i := range grid {
		grid[i] = make([]int, 3);
	}
}