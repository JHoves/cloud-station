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
	bucket, err := client.Bucket(BucketName)
	if err != nil {
		t.Log(err)
	}

	//上传文件到bucket中
	//常见我们文件，需要创建一个文件夹（但是云商ossServer会根据你的key路径结果，自动帮你创建目录）
	//objectKey：上传到bucket里面的对象的名称（包含路径）
	//mydir/test.go，oss server会自动创建一个mydir目录，相当于mkdir -pv
	//把当前这个文件上传到了mydir下
	err = bucket.PutObjectFromFile("mydir/test.go", "oss_test.go")
	if err != nil {
		t.Log(err)
	}
}

// 初始化一个Oss Client 等下其他测试用例可以直接使用
func init() {
	c, err := oss.New(OssEndpoint, AccessKey, AccessSecret)
	if err != nil {
		panic(err)
	}
	client = c
}
