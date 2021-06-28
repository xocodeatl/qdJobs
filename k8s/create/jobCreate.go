package k8s

import (
    k8s "github.com/xocodeatl/qdjobs/k8s"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"context"
	"strings"
)

func K8sJobs(name string, image string, cmd string) {
  	clientset, _ := k8s.K8sConfig()
	jobs := clientset.BatchV1().Jobs("default")
	var jobbackoff int32 = 0

	jobSpec := &batchv1.Job{
        ObjectMeta: metav1.ObjectMeta{
            Name:      name,
        },
        Spec: batchv1.JobSpec{
            Template: v1.PodTemplateSpec{
                Spec: v1.PodSpec{
                    Containers: []v1.Container{
                        {
                            Name:    name,
                            Image:   image,
                            Command: strings.Split(cmd, " "),
                        },
                    },
                    RestartPolicy: v1.RestartPolicyNever,
                },
            },
            BackoffLimit: &jobbackoff,
        },
    }

    _, err := jobs.Create(context.TODO(), jobSpec, metav1.CreateOptions{})
    if err != nil {
        log.Println(err)
        log.Fatalln("Failed to create K8s job.")
    }

    //print job details
    log.Println("Created K8s job successfully")
    pod_name := k8s.JobDetails(name)
    log.Println("Pod name: ", pod_name)
}

