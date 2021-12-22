package k8s

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/klog"
)

func InitClient() *kubernetes.Clientset {
	config, err := rest.InClusterConfig()
	if err != nil {
		klog.Fatal(err)
	}
	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatal(err)
	}
	return kubeClient
}
