// Copyright 2018 Gustavo Maurizio
// Permission is hereby granted, free of charge, to any person obtaining a 
// copy of this software and associated documentation files (the "Software"), 
// to deal in the Software without restriction, including without limitation 
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
// OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL
// THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.
//

package main

import (
	"flag"
//        "fmt"
        "golang.org/x/text/language"
        "golang.org/x/text/message"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"syscall"
	"time"
//        "types/stack"
//        "strings"
	"strconv"
)

// This is what gets loaded from the -f .yaml configuration file
type Config struct {
        A string			`yaml:"a"`
	DefaultUnit    string		`yaml:"defaulttimeunit"`
	DefaultTick    int		`yaml:"defaulttimetick"`
	PrometheusPort int		`yaml:"prometheusport"`
	PrometheusHandle string		`yaml:"prometheushandle"`
        Plugins []struct {
		PluginName   string	`yaml:"pluginname"`
		PluginModule string	`yaml:"pluginmodule"`
		PluginUnit   string	`yaml:"plugintimeunit"`
		PluginTick   int	`yaml:"plugintimetick"`
        }
}


var p = message.NewPrinter(language.English)

func cleanup() {
	log.Print("Program Cleanup Started")
}


func main() {
	// get the program name and directory where it is loaded from
	// also create a properly formatted (language aware) printer object
        myName    := filepath.Base(os.Args[0])
	myExecDir := filepath.Dir(os.Args[0])
        p         := message.NewPrinter(language.English)

	//--------------------------------------------------------------------------//
	// good practice to initialize what we want 
        rand.Seed(time.Now().UTC().UnixNano())
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile | log.LUTC)

        yamlPtr   := flag.String("f", "./agent.yaml",  "Agent configuration YAML file")
        debugPtr  := flag.Bool("d", false,  "Agent debug mode - verbose")
        numPtr    := flag.Int("n", 100,  "number of records")
        flag.Parse()

        logrecord := p.Sprintf("%s [from %s] will read config from %s in debug %v and generate %d \n",
                myName, myExecDir, *yamlPtr, *debugPtr, *numPtr)
	log.Print("Program Started")
	log.Print(logrecord)

	//--------------------------------------------------------------------------//
	// read the yaml configuration into the Config structure
	config 	  := Config{}
	yamlFile, err := ioutil.ReadFile(*yamlPtr)
	if err != nil {
		log.Fatalf("config YAML file Get err  #%v ", err)
	}
	err	  = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
        	log.Fatalf("error: %v", err)
        }
	logrecord = p.Sprintf("config: %#v\n", config)
	log.Print(logrecord)

	//--------------------------------------------------------------------------//
	// time to start a prometheus metrics server
        // and export any metrics on the /metrics endpoint.
        http.Handle(config.PrometheusHandle, promhttp.Handler())
	go func() {
        	log.Printf("Beginning to serve on port %d at %s", config.PrometheusPort, config.PrometheusHandle)
        	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(config.PrometheusPort), nil))
	}()

	//--------------------------------------------------------------------------//
	// now get ready to finish if some signals are received
	log.Println("Setting signal handlers")
	csignal   := make(chan os.Signal, 3)
	signal.Notify(csignal, syscall.SIGINT)
	signal.Notify(csignal, syscall.SIGTERM)
	log.Println("Waiting for a signal to end")

	s         := <-csignal

	log.Println("Got signal:", s)	
	cleanup()
	log.Println("Program Ended")
	os.Exit(4)
}
