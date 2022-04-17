package config

type Config struct {
	Web      Web        `yaml:"Web"`
	JWT      JWTCfg     `yaml:"JWTCfg"`
	Log      LogCfg     `yaml:"Log"`
	Redis    RedisCfg   `yaml:"Redis"`
	Mongodb  MongodbCfg `yaml:"Mongodb"`
	Tracking string     `yaml:"Tracking"`
}

type Web struct {
	WebHost    string   `yaml:"WebHost"`
	RateLimit  int64    `yaml:"RateLimit"`
	IgnoreUrls []string `yaml:"IgnoreURLs"`
}

type RedisCfg struct {
	RedisHost   string `yaml:"RedisHost"`   //redis服务IP, 逗号分割
	RedisPort   int    `yaml:"RedisPort"`   // 该字段保留
	RedisPasswd string `yaml:"RedisPasswd"` //redis密码
	Cluster     bool   `yaml:"Cluster"`     //	是否启用集群模式
}

type MongodbCfg struct {
	Host      string `yaml:"Host"`
	Port      int    `yaml:"Port"`
	User      string `yaml:"User"`
	Passwd    string `yaml:"Passwd"`
	DbName    string `yaml:"DbName"`
	PoolLimit int    `yaml:"PoolLimit"`
}

type LogCfg struct {
	LogPath    string `yaml:"LogPath"`    //日志存放路径
	LogLevel   string `yaml:"LogLevel"`   //日志记录级别
	MaxSize    int    `yaml:"MaxSize"`    //日志分割的尺寸 MB
	MaxAge     int    `yaml:"MaxAge"`     //分割日志保存的时间 day
	Stacktrace string `yaml:"Stacktrace"` //记录堆栈的级别
	IsStdOut   string `yaml:"IsStdOut"`   //是否标准输出console输出
}

type JWTCfg struct {
	JwtRenewSwitch bool   `yaml:"JwtRenewSwitch"`
	JwtTimeOut     int64  `yaml:"JwtTimeOut"`
	JwtLogLevel    string `yaml:"JwtLogLevel"`
	JwtSecret      string `yaml:"JwtSecret"`
}
