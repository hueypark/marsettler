package data

import "sort"

func Building(id int) *BuildingData {
	return buildingMap[id]
}

func Buildings() []*BuildingData {
	return buildings
}

type BuildingData struct {
	ActorID int
}

var buildingMap = map[int]*BuildingData{
	1: {
		1,
	},
}

var buildings []*BuildingData

func initBuilding() {
	var ids []int
	for id := range buildingMap {
		ids = append(ids, id)
	}

	sort.Ints(ids)

	for _, id := range ids {
		buildings = append(buildings, buildingMap[id])
	}
}
