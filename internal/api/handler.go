package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/jhachmer/imgo/internal/config"
	"github.com/jhachmer/imgo/internal/utils"
	"github.com/jhachmer/imgo/pkg/img"
	"github.com/jhachmer/imgo/pkg/transform"
)

func fileHandler(directory string) http.Handler {
	_, err := os.Stat(directory)
	if os.IsNotExist(err) {
		log.Fatalf("Directory '%s' not found.\n", directory)
		return nil
	}
	return http.FileServer(http.Dir(directory))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	info := map[string]string{"application": "imgo", "version": config.VERSION}
	jData, err := json.Marshal(info)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Error: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jData)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Error: err.Error()})
		return
	}
}

type Upload struct {
	Filename string `json:"filename,omitempty"`
}

type UploadResponse struct {
	Filename     string `json:"filename,omitempty"`
	BytesWritten int64  `json:"bytes_written,omitempty"`
	//Filetype     string `json:"filetype,omitempty"`
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Error: err.Error()})
	}
	file, header, err := r.FormFile("image")
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Error: err.Error()})
		return
	}
	if header == nil {
		http.Error(w, "No header", http.StatusInternalServerError)
		return
	}
	filename, nbBytes, err := utils.SaveUpload(file, header)
	if err != nil {
		http.Error(w, "Error while saving uploaded file", http.StatusInternalServerError)
	}
	var ur = UploadResponse{
		Filename:     filename,
		BytesWritten: nbBytes,
	}
	err = utils.WriteJSON(w, http.StatusOK, ur)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Error: err.Error()})
		return
	}
	log.Println("File", ur.Filename, "uploaded.", utils.SizeToString(ur.BytesWritten), "written")
}

type FourierResponse struct {
	MagnitudeFile string `json:"magnitude_file,omitempty"`
	PhaseFile     string `json:"phase_file,omitempty"`
}

func handleFourier(w http.ResponseWriter, r *http.Request) {
	filename := r.PathValue("filename")
	image, err := img.NewImageGray(filepath.Join(config.IMAGELOCATION, filename))
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Error: err.Error()})
		return
	}
	dft := transform.NewDFT(image.Pixels)
	magFN := config.IMAGELOCATION + "dftMag" + utils.CutFileExtension(filename)
	phaseFN := config.IMAGELOCATION + "dftPha" + utils.CutFileExtension(filename)
	err = img.ToPNG(magFN, img.ToImage(dft.Magnitude))
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Error: err.Error()})
		return
	}
	err = img.ToPNG(phaseFN, img.ToImage(dft.Phase))
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Error: err.Error()})
		return
	}
	fr := FourierResponse{
		MagnitudeFile: magFN,
		PhaseFile:     phaseFN,
	}
	err = utils.WriteJSON(w, http.StatusOK, fr)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{Error: err.Error()})
		return
	}
}
