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

// Dish in the menu
type Dish struct {
	ID int64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	// price of dish
	Price int64 `json:"price,omitempty"`
}
