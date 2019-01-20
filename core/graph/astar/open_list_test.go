package astar

import (
	"fmt"
)

func ExampleOpenLisPop() {
	pq := NewOpenList()
	pq.Push(
		&Score{
			1,
			1,
			0,
			0,
			0})
	pq.Push(
		&Score{
			2,
			3,
			0,
			0,
			0})
	pq.Push(
		&Score{
			3,
			2,
			0,
			0,
			0})
	pq.Push(
		&Score{
			4,
			4,
			0,
			0,
			0})

	for !pq.Empty() {
		fmt.Println(pq.Pop().F)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
}

func ExampleOpenLisGet() {
	pq := NewOpenList()
	pq.Push(
		&Score{
			1,
			10,
			0,
			0,
			0})
	fmt.Println(pq.Get(1).F)
	// Output:
	// 10
}
