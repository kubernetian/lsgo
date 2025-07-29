package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/nexidian/gocliselect"
)

const (
	Blue  = "\033[34m"
	Green = "\033[32m"
	Reset = "\033[0m"
)

func main() {

	menuTitle := "Choose a folder to list the content"
	userHomeDir, _ := os.UserHomeDir()
	foldernames := os.Args[1:]

	if len(foldernames) == 0 {
		fmt.Println("Please provide at least one folder name as an argument.")
		return
	}

	var allPaths []string

	for _, folder := range foldernames {
		paths, err := findFolders(userHomeDir+"/Documents", folder)
		if err != nil {
			fmt.Printf("Error finding folder '%s': %v\n", folder, err)
			continue
		}
		allPaths = append(allPaths, paths...)
	}

	if len(allPaths) == 0 {
		fmt.Println("No folders found matching your search.")
		return
	}

	const pageSize = 10
	page := 0

	for {
		menu := gocliselect.NewMenu(menuTitle)
		start := page * pageSize
		end := min(start+pageSize, len(allPaths))
		for _, path := range allPaths[start:end] {
			menu.AddItem(path, path)
		}
		if end < len(allPaths) {
			menu.AddItem("-- Next page --", "__NEXT__")
		}
		if page > 0 {
			menu.AddItem("-- Previous page --", "__PREV__")
		}

		choice := menu.Display()

		if choice == "__NEXT__" {
			page++
			continue
		}
		if choice == "__PREV__" {
			page--
			continue
		}

		fmt.Printf("%sContents of '%s':%s\n", Green, choice, Reset)
		files, err := os.ReadDir(choice)
		if err != nil {
			fmt.Printf("Error reading directory '%s': %v\n", choice, err)
		}
		for _, file := range files {
			if file.IsDir() {
				fmt.Printf(" - %s%s%s\n", Blue, file.Name(), Reset)
			} else {
				fmt.Printf(" - %s\n", file.Name())
			}
		}
		break
	}
}

func findFolders(dirname, foldername string) ([]string, error) {
	var foundPaths []string
	err := filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && info.Name() == foldername {
			foundPaths = append(foundPaths, path)
		}
		return nil
	})
	return foundPaths, err
}
