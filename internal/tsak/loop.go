package tsak

import (
	"os"
	"time"

	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/TSAK-3/internal/conf"
	"github.com/vulogov/TSAK-3/internal/dsl"
	"github.com/vulogov/TSAK-3/internal/signal"
)

func Loop() {
	Init()
	log.Debug("[ TSAK-3 ] tsak.Loop() is reached")
	cfg := dsl.InitDSL()
	dsl.Cfg = cfg
	env := dsl.MakeEnvironment(cfg)
	dsl.Env = env
	if *conf.BootStrap != "" {
		log.Debugf("TSAK_script bootstrap file: %v", *conf.BootStrap)
		file, err := os.Open(*conf.BootStrap)
		err = env.LoadFile(file)
		dsl.PanicOn(err)
		_, err = env.Run()
		dsl.PanicOn(err)
	}
	log.Debug("ExitRequest event loop reached")
	for !signal.ExitRequested() {
		time.Sleep(100 * time.Millisecond)
	}
	signal.ExitRequest()
}
