package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
)

type ReactServer struct {
	DistPath     string
	BuildCommand string
	Port         string
}

func NewReactServer() *ReactServer {
	return &ReactServer{
		DistPath:     "dist",
		BuildCommand: "npm run build",
		Port:         "3005",
	}
}

func (rs *ReactServer) Build() {
	fmt.Println("Building React app...")
	// You can run the build command using exec.Command if needed
	//this is not needed as code files wont be bundled in
}

func (rs *ReactServer) openBrowser() {
	url := "http://localhost:" + rs.Port
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default: // Linux
		cmd = "xdg-open"
		args = []string{url}
	}

	err := exec.Command(cmd, args...).Start()
	if err != nil {
		fmt.Println("Failed to open browser:", err)
	}
}

func (rs *ReactServer) Serve() {
	fs := http.FileServer(http.Dir(rs.DistPath))
	http.Handle("/", http.StripPrefix("/", fs))

	fmt.Printf("Serving React app on http://localhost:%s\n", rs.Port)
	go rs.openBrowser()
	err := http.ListenAndServe(":"+rs.Port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func main() {
	server := NewReactServer()
	server.Serve()
}
