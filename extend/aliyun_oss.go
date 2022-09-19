package extend

import (
	"errors"
	"fmt"
	"gin/config"
	"mime/multipart"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type AliyunOss struct{}

func (*AliyunOss) UploadFile(file *multipart.FileHeader) (string, string, error) {
	bucket, err := NewBucket()
	if err != nil {
		return "", "", errors.New("function Aliyunoss.NewBucket() Failed, err:" + err.Error())
	}

	// 读取本地文件。
	f, openError := file.Open()
	if openError != nil {
		return "", "", errors.New("function file.Open() Failed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭
	// 上传阿里云路径 文件名格式 自己可以改 建议保证唯一性
	// yunFileTmpPath := filepath.Join("uploads", time.Now().Format("2006-01-02")) + "/" + file.Filename
	yunFileTmpPath := config.Config.Aliyunoss.BasePath + "/" + "uploads" + "/" + time.Now().Format("2006-01-02") + "/" + file.Filename

	// 上传文件流。
	fmt.Println(yunFileTmpPath)
	err = bucket.PutObject(yunFileTmpPath, f)
	if err != nil {
		return "", "", errors.New("function formUploader.Put() Failed, err:" + err.Error())
	}

	return config.Config.Aliyunoss.BucketUrl + "/" + yunFileTmpPath, yunFileTmpPath, nil
}

func (*AliyunOss) DeleteFile(key string) error {
	bucket, err := NewBucket()
	if err != nil {
		return errors.New("function Aliyunoss.NewBucket() Failed, err:" + err.Error())
	}

	// 删除单个文件。objectName表示删除OSS文件时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// 如需删除文件夹，请将objectName设置为对应的文件夹名称。如果文件夹非空，则需要将文件夹下的所有object删除后才能删除该文件夹。
	err = bucket.DeleteObject(key)
	if err != nil {
		return errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
	}

	return nil
}

func NewBucket() (*oss.Bucket, error) {
	// 创建OSSClient实例。
	client, err := oss.New(config.Config.Aliyunoss.Endpoint, config.Config.Aliyunoss.AccessKeyId, config.Config.Aliyunoss.AccessKeySecret)
	if err != nil {
		return nil, err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(config.Config.Aliyunoss.BucketName)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}
