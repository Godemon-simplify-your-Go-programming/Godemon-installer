package prepareDirs

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func PrepareDirs() {
	localPath, err := os.Getwd()
	home := os.Getenv("HOME")
	cmd := exec.Command("mkdir", home+"/.godemon")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	cmd = exec.Command("mkdir", home+"/.godemon/bin")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	cmd = exec.Command("mkdir", home+"/.godemon/logs")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	cmd = exec.Command("cp", "-r", "./godemon-21.06/CHANGELOGS", home+"/.godemon/")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	cmd = exec.Command("mkdir", home+"/.godemon/.infos/")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	v := []byte("21.06")
	ioutil.WriteFile(home+"/.godemon/.infos/version.txt", v, 0644)

	err = os.Chdir(localPath)
	if err != nil {
		log.Fatal(err)
	}
	cmd = exec.Command("ls")
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
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func RmDirs() {
	cmd := exec.Command("rm", "-r", "./godemon-21.06")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
