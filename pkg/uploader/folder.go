package uploader

import (
	"fmt"
	"os"
)

// Function to check if the folder exists
func FolderExist(folderPath string) bool {
	info, err := os.Stat(folderPath)
	return err == nil && info.IsDir()
}

// Function to create the folder if it does not exist
func CreateFolder(folderPath string) error {
	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		return err
	}
	return nil
}

// CheckSourceFolder checks if the source folder exists.
// If it does not exist, it will create one at the specified location.
// If creation fails, it will return false.
// If the folder exists or creation succeeds, it will return true.
func CheckSourceFolder() bool {
	folderPath := "./source"

	if !FolderExist(folderPath) {
		fmt.Printf("`source` folder does not exist. Creating folder at %s...\n", folderPath)

		if err := CreateFolder(folderPath); err != nil {
			fmt.Printf("Error creating folder: %v\n", err)
			return false
		}

		fmt.Println("Folder created successfully.\nPut file for upload inside and re-run again.")
	}
	return true
}
