package infra

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/HyperloopUPV-H8/Backend-H8/Shared/server/domain"
)

func (server HTTPServer[D, O, M]) HandlePodData(route string, podData domain.PodData) {
	server.router.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		r.Body.Close()

		encodedPodData, err := json.Marshal(podData)
		if err != nil {
			log.Fatalln(err)
		}

		w.Write(encodedPodData)
	})
}
