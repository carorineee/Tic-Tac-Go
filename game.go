package main 

import(
	"fmt"
	"log"
	"strconv"
	"github.com/jroimartin/gocui"
)

var	grid = make([][]string, 3)
var player = "X"
var turns = 0

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
	g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, takeTurn)

	g.SetCurrentView("11")
	g.CurrentView().Highlight = true
	g.CurrentView().SelBgColor = gocui.ColorGreen
	g.CurrentView().SelFgColor = gocui.ColorBlack

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}


func keyLeft(g *gocui.Gui, v *gocui.View) error {
	switch(v.Name()) {
		case "11":
			v.Highlight = false
			g.SetCurrentView("10")
		case "12":
			v.Highlight = false
			g.SetCurrentView("11")
		case "01":
			v.Highlight = false
			g.SetCurrentView("00")
		case "02":
			v.Highlight = false
			g.SetCurrentView("01")
		case "21":
			v.Highlight = false
			g.SetCurrentView("20")
		case "22":
			v.Highlight = false
			g.SetCurrentView("21")
		default:	
	}

	g.CurrentView().Highlight = true
	g.CurrentView().SelBgColor = gocui.ColorGreen
	g.CurrentView().SelFgColor = gocui.ColorBlack

	return nil
}

func keyUp(g *gocui.Gui, v *gocui.View) error {

	switch(v.Name()) {
		case "10":
			v.Highlight = false
			g.SetCurrentView("00")			
		case "11":
			v.Highlight = false
			g.SetCurrentView("01")
		case "12":
			v.Highlight = false
			g.SetCurrentView("02")
		case "20":
			v.Highlight = false
			g.SetCurrentView("10")
		case "21":
			v.Highlight = false
			g.SetCurrentView("11")
			g.CurrentView().Highlight = true
		case "22":
			v.Highlight = false
			g.SetCurrentView("12")
		default:

	}
	g.CurrentView().Highlight = true
	g.CurrentView().SelBgColor = gocui.ColorGreen
	g.CurrentView().SelFgColor = gocui.ColorBlack

	return nil
}

func keyRight(g *gocui.Gui, v *gocui.View) error {
	switch(v.Name()) {
		case "10":
			v.Highlight = false
			g.SetCurrentView("11")
		case "11":
			v.Highlight = false
			g.SetCurrentView("12")
		case "00":
			v.Highlight = false
			g.SetCurrentView("01")
		case "01":
			v.Highlight = false
			g.SetCurrentView("02")
		case "20":
			v.Highlight = false
			g.SetCurrentView("21")
		case "21":
			v.Highlight = false
			g.SetCurrentView("22")
		default:

	}
	
	g.CurrentView().Highlight = true
	g.CurrentView().SelBgColor = gocui.ColorGreen
	g.CurrentView().SelFgColor = gocui.ColorBlack

	return nil
}

func keyDown(g *gocui.Gui, v *gocui.View) error {
	switch(v.Name()) {
		case "10":
			v.Highlight = false
			g.SetCurrentView("20")
		case "11":
			v.Highlight = false
			g.SetCurrentView("21")
		case "12":
			v.Highlight = false
			g.SetCurrentView("22")
		case "00":
			v.Highlight = false
			g.SetCurrentView("10")
		case "01":
			v.Highlight = false
			g.SetCurrentView("11")
		case "02":
			v.Highlight = false
			g.SetCurrentView("12")
		default:

	}

	g.CurrentView().Highlight = true
	g.CurrentView().SelBgColor = gocui.ColorGreen
	g.CurrentView().SelFgColor = gocui.ColorBlack

	return nil
}


func takeTurn(g *gocui.Gui, v *gocui.View) error {
	success := false
	clickedSpot := v.Name()
	i, erri := strconv.Atoi(string(clickedSpot[0]))
	j, errj := strconv.Atoi(string(clickedSpot[1]))
	if erri != nil {
		log.Panicln(erri)
	}
	if errj != nil {
		log.Panicln(errj)
	}

	if grid[i][j] == "-" {
		grid[i][j] = player;
		v.Clear()
		fmt.Fprintln(v, player)
		success = true
		turns++;
	}

	hasWinner := checkForWinner()
	mX, mY := g.Size()
	if hasWinner {
		if v, err := g.SetView("Winner", 0, 0, mX, mY); err != nil {
			v.Frame = false
			fmt.Fprintln(v, "Congrats player", player + ",", "you won!");
		}
		g.SetViewOnTop("Winner")
	} else if turns >= 9 {
		if v, err := g.SetView("Cats", 0, 0, mX, mY); err != nil {
			v.Frame = false
			fmt.Fprintln(v, "Cat's Game...");
		}
		g.SetViewOnTop("Cats")
	}

	if success && player == "X" {
		player = "O"
	} else if success {
		player = "X"
	}

	return nil;
}

func checkForWinner() bool {
	//Check rows
	for i := 0; i < 3; i++ {
		p := grid[i][0]
		for j := 1; j <= 3; j++ {
			if j == 3 {
				return true
			}
			if grid[i][j] == "-" || grid[i][j] != p {
				break;
			} 
		}
	}

	//Check cols
	for j := 0; j < 3; j++ {
		p := grid[0][j];
		for i := 1; i <= 3; i++ {
			if i == 3 {
				return true
			}
			if grid[i][j] == "-" || grid[i][j] != p {
				break;
			} 
		}
	}

	//Check diagonals
	p1 := grid[0][0]
	for z := 1; z <= 3; z++ {
		if z == 3 {
			return true
		}
		if grid[z][z] == "-" || grid[z][z] != p1 {
			break;
		} 
	}

	player1 := grid[0][2]
	player2 := grid[1][1]
	player3 := grid[2][0]

	if player1 != "-" && player1 == player2 && player1 == player3 {
		return true
	}

	return false
}

func setupGrid(g *gocui.Gui) {

	for i := range grid {
		grid[i] = make([]string, 3);
		for y := range grid[i] {
			grid[i][y] = "-"
		}
	}

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

	if v, err := g.SetView("00", x1, yt, z1, jt); err != nil {
		fmt.Fprintln(v, grid[0][0]);
	}
	if v, err := g.SetView("01", x2, yt, z2, jt); err != nil {
		fmt.Fprintln(v, grid[0][1]);
	}
	if v, err := g.SetView("02", x3, yt, z3, jt); err != nil {
		fmt.Fprintln(v, grid[0][2]);
	}
	if v, err := g.SetView("10", x1, ym, z1, jm); err != nil {
		fmt.Fprintln(v, grid[1][0]);
	}
	if v, err := g.SetView("11", x2, ym, z2, jm); err != nil {
		fmt.Fprintln(v, grid[1][1]);
	}
	if v, err := g.SetView("12", x3, ym, z3, jm); err != nil {
		fmt.Fprintln(v, grid[1][2]);
	}
	if v, err := g.SetView("20", x1, yb, z1, jb); err != nil {
		fmt.Fprintln(v, grid[2][0]);
	}
	if v, err := g.SetView("21", x2, yb, z2, jb); err != nil {
		fmt.Fprintln(v, grid[2][1]);
	}
	if v, err := g.SetView("22", x3, yb, z3, jb); err != nil {
		fmt.Fprintln(v, grid[2][2]);
	}
}