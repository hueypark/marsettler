package behavior_tree

func init() {

}

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

// GetInt returns int value in key.
func (blackboard *Blackboard) GetInt(key BlackboardKey) *int {
	val, ok := blackboard.datas[key]
	if !ok {
		return nil
	}

	return val.(*int)
}

// GetInt64 returns int64 value in key.
func (blackboard *Blackboard) GetInt64(key BlackboardKey) *int64 {
	val, ok := blackboard.datas[key]
	if !ok {
		return nil
	}

	return val.(*int64)
}

// Set sets value in key.
func (blackboard *Blackboard) Set(key BlackboardKey, value interface{}) {
	blackboard.datas[key] = value
}

// SetInt sets int value in key.
func (blackboard *Blackboard) SetInt(key BlackboardKey, val int) {
	blackboard.datas[key] = &val
}

// SetInt64 sets int64 value in key.
func (blackboard *Blackboard) SetInt64(key BlackboardKey, val int64) {
	blackboard.datas[key] = &val
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
