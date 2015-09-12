package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var cmdInit = &Command{
	UsageLine: "init",
	Short:     "initializes the dotf directory",
	Long:      `initializes the dotf directory`,
}

func init() {
	cmdInit.Run = runInit // break init loop
}

func runInit(cmd *Command, args []string) {
	dotfpath := getDotfPath()
	_, err := os.Stat(dotfpath)
	if err != nil {
		if os.IsExist(err) {
			return
		}
		if !os.IsNotExist(err) {
			log.Fatal(err)
		}
	}
	log.Printf("setting %s as DOTFPATH", dotfpath)
	err = os.MkdirAll(dotfpath, 0755)
	if err != nil {
		log.Fatal(err)
	}

	gitCmd := exec.Command("git", "init")
	gitCmd.Dir = dotfpath
	err = gitCmd.Run()
	if err != nil {
		log.Fatalf("could not initialize git repository: %s", err)
	}
}

func getDotfPath() string {
	dotfpath := os.Getenv("DOTFPATH")
	if dotfpath == "" {
		log.Fatal("DOTFPATH not set")
	}
	abs, err := filepath.Abs(dotfpath)
	if err != nil {
		log.Fatal(err)
	}
	return abs
}
