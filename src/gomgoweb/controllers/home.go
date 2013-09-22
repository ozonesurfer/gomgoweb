// home
package controllers

import (
	"fmt"
	"github.com/QLeelulu/goku"
	"gomgoweb"
	"gomgoweb/models"
)

var HomeController = goku.Controller("home").
	Get("index", func(ctx *goku.HttpContext) goku.ActionResulter {
	ctx.ViewData["Title"] = "CD Catalog Site"
	bands := models.GetAll(gomgoweb.BAND_COL)
	DocDump()
	return ctx.View(bands)
}).Get("genrelist", func(ctx *goku.HttpContext) goku.ActionResulter {
	ctx.ViewData["Title"] = "List of Genres"
	genres := models.GetAll(gomgoweb.GENRE_COL)
	return ctx.View(genres)
}).Get("bygenre", func(ctx *goku.HttpContext) goku.ActionResulter {
	rawId := ctx.RouteData.Params["id"]
	genreId := ToObjectId(rawId)
	fmt.Println("genreId =", genreId)
	genreName := models.GetGenreName(genreId)
	ctx.ViewData["Title"] = fmt.Sprintf("%s Albums", genreName)
	ctx.ViewData["GenreId"] = genreId
	bands := models.GetBandsByGenre(genreId)
	return ctx.View(bands)
})
