package cli

import (
	"fmt"
	"github.com/JHoves/cloud-station/store"
	"github.com/JHoves/cloud-station/store/aliyun"
	"github.com/spf13/cobra"
)

var (
	ossProvider     string
	ossEndpoint     string
	accessKey       string
	accessSecretKey string
	bucketName      string
	uploadFile      string
)

var UploadCmd = &cobra.Command{
	Use:     "upload",
	Long:    "upload 文件上传",
	Short:   "upload 文件上传",
	Example: "upload -f filename",
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			uploader store.Uploader
			err      error
		)
		switch ossProvider {
		case "aliyun":
			uploader, err = aliyun.NewAliOssStore(&aliyun.Options{
				Endpoint:        ossEndpoint,
				AccessKey:       accessKey,
				AccessSecretKey: accessSecretKey,
			})
		case "tx":
			//uploader = tx.NewTxOssStore()
		case "aws":
			//uploader = tx.NewAwsOssStore()
		default:
			return fmt.Errorf("not support oss store provider")
		}
		if err != nil {
			return err
		}

		//使用uploader上传文件
		return uploader.Upload(bucketName, uploadFile, uploadFile)
	},
}

func init() {
	f := UploadCmd.PersistentFlags()
	f.StringVarP(&ossProvider, "provider", "p", "aliyun", "oss storage provider [aliyun/tx/aws]")
	f.StringVarP(&ossEndpoint, "endpoint", "e", "oss-cn-shenzhen.aliyuncs.com", "oss storage endpoint")
	f.StringVarP(&accessKey, "access_key", "k", "", "oss storage provider ak")
	f.StringVarP(&accessSecretKey, "accessSecret_key", "s", "", "oss storage provider sk")
	f.StringVarP(&bucketName, "bucket_name", "b", "jh-devcloud-station", "oss storage provider bucketName")
	f.StringVarP(&uploadFile, "upload_file", "f", "", "upload file name")
	RootCmd.AddCommand(UploadCmd)
}
