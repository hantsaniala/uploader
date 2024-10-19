package uploader

import (
	"fmt"
	"log"
)

// Run starts the uploader process.
//
// It checks if the config file exists, parses it,
// checks if the source folder exists, and if the uploaded folder does not exist,
// creates it.
//
// Then it reads the source folder, and for each file, it tries to upload it to
// all the hosts specified in the config file. If all the uploads succeed, it
// moves the file to the uploaded folder.
func Run() {
	// Check if the config file exists
	if ConfigFileExist() {
		// Load the configuration
		conf, _ := LoadConfig()
		log.Println(conf.Destination)

		// Check if the source folder exists
		if CheckSourceFolder() {
			// Create the uploaded folder if it does not exist
			if !FolderExist(fmt.Sprintf("./%s", UPLOADED_FORLDER)) {
				CreateFolder(fmt.Sprintf("./%s", UPLOADED_FORLDER))
			}

			// Get the list of files in the source folder
			fileList, _ := GetFileList(conf.Source)
			for _, f := range fileList {
				fmt.Printf("Processing %s\n", f)
				var failed bool
				// Upload the file to each host in the configuration
				for _, h := range conf.Server {
					fmt.Printf("-- [start] Uploading to %s\n", h.Host)
					if err := Push(conf, h, f); err != nil {
						failed = true
						fmt.Printf("-- [failed] Uploading to %s, err: %s\n", h.Host, err)
						continue
					}
					fmt.Printf("-- [finished] Uploading to %s\n", h.Host)
				}

				// Move the file to the uploaded folder if all uploads succeed
				if !failed {
					if err := MoveFile(
						fmt.Sprintf("%s/%s", conf.Source, f),
						fmt.Sprintf("%s/%s", UPLOADED_FORLDER, f),
					); err != nil {
						fmt.Println("err", err)
					}
				}
			}
		}
	}
}
