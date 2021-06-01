# gcp-instance-scheduler

Start and stop scheduler for compute engines and cloud sql on GCP

[![License](https://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://github.com/andypangaribuan/gcp-instance-scheduler/blob/main/LICENSE)



## Environment
- PRINT_LOG : 1/0
- SCHEDULER_DELAY : in second
- SCHEDULER_CONFIG_FILE : must be existing path
```shell
file at ./env/sample.env
```

## Config File
- DAYS : separate by a comma. Sunday:0, Monday:1 ...
- TYPE : VM or SQL
- TIME : 00:00 to 23:59

If TYPE == SQL, then ZONE can be empty
```shell
configure this file on env: SCHEDULER_CONFIG_FILE
```


## Watch the log file  
```shell
On Linux/Mac
$ tail -f /your-directory/action.log  
$ tail -f /your-directory/error.log  
$ lnav /your-directory/action.log /your-directory/error.log

On Windows
using windows power shell
$ type -wait {path to action.log or error.log}
```

## Build the project  
```shell
$ GOOS=windows GOARCH=amd64 go build -o gcp-instance-scheduler.exe
$ GOOS=linux GOARCH=amd64 go build -o gcp-instance-scheduler
$ GOOS=darwin GOARCH=amd64 go build -o gcp-instance-scheduler
```

## Load environment file  
```shell
On windows
download the ./other/load-env.bat file
$ load-env.bat {.env file path}

On linux
$ set -o allexport; source {.env file path}; set +o allexport
```

## Run the gcp-instance-scheduler  
```shell
On windows
$ gcp-instance-scheduler.exe

On linux
$ ./gcp-instance-scheduler
```

## API Endpoint
POST Method
- /private/time
- /private/day
- /private/clear-console
- /private/log-status
- /private/reverse-log-status
