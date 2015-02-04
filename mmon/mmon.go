package mmon

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
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
	Idle Resource
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

	m.Cluster.Idle.Cpus = m.Cluster.Resource.Cpus - m.Cluster.Used.Cpus
	m.Cluster.Idle.Mem = m.Cluster.Resource.Mem - m.Cluster.Used.Mem
	m.Cluster.Idle.Disk = m.Cluster.Resource.Disk - m.Cluster.Used.Disk
}

func (m *MesosState) Interval() string {
	return Config.Interval
}

func postES(c Cluster) {
	timestamp := time.Now().Format("2006.01.02")
	var postData []byte
	w := bytes.NewBuffer(postData)
	json.NewEncoder(w).Encode(c)
	url := Config.ES + "/logstash-" + timestamp + "/mmon/"
	if resp, err := http.Post(url, "application/json", w); err == nil {
		resp.Body.Close()
	} else {
		log.Printf("error sending %v to ElasticSearch(%v) - %v \n", c, url, err)
	}
}
