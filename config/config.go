package config

import (
	"flag"
	"gin-vue-template/pkg/utils/gptr"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	Proj = "gin-vue-template"
)

type Config struct {
	CommandFlags CommandFlags
	App          AppConfig        `yaml:"app"`
	HTTPServer   HTTPServerConfig `yaml:"http_server"`
	Mysql        MysqlConfig      `yaml:"mysql"`
	Redis        RedisConfig      `yaml:"redis"`
	// add your config here
}

type CommandFlags struct {
	ConfigPath string `yaml:"config_path"`
}

type AppConfig struct {
	RunMode string `yaml:"run_mode"`
}

type HTTPServerConfig struct {
	Port         int   `yaml:"port"`
	ReadTimeout  int64 `yaml:"read_timeout"`
	WriteTimeout int64 `yaml:"write_timeout"`
}

type MysqlConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

func getConfigPath(cmdConfPath string) string {
	p := cmdConfPath

	// check if p exists
	if _, err := os.Stat(p); err == nil {
		return p
	}

	// try to get from env Proj
	p = os.Getenv(Proj)
	if p != "" {
		return p
	}

	// try to get from relative path ./conf/app.yml of executable file
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Join(filepath.Dir(exePath), "conf/app.yaml")
}

func LoadConfig(cmdConfPath string) *Config {
	configPath := getConfigPath(cmdConfPath)
	log.Printf("configPath: %v", configPath)

	// 读取文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("read config file(%v) error: %v", configPath, err)
	}

	// 解析YAML
	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		log.Fatalf("yaml unmarshal error: %v", err)
	}

	return config
}

func NewConfig() *Config {
	configPath := flag.String("c", "", "config file path")
	flag.Parse()

	cmdFlags := CommandFlags{
		ConfigPath: gptr.Indirect(configPath),
	}

	cfg := LoadConfig(cmdFlags.ConfigPath)
	cfg.CommandFlags = cmdFlags
	return cfg
}
