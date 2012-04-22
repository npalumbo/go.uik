package main

import (
	"fmt"
<<<<<<< HEAD
	"github.com/skelterjohn/go.wde"
=======
	"github.com/skelterjohn/geom"
>>>>>>> upstream/master
	"github.com/skelterjohn/go.uik"
	"github.com/skelterjohn/go.uik/layouts"
	"github.com/skelterjohn/go.uik/widgets"
	"image/color"
)

func (b *Button) handleState() {
        for {
                select {
                case <-b.click:
                        b.Label.TextCh <- "clickedLoco!"
                }
        }
}


func main() {
	wbounds := geom.Rect{
		Max: geom.Coord{480, 320},
	}
	w, err := uik.NewWindow(nil, int(wbounds.Max.X), int(wbounds.Max.Y))
	if err != nil {
		fmt.Println(err)
		return
	}
	w.W.SetTitle("GoUI")

	// Here we create a flow layout, which just lines up its blocks from
	// left to right.
	fl := layouts.NewFlow(wbounds.Max)
	// We add it to the window, taking up the entire space the window has.
	// At some point we'd need to create a mechanism for blocks and foundations
	// to resize themselves by sending hints up to their parents.
	w.SetBlock(&fl.Block)

	// Create a button with the given size and label
	b := widgets.NewButton(geom.Coord{100, 50}, "Hi")
	// Here we get the button's label's data
	ld := <-b.Label.GetConfig
	// we modify the copy for a special message to display
	ld.Text = "clicked!"
	// the widget.Buttton has a special channels that sends out wde.Buttons
	// whenever its clicked. Here we set up something that changes the
	// label's text every time a click is received.
	clicker := make(widgets.Clicker)
	b.AddClicker <- clicker
	go func() {
		for _ = range clicker {
			b.Label.SetConfig <- ld
		}
	}()

	l := widgets.NewLabel(geom.Coord{100, 50}, widgets.LabelData{"text", 14, color.Black})

	b2 := widgets.NewButton(geom.Coord{70, 30}, "there")
	ld2 := <-b2.Label.GetConfig
	ld2.Text = "BAM"
	clicker2 := make(widgets.Clicker)
	b2.AddClicker <- clicker2
	go func() {
		for _ = range clicker2 {
			b.Label.SetConfig <- ld2
			b2.Label.SetConfig <- ld
			l.SetConfig <- widgets.LabelData{"oops", 14, color.Black}
		}
	}()

	cb := widgets.NewCheckbox(geom.Coord{50, 50})

	kg := widgets.NewKeyGrab(geom.Coord{50, 50})
	kg2 := widgets.NewKeyGrab(geom.Coord{50, 50})

	fl.Add <- &b.Block
	fl.Add <- &l.Block
	fl.Add <- &kg.Block
	fl.Add <- &b2.Block
	fl.Add <- &cb.Block
	fl.Add <- &kg2.Block

	w.Show()

	// Here we set up a subscription on the window's close events.
	done := make(chan interface{})
	isDone := func(e interface{}) (accept, done bool) {
		_, accept = e.(uik.CloseEvent)
		done = accept
		return
	}
	w.Block.Subscribe <- uik.Subscription{isDone, done}

	// once a close event comes in on the subscription, end the program
	<-done

	w.W.Close()

}
