# qdJobs

qdJobs is a cli tool to create quick and dirty jobs in kubernetes without too much hassle. 
This is a debugging and troubleshotting tool, not for use in production environments.
## Usage
```
USAGE:
   qdJobs command [command options]

COMMANDS:
   run, r  qdJobs create 'JobName' 'Image' 'command' : qdJobs create firstJob ubuntu "ls -ltr"
   create, c  qdJobs create 'JobName' 'Image' 'command' : qdJobs create firstJob ubuntu "ls- ltr"
   status, s  qdJobs status 'JobName' : qdJobs status firstJob
   logs, l    qdJobs  logs 'JobName' : qdJobs logs firstJob
   help, h    Shows a list of commands or help for one command
```
## Installation

``` 
go build .
mv qdJobs /usr/bin
```
