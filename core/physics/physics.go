package physics

// World represents physics world.
type World struct {
	bodies map[int64]Body
}

func (world *World) Init() {
	world.bodies = make(map[int64]Body)
}

// Update updates world.
func (world *World) Update(delta float64) {
	contacts := broadphase(world.bodies)

	for _, contact := range contacts {
		contact.SolveCollision()
	}
}

// AddRigidBody adds body to world.
func (world *World) AddBody(b Body) {
	world.bodies[b.Id()] = b
}

func (world *World) RemoveBody(id int64) {
	delete(world.bodies, id)
}
