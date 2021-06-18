package k8s

import (
	"log"
	"os"
	"k8s.io/client-go/tools/clientcmd"
	"path/filepath"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var clientset *kubernetes.Clientset

func K8sConfig() (*kubernetes.Clientset, error) {
	//kubeconfig file path for linux
	defaultCfg := filepath.Join(os.Getenv("HOME"), ".kube", "config") 
	
	var config *rest.Config
	
	//Outside cluster access
	config, err := clientcmd.BuildConfigFromFlags("", defaultCfg)
	if err != nil {
	 	log.Fatal(err)
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return clientset, nil 
}
