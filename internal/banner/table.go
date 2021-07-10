package banner

import (
	"fmt"
	"os"

	"github.com/mgutz/ansi"
	"github.com/tomlazar/table"

	"github.com/vulogov/TSAK-3/internal/conf"
)

func Table() {
	var cfg table.Config

	if !*conf.VTable {
		return
	}

	cfg.ShowIndex = true
	if *conf.Color {
		cfg.Color = true
		cfg.AlternateColors = true
		cfg.TitleColorCode = ansi.ColorCode("white+buf")
		cfg.AltColorCodes = []string{"", ansi.ColorCode("white:grey+h")}
	} else {
		cfg.Color = false
		cfg.AlternateColors = false
		cfg.TitleColorCode = ansi.ColorCode("white+buf")
		cfg.AltColorCodes = []string{"", ansi.ColorCode("white:grey+h")}
	}
	if *conf.VTable {
		tab := table.Table{
			Headers: []string{"Description", "Value"},
			Rows: [][]string{
				{"Version", conf.BVersion},
				{"Application ID", *conf.ID},
				{"Application name", *conf.Name},
				{"SNMP community", *conf.SNMPCommunity},
				{"SNMP agent", fmt.Sprintf("%v", *conf.IsSNMPAgent)},
				{"SNMP trapd", fmt.Sprintf("%v", *conf.IsSNMPTrap)},
			},
		}
		tab.WriteTable(os.Stdout, &cfg)
	}
}
