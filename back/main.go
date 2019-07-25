package main

import (
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

	roads := phase1.ConnectOnRoadPoints(datas.Roads, datas.Places)
	roads, places := phase1.EnumerateCrossPoints(roads)

	datas.Roads = roads
	datas.Places = append(datas.Places, places...)

	fmt.Fprint(w, utils.DatasToQuerys(*datas))
}

func recomendClossPoints(w http.ResponseWriter, query string) {
	datas, err := utils.ParseData(query)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	recomendPoints := phase2.RecomendClossPoint(datas.Roads, datas.Places)

	recomendPlaces := []*model.Place{}
	for i, rp := range recomendPoints {
		p := &model.Place{
			Coord: *rp,
			Id:    fmt.Sprintf("R%d", i),
		}

		recomendPlaces = append(recomendPlaces, p)
	}
	datas.Places = recomendPlaces

	fmt.Fprint(w, utils.DatasToQuerys(*datas))
}
