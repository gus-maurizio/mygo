package main
  
import  (
	"log"
        "time"
        )

type funcPlugin func(string, *time.Ticker)

func pluginMaker(duration time.Duration, pName string, plugin funcPlugin) {
	logrecord := p.Sprintf("PluginMaker with duration %v name: %s and function %#v\n", duration, pName, plugin)
        log.Print(logrecord)
	pRuntime := PluginRuntime{ ticker: time.NewTicker(duration), pluginName: pName}
	PluginSlice = append(PluginSlice, pRuntime)
	go plugin(pRuntime.pluginName, pRuntime.ticker)
}



func baseChannelPlugin(myName string, ticker *time.Ticker) {
	defer ticker.Stop()
        log.Printf("%s started", myName)
	for t := range ticker.C {
		log.Printf("%20s Tick at %v", myName, t)
	}
}


func baseMutexPlugin(myName string, ticker *time.Ticker) {
        defer ticker.Stop()
        log.Printf("%s started", myName)
        for t := range ticker.C {
		log.Printf("%20s Tick at %v", myName, t)
        }
}

