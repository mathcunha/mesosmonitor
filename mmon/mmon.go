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
	ES       string
}

type MesosState struct {
	Activated  int          `json:"activated_slaves"`
	Frameworks []Frameworks `json:"frameworks"`
	Slaves     []Slave      `json:"slaves"`
	Cluster
}

type Cluster struct {
	Resource
	Used Resource
}

type Frameworks struct {
	Active   bool
	Resource `json:"resources"`
	Used     Resource `json:"used_resources"`
}

type Slave struct {
	Resource `json:"resources"`
	Id       string
	Pid      string
}

type Resource struct {
	Cpus float32 `json:"cpus"`
	Mem  float32 `json:"mem"`
	Disk float32 `json:"disk"`
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
		m.updateCluster()
		log.Printf("%v", *m)
		go postES(m.Cluster)
		return
	} else {
		log.Printf("error reading response body (%v) - %v \n", resp, err)
	}
}

func (m *MesosState) updateCluster() {
	m.Cluster.Resource = Resource{0.0, 0.0, 0.0}
	m.Cluster.Used = Resource{0.0, 0.0, 0.0}

	for _, v := range m.Frameworks {
		if v.Active {
			m.Cluster.Used.Cpus += v.Resource.Cpus
			m.Cluster.Used.Mem += v.Resource.Mem
			m.Cluster.Used.Disk += v.Resource.Disk
		}
	}

	for _, v := range m.Slaves {
		m.Cluster.Resource.Cpus += v.Resource.Cpus
		m.Cluster.Resource.Mem += v.Resource.Mem
		m.Cluster.Resource.Disk += v.Resource.Disk
	}
}

func (m *MesosState) Interval() string {
	return Config.Interval
}

func postES(c Cluster) {
}
