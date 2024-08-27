package utils

import (
	"fmt"

	uuid2 "github.com/google/uuid"
)

// FilenameWithUUID returns given filename with a UUID as prefix
// Format: "uuidprefix_filename.extension"
func FilenameWithUUID(filename string) string {
	uuid := uuid2.New()
	return fmt.Sprintf("%s_%s", uuid, filename)
}

// IsValidUUID checks if given string is a valid UUID string
func IsValidUUID(uuid string) bool {
	_, err := uuid2.Parse(uuid)
	return err == nil
}
