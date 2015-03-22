package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/fsouza/go-dockerclient"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "docker-clone"
	app.Usage = "Create a similar container as specified"

	app.Action = func(c *cli.Context) {
		if len(c.Args()) < 1 {
			cli.ShowAppHelp(c)
		} else {
			containerName := c.Args()[0]
			fmt.Println("Recreating", containerName)

			// connect to Docker daemon
			endpoint := "unix:///var/run/docker.sock"
			client, _ := docker.NewClient(endpoint)

			// inspect container
			container, error := client.InspectContainer(containerName)
			if error != nil {
				panic(error)
			}

			serializedBinding, _ := json.Marshal(container.HostConfig.PortBindings)
			fmt.Println("Binding", string(serializedBinding))

			// Actually inspect every one of those bindings correctly
			// todo check for allocated ports to avoid conflict
			container.HostConfig.PortBindings["8080/tcp"][0] = docker.PortBinding{HostPort: "8081"}

			cloneContainerOptions := docker.CreateContainerOptions{
				HostConfig: container.HostConfig,
				// todo check for names to avoid conflict
				Name:   container.Name + "_copy",
				Config: container.Config,
			}

			newContainer, error := client.CreateContainer(cloneContainerOptions)
			if error != nil {
				panic(error)
			}

			error = client.StartContainer(newContainer.ID, newContainer.HostConfig)

			if error != nil {
				panic(error)
			}

			fmt.Println("Container created")
		}
	}

	app.Run(os.Args)
}
