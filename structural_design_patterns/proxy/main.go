package main

// Proxy is a structural design pattern that lets you provide a substitute or placeholder for another object. A proxy controls access to the original object, allowing you to perform something either before or after the request gets through to the original object.
// Examples: Access Restriction, Caching, PreProcessing & PostProcessing (logging, publish event), Lazy initialization

import (
	"fmt"
	"lld/structural_design_patterns/proxy/server"
)

const (
	appStatusURL  = "/app/status"
	createUserURL = "/create/user"

	getRequestType  = "GET"
	postRequestType = "POST"
)

func main() {
	nginxServer := server.NewNginxServer()

	httpCode, body := nginxServer.HandleRequest(appStatusURL, getRequestType)
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(appStatusURL, getRequestType)
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(appStatusURL, getRequestType)
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(createUserURL, postRequestType)
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createUserURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(appStatusURL, getRequestType)
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(createUserURL, postRequestType)
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createUserURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(createUserURL, postRequestType)
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createUserURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(appStatusURL, getRequestType)
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.HandleRequest(createUserURL, postRequestType)
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createUserURL, httpCode, body)
}
