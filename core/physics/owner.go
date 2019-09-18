package physics

import "github.com/hueypark/marsettler/core/physics/body"

type owner interface {
	ID() int64
	Body() *body.Body
}
