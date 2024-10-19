package uploader

import (
	"os"
	"path/filepath"
)

// Function to check if the config file exists
func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// GetFileList reads the directory specified by dirPath and returns a list of file names excluding directories.
func GetFileList(dirPath string) ([]string, error) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	var fileNames []string
	// Loop through the files and collect their names
	for _, file := range files {
		if !file.IsDir() { // Check if it's not a directory
			fileNames = append(fileNames, file.Name())
		}
	}
	return fileNames, nil
}

// MoveFile moves a file from sourcePath to destinationPath.
// It first ensures that the destination directory exists by calling os.MkdirAll.
// Then it calls os.Rename to move the file.
// If any errors occur during this process, they are returned.
func MoveFile(sourcePath string, destinationPath string) error {
	err := os.MkdirAll(filepath.Dir(destinationPath), os.ModePerm)
	if err != nil {
		return err
	}

	err = os.Rename(sourcePath, destinationPath)
	if err != nil {
		return err
	}

	return nil
}
