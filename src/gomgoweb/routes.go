// routes
package gomgoweb

import (
	//	"fmt"
	"github.com/QLeelulu/goku"
)

// routes
var Routes []*goku.Route = []*goku.Route{
	// static file route
	&goku.Route{
		Name:     "static",
		IsStatic: true,
		Pattern:  "/static/(.*)",
	},
	// default controller and action route
	&goku.Route{
		Name:    "default",
		Pattern: "/{controller}/{action}/{id}",
		Default: map[string]string{"controller": "home", "action": "index",
			"id": "0"},
	},
}
