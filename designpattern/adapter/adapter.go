package adapter

type MovementAdapter struct {
	old OldMovement
}

func (movement *MovementAdapter) up(p *Position) {
	movement.old.up(p)
}
func (movement *MovementAdapter) down(p *Position) {
	movement.old.down(p)
}
func (movement *MovementAdapter) left(p *Position) {
	movement.old.left(p)
}
func (movement *MovementAdapter) right(p *Position) {
	movement.old.right(p)
}
func (movement *MovementAdapter) leftUp(p *Position) {
	movement.old.left(p)
	movement.old.up(p)
}
func (movement *MovementAdapter) leftDown(p *Position) {
	movement.old.left(p)
	movement.old.down(p)
}
func (movement *MovementAdapter) rightUp(p *Position) {
	movement.old.right(p)
	movement.old.up(p)
}
func (movement *MovementAdapter) rightDown(p *Position) {
	movement.old.right(p)
	movement.old.down(p)
}
