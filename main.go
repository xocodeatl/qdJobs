package main

import (
  "github.com/urfave/cli"
  "os"
  "log"
  k8s "github.com/xocodeatl/qdjobs/k8s/create"
  status "github.com/xocodeatl/qdjobs/k8s/status"
  logs "github.com/xocodeatl/qdjobs/k8s/logs"
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
      Name: "status",
      Aliases: []string{"s"},
      Usage: "qdJobs status 'JobName' : qdJobs status firstJob",
      Action: func(c *cli.Context) {
        jobName  := c.Args().Get(0)
        status.GetJobsStatus(jobName)
      },
    },
    {
      Name: "logs",
      Aliases: []string{"l"},
      Usage: "qdJobs  logs 'JobName' : qdJobs logs firstJob",
      Action: func(c *cli.Context) {
        jobName  := c.Args().Get(0)
        logs.LogsJobs(jobName)
      },
    },
    {
      Name: "run",
      Aliases: []string{"r"},
      Usage: "qdJobs  run 'JobName' 'Image' 'command' : firstJob",
      Action: func(c *cli.Context) {
        jobName  := c.Args().Get(0)
        jobImage := c.Args().Get(1)
        jobCmd := c.Args().Get(2)
        RunJob(jobName, jobImage, jobCmd)
      },
    },
  }
  err := app.Run(os.Args)
  if err != nil {
    log.Println(err)
  }
}


