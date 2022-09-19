package config

import (
	"gin/common/env"
	"github.com/jinzhu/configor"
	"go/build"
	"os"
)

var ENV string
var CONFIG_PATH string

var Config = struct {
	APPName   string `default:"test" yaml:"appname"`
	APPSecret string `default:"test"`

	DB struct {
		Host     string
		Database string
		User     string `default:"root"`
		Password string `required:"true" env:""`
		Port     string `default:"3306"`
		Prefix   string `default:""`
	}

	REDIS struct {
		Host     string
		User     string `default:"root"`
		Password string
		Port     string `default:"6379"`
	}

	//钉钉推送
	DINGding struct {
		Errorpush             string
		Xueqiucommentpush     string
		Xueqiuselectpush      string
		Xueqiucombinationpush string
		Ciyunpush             string
	}

	//企业微信推送
	Weixin struct {
		Errorpush             string
		Xueqiucommentpush     string
		Xueqiuselectpush      string
		Xueqiucombinationpush string
		Ciyunpush             string
	}

	Queue struct {
		Isopenmultithreading int `default:0`
	}

	Domain struct {
		Domain string
	}

	//阿里云OSS配置
	Aliyunoss struct {
		Endpoint        string `yaml:"endpoint"`
		AccessKeyId     string `yaml:"accessKeyId"`
		AccessKeySecret string `yaml:"accessKeySecret"`
		BucketName      string `yaml:"bucketName"`
		BucketUrl       string `yaml:"bucketUrl"`
		BasePath        string `yaml:"basePath"`
	}

	//JWT配置
	Jwt struct {
		SigningKey  string `yaml:"signingkey"`  // jwt签名
		ExpiresTime string `yaml:"expirestime"` // 过期时间
		BufferTime  string `yaml:"buffertime"`  // 缓冲时间
		Issuer      string `yaml:"issuer"`      // 签发者
	}
}{}

func init() {
	env := env.Active().Value()
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	path := gopath + "/src/gin-hong-api/config"
	configor.New(&configor.Config{Debug: true}).Load(&Config, path+"/settings."+env+".yml")
}

func GetProjectTruePath() string {
	return build.Default.GOPATH + "/bin"
}
