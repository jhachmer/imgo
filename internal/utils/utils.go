package utils

import (
	"encoding/json"
	"io"
	"math"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/jhachmer/imgo/internal/config"
)

// WriteJSON encodes values and sends them over ResponseWriter with status code
// Returns error if encoding fails
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return err
	}
	return nil
}

// TODO: user
func SaveUpload(file multipart.File, header *multipart.FileHeader, w http.ResponseWriter) (string, int64) {
	filename := FilenameWithUUID(header.Filename)
	dst, err := os.Create(filepath.Join(config.IMAGELOCATION, filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "", 0
	}
	defer dst.Close()
	nbBytes, err := io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "", 0
	}
	return filename, nbBytes
}

func checkDir(path string) bool {
	folderInfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return folderInfo.IsDir()
}

func createDir(path string) error {
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func SetupDir(path string) error {
	if !checkDir(path) {
		err := createDir(path)
		if err != nil {
			return err
		}
	}
	return nil
}

// SizeToString converts number of bytes to binary prefix (Kibibyte, Mebibyte, ...)
// Returns number of bytes formated as [dddd.dcc] d:digit, c:char
func SizeToString(n int64) string {
	if n == 0 {
		return "0B"
	}
	bytes := float64(n)
	units := []string{"B", "KB", "MB", "GB", "TB", "PB"}
	unitPow := math.Floor(math.Log(bytes) / math.Log(1024))
	return strconv.FormatFloat((bytes/math.Pow(1024, unitPow))*1, 'f', 1, 64) + units[int(unitPow)]
}

func CutFileExtension(filename string) string {
	ext := filepath.Ext(filename)
	return strings.TrimSuffix(filename, ext)
}
