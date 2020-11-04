package behavior_tree

import "sync"

var (
	nextKey BlackboardKey
	keys    map[string]BlackboardKey
	strKeys map[BlackboardKey]string
	mux     sync.Mutex
)

func init() {
	nextKey = 0
	keys = make(map[string]BlackboardKey)
	strKeys = make(map[BlackboardKey]string)
}

// Key returns key from string key.
func Key(strKey string) BlackboardKey {
	mux.Lock()
	defer mux.Unlock()

	if key, ok := keys[strKey]; ok {
		return key
	}

	curKey := nextKey
	nextKey++
	keys[strKey] = curKey
	strKeys[curKey] = strKey
	nextKey++

	return curKey
}

// StrinsKey returns string key from key.
func StringKey(key BlackboardKey) string {
	mux.Lock()
	defer mux.Unlock()

	if strKey, ok := strKeys[key]; ok {
		return strKey
	}

	return ""
}
