package lib

import (
	"bufio"
	"flag"
	"github.com/binganao/breachkit/config"
	"github.com/binganao/breachkit/pkg/logger"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

var AliveHost chan ParseResult

//var PocHost chan string

func Core() {
	banner()
	flag.Parse()
	config.GetRule()

	//fmt.Println(config.PocRules)

	wg := &sync.WaitGroup{}
	lock := &sync.Mutex{}

	AliveHost = make(chan ParseResult, 10000000)

	// show helps
	if Target == "" && Targets == "" {
		help()
		os.Exit(0)
	}

	// Add the header to output
	_, err := os.Stat(OutPut)
	if err != nil {
		var BreachkitHeader string
		BreachkitHeader = "4a1445fc600f53bab725f243d30d433a\r\n"
		f, _ := os.Create(OutPut)
		defer f.Close()
		_, err := f.WriteString(BreachkitHeader)
		if err != nil {
			panic(err)
		}
	}

	// alive host test
	if Targets == "" {
		host := Parse(Target)
		if NoIcmp == false {
			logger.Info("Try to Connect the Target" + " [ " + logger.LightCyan(host.Ip) + " ] ")
			if PingHost(host.Ip, TimeOut) {
				AliveHost <- host
			}
		} else {
			AliveHost <- host
			logger.Warning("The host is set to live" + " [ " + logger.LightCyan(host.Ip) + " ] ")
		}
	} else {
		if NoIcmp == false {
			logger.Info("Try to Connect the Targets")
			for _, host := range GetHosts() {
				wg.Add(1)
				go func(host ParseResult) {
					if PingHost(host.Ip, TimeOut) {
						lock.Lock()
						AliveHost <- host
						lock.Unlock()
					}
					wg.Done()
				}(host)
			}
			wg.Wait()
		} else {
			logger.Warning("All hosts is set to live")
			for _, host := range GetHosts() {
				AliveHost <- host
			}
		}
	}

	close(AliveHost)
	if Targets != "" {
		logger.Info("There are [" + strconv.Itoa(len(AliveHost)) + "/" + strconv.Itoa(len(GetHosts())) + "] hosts alive")
	}

	// vuln poc
	logger.Info("Start vulnerability detection")
	for host := range AliveHost {
		wg.Add(1)
		go func(host ParseResult) {
			veritypoc(host)
			wg.Done()
		}(host)
	}

	wg.Wait()
}

func GetHosts() []ParseResult {
	var Hosts []ParseResult

	_, err := os.Stat(Targets)

	if err != nil {
		logger.Error("Fail to open file " + Targets)
		os.Exit(0)
	}

	targetsf, err := os.OpenFile(Targets, os.O_RDONLY, 0666)

	if err != nil {
		logger.Error("Fail to open file " + Targets)
		os.Exit(0)
	}

	buf := bufio.NewReader(targetsf)

	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			return Hosts
		}
		Hosts = append(Hosts, Parse(strings.Replace(line, "\n", "", -1)))
	}
	return Hosts
}

func PingHost(host string, timeout int) bool {
	var to = strconv.Itoa(timeout)
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("ping", host, "-n", "1", "-w", to)
	case "linux":
		cmd = exec.Command("ping", host, "-c", "1", "-w", to, "-W", to)
	case "darwin":
		cmd = exec.Command("ping", host, "-c", "1", "-W", to)
	}
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}
