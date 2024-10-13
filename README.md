# A REST API in GoLang

## Description
This project uses gorilla/mux to run a simple REST API in Go. The entry point to the application is the cmd/server/main.go file. From there is creates a new instance of App which will initialize all the required components (configs, handlers, services, repos, etc.). Finally, it will start the server in a separate go routine and use a channel to wait for a termination signal from the OS. That way, when we kill the server, we can gracefully shutdown the application by cleaning up resources.