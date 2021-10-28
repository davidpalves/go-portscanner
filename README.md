# Port Scanner implemented in Golang

This is only for educational purposes use it on your own risk.

![go-portscanner](https://user-images.githubusercontent.com/14933043/139306108-8a42684b-b0c7-4cba-9437-c15a67ae46a2.gif)

## install all dependencies using
```
go get
```
------------

## Basic Usage
```shell
go run main.go
```

#### Arguments accepted

For getting the all arguments you can use, execute `go run main.go -h`

```
usage: main.go        [-h]
                      [--hostname Host that will be scanned. E.g.: 127.0.0.1 (default: 127.0.0.1)]
                      [--protocol Protocol used on the scan. E.g.: tpc (default "tcp")]
                      [--lowest-port Lowest port used on the scan. E.g.: 0 (default 0)]
                      [--highest-port Highest port used on the scan. E.g.: 65535 (default 65535)]
                      [--concurrent-operations How many operations will occur concurrently. E.g.: 32 (default 32)]

Scan open ports easily and fastly

optional arguments:
  -h, --help            show help message and exit
  --hostname, -hostname Hostname
                        Host that will be scanned
  --protocol, -protocol Protocol
                        Protocol used on the scan
  --lowest-port, -lowest-port Lowest Port
                        Define which port (0-65535) should start being scanned
  --highest-port, -highest-port Highest Port
                        Define which port (0-65535) should finish being
                        scanned
  --concurrent-operations, -concurrent-operations Concurrent Operations
                        Define with how many concurrent processes the scan 
                        will be executed
```
