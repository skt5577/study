package adapter

type Position struct {
	x int
	y int
}
type OldMovement struct {
}

func (movement *OldMovement) up(p *Position) {
	p.y++
}
func (movement *OldMovement) down(p *Position) {
	p.y--
}
func (movement *OldMovement) left(p *Position) {
	p.x--
}
func (movement *OldMovement) right(p *Position) {
	p.x++
}

type NewMovement interface {
	up(p *Position)
	down(p *Position)
	left(p *Position)
	right(p *Position)
	leftUp(p *Position)
	leftDown(p *Position)
	rightUp(p *Position)
	rightDown(p *Position)
}
