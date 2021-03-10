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

func rmZips() {
	cmd := exec.Command("rm", "-r", "21.04-LTS.zip", "godemon-update-21-04-LTS.zip")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func buildGodemon() {
	err := os.Chdir("./godemon-21.04-LTS")
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command("go", "build")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	cmd = exec.Command("mv", "godemon", "../")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	err = os.Chdir("../")
	if err != nil {
		log.Fatal(err)
	}
}

func prepareDirs() {
	localPath, err := os.Getwd()
	err = os.Chdir(os.Getenv("HOME"))
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command("mkdir", ".godemon")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	cmd = exec.Command("mkdir", ".godemon/bin")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	cmd = exec.Command("mkdir", ".godemon/logs")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	cmd = exec.Command("cd", "/")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	cmd = exec.Command("ls")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	err = os.Chdir(localPath)
	if err != nil {
		log.Fatal(err)
	}
}

func move() {
	home := os.Getenv("HOME")
	cmd := exec.Command("mv", "godemon", home+"/.godemon/bin")
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
	rmZips()
	buildGodemon()
	prepareDirs()
	move()
}
