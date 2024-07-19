package main

import (
    "fmt"
    "os/exec"
)

func main() {
    fmt.Println("Starting Docker containers...")

    cmd := exec.Command("docker-compose", "-f", "./docker/docker-compose.yml", "up", "-d")
    cmd.Stdout = nil
    cmd.Stderr = nil
    err := cmd.Run()
    if err != nil {
        fmt.Println("Error starting Docker containers:", err)
        return
    }

    fmt.Println("Docker containers started successfully.")
}
