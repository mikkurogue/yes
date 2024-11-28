package entity

type Player struct {
	Health    int
	XPosition int
	YPosition int
	// ZPosition int // maybe think about weird Z positions, maybe for map layering?
	Spawned bool
	drawable uint32
}

func (*p Player) Draw() {
 if p.Spawned || p.Health > 0 {
   return;
 }

 // TODO: Make the correct vertex array and points.
 	gl.BindVertexArray(p.drawable)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(square)/3))
}

func Spawn() Player {
	return Player{
		Health:    100,
		Spawned: true,
		XPosition: 0,
		YPosition: 0,
		
	}
}

func (p *Player) Move() {

	// TODO: Figure out how we can move the player on the canvas

}
