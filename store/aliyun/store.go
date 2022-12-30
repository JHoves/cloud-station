package aliyun

import (
	"fmt"
	"github.com/JHoves/cloud-station/store"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

var (
	//对象是否实现了接口的约束
	_ store.Uploader = &AliOssStore{}
)

// 抽象出去，到时可以随时添加参数而不影响接口的参数
type Options struct {
	Endpoint        string
	AccessKey       string
	AccessSecretKey string
}

func (o *Options) Validate() error {
	//参数检验
	if o.Endpoint == "" || o.AccessKey == "" || o.AccessSecretKey == "" {
		return fmt.Errorf("endpoint,accessKey,accessSecretKey has one empty")
	}
	return nil
}

func NewDefaultAliOssStore() (*AliOssStore, error) {
	return NewAliOssStore(&Options{
		Endpoint:        os.Getenv("ALI_OSS_ENDPOINT"),
		AccessKey:       os.Getenv("ALI_AK"),
		AccessSecretKey: os.Getenv("ALI_SK"),
	})
}

// AliOssStore对象的构造函数
func NewAliOssStore(o *Options) (*AliOssStore, error) {
	//参数检验
	o.Validate()
	client, err := oss.New(o.Endpoint, o.AccessKey, o.AccessSecretKey)
	if err != nil {
		return nil, err
	}
	return &AliOssStore{
		client:   client,
		listener: NewDefaultProgressListener(),
	}, nil
}

// 创建一个对象
type AliOssStore struct {
	client *oss.Client
	//依赖listener的实现
	listener oss.ProgressListener
}

// 实现接口
func (s *AliOssStore) Upload(bucketName string, objectKey string, filename string) error {
	//1、获取bucket对象
	bucket, err := s.client.Bucket(bucketName)
	if err != nil {
		return err
	}

	//2、上传文件到bucket
	if err := bucket.PutObjectFromFile(objectKey, filename, oss.Progress(s.listener)); err != nil {
		return err
	}

	//3、打印下载信息
	downloadURL, err := bucket.SignURL(objectKey, oss.HTTPGet, 60*60*24)
	if err != nil {
		return err
	}
	fmt.Printf("文件下载链接: %s\n", downloadURL)
	fmt.Println("\n注意: 文件下载有效期为1天, 中转站保存时间为3天, 请及时下载")

	return nil
}
