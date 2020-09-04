package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
			fmt.Printf("[" + time.Now().Format(time.UnixDate) + "] " + r.RemoteAddr + " [" + r.Method + "] " + /*"[" +  + "] " +*/ /*r.URL.String()*/ r.Host + r.URL.String() + "\n")
		},
	)
}