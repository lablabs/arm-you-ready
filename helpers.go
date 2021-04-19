package main

import (
	"fmt"
	"os"

	"k8s.io/client-go/kubernetes"

	// This is required to auth to cloud providers (i.e. GKE)
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func getKubeClient() kubernetes.Interface {
	kubeConf, err := config.GetConfig()
	if err != nil {
		fmt.Println("Error getting kubeconfig:", err)
		os.Exit(1)
	}
	clientset, err := kubernetes.NewForConfig(kubeConf)
	if err != nil {
		fmt.Println("Error creating kubernetes client:", err)
		os.Exit(1)
	}
	return clientset
}
