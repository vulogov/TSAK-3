package tsak

import (
	"runtime/pprof"

	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/TSAK-3/internal/banner"
)

func Fin() {
	log.Debug("[ TSAK-3 ] tsak.Fin() is reached")
	log.Debug("[ TSAK-3 ] Stopping CPU profile")
	pprof.StopCPUProfile()
	banner.Banner("[ Zay Gezunt ]")
}
