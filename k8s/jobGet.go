package k8s

import (

	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"fmt"
	"context"
	"strings"
)

func GetJobs(name string) {
  	clientset, _ := K8sConfig()
	jobs := clientset.BatchV1().Jobs("")

    job, err := jobs.Create(context.TODO(), jobSpec, metav1.GetOptions{})
	fmt.Println(job)
    if err != nil {
        log.Fatalln("Failed to get K8s job.")
    }
}
