package main

import (
	"Godemon-installer/building"
	"Godemon-installer/get"
	"Godemon-installer/prepareDirs"
	"Godemon-installer/zipDirs"
)

func main() {
	get.GetGodemonUpdate()
	zipDirs.UnzipGodemonUpdate()
	get.GetGodemon()
	zipDirs.UnzipGodemon()
	zipDirs.RmZips()
	building.BuildGodemon()
	building.BuildGodemonUpdate()
	prepareDirs.PrepareDirs()
	prepareDirs.Move()
	prepareDirs.RmDirs()
}
