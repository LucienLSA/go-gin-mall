package conf

import (
	"fmt"
	"ginmall/pkg/util/logging"
	"os"
	"time"

	"github.com/spf13/viper"
)

var Config *Conf

type Conf struct {
	System        *System           `yaml:"system"`
	Oss           *Oss              `yaml:"oss"`
	MySql         map[string]*MySql `yaml:"mysql"`
	Email         *Email            `yaml:"email"`
	Redis         *Redis            `yaml:"redis"`
	EncryptSecret *EncryptSecret    `yaml:"encryptSecret"`
	Cache         *Cache            `yaml:"cache"`
	// KafKa         map[string]*KafkaConfig `yaml:"kafKa"`
	// RabbitMq      *RabbitMq               `yaml:"rabbitMq"`
	// Es            *Es                     `yaml:"es"`
	PhotoPath *LocalPhotoPath `yaml:"photoPath"`
}

type System struct {
	RunMode      string        `yaml:"runmode"`
	Domain       string        `yaml:"domain"`
	Version      string        `yaml:"version"`
	HttpPort     string        `yaml:"httpPort"`
	Host         string        `yaml:"host"`
	UploadModel  string        `yaml:"uploadModel"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
}

type Oss struct {
	BucketName      string `yaml:"bucketName"`
	AccessKeyId     string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	Endpoint        string `yaml:"endPoint"`
	EndpointOut     string `yaml:"endpointOut"`
	QiNiuServer     string `yaml:"qiNiuServer"`
}

type MySql struct {
	Dialect  string `yaml:"dialect"`
	DbHost   string `yaml:"dbHost"`
	DbPort   string `yaml:"dbPort"`
	DbName   string `yaml:"dbName"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}

type Email struct {
	ValidEmail string `yaml:"validEmail"`
	SmtpHost   string `yaml:"smtpHost"`
	SmtpEmail  string `yaml:"smtpEmail"`
	SmtpPass   string `yaml:"smtpPass"`
}

type Redis struct {
	RedisHost        string        `yaml:"redisHost"`
	RedisPort        string        `yaml:"redisPort"`
	RedisUsername    string        `yaml:"redisUsername"`
	RedisPassword    string        `yaml:"redisPwd"`
	RedisDbName      int           `yaml:"redisDbName"`
	RedisNetwork     string        `yaml:"redisNetwork"`
	RedisMaxIdle     int           `yaml:"redisMaxIdle"`     // 最大空闲连接数
	RedisMinIdle     int           `yaml:"redisMinIdle"`     // 最小空闲连接数
	RedisIdleTimeout time.Duration `yaml:"redisIdleTimeout"` // 空闲连接超时时间（秒）
	RedisMaxRetries  int           `yaml:"redisMaxRetries"`  // 最大重试次数
}

// EncryptSecret 加密的东西
type EncryptSecret struct {
	JwtSecret   string `yaml:"jwtSecret"`
	EmailSecret string `yaml:"emailSecret"`
	PhoneSecret string `yaml:"phoneSecret"`
	MoneySecret string `yaml:"moneySecret"`
}

type LocalPhotoPath struct {
	PhotoHost   string `yaml:"photoHost"`
	ProductPath string `yaml:"productPath"`
	AvatarPath  string `yaml:"avatarPath"`
}

type Cache struct {
	CacheType    string `yaml:"cacheType"`
	CacheExpires int64  `yaml:"cacheExpires"`
	CacheWarmUp  bool   `yaml:"cacheWarmUp"`
	CacheServer  string `yaml:"cacheServer"`
}

func ReadConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/conf/locales")
	viper.AddConfigPath(workDir)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		logging.LogrusObj.Infoln("读取配置文件失败", err)
		panic(err)
	}
	fmt.Println("读取配置文件成功")

}

// 获取过期时间
func GetExpiresTime() int64 {
	if Config.Cache.CacheExpires == 0 {
		return int64(30 * time.Minute)
	}
	if Config.Cache.CacheExpires == -1 {
		return -1
	}
	return int64(time.Duration(Config.Cache.CacheExpires) * time.Minute)
}
