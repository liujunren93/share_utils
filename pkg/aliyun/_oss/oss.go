package _oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"net/url"
	"strings"
)

type OSS struct {
	endpoint string
	client   *oss.Client
}

//NewOSS
func NewOSS(endpoint, accessKeyId, secret string, option ...oss.ClientOption) (*OSS, error) {
	client, err := oss.New(endpoint, accessKeyId, secret, option...)
	if err != nil {
		return nil, err
	}
	return &OSS{client: client, endpoint: endpoint}, err
}
func (s *OSS) UploadStream() {

}

//Upload
func (s *OSS) Upload(objectName, bucketName string, data io.Reader, options ...oss.Option) (string, error) {
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
func (s *OSS) MvFile(bucketName, destPATH, srcPath string, srcFiles ...string) ([]string, error) {
	var result []string
	bucket, err := s.client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}
	for _, src := range srcFiles {
		srcFilePath, _ := url.Parse(src)
		newFile := strings.Replace(srcFilePath.Path, srcPath, destPATH, 1)
		result = append(result, srcFilePath.Scheme+"://"+srcFilePath.Host+newFile)
		_, err := bucket.CopyObject(srcFilePath.Path[1:], newFile[1:])
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (s *OSS) DeleteObjects(bucketName string, files ...string) (int, error) {
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
func (s *OSS) SignURL(objectName, bucketName string, expire int64, options ...oss.Option) (string, error) {
	bucket, err := s.client.Bucket(bucketName)
	if err != nil {
		return "", err
	}
	return bucket.SignURL(objectName, oss.HTTPGet, expire, options...)
}
