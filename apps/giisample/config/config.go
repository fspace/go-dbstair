package config

import (
	"gopkg.in/ini.v1"
	"strings"
)

type Config struct {
	DbUser string // = "root"
	DbPass string // = ""
	DbHost string //= "127.0.0.1"
	DbPort int    //= 3306

	DbDriver      string // mysql
	DataSourceFmt string // TODO  后续提供根据不同的驱动  生成不同的数据源功能 每个驱动类型的数据源格式还不一样

	APPName string `default:"app name"`
	// 原始配置对象 Original Configor
	// Raw *config.Config
	raw         *ini.File // TODO 这里保存原始配置对象是为了以后用 可以考虑改为提供Load|Populate|Configure 方法
	configPaths []string
}

// Configure populate the cfg object from the original configuration
// 以此可以实现 分次提取配置
func (c *Config) Configure(cfg interface{}, section ...string) error {
	var sec string
	if len(section) > 0 {
		sec = strings.Join(section, ".")
		return c.raw.Section(sec).MapTo(cfg)
	}
	return c.raw.MapTo(cfg)
}

var DefaultConfig = Config{
	DbUser: "root",
	DbPass: "",
	DbHost: "127.0.0.1",
	DbPort: 3306,
}

func (c Config) Validate() error {
	//return validation.ValidateStruct(&config,
	//	validation.Field(&config.DSN, validation.Required),
	//	validation.Field(&config.JWTSigningKey, validation.Required),
	//	validation.Field(&config.JWTVerificationKey, validation.Required),
	//)
	return nil
}

// LoadConfig loads configuration from the given list of paths and populates it into the Config variable.
// The configuration file(s) should be named as app.yaml.
func LoadConfig(configPaths ...interface{}) (*Config, error) {
	conf := &Config{}

	*conf = DefaultConfig

	//c := configor.New(nil)
	//err := c.Load(conf, configPaths...)
	//if err != nil {
	//	return nil, err
	//}
	//// 保存起来重复可以重复使用哦
	//conf.configor = c
	//
	var cfg *ini.File
	var err error
	if len(configPaths) > 1 {
		otherConfFiles := configPaths[1:]
		cfg, err = ini.Load(configPaths[0], otherConfFiles...)
		if err != nil {
			return nil, err
		}
	} else {
		cfg, err = ini.Load(configPaths[0])
		if err != nil {
			return nil, err
		}
	}

	err = cfg.MapTo(conf)
	if err != nil {
		return nil, err
	}

	// conf.configPaths = configPaths
	conf.raw = cfg

	return conf, nil
}
