package tsak

import (
	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/TSAK-3/internal/conf"
	tlog "github.com/vulogov/TSAK-3/internal/log"
	"github.com/vulogov/TSAK-3/internal/signal"
	"github.com/vulogov/TSAK-3/internal/snmp"
)

func Init() {
	tlog.Init()
	log.Debug("[ TSAK-3 ] tsak.Init() is reached")
	signal.InitSignal()
	if *conf.IsSNMPAgent {
		snmp.InitSNMPAgent()
	} else {
		log.Debug("[ TSAK-3 ] SNMP Agent is disabled")
	}
	if *conf.IsSNMPTrap {
		snmp.InitSNMPTrapReceiver()
	} else {
		log.Debug("[ TSAK-3 ] SNMP Trapd is disabled")
	}
}
