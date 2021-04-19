package main

import (
	"flag"
	"fmt"
	"sync"

	dockerClient "github.com/docker/docker/client"
	v1 "k8s.io/api/apps/v1"
)

type T []Output

func main() {
  var outputs []Output

	targetArchitecture := flag.String("arch", "arm64", "Target architecture")
	targetNamespace := flag.String("namespace", "", "Target namespace")
	flag.Parse()

	client := getKubeClient()
	cli, err := dockerClient.NewClientWithOpts(dockerClient.FromEnv, dockerClient.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	deployments := getDeployments(client, *targetNamespace)

	var wg sync.WaitGroup
	queue := make(chan T, 1)
	wg.Add(len(deployments))

	for _, d := range deployments {
		go func (d v1.Deployment) {
			containers := getAllDeploymentContainerImages(d)
			var outputs []Output
			for _, c := range containers {
				architectures, err := getImageArch(cli, c)
				if (err != nil) {
					fmt.Println(err)
					continue
				}
				outputs = append(outputs, Output{d.Name, d.Namespace, "Deployment", c, architectures})
			}
			queue <- T(outputs)
		}(d)
	}

	go func() {
		for t:= range queue {
			outputs = append(outputs, t...)
			wg.Done()
		}
	}()

	wg.Wait()

	renderTable(*targetArchitecture, outputs)
}
