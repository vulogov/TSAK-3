package tsak

import (
	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/TSAK-3/internal/banner"
)

func Fin() {
	log.Debug("[ TSAK-3 ] tsak.Fin() is reached")
	banner.Banner("[ Zay Gezunt ]")
}
