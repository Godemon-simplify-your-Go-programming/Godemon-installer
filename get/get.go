package get

import (
	"log"
	"os"
	"os/exec"
)

func GetGodemonUpdate() {
	cmd := exec.Command("wget", "https://github.com/Godemon-simplify-your-Go-programming/Godemon-update/archive/godemon-update-21-04-LTS.zip")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func GetGodemon() {
	cmd := exec.Command("wget", "https://github.com/Godemon-simplify-your-Go-programming/godemon/archive/21.04-LTS.zip")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
