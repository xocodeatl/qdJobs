package main

import (
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
        k8s.K8sJobs(jobName, jobImage, jobCmd)
      },
    },
    {
      Name: "get",
      Aliases: []string{"g"},
      Usage: "qdJobs get 'JobName' ",
      Action: func(c *cli.Context) {
        jobName  := c.Args().Get(0)
      
        k8s.K8sGetJobs(jobName)
      },
    },
  }
  err := app.Run(os.Args)
  if err != nil {
    log.Println(err)
  }
}


