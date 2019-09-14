package physics

import "github.com/hueypark/marsettler/core/physics/body"

type owner interface {
	Body() *body.Body
}
