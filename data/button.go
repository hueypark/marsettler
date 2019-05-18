package data

import (
	"log"
	"sort"

	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/client/asset"
)

func Button(id int) *ButtonData {
	return buttons[id]
}

func Buttons() []*ButtonData {
	var ids []int
	for id := range buttons {
		ids = append(ids, id)
	}
	sort.Ints(ids)

	var buttonSlice []*ButtonData
	for _, id := range ids {
		buttonSlice = append(buttonSlice, buttons[id])
	}

	return buttonSlice
}

type ButtonData struct {
	Name    string
	Image   *ebiten.Image
	OnClick func()
}

var buttons = map[int]*ButtonData{
	1: {
		"CityHall",
		asset.Building,
		func() {
			log.Println("City hall button clicked.")
		},
	},
	2: {
		"LogHouse",
		asset.Building,
		func() {
			log.Println("Log house button clicked.")
		},
	},
}
