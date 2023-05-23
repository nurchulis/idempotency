package contract

import (
	"github.com/nurchulis/idempotency/internal/presentations"
	"net/http"
)

// UseCase is a use case contract
type UseCase interface {
	Serve(w http.ResponseWriter, r *http.Request) presentations.Response
}
