package fileupload

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gitlab.id.vin/gami/ps2-gami-common/configs"
	"io"
	"time"

	"gitlab.id.vin/gami/ps2-gami-common/models"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"google.golang.org/api/option"
)

type googleStorageAdapter struct {
	client            *storage.Client
	storageURL        string
	folderName        string
	bucketName        string
	googleCredentials models.GoogleCredentials
}

// NewGoogleStorageAdapter returns a new instance of Adapter
func NewGoogleStorageAdapter(storageURL, folderName string, credential []byte, bucketName string) (Adapter, error) {
	client, err := storage.NewClient(context.Background(), option.WithCredentialsJSON(credential))
	if err != nil {
		return nil, err
	}

	googleCre := models.GoogleCredentials{}
	err = json.Unmarshal(credential, &googleCre)
	if err != nil {
		return nil, err
	}

	a := googleStorageAdapter{
		client:            client,
		storageURL:        storageURL,
		folderName:        folderName,
		bucketName:        bucketName,
		googleCredentials: googleCre,
	}
	return &a, nil
}

func (a *googleStorageAdapter) Upload(uploadFile io.Reader, uploadFileType, folderName string) (string, error) {
	return a.UploadWithName(uploadFile, uploadFileType, folderName, uuid.New().String())
}
func (a *googleStorageAdapter) UploadWithName(uploadFile io.Reader, uploadFileType, folderName string, fileName string) (string, error) {
	ctx := context.Background()
	if len(fileName) == 0 {
		fileName = uuid.New().String()
	}
	outputFileName := fmt.Sprintf("%v/%v.%v",
		folderName,
		fileName,
		uploadFileType)

	if a.client == nil {
		return "", errors.New("key not found")
	}

	w := a.client.Bucket(a.bucketName).Object(outputFileName).NewWriter(ctx)

	w.ACL = []storage.ACLRule{{
		Entity: storage.AllAuthenticatedUsers, Role: storage.RoleReader}}

	w.ContentType = "[*/*]"

	// Entries are immutable, be aggressive about caching (1 day).
	//w.CacheControl = "public, max-age=86400"
	w.CacheControl = fmt.Sprintf("public, max-age=%v", configs.AppConfig.UploadedFileExpirationTime)

	if _, err := io.Copy(w, uploadFile); err != nil {
		return "", err
	}
	if err := w.Close(); err != nil {
		return "", err
	}
	url := fmt.Sprintf("%v/%v/%v",
		a.storageURL,
		a.bucketName,
		outputFileName)

	return url, nil
}

func (a *googleStorageAdapter) GetURL(url string) (string, error) {
	expires := time.Now().Add(time.Minute * 60)
	r, err := storage.SignedURL(a.bucketName, url, &storage.SignedURLOptions{
		GoogleAccessID: a.googleCredentials.ClientEmail,
		PrivateKey:     []byte(a.googleCredentials.PrivateKey),
		Method:         "GET",
		Expires:        expires,
	})
	return r, err
}
