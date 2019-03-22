package behavior_tree

// Blackboard is data storage for behavior tree.
type Blackboard struct {
	datas map[BlackboardKey]interface{}
}

// NewBlackboard creates new blackboard.
func NewBlackboard() *Blackboard {
	blackboard := &Blackboard{
		make(map[BlackboardKey]interface{}),
	}

	return blackboard
}

// Get returns value in key.
func (blackboard *Blackboard) Get(key BlackboardKey) interface{} {
	return blackboard.datas[key]
}

// Set sets value in key.
func (blackboard *Blackboard) Set(key BlackboardKey, value interface{}) {
	blackboard.datas[key] = value
}

// Delete deletes value in key
func (blackboard *Blackboard) Delete(key BlackboardKey) {
	delete(blackboard.datas, key)
}

// GetInt64s returns value in key.
func (blackboard *Blackboard) GetInt64s(key BlackboardKey) *[]int64 {
	return blackboard.datas[key].(*[]int64)
}

// SetInt64s sets value in key.
func (blackboard *Blackboard) SetInt64s(key BlackboardKey, value *[]int64) {
	blackboard.datas[key] = value
}

// BlackboardKey represents blackboard key.
type BlackboardKey int
