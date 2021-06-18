package main

import (
  "fmt"
  "github.com/urfave/cli"
  "os"
  "log"
  "github.com/xocodeatl/qdjobs/k8s"
)

func main() {
  app := cli.NewApp()
  app.Name = "goJobs, a simple wrapper around k8s to deploy and monitor jobs"
  app.Version = "0.0.1"
  app.Author = "Ramon Esparza Rodriguez"
  app.Commands = []cli.Command{
    {
      Name: "create",
      Aliases: []string{"c"},
      Usage: "qdJobs create 'JobName' 'Image' 'command, you can use double quotes for commands that requiere arguments': qdJobs create firstJob ubuntu ls",
      Action: func(c *cli.Context) {
        jobName  := c.Args().Get(0)
        jobImage := c.Args().Get(1)
        jobCmd := c.Args().Get(2)
        fmt.Println("Hello one")
        fmt.Println(jobName, jobImage, jobCmd)
        k8s.K8sJobs(jobName, jobImage, jobCmd)
      },
    },
  }
  err := app.Run(os.Args)
  if err != nil {
    log.Println(err)
  }
}


