package uploader

import (
	"fmt"

	"github.com/pterm/pterm"
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
	multi := pterm.DefaultMultiPrinter
	multi.Start()

	// Check if the config file exists
	if ConfigFileExist() {
		// Load the configuration
		conf, _ := LoadConfig()

		// Check if the source folder exists
		if CheckSourceFolder() {
			// Create the uploaded folder if it does not exist
			if !FolderExist(fmt.Sprintf("./%s", UPLOADED_FORLDER)) {
				CreateFolder(fmt.Sprintf("./%s", UPLOADED_FORLDER))
			}

			for _, h := range conf.Server {

				// Get the list of files in the source folder
				fileList, _ := GetFileList(conf.Source)

				sp, _ := pterm.DefaultSpinner.
					WithWriter(multi.NewWriter()).
					Start(fmt.Sprintf("Uploading to '%s'", h.Host))

				// Upload the file to each host in the configuration
				var err error
				var failed bool
				for _, f := range fileList {
					pb, _ := pterm.DefaultProgressbar.WithTotal(1).
						WithWriter(multi.NewWriter()).
						Start(fmt.Sprintf("Uploading '%s'", truncateMiddle(f, 25)))

					if err = Push(conf, h, f); err != nil {
						failed = true
						pb.WithTitleStyle(pterm.NewStyle(pterm.FgRed))
						continue
					}

					// Move the file to the uploaded folder if all uploads succeed
					if err = MoveFile(
						fmt.Sprintf("%s/%s", conf.Source, f),
						fmt.Sprintf("%s/%s", UPLOADED_FORLDER, f),
					); err != nil {
						failed = true
						continue
					}

					pb.Increment()
				}

				if failed {
					sp.Fail(err)
					continue
				}

				sp.Success()
			}

			multi.Stop()
		}
	}
}
