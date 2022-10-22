package tls

import (
	"net/http"

	"github.com/roadrunner-server/http/v3/attributes"
)

func tlsAttributes(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = attributes.Init(r)
		attributes.Set(r, "davekey", "gebvalue")
		next.ServeHTTP(w, r)
	})
}
