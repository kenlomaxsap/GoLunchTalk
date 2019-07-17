package main

import (
	"fmt"
	"time"
)

type piece struct {
	x  int
	y  int
	vx int
	vy int
	id string
	in chan (msg)
}

type msg struct {
	cmd      string
	dx       int
	response chan msg
}

var register = make(chan piece)

func main() {
	// create a new piece
	p := piece{x: 1, y: 1, vx: 2, vy: 0, id: "alien1", in: make(chan msg)}
	// Kick off its conductor routine
	go p.conductor()
	go p.listenAndAct()

	// Wait for status updates
	var ticker = time.NewTicker(200 * time.Millisecond)
	for _ = range ticker.C {
		p.in <- msg{cmd: "Report"}
		status := <-register
		fmt.Printf("Status of piece is %+v \n", status)
	}

	<-make(chan bool)
}

func (p *piece) conductor() {
	var ticker = time.NewTicker(500 * time.Millisecond)
	for _ = range ticker.C {
		p.in <- msg{cmd: "Move", dx: 2}

	}
}

func (p *piece) listenAndAct() {
	// The state of the piece is "here"

	for {
		m := <-p.in
		if m.cmd == "Move" {
			p.x = p.x + m.dx
		}
		if m.cmd == "Report" {
			register <- *p
		}
	}
}
