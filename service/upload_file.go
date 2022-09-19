package service

import (
	"fmt"
	"gin/extend"
	"mime/multipart"
)

// 上传文件例
func UploadFile(header *multipart.FileHeader) {
	oss := &extend.AliyunOss{}
	filePath, yunFileTmpPath, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		panic(uploadErr)
	}

	fmt.Println(filePath)
	fmt.Println(yunFileTmpPath)
}
