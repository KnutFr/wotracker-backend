## Fly.io
From the Fly.io docs :

### Build, Deploy And Run A Go Application

Getting an application running on Fly is essentially working out how to package it as a deployable image. Once packaged it can be deployed to the Fly infrastructure to run on the global application platform.

In this guide we'll learn how to deploy a Go application on Fly.

#### The Example Application

You can get the code for the example from the GitHub repository. Just git clone https://github.com/fly-apps/go-example to get a local copy.

The go-example application is, as you'd expect for an example, small. It's a Go application that uses the http server and templates from the standard library. Here's all the code from main.go:

```go
package main

import (
"embed"
"html/template"
"log"
"net/http"
"os"
)

//go:embed templates/*
var resources embed.FS

var t = template.Must(template.ParseFS(resources, "templates/*"))

func main() {
port := os.Getenv("PORT")
if port == "" {
port = "8080"
}

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := map[string]string{
            "Region": os.Getenv("FLY_REGION"),
        }

        t.ExecuteTemplate(w, "index.html.tmpl", data)
    })

    log.Println("listening on", port)
    lo

    g.Fatal(http.ListenAndServe(":"+port, nil))
}
```
The main function starts a server that responds with a html page showing the fly region that served the request. The template lives in ./templates/ and is embedded into the binary using the embed package in go 1.16+.

The template itself, index.html.tmpl, is very simple too:

```html
<!DOCTYPE html>
<html lang="en">
<head>
</head>
<body>
<h1>Hello from Fly</h1>
{{ if .Region }}
<h2>I'm running in the {{.Region}} region</h2>
{{end}}
</body>
</html>
```
Building the Application

As with most Go applications, a simple go build will create a hellofly binary which we can run. It'll default to using port 8080 and you can view it on localhost:8080 with your browser. So, the raw application works. Now to package it up for Fly.

Install Flyctl and Login

We are ready to start working with Fly and that means we need flyctl, our CLI app for managing apps on Fly. If you've already installed it, carry on. If not, hop over to our installation guide. Once thats installed you'll want to log in to Fly.

Launch the App on Fly

To launch an app on fly, run flyctl launch in the directory with your source code. This will create and configure a fly app for you by inspecting your source code, then prompt you to deploy.

```shell

flyctl launch

Scanning source code
Detected Go app
Using the following build configuration
Builder: paketobuildpacks/builder:base
Buildpacks: gcr.io/paketo-buildpacks/go
? Select organization: Demo (demo)
? Select region: ord (Chicago, Illinois (US))
Created app hellofly in organization personal
Wrote config file fly.toml
Your app is ready. Deploy with `flyctl deploy`
? Would you like to deploy now? Yes
```
Deploying hellofly
...
First, this command scans your source code to determine how to build a deployment image as well as identify any other configuration your app needs, such as secrets and exposed ports.

After your source code is scanned and the results are printed, you'll be prompted for an organization. Organizations are a way of sharing application and resources between Fly users. Every fly account has a personal organization, called personal, which is only visible to your account. Let's select that for this guide.

Next, you'll be prompted to select a region to deploy in. The closest region to you is selected by default. You can use this or change to another region.

At this point, flyctl creates an app for you and writes your configuration to a fly.toml file. You'll then be prompted to build and deploy your app. Once complete, your app will be running on fly.

Inside
fly.toml

The fly.toml file now contains a default configuration for deploying your app. In the process of creating that file, flyctl has also created a Fly-side application slot of the same name, "hellofly". If we look at the fly.toml file we can see the name in there:

```yaml

app = "hellofly"

[build]
builder = "paketobuildpacks/builder:base"
buildpacks = ["gcr.io/paketo-buildpacks/go"]

[[services]]
internal_port = 8080
protocol = "tcp"

[services.concurrency]
hard_limit = 25
soft_limit = 20

[[services.ports]]
handlers = ["http"]
port = "80"

[[services.ports]]
handlers = ["tls", "http"]
port = "443"

[[services.tcp_checks]]
interval = 10000
timeout = 2000


```
The flyctl command will always refer to this file in the current directory if it exists, specifically for the app name value at the start. That name will be used to identify the application to the Fly service. The rest of the file contains settings to be applied to the application when it deploys.

You can see in the [build] section the builder image and the list of buildpacks that will be used to build the app for deployment. You can also add build arguments to configure the build process using the [build.args] section.

For example, the following will instruct the Go buildpack to include files in the resources directory in the final image:

```yaml


[build.args]
BP_KEEP_FILES = "assets/*:resources/*"
```
See the Paketo Go Buildpack documentation for more options.

Deploying to Fly

To deploy your app, just run:


flyctl deploy
This will lookup our fly.toml file, and get the app name hellofly from there. Then flyctl will start the process of deploying our application to the Fly platform. Flyctl will return you to the command line when it's done.

Viewing the Deployed App

Now the application has been deployed, let's find out more about its deployment. The command flyctl status will give you all the essential details.

```shell


flyctl status

App
Name     = hellofly
Owner    = demo
Version  = 0
Status   = running
Hostname = hellofly.fly.dev

Allocations
ID       VERSION REGION DESIRED STATUS  HEALTH CHECKS      RESTARTS CREATED
0ac9ed79 0       fra    run     running 1 total, 1 passing 0        44s ago
$
```
As you can see, the application has been with a DNS hostname of hellofly.fly.dev, and an instance is running in Frankfurt. Your deployment's name will, of course, be different.

Connecting to the App

The quickest way to connect to your deployed app is with the flyctl open command. This will open a browser on the HTTP version of the site. That will automatically be upgraded to an HTTPS secured connection (for the fly.dev domain).

to connect to it securely. Add /name to flyctl open and it'll be appended to the URL as the path and you'll get an extra greeting from the hellofly application.

Bonus Points

If you want to know what IP addresses the app is using, try flyctl ips list:

```shell


flyctl ips list

TYPE ADDRESS                              CREATED AT
v4   50.31.246.73                         23m42s ago
v6   2a09:8280:1:3949:7ac8:fe55:d8ad:6b6f 23m42s ago
Arrived at Destination
```
You have successfully built, deployed, and connected to your first Go application on Fly.