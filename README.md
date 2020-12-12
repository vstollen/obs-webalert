# OBS Webalert
A tool to add a message feed overlay as a browser source in [OBS Studio](https://github.com/obsproject/obs-studio).

## Usage
Start the webserver the directory containing the `static` folder. The page for sending alerts will then be reachable
at `<your-ip-address>:8080/send.html` and the message feed will be on the index page: `<your-ip-address>:8080`.

## Building and running from source
### Prerequisites
- You need to have setup and installed (Go)[https://golang.org/].

## Running from source
To build and run the project, use the `go run` command:
```bash
$ go run main.go
```

## Building a binary
To create an executable, use the `go build` command in the project directory:
```bash
obs-webalert$ go build
```
