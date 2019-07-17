package main

import "fmt"

type piece struct {
	x  int
	y  int
	vx int
	vy int
	id string
}

func (p *piece) move() {
	// Do something
	p.x = p.x + 1
}

func main() {
	p := piece{x: 1, y: 1, vx: 2, vy: 0, id: "piece1"}
	fmt.Printf("Piece %+v\n", p)
	p.move()
	fmt.Printf("Piece %+v\n", p)
}

/* Equivalent to Java

class Alien {
	int x;
	int y;
	int vx;
	int vy;
	String id;

	public void move(){
		// Do something
	}
}

**/
