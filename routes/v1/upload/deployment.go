package upload

import (
	"archive/zip"
	"errors"
	"fmt"
	"infra-server/config"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

type FileManager interface {
	PromoteVersion(deploymentID string, environmentID string, deployID string) error
	SaveNewVersion(deploymentID string, environmentID string, file *multipart.FileHeader) (*CodeVersion, error)
}

type v1FileManager struct {
	cfg        config.ServerConfig
	savedFiles map[string]CodeVersion
}

type V1FileMangerDeps struct {
	Config config.ServerConfig
}

func NewV1Manager(deps V1FileMangerDeps) *v1FileManager {
	return &v1FileManager{
		cfg:        deps.Config,
		savedFiles: map[string]CodeVersion{},
	}
}

type BindFile struct {
	Name        string                `form:"name" binding:"required"`
	File        *multipart.FileHeader `form:"file" binding:"required"`
	Environment string                `form:"environment" binding:"required"`
	Service     string                `form:"service" binding:"required"`
}

type CodeVersion struct {
	ID           string     `json:"id"`
	FileName     string     `json:"-"`
	SizeInKB     uint       `json:"sizeKB"`
	DateUploaded *time.Time `json:"dateUploaded"`
}

func (v1 v1FileManager) SaveNewVersion(deploymentID string, environmentID string, file *multipart.FileHeader) (*CodeVersion, error) {
	id := uuid.NewString()
	directory := fmt.Sprintf("%s/%s", v1.cfg.TempDir, deploymentID)
	fullPath := fmt.Sprintf("%s/%s%s", directory, id, filepath.Ext(file.Filename))

	rawFile, err := file.Open()

	defer func() {
		err := rawFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	if err != nil {
		return nil, err
	}

	buff := make([]byte, 512)
	_, err = rawFile.Read(buff)
	if err != nil {
		return nil, err
	}

	filetype := http.DetectContentType(buff)
	if filetype != "application/zip" && filetype != "application/x-gzip" {
		fmt.Println("Bad upload. Got ", filetype)
		return nil, errors.New("detected bad file type")
	}
	_, err = rawFile.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}

	err = os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return nil, err
	}

	f, err := os.Create(fullPath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	_, err = io.Copy(f, rawFile)
	if err != nil {
		return nil, err
	}

	uploaded := time.Now()

	version := &CodeVersion{
		ID:           id,
		FileName:     fullPath,
		SizeInKB:     uint(file.Size) / 1024, // KB
		DateUploaded: &uploaded,
	}
	v1.savedFiles[id] = *version

	return version, nil
}

func (v1 v1FileManager) PromoteVersion(deploymentID string, environmentID string, deployID string) error {
	deploy, exists := v1.savedFiles[deployID]
	if !exists {
		return errors.New("deployment does not exist")
	}
	inputPath := deploy.FileName
	outputPath := fmt.Sprintf("%s/%s/%s", v1.cfg.StaticSiteDirectory, deploymentID, environmentID)
	os.RemoveAll(outputPath)

	archive, err := zip.OpenReader(inputPath)
	if err != nil {
		return err
	}
	defer archive.Close()

	for _, f := range archive.File {
		parts := strings.Split(f.Name, string(os.PathSeparator))
		rootDirLen := len(parts[0])
		// pretty much --strip-components=1
		newFileAppend := f.Name[rootDirLen:]
		filePath := filepath.Join(outputPath, newFileAppend)
		fmt.Println("unzipping file ", filePath)

		if f.FileInfo().IsDir() {
			fmt.Println("creating directory...")
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		fileInArchive, err := f.Open()
		if err != nil {
			return err
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			return err
		}

		dstFile.Close()
		fileInArchive.Close()
	}

	return nil
}
