package oss

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/itachilee/gin/global"
	"github.com/tencentyun/cos-go-sdk-v5"
)

// 设置腾讯cos
func SetupCos() {
	u, _ := url.Parse(global.TencentOssSetting.Url)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  global.TencentOssSetting.SecretId,  // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			SecretKey: global.TencentOssSetting.SecretKey, // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
		},
	})
	global.CosCli = c
}
func GetCosCli() *cos.Client {
	if global.CosCli == nil {
		SetupCos()
	}
	return global.CosCli
}

// 查询储存桶
func QueryBuckets() {
	s, _, err := global.CosCli.Service.Get(context.Background())
	if err != nil {
		panic(err)
	}

	for _, b := range s.Buckets {
		fmt.Printf("%#v\n", b)
	}
}

// 上传对象
func UploadObject() {

	// 对象键（Key）是对象在存储桶中的唯一标识。
	// 例如，在对象的访问域名 `examplebucket-1250000000.cos.COS_REGION.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go
	name := "test/objectPut.go"
	// 1.通过字符串上传对象
	f := strings.NewReader("test")

	_, err := global.CosCli.Object.Put(context.Background(), name, f, nil)
	if err != nil {
		panic(err)
	}
	// 2.通过本地文件上传对象
	_, err = global.CosCli.Object.PutFromFile(context.Background(), name, "../test", nil)
	if err != nil {
		panic(err)
	}
	// 3.通过文件流上传对象
	fd, err := os.Open("./test")
	if err != nil {
		panic(err)
	}
	defer fd.Close()
	_, err = global.CosCli.Object.Put(context.Background(), name, fd, nil)
	if err != nil {
		panic(err)
	}
}

// 查询储存桶对象
// opt := &cos.BucketGetOptions{
//	Prefix:  "test",
//	MaxKeys: 3,
// }
func QueryObjects(opt *cos.BucketGetOptions) {

	v, _, err := global.CosCli.Bucket.Get(context.Background(), opt)
	if err != nil {
		panic(err)
	}

	for _, c := range v.Contents {
		fmt.Printf("%s, %d\n", c.Key, c.Size)
	}
}

func DownloadObject() {
	// 1.通过响应体获取对象
	name := "test/objectPut.go"
	resp, err := global.CosCli.Object.Get(context.Background(), name, nil)
	if err != nil {
		panic(err)
	}
	bs, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("%s\n", string(bs))
	// 2.获取对象到本地文件
	_, err = global.CosCli.Object.GetToFile(context.Background(), name, "exampleobject", nil)
	if err != nil {
		panic(err)
	}
}

// 删除对象
func DeleteObject(name string) {
	// name := "test/objectPut.go"
	_, err := global.CosCli.Object.Delete(context.Background(), name)
	if err != nil {
		panic(err)
	}
}
