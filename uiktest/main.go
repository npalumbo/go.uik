package main

import (
	"fmt"
	"github.com/skelterjohn/go.wde"
	"github.com/skelterjohn/go.uik"
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
	width := 480.0
	height := 320.0
	w, err := uik.NewWindow(nil, int(width), int(height))
	if err != nil {
		fmt.Println(err)
		return
	}
	w.W.SetTitle("GoUI")
	
	bw := width/4
	bh := width/4
	bsize := uik.Coord{bw, bh}
	
	b := uik.NewButton(bsize, "Hi")
	w.AddBlock(&b.Block)
	w.Show()
	<-w.Done
}
