package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/stakkato95/service-engineering-go-lib/logger"
)

func WriteResponse(w http.ResponseWriter, code int, data interface{}) {
	//1 Content-Type - always first
	w.Header().Add("Content-Type", "application/json")
	//2 HTTP status code
	w.WriteHeader(code)
	//3 body
	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.Panic("can not wrote json: " + err.Error())
	}
}
