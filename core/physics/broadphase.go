package physics

import "github.com/hueypark/marsettler/core/math/vector"

func broadphase(bodies map[int64]Body) (contacts []Contact) {
	for _, lhs := range bodies {
		for _, rhs := range bodies {
			if lhs.Id() <= rhs.Id() {
				continue
			}

			contacts = append(contacts, Contact{lhs, rhs, vector.Zero(), 0})
		}
	}

	return contacts
}
