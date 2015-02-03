package mmon

import (
	"encoding/json"
	"log"
	"net/http"
)

var Config config

type config struct {
	Mesos    string
	Interval string
}

type MesosState struct {
	activated  int        `json:"activated_slaves"`
	frameworks []Resource `json:"frameworks"`
	slaves     []Slave    `json:"slaves"`
}

type Cluster struct {
	Resource
	Used Resource
}

type Slave struct {
	Resource
	Id  string
	Pid string
}

type Resource struct {
	cpus int
	mem  int
	disk int
}

func (m *MesosState) Run() {
	resp, err := http.Get(Config.Mesos)
	if err != nil {
		log.Printf("error calling meso master (%v) - %v \n", Config.Mesos, err)
		return
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(m); err == nil {
		log.Printf("%v", *m)
		return
	} else {
		log.Printf("error reading response body (%v) - %v \n", resp, err)
	}
}

func (m *MesosState) Interval() string {
	return Config.Interval
}
