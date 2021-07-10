package tsak

import (
	"fmt"

	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/TSAK-3/internal/banner"
	"github.com/vulogov/TSAK-3/internal/conf"
)

func Version() {
	Init()
	log.Debug("[ TSAK-3 ] tsak.Version() is reached")
	banner.Banner(fmt.Sprintf("[ TSAK %v ]", conf.EVersion))
	banner.Table()
}
