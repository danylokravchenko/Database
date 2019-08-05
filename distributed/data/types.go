package data

import "sync"

// Database struct
type OneAndOnlyNumber struct {
	num        int
	generation int
	numMutex   sync.RWMutex
}

/**
 * Init the db
 * @val number to init with
 * @returns a pointer to database
 */
func InitTheNumber(val int) *OneAndOnlyNumber {

	return &OneAndOnlyNumber{
		num: val,
	}

}

/**
 * Set new value to the db and update current generation
 * @val number to setup
 */
func (n *OneAndOnlyNumber) SetValue(newVal int) {

	n.numMutex.Lock()
	defer n.numMutex.Unlock()
	n.num = newVal
	n.generation++

}

/**
 * Get value and generation of database
 * @returns number value and generation of the database
 */
func (n *OneAndOnlyNumber) GetValue() (int, int) {

	n.numMutex.RLock()
	defer n.numMutex.RUnlock()
	return n.num, n.generation

}

/**
 * Check if database is outdated and update value if it is
 * @curVal current value of the cluster
 * @curGeneration current generation of the cluster
 * @returns a boolean if data was updated
 */
func (n *OneAndOnlyNumber) NotifyValue(curVal int, curGeneration int) bool {

	if curGeneration > n.generation {
		n.numMutex.Lock()
		defer n.numMutex.Unlock()
		n.generation = curGeneration
		n.num = curVal
		return true
	}

	return false

}
