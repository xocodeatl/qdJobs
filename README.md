# qdJobs

qdJobs is a cli tool to create quick and dirty jobs in kubernetes without too much hassle. 

## Usage
```
USAGE:
   qdJobs command [command options]

COMMANDS:
   create, c  qdJobs create 'JobName' 'Image' 'command, you can use double quotes for commands that requiere arguments': qdJobs create firstJob ubuntu ls
   help, h    Shows a list of commands or help for one command
```
## Installation

``` 
go build .
mv qdJobs /usr/bin
```