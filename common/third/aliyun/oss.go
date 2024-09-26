package aliyun

import (
	"errors"
	"github.com/mooncake9527/orange/common/config"
	"mime/multipart"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type AliyunOSS struct{}

func NewBucket() (bucket *oss.Bucket, err error) {
	// 创建OSSClient实例。
	client, err := oss.New(config.Ext.AliOSS.Endpoint, config.Ext.AliOSS.AccessKeyId, config.Ext.AliOSS.AccessKeySecret)
	if err != nil {
		return
	}

	// yourBucketName填写存储空间名称。
	bucket, err = client.Bucket(config.Ext.AliOSS.BucketName)
	return
}

func (*AliyunOSS) UploadFile(file *multipart.FileHeader, name, userId string) (string, string, error) {
	bucket, err := NewBucket()
	if err != nil {
		return "", "", errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
	}

	// 读取本地文件。
	f, openError := file.Open()
	if openError != nil {
		return "", "", errors.New("function file.Open() Failed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	fileTmpPath := config.Ext.AliOSS.BasePath + "/" + userId + "/" + name

	// 上传文件流。
	err = bucket.PutObject(fileTmpPath, f)
	if err != nil {
		return "", "", errors.New("function formUploader.Put() Failed, err:" + err.Error())
	}

	return config.Ext.AliOSS.BucketUrl + "/" + fileTmpPath, fileTmpPath, nil
}
