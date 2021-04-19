package main

import (
	"context"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// List all deployments in a namespace
func getDeployments(client kubernetes.Interface, namespace string) []v1.Deployment {
	deploymentsClient := client.AppsV1().Deployments(namespace)
	deployments, err := deploymentsClient.List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	return deployments.Items
}

// Returns all images in a deployment spec, including InitContainers
func getAllDeploymentContainerImages(deployment v1.Deployment) []string {
	var images []string

	for _, initContainer := range deployment.Spec.Template.Spec.InitContainers {
		images = append(images, initContainer.Image)
	}

	for _, container := range deployment.Spec.Template.Spec.Containers {
		images = append(images, container.Image)
	}

	return images
}
