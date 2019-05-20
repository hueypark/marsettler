package physics

// World represents physics world.
type World struct {
	bodies map[int64]Body
}

func NewWorld() *World {
	world := &World{
		bodies: make(map[int64]Body),
	}

	return world
}

// Tick updates world.
func (world *World) Tick() {
	contacts := broadphase(world.bodies)

	for _, contact := range contacts {
		contact.SolveCollision()
	}
}

// AddRigidBody adds body to world.
func (world *World) AddBody(b Body) {
	world.bodies[b.ID()] = b
}

func (world *World) RemoveBody(id int64) {
	delete(world.bodies, id)
}
