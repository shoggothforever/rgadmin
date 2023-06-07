package middleware

import "net/http"

type CoorsMiddleware struct {
}

func NewCoorsMiddleware() *CoorsMiddleware {
	return &CoorsMiddleware{}
}

func (m *CoorsMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		w.Header().Set("Access-Control-Allow-Origin", "*") //"https://localhost:8080"
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE,PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, X-Extra-Header, Content-Type, Accept, Authorization , User-Agent, Cache-Control, Referer")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		} else if r.Method == "" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		next(w, r)
	}
}
