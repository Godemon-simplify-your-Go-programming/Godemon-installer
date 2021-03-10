package prepareDirs

import (
	"log"
	"os"
	"os/exec"
)

func PrepareDirs() {
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
	err = os.Chdir(localPath)
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

func Move() {
	home := os.Getenv("HOME")
	cmd := exec.Command("mv", "godemon", home+"/.godemon/bin")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	cmd = exec.Command("mv", "godemon-update", home+"/.godemon/bin")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func RmDirs() {
	cmd := exec.Command("rm", "-r", "godemon-21.04-LTS", "Godemon-update-godemon-update-21-04-LTS")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
