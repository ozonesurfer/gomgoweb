// config
package gomgoweb

import (
	"github.com/QLeelulu/goku"
	"path"
	"runtime"
	"time"
)

var Config *goku.ServerConfig = &goku.ServerConfig{
	Addr:           "localhost:8080",
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 20,
	StaticPath:     "static",
	ViewPath:       "views",
	LogLevel:       goku.LOG_LEVEL_LOG,
	Debug:          true,
}

const (
	DATABASE     = "gomgoweb-8"
	BAND_COL     = "bands"
	LOCATION_COL = "locations"
	GENRE_COL    = "genres"
	ALBUM_COL    = "albums"
	DATABASE_DSN = "localhost"
)

func init() {
	// project root directory
	_, filename, _, _ := runtime.Caller(1)
	Config.RootDir = path.Dir(filename)
}
