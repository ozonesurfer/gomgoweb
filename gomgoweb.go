// gomgoweb
package main

import (
	//	"fmt"
	"github.com/QLeelulu/goku"
	"gomgoweb"
	"gomgoweb/controllers"
	//	"gomgoweb/models"
	//	"labix.org/v2/mgo/bson"
	"log"
)

func main() {
	//	fmt.Println("Hello World!")
	rt := &goku.RouteTable{Routes: gomgoweb.Routes}
	middlewares := []goku.Middlewarer{}
	s := goku.CreateServer(rt, middlewares, gomgoweb.Config)
	goku.Logger().Logln("Server start on", s.Addr)
	log.Fatal(s.ListenAndServe())
}

var home = controllers.HomeController
var band = controllers.BandController
var album = controllers.AlbumController
