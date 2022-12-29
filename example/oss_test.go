package example_test

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"testing"
)

var (
	//全局client实例，在包加载的时候初始化
	client *oss.Client
)

var (
	AccessKey    = os.Getenv("ALI_AK")
	AccessSecret = os.Getenv("ALI_SK")
	OssEndpoint  = os.Getenv("ALI_OSS_ENDPOINT")
	BucketName   = os.Getenv("ALI_BUCKET_NAME")
)

// 测试阿里云OssSDK
func TestBucketList(t *testing.T) {
	lsRes, err := client.ListBuckets()
	if err != nil {
		t.Log(err)
	}

	for _, bucket := range lsRes.Buckets {
		fmt.Println("Buckets:", bucket.Name)
	}
}

// 测试上传文件的功能是否正常
func TestUploadFile(t *testing.T) {
	bucket, err := client.Bucket("my-bucket")
	if err != nil {
		t.Log(err)
	}

	err = bucket.PutObjectFromFile("my-object", "LocalFile")
	if err != nil {
		t.Log(err)
	}
}

// 初始话一个Oss Client 等下其他测试用例可以直接使用
func init() {
	c, err := oss.New(OssEndpoint, AccessKey, AccessSecret)
	//c, err := oss.New("oss-cn-shenzhen.aliyuncs.com", "LTAI5tKHCfxqqM8jCaJBHNEv", "q8wmnzK5jhXux9yhbMJHNRShWg0MNu")
	if err != nil {
		panic(err)
	}
	client = c
}
