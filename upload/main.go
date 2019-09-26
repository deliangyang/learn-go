package main

import (
	"fmt"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"

)

func main() {
	fmt.Println(upload())
}


func upload() string {
	accessKey := "IknBDJgYrCim-nYKm3Tfqbnw6He_-eiKCR0C3cDp"
	secretKey := "aKuccgfplZ6ZK9A08YPPUvApPrKqy4Idv06YtQTT"

	bucket := "sourcedev"

	putPolicy := storage.PutPolicy{
		Scope: bucket,
		Expires: 7200,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	uploadToken := putPolicy.UploadToken(mac)

	return uploadToken
}
