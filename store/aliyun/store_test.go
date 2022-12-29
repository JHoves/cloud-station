package aliyun_test

import (
	"github.com/JHoves/cloud-station/store"
	"github.com/JHoves/cloud-station/store/aliyun"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	AccessKey    = os.Getenv("ALI_AK")
	AccessSecret = os.Getenv("ALI_SK")
	OssEndpoint  = os.Getenv("ALI_OSS_ENDPOINT")
	BucketName   = os.Getenv("ALI_BUCKET_NAME")
)

var (
	uploader store.Uploader
)

// Aliyun Oss store Upload测试用例
func TestUpload(t *testing.T) {
	//使用assert 编写测试用例的断言语句
	should := assert.New(t)
	err := uploader.Upload(BucketName, "test.txt", "store_test.go")
	if should.NoError(err) {
		//没有error 开启下一个步骤
		t.Log("upload ok")
	}
}

// 通过init 编写uploader 实例化逻辑
func init() {
	ali, err := aliyun.NewDefaultAliOssStore()
	if err != nil {
		panic(err)
	}
	uploader = ali
}
