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
	http.HandleFunc("/enumCrossPoints", cors(guard(parser(enumerateCrossPointsHandler))))
	http.HandleFunc("/recomendCrossPoints", cors(guard(parser(recomendClossPointsHandler))))
	http.HandleFunc("/detectionHighWays", cors(guard(parser(detectionHighWays))))
	http.HandleFunc("/shortestPath", cors(guard(parser(shortestPathHandler))))

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

func parser(handler func(http.ResponseWriter, *utils.Datas)) http.HandlerFunc {
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

		datas, err := utils.ParseData(string(body))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		handler(w, datas)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server running")
}

func enumerateCrossPointsHandler(w http.ResponseWriter, datas *utils.Datas) {
	enumerateCrossPoints(datas)
	idReregistration(datas)

	fmt.Fprint(w, utils.DatasToQuerys(*datas))
}

func recomendClossPointsHandler(w http.ResponseWriter, datas *utils.Datas) {
	enumerateCrossPoints(datas)
	recomendClossPoints(datas)
	idReregistration(datas)

	fmt.Fprint(w, utils.DatasToQuerys(*datas))
}

func detectionHighWays(w http.ResponseWriter, datas *utils.Datas) {
	enumerateCrossPoints(datas)
	recomendClossPoints(datas)
	idReregistration(datas)

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

func shortestPathHandler(w http.ResponseWriter, datas *utils.Datas) {
	enumerateCrossPoints(datas)
	recomendClossPoints(datas)
	idReregistration(datas)

	shortestPaths := make(map[string][][]*model.Road)
	for _, q := range datas.Queries {
		key := fmt.Sprintf("%s %s %d", q.Start, q.Dest, q.Num)
		shortestPaths[key] = phase2.CalcKthShortestPath(*q, datas.Places, datas.Roads)
	}

	type jsonData struct {
		Paths map[string][][][]string `json:"paths"`
		Query string                  `json:"query"`
	}

	result := new(jsonData)
	result.Query = utils.DatasToQuerys(*datas)
	result.Paths = make(map[string][][][]string)
	for key, paths := range shortestPaths {
		result.Paths[key] = make([][][]string, len(paths))
		for i, path := range paths {
			result.Paths[key][i] = make([][]string, len(path))
			for j, road := range path {
				var r []string
				if road.From.Id < road.To.Id {
					r = []string{road.From.Id, road.To.Id}
				} else {
					r = []string{road.To.Id, road.From.Id}
				}
				result.Paths[key][i][j] = r
			}
		}
	}

	json, err := json.Marshal(*result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}
