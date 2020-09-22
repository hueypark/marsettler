package physics

import "github.com/hueypark/marsettler/pkg/internal/physics/body"

type owner interface {
	ID() int64
	Body() *body.Body
}
