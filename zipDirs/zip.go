package zipDirs

import (
	"log"
	"os"
	"os/exec"
)

func UnzipGodemon() {
	cmd := exec.Command("unzip", "21.04-LTS.zip")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func UnzipGodemonUpdate() {
	cmd := exec.Command("unzip", "godemon-update-21-04-LTS.zip")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func RmZips() {
	cmd := exec.Command("rm", "-r", "21.04-LTS.zip", "godemon-update-21-04-LTS.zip")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
