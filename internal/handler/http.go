package handler

import (
	"encoding/json"
	"github.com/nurchulis/idempotency/internal/ucase/contract"
	"net/http"
)

// Http handler func wrapper
func Http(svc contract.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := svc.Serve(w, r)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.Code)
		json.NewEncoder(w).Encode(response)
		return
	}
}
