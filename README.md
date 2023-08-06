# Terraformer

An application that allows you to create terraform files from a simple yaml file. It will take the 
yaml file and parse it and create the terraform files for you. Your modules will have to be created
to support this application.  

The application is supposed to be a simple interface for developers to define their infrastructure
without having to know the ins and outs of terraform and should allow you to create a good standard
for your infrastructure. It supports both json and yaml files at this point.  

**Note:** This is a work in progress and is not ready for use. Should be used as a reference for now.

## Run

Running the application is simple. You can run it with the following command:

```shell
task build
./builds/terraformer-darwin-amd64 --file ./examples/project.yaml --output ./examples
```

You can also run with `go run main.go` if you have go installed.

## Why its created?

I created this application to show how you can create a simple interface for developers to create
their infrastructure without having to know the ins and outs of terraform. It also allows you to
create a standard for your infrastructure. You should use this code as a reference to build and 
extend it to your needs. 