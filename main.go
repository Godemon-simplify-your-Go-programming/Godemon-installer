package main

import (
	"log"
	"os"
	"os/exec"
)

func getGodemonUpdate() {
	cmd := exec.Command("wget", "https://github.com/Godemon-simplify-your-Go-programming/Godemon-update/archive/godemon-update-21-04-LTS.zip")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func unzipGodemonUpdate() {
	cmd := exec.Command("unzip", "godemon-update-21-04-LTS.zip")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func getGodemon() {
	cmd := exec.Command("wget", "https://github.com/Godemon-simplify-your-Go-programming/godemon/archive/21.04-LTS.zip")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func unzipGodemon() {
	cmd := exec.Command("unzip", "21.04-LTS.zip")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	getGodemonUpdate()
	unzipGodemonUpdate()
	getGodemon()
	unzipGodemon()
}
