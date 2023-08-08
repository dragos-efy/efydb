package config

import (
	"fmt"
	"os"
)

// The directory where the database and files are stored
func RootDir() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		dirname = "."
	}
	return fmt.Sprintf("%s/efydb/", dirname)
}
func MkDirIfNotExists(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, os.ModePerm)
	}
}