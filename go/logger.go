/*
 * Order Lunch Ginno
 *
 * This is a order lunch server for everyone worked at Ginnovation. APIs are writed by swagger [http://swagger.io](http://swagger.io) or on  [irc.freenode.net, #swagger](http://swagger.io/irc/).
 *
 * API version: 1.0.0
 * Contact: apiteam@ginno.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"log"
	"net/http"
	"time"
)

// Logger Logger
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
