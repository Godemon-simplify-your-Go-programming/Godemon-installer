package building

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func BuildGodemon() {
	err := os.Chdir("./godemon-21.06")
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
	cmd = exec.Command("sudo", "chmod", "777", "godemon")
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

func BuildGodemonUpdate() {
	err := os.Chdir("./Godemon-update-godemon-update-21-04-LTS")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(os.Getwd())
	cmd := exec.Command("g++", "src/godemon_update.cpp", "-o", "godemon-update")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	cmd = exec.Command("sudo", "chmod", "777", "godemon-update")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	cmd = exec.Command("mv", "godemon-update", "../")
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
