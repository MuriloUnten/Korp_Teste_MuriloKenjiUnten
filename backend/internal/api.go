package internal

import (
	"net/http"
	"fmt"
	"encoding/json"
	"errors"
	"strconv"
)

type APIFunc func(w http.ResponseWriter, r *http.Request) error

func MakeHandler(handler APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request from %s\t%s %s\n", r.RemoteAddr, r.Method, r.URL.Path)

		err := handler(w, r)
		if err != nil {
			if e, ok := err.(APIError); ok {
				fmt.Println("API error:", e.Msg)
				WriteJSON(w, e.StatusCode, e)
			} else {
				fmt.Println("error:", err)
				WriteJSON(w, http.StatusInternalServerError, "Internal Error")
			}
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func GetPathId(wildcard string, r *http.Request) (int, error) {
	v := r.PathValue(wildcard)
	if v == "" {
		return 0, errors.New("unable to get path id")
	}

	id, err := strconv.Atoi(v)
	return id, err
}

