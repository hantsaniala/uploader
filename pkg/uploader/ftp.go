package uploader

import (
	"fmt"
	"os"
	"path/filepath"

	ftp "github.com/moov-io/go-ftp"
)

// Push pushes a file to the FTP server.
//
// It takes in a Config, Server, and file name, and returns an error if any occurs.
//
// The function first creates an FTP client using the given Server.
// Then, it logs in to the FTP server and uploads the file to the destination path.
// Lastly, it closes the FTP connection.
func Push(c Config, s Server, f string) error {
	// Create an FTP client using the server's host and port
	clientConfig := ftp.ClientConfig{
		Hostname: fmt.Sprintf("%s:%d", s.Host, s.Port),
		Username: s.Username,
		Password: s.Password,
	}

	// Create a new FTP client
	client, err := ftp.NewClient(clientConfig)
	if err != nil {
		return err
	}

	// Check if the FTP client is reachable
	if err := client.Ping(); err != nil {
		return err
	}

	// Close the FTP connection when done
	defer client.Close()

	// Open the file to be uploaded
	fileData, err := os.Open(filepath.Join(c.Source, f))
	if err != nil {
		return err
	}

	// Upload the file to the destination path
	if err := client.UploadFile(fmt.Sprintf("%s/%s", c.Destination, f), fileData); err != nil {
		return err
	}
	return nil
}
