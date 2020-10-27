package _oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
)

type _oss struct {
	client *oss.Client
}

//NewOSS
func NewOSS(endpoint, accessKeyId, secret string) (*_oss, error) {
	client, err := oss.New(endpoint, accessKeyId, secret)
	if err != nil {
		return nil, err
	}
	return &_oss{client: client}, err
}

//Upload
func (s *_oss) Upload(objectName, bucketName string, data io.Reader, options ...oss.Option) error {
	bucket, err := s.client.Bucket(bucketName)
	if err != nil {
		return err
	}
	err = bucket.PutObject(objectName, data, options...)
	if err != nil {
		return err
	}
	return nil
}

//SignURL
func (s *_oss) SignURL(objectName, bucketName string, expire int64, options ...oss.Option) (string, error) {
	bucket, err := s.client.Bucket(bucketName)
	if err != nil {
		return "", err
	}
	return bucket.SignURL(objectName, oss.HTTPGet, expire, options...)
}
