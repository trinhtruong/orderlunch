/*
 * Order Lunch Ginno
 *
 * This is a order lunch server for everyone worked at Ginnovation. APIs are writed by swagger [http://swagger.io](http://swagger.io) or on  [irc.freenode.net, #swagger](http://swagger.io/irc/).
 *
 * API version: 1.0.0
 * Contact: apiteam@ginno.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"log"

	sw "./go"
)

func main() {
	log.Printf("Server started")

	sw.Initialize("postgres", "1", "orderlunch")
	sw.Run(":8083")
}
