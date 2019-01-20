package decorator

import (
	"gitlab.com/legionary/legionary/core/ai"
)

type Repeat struct {
	ai.Decorator

	limit int
	count int
}

func NewRepeat(limit int) *Repeat {
	decorator := &Repeat{}
	decorator.limit = limit
	decorator.count = 0

	return decorator
}

func (decorator *Repeat) Init() {
	decorator.Decorator.Init()

	decorator.count = 0
}

func (decorator *Repeat) Update(delta float64) ai.State {
	state := decorator.Child().Update(delta)

	if state != ai.Success {
		return state
	}

	decorator.count++

	if decorator.count != decorator.limit {
		decorator.Child().Init()

		return ai.Running
	} else {
		return ai.Success
	}
}
