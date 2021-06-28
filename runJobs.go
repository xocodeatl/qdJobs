package main

import (
	"time"
	"log"

	create "github.com/xocodeatl/qdjobs/k8s/create"
	status "github.com/xocodeatl/qdjobs/k8s/status"
	logs "github.com/xocodeatl/qdjobs/k8s/logs"
)

func RunJob(name string, image string, cmd string) {
	create.K8sJobs(name, image, cmd)
	deadline := time.Now().Add(10 * time.Second)

	//time.Sleep(8 * time.Second)
	for {
		statusJob, _ := status.GetJobsStatus(name)
		if statusJob == 0 || time.Now().After(deadline) {
			log.Println("Success!! :)")
			logs.LogsJobs(name)
			break
		}
		if statusJob == 2 || time.Now().After(deadline) {
			log.Println("Fail!! :(")
			logs.LogsJobs(name)
			break
		} 
	}	
}
