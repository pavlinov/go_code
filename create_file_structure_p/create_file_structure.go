package main

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strconv"
	"sync"
)

func main() {
	var folderSrc string = "folder_src"
	var remove string = ""
	var foldersCnt int
	var filesCnt int
	var help bool

	flag.BoolVar(&help, "help", false, "Usage: ./create_file_structure --foldersCnt=<foldersCnt> --filesCnt=<filesCnt>")
	flag.IntVar(&foldersCnt, "foldersCnt", getIntEnv("FOLDERS_CNT", 2), "<foldersCnt>")
	flag.IntVar(&filesCnt, "filesCnt", getIntEnv("FILES_CNT", 2), "<filesCnt>")
	flag.StringVar(&remove, "remove", "", "<folder_name>")

	args := os.Args[1:]
	if len(args) == 0 {
		flag.Usage()
		fmt.Println("At least one argument is required")
		os.Exit(1)
	}

	flag.Parse()
	if help {
		flag.Usage()
		os.Exit(1)
	}

	if foldersCnt <= 1 || filesCnt <= 1 {
		fmt.Println("Err: values <foldersCnt>, <filesCnt> must be greater than 1")
		flag.PrintDefaults()
		os.Exit(1)
	}

	err := createFileStructure(folderSrc, foldersCnt, filesCnt)
	if err != nil {
		fmt.Printf("Error creating file structure: %+v", err)
		return
	}

	if remove != "" {
		folderPath := fmt.Sprintf("%s/%s", folderSrc, remove)
		err := removeFolder(folderPath)
		if err != nil {
			fmt.Printf("Error in remove: %+v", err)
			return
		}
	}

}

func createFileStructure(folderSrc string, foldersCnt int, filesCnt int) error {
	if foldersCnt == 0 {
		return fmt.Errorf("foldersCnt should be greater than 1, but got %d", foldersCnt)
	}
	if filesCnt == 0 {
		return fmt.Errorf("filesCnt should be greater than 1, but got %d", filesCnt)
	}
	fmt.Printf("Going to create %d folders with %d files\n", foldersCnt, filesCnt)
	var wg sync.WaitGroup
	for i := 1; i <= foldersCnt; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			folderName := fmt.Sprintf("folder_%d", i)

			err := os.MkdirAll(fmt.Sprintf("%s/%s", folderSrc, folderName), 0755)
			if err != nil {
				return
			}

			for j := 1; j <= filesCnt; j++ {
				fileName := fmt.Sprintf("file_%d-%d", i, j)
				file, err := os.Create(fmt.Sprintf("%s/%s/%s", folderSrc, folderName, fileName))
				if err != nil {
					return
				}
				file.Close()
			}
		}(i)
	}
	wg.Wait()

	fmt.Fprintf(os.Stdout, "Created: %d folders and %d files structure created successfully!\n", foldersCnt, filesCnt)
	return nil
}

func getIntEnv(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}

func removeFolder(folderPath string) error {
	info, err := os.Stat(folderPath)
	if err != nil {
		return err
	}

	if info.IsDir() {
		err := os.RemoveAll(folderPath)
		if err != nil {
			fmt.Printf("Error removing %s: %+v", folderPath, err)
			return err
		}
		fmt.Printf("Folder %s removed successfully!", folderPath)
	} else {
		fmt.Printf("Path: "+folderPath, "is not directory")
		return errors.Errorf("Path: %s is not directory", folderPath)
	}
	return nil
}
