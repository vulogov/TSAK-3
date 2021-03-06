package tsak

import (
	"fmt"
	"os"
	"path/filepath"

	zygo "github.com/glycerine/zygomys/zygo"
	"github.com/peterh/liner"
	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/TSAK-3/internal/conf"
	"github.com/vulogov/TSAK-3/internal/dsl"
	"github.com/vulogov/TSAK-3/internal/signal"
)

func Shell() {
	Init()
	var fnHistory = filepath.Join(os.TempDir(), fmt.Sprintf(".tsak_history"))
	log.Debug("[ TSAK-3 ] tsak.Shell() is reached")
	log.Debugf("[ TSAK-3 ] Shell history is stored in %v", fnHistory)

	line := liner.NewLiner()

	if f, err := os.Open(fnHistory); err == nil {
		line.ReadHistory(f)
		f.Close()
	} else {
		log.Errorf("[ TSAK-3 ] Error accessing to a shell history file:  %v", fnHistory)
	}
	cfg := dsl.InitDSL()
	dsl.Cfg = cfg
	env := dsl.MakeEnvironment(cfg)
	dsl.Env = env
	log.Debugf("TSAK_script bootstrap file: %v", *conf.BootStrap)
	file, err := os.Open(*conf.BootStrap)
	err = env.LoadFile(file)
	dsl.PanicOn(err)
	_, err = env.Run()
	dsl.PanicOn(err)
	log.Info("Entering REPL loop. Ctrl-D to exit")
	for {
		if value, err := line.Prompt(cfg.Prompt); err == nil {
			log.Debug("COMMAND: ", value)
			line.AppendHistory(value)
			res, err := env.EvalString(value)
			switch err {
			case nil:
			case zygo.NoExpressionsFound:
				env.Clear()
				continue
			default:
				fmt.Print(env.GetStackTrace(err))
				env.Clear()
				continue
			}
			log.Infof("RETURNED: %v", res.SexpString(nil))
			if signal.ExitRequested() {
				log.Infof("EXIT REQUESTED: %v", signal.Len())
				break
			}
		} else if err == liner.ErrPromptAborted {
			log.Error("Aborted")
			signal.ExitRequest()
			break
		} else {
			log.Error("Error reading line: ", err)
			signal.ExitRequest()
			break
		}
	}
	if f, err := os.Create(fnHistory); err != nil {
		log.Errorf("Error writing history file: %v", err)
	} else {
		line.WriteHistory(f)
		f.Close()
	}
	dsl.CloseEnvironment(cfg, env)
}
