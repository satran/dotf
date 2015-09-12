package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

var cmdAdd = &Command{
	UsageLine: "add <file|folder>",
	Short:     "adds file/folder to the dotf database",
	Long:      `adds file/folder to the dotf database.`,
}

func init() {
	cmdAdd.Run = runAdd // break init loop
}

func runAdd(cmd *Command, args []string) {
	for _, file := range args {
		addFile(file)
	}
}

func addFile(name string) {
	name, err := filepath.Abs(name)
	if err != nil {
		log.Fatal(err)
	}
	dotfpath := getDotfPath()
	_, err = os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("%s does not exist", name)
			return
		}
	}
	home := os.Getenv("HOME")
	newpath := name
	if strings.HasPrefix(name, home) {
		newpath = strings.Replace(name, home, "home", 1)
	}
	newpath = filepath.Join(dotfpath, newpath)
	_, err = os.Stat(newpath)
	if err != nil {
		if os.IsExist(err) {
			log.Printf("%s already being tracked", name)
			return
		}
		if !os.IsNotExist(err) {
			log.Println(err)
			return
		}
	}
	parent := filepath.Dir(newpath)
	err = os.MkdirAll(parent, 0755)
	if err != nil {
		log.Println(err)
		return
	}
	err = os.Rename(name, newpath)
	if err != nil {
		log.Println(err)
		return
	}
	err = os.Symlink(newpath, name)
	if err != nil {
		log.Printf("could not create symlink: %s", err)
		os.Rename(newpath, name)
		return
	}
}
