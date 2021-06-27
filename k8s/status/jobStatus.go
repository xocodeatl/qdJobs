package k8s

import (
    k8s "github.com/xocodeatl/qdjobs/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"context"
)

func GetJobsStatus(name string) (int, error) {
    clientset, _ := k8s.K8sConfig()
    jobs := clientset.BatchV1().Jobs("default")
    job, err := jobs.Get(context.TODO(), name, metav1.GetOptions{})
    
    if job.Status.Active > 0 {
        log.Println("Job is still running")
        return 1, nil
    } else {
        if job.Status.Succeeded >0 {
            log.Println("Job was succesful")
            return 0, nil
        }
        log.Println("Job failed")
        return 2, nil
    }   
        
    if err != nil {
        log.Println("Error: ", err)
        return 2, err
    }    
    return 0, nil
}