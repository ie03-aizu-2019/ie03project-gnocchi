package main

import (
	"fmt"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
	"github.com/uzimaru0000/ie03project-gnocchi/back/phase1"
	"github.com/uzimaru0000/ie03project-gnocchi/back/phase2"
	"github.com/uzimaru0000/ie03project-gnocchi/back/utils"
)

func enumerateCrossPoints(datas *utils.Datas) {
	roads, crossPoints := phase1.EnumerateCrossPoints(datas.Roads)
	places := append(datas.Places, crossPoints...)
	roads = phase1.ConnectOnRoadPoints(roads, places)

	datas.Roads = roads
	datas.Places = places
}

func recomendClossPoints(datas *utils.Datas) {
	recomendRoads, recomendPlaces := phase2.CreateRecomendRoads(datas.Places, datas.Roads, datas.AddPlaces)

	recomendPlaces = append(recomendPlaces, datas.AddPlaces...)
	datas.Places = append(datas.Places, recomendPlaces...)
	datas.Roads = append(datas.Roads, recomendRoads...)
	datas.AddPlaces = []*model.Place{}

	roads := phase1.ConnectOnRoadPoints(datas.Roads, datas.Places)
	roads, places := phase1.EnumerateCrossPoints(roads)
	datas.Roads = roads
	datas.Places = append(datas.Places, places...)
}

func idReregistration(datas *utils.Datas) {
	placeIDTable := make(map[string]string)
	for i, place := range datas.Places {
		placeIDTable[place.Id] = fmt.Sprint(i + 1)
		place.Id = fmt.Sprint(i + 1)
	}
	for i, road := range datas.Roads {
		road.Id = i
	}
	for _, query := range datas.Queries {
		query.Start = placeIDTable[query.Start]
		query.Dest = placeIDTable[query.Dest]
	}
}
