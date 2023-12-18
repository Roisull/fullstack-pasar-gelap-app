package log

import (
	"log"
	"net/http"
	"time"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        // Izinkan permintaan untuk dilanjutkan ke handler selanjutnya
        next.ServeHTTP(w, r)

        // Logging setelah permintaan selesai
        log.Printf(
            "[%s] %s %s %s",
            r.Method,
            r.RequestURI,
            time.Since(start),
            r.RemoteAddr,
        )
    })
}