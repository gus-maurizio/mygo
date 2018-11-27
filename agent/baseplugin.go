package main
  
import  (
	"log"
	"runtime"
        "time"
        )

type fMeasure func() string
type fPlugin  func(string, *time.Ticker, fMeasure)


func baseMeasure()  string {
        caller := "not available"
        whoami := "not available"

	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil { caller = details.Name() }
	
	me, _, _, ok := runtime.Caller(0)
	mydetails := runtime.FuncForPC(me)
	if ok && mydetails != nil { whoami = mydetails.Name() }
        return(p.Sprintf("sample %30s called by %30s at %f", whoami, caller, float64(time.Now().UnixNano())/1e9))
}


func pluginMaker(duration time.Duration, pName string, plugin fPlugin, measure fMeasure) {
	logrecord := p.Sprintf("PluginMaker with duration %v name: %s and function %#v with function_measure %#v\n", duration, pName, plugin, measure)
        log.Print(logrecord)
	pRuntime := PluginRuntime{ ticker: time.NewTicker(duration), pluginName: pName}
	PluginSlice = append(PluginSlice, pRuntime)
	go plugin(pRuntime.pluginName, pRuntime.ticker, measure)
}


func baseChannelPlugin(myName string, ticker *time.Ticker, measure fMeasure) {
	// make sure we Stop at end
	defer ticker.Stop()
        log.Printf("%s started", myName)
	for t := range ticker.C {
		myMeasure := measure()
		log.Printf("%20s Tick at %v measure: [%v]\n", myName, t, myMeasure)
	}
}


func baseMutexPlugin(myName string, ticker *time.Ticker, measure fMeasure) {
        defer ticker.Stop()
        log.Printf("%s started", myName)
        for t := range ticker.C {
		myMeasure := measure()
		log.Printf("%20s Tick at %v measure: [%v]\n", myName, t, myMeasure)
        }
}

