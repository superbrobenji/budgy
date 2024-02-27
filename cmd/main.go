package main

import (
	"net/http"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
	"github.com/superbrobenji/budgy/infrastructure/cdk"
)

func main() {

	defer jsii.Close()

	app := awscdk.NewApp(nil)

	cdk.NewGoCdkStack(app, "GoCdkStack", &cdk.GoCdkStackProps{
		StackProps: awscdk.StackProps{
			Env: cdk.Env(),
		},
	})

	app.Synth(nil)
	// Create a new request multiplexer
	// Take incoming requests and dispatch them to the matching handlers
	mux := http.NewServeMux()

	// Register the routes and handlers
	mux.Handle("/", &homeHandler{})

	// Run the server
	http.ListenAndServe(":8080", mux)
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}
