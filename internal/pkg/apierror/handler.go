package apierror

import (
	"errors"
	"net/http"
)

type ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)

func ResponseErrorHandler() ResponseErrorHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, err error) {
		var apiErr *ApiError
		switch {
		case errors.As(err, &apiErr):
			handleAppErr(w, err)
		default:
			w.WriteHeader(http.StatusTeapot)
			w.Write(internalError(err).Marshal())
		}
	}
}

func handleAppErr(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, ErrNotFound):
		http.Error(w, string(ErrNotFound.Marshal()), http.StatusNotFound)
		return
	default:
		err := err.(*ApiError)
		http.Error(w, string(err.Marshal()), http.StatusInternalServerError)
	}
}
