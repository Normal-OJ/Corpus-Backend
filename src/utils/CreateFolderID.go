package utils

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// CreateFolderID create an unique folder name
func CreateFolderID() string {
	id := uuid.Must(uuid.NewV4(), nil)
	return fmt.Sprintf("%x", id.Bytes())
}
