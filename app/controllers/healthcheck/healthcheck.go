package healthcheck

import (
	"encoding/json"
	"lexes_learn_server/pkg/models"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	//_, err := fmt.Fprint(w, "Hello, World!")
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//}

	health := &models.HealthCheck{
		ServerName: "Lexes Learn Server",
		Author:     "Prince Bobby",
		Version:    "1.0.0",
		Health:     "Alive",
	}


	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(health)
	if err != nil {
		log.Printf("Unlable to check health of server")
	}
}