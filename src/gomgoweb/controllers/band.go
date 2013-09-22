// band
package controllers

import (
	//	"encoding/json"
	"fmt"
	"github.com/QLeelulu/goku"
	"gomgoweb"
	"gomgoweb/models"
	"labix.org/v2/mgo/bson"
	"strings"
)

var BandController = goku.Controller("band").
	Get("add", func(ctx *goku.HttpContext) goku.ActionResulter {
	ctx.ViewData["Title"] = "Adding A Band"
	locations := models.GetAll(gomgoweb.LOCATION_COL)
	return ctx.View(locations)
	//return ctx.Html("not implemented")
}).Post("verify", func(ctx *goku.HttpContext) goku.ActionResulter {
	ctx.ViewData["Title"] = "Verifying Band"
	name := strings.TrimSpace(ctx.Request.FormValue("name"))
	loctype := ctx.Request.FormValue("loctype")
	var locationId bson.ObjectId
	errorString := "no errors"
	switch loctype {
	case "existing":
		if ctx.Request.FormValue("location_id") == "" {
			errorString = "No location was selected"
		} else {
			locationId = ToObjectId(ctx.Request.FormValue("location_id"))
		}
		break
	case "new":
		if ctx.Request.FormValue("country") != "" {
			locationId = models.GenerateId()
			city := ctx.Request.FormValue("city")
			state := ctx.Request.FormValue("state")
			country := ctx.Request.FormValue("country")
			location := models.Location{
				City:    city,
				State:   state,
				Country: country}
			//		var q interface{} = location

			doc := models.MyDoc{Id: locationId,
				Value: bson.M{"City": location.City,
					"State": location.State, "Country": location.Country}}
			//		json.Unmarshal(q.([]byte), &doc.Value)
			err := models.AddDoc(doc, gomgoweb.LOCATION_COL)
			if err != nil {
				errorString = "error on location add: " + err.Error()
			}
		} else {
			errorString = "Country is required"
		}
		break
	}
	if errorString == "no errors" {
		var albums []models.Album
		//		id := models.GenerateId()
		band := models.Band{Name: name, LocationId: locationId, Albums: albums}

		doc := models.MyDoc{Id: models.GenerateId(), Value: bson.M{"Name": band.Name,
			"LocationId": band.LocationId, "Albums": band.Albums}}
		/*	var q interface{} = band
			json.Unmarshal(q.([]byte), &doc.Value) */

		err := models.AddDoc(doc, gomgoweb.BAND_COL)
		if err != nil {
			errorString = fmt.Sprintf("Band add: %s", err.Error())
		}

	}
	ctx.ViewData["Message"] = errorString
	return ctx.View(nil)
})
