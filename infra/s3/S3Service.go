package s3

import (
	"base-api/config"
	"base-api/constants"
	"bytes"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

type S3Object struct {
	Params      *s3.PutObjectInput
	Key         string
	ContentType string
}

type s3Configuration struct {
	cfg *config.S3Configuration
}

type S3Interface interface {
	S3Upload(r *http.Request, folder, param string) (string, string, error)
	NewS3Object(file *os.File, folder string, filename string) (*S3Object, error)
	S3Url(path string) string
}

func NewS3Configuration(cfg *config.S3Configuration) S3Interface {
	return &s3Configuration{
		cfg: cfg,
	}
}

func (s *s3Configuration) S3Upload(r *http.Request, folder, param string) (string, string, error) {
	file, h, _ := r.FormFile("file")
	defer func() {
		err := file.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	f, err := os.OpenFile(constants.TempDownloadedFileDir+h.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", "", err
	}

	defer func() {
		err := f.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	_, err = io.Copy(f, file)
	if err != nil {
		return "", "", err
	}

	out, err := os.Open(constants.TempDownloadedFileDir + h.Filename)
	if err != nil {
		return "", "", err
	}
	s3Obj, err := s.NewS3Object(out, folder, h.Filename)
	if err != nil {
		return "", "", err
	}
	_, err = s3Obj.Send(s)
	if err != nil {
		return "", "", err
	}
	os.Remove(constants.TempDownloadedFileDir + h.Filename)

	return *s3Obj.Params.Key, *s3Obj.Params.ContentType, nil
}

func (s S3Object) Send(cfg *s3Configuration) (output *s3.PutObjectOutput, err error) {
	aws_access_key_id := cfg.cfg.Key
	aws_secret_access_key := cfg.cfg.Secret

	token := ""
	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)
	_, err = creds.Get()
	if err != nil {
		return nil, err
	}
	awsConfig := aws.NewConfig().WithRegion(cfg.cfg.Region).WithCredentials(creds).WithS3ForcePathStyle(true)
	newSession, err := session.NewSession(awsConfig)
	if err != nil {
		return nil, err
	}
	svc := s3.New(newSession)
	output, err = svc.PutObject(s.Params)
	return output, err
}

func (s *s3Configuration) NewS3Object(file *os.File, folder string, filename string) (*S3Object, error) {
	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)
	_, err := file.Read(buffer)
	if err != nil {
		return &S3Object{}, err
	}
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)
	path := s.cfg.RootFolder + "/" + bson.NewObjectId().Hex() + filepath.Ext(filename)
	if folder != "" {
		path = s.cfg.RootFolder + "/" + folder + "/" + bson.NewObjectId().Hex() + filepath.Ext(filename)
	}
	return &S3Object{
		Params: &s3.PutObjectInput{
			Bucket:        aws.String(s.cfg.Bucket),
			Key:           aws.String(path),
			Body:          fileBytes,
			ContentLength: aws.Int64(size),
			ContentType:   aws.String(fileType),
			ACL:           aws.String("public-read"),
		},
	}, nil
}

func (s *s3Configuration) S3Url(path string) string {
	if len(path) > 0 {
		return s.cfg.PublicUrl + path
	}
	return ""
}
