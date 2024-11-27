package conf

import (
	"ginmall/util/log"
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
	AppEnv      string `yaml:"appEnv"`
	Domain      string `yaml:"domain"`
	Version     string `yaml:"version"`
	HttpPort    string `yaml:"httpPort"`
	Host        string `yaml:"host"`
	UploadModel string `yaml:"uploadModel"`
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
	RedisHost     string `yaml:"redisHost"`
	RedisPort     string `yaml:"redisPort"`
	RedisUsername string `yaml:"redisUsername"`
	RedisPassword string `yaml:"redisPwd"`
	RedisDbName   int    `yaml:"redisDbName"`
	RedisNetwork  string `yaml:"redisNetwork"`
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

// Init 初始化配置项
func InitLodding() {
	// 从本地读取环境变量
	// godotenv.Load()
	// 读取配置文件
	ReadConfig()
	// 初始化Mysql

	// 初始化日志
	log.InitLog()
	// 设置日志级别
	// util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		log.LogrusObj.Panic("翻译文件加载失败", err)
		// util.Log().Panic("翻译文件加载失败", err)
	}

	// 连接数据库
	// model.Database(os.Getenv("MYSQL_DSN"))
	// cache.Redis()
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
		panic(err)
	}
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