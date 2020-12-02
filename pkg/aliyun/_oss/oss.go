package _oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"strings"
)

type _oss struct {
	endpoint string
	client   *oss.Client
}

//NewOSS
func NewOSS(endpoint, accessKeyId, secret string) (*_oss, error) {
	client, err := oss.New(endpoint, accessKeyId, secret)
	if err != nil {
		return nil, err
	}
	return &_oss{client: client, endpoint: endpoint}, err
}

//Upload
func (s *_oss) Upload(objectName, bucketName string, data io.Reader, options ...oss.Option) (string, error) {
	bucket, err := s.client.Bucket(bucketName)
	if err != nil {
		return "", err
	}
	err = bucket.PutObject(objectName, data, options...)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://%s.%s/%s", bucketName, s.endpoint, objectName), nil
}

//MvFile
// destPATH:目标目录
// srcFile：原文件
// keepOld:是否保留原文件
func (s *_oss) MvFile(bucketName,destPATH ,srcPath string, srcFiles ...string) ([]string,error) {
	var result []string
	bucket, err := s.client.Bucket(bucketName)
	if err != nil {
		return nil,err
	}
	for _, src := range srcFiles {
		newFile := strings.Replace(src, srcPath, destPATH,1)
		result=append(result, newFile)
		_, err = bucket.CopyObject(src, newFile)
		if err != nil {
			return nil,err
		}
	}
	return result,nil
}

func (s *_oss) DeleteObjects(bucketName string, files ...string) (int, error) {
	bucket, err := s.client.Bucket(bucketName)
	if err != nil {
		return 0, err
	}
	delRes, err := bucket.DeleteObjects(files)
	if err != nil {
		return 0, err
	}
	return len(delRes.DeletedObjects), nil
}

//SignURL
func (s *_oss) SignURL(objectName, bucketName string, expire int64, options ...oss.Option) (string, error) {
	bucket, err := s.client.Bucket(bucketName)
	if err != nil {
		return "", err
	}
	return bucket.SignURL(objectName, oss.HTTPGet, expire, options...)
}
