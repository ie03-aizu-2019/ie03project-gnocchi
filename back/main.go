package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
	"github.com/uzimaru0000/ie03project-gnocchi/back/phase1"
	"github.com/uzimaru0000/ie03project-gnocchi/back/phase2"

	"github.com/uzimaru0000/ie03project-gnocchi/back/utils"
)

func main() {
	http.HandleFunc("/", healthCheck)
	http.HandleFunc("/enumCrossPoints", cors(guard(parser(enumerateCrossPoints))))
	http.HandleFunc("/recomendCrossPoints", cors(guard(parser(recomendClossPoints))))
	http.HandleFunc("/detectionHighWays", cors(guard(parser(detectionHighWays))))

	log.Printf("The server is running at http://localhost:5000")
	http.ListenAndServe(":5000", nil)
}

func cors(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		handler(w, req)
	}
}

func guard(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		//Validate request
		if req.Method != "POST" {
			log.Printf("Method is %s", req.Method)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if req.Header.Get("Content-Type") != "text/plain" {
			log.Printf("Content-Type is %s", req.Header.Get("Content-Type"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		handler(w, req)
	}
}

func parser(handler func(http.ResponseWriter, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		//To allocate slice for request body
		length, err := strconv.Atoi(req.Header.Get("Content-Length"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		//Read body data to parse json
		body := make([]byte, length)
		length, err = req.Body.Read(body)
		if err != nil && err != io.EOF {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/plain")

		handler(w, string(body))
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server running")
}

func enumerateCrossPoints(w http.ResponseWriter, query string) {
	datas, err := utils.ParseData(query)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	roads, crossPoints := phase1.EnumerateCrossPoints(datas.Roads)
	places := append(datas.Places, crossPoints...)
	roads = phase1.ConnectOnRoadPoints(roads, places)

	datas.Roads = roads
	datas.Places = places

	fmt.Fprint(w, utils.DatasToQuerys(*datas))
}

func recomendClossPoints(w http.ResponseWriter, query string) {
	datas, err := utils.ParseData(query)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	recomendPoints := phase2.RecomendClossPoint(datas.Roads, datas.AddPlaces)

	recomendPlaces := []*model.Place{}
	recomendRoads := []*model.Road{}
	for i, rp := range recomendPoints {
		p := &model.Place{
			Coord: *rp,
			Id:    fmt.Sprintf("R%d", i),
		}
		recomendPlaces = append(recomendPlaces, p)

		road := &model.Road{
			Id:   len(recomendRoads) + len(datas.Roads),
			From: p,
			To:   datas.AddPlaces[i],
		}
		recomendRoads = append(recomendRoads, road)
	}

	recomendPlaces = append(recomendPlaces, datas.AddPlaces...)
	datas.Places = append(datas.Places, recomendPlaces...)
	datas.Roads = append(datas.Roads, recomendRoads...)
	datas.AddPlaces = []*model.Place{}

	roads := phase1.ConnectOnRoadPoints(datas.Roads, datas.Places)
	roads, places := phase1.EnumerateCrossPoints(roads)
	datas.Roads = roads
	datas.Places = append(datas.Places, places...)

	fmt.Fprint(w, utils.DatasToQuerys(*datas))
}

func detectionHighWays(w http.ResponseWriter, query string) {
	datas, err := utils.ParseData(query)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	roads, _ := phase1.EnumerateCrossPoints(datas.Roads)
	highWays := phase2.DetectBridge(roads)

	preData := [][]string{}
	for from, dests := range highWays {
		for _, to := range dests {
			var highWay []string
			if from.Id < to.Id {
				highWay = []string{from.Id, to.Id}
			} else {
				highWay = []string{to.Id, from.Id}
			}

			preData = append(preData, highWay)
		}
	}

	json, err := json.Marshal(preData)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}
