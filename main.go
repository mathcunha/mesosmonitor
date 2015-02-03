package main

import (
	"flag"
	"github.com/mathcunha/amon/scheduler"
	"github.com/mathcunha/mesosmonitor/mmon"
	"sync"
)

func init() {
	flag.StringVar(&mmon.Config.Mesos, "mesos", "http://127.0.0.1:5050/master/state.json", "Apache Mesos master url")
	flag.StringVar(&mmon.Config.Interval, "interval", "1m", "Interval to update the cluster resource info")
}

func main() {
	flag.Parse()
	m := mmon.MesosState{}
	var wg sync.WaitGroup
	scheduler.ScheduleOne(&m)
	wg.Add(1)
	wg.Wait()
}
