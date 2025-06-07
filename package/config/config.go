package config

import (
	"errors"
	"fmt"
	"os"
	"training-frontend/package/log"
	"training-frontend/package/util"

	"github.com/spf13/viper"
)

// Config defines configurations to be exportes
type Config struct {
	WebServer   WebServerConfig
	Database    DatabaseConfig
	Subsystems  []*Subsystem
	Secret      SecretKey
	PrivateKeys []Key `yaml:"privateKeys"`
	PublicKeys  []Key `yaml:"publicKeys"`
}

type Key struct {
	SystemName string `yaml:"systemName"`
	KeyPath    string `yaml:"keyPath"`
}

// ServerConf defines server configurations
type WebServerConfig struct {
	Host    string
	BaseUrl string
	Port    int
}

// DicomServerConfig defines server configurations
type Subsystem struct {
	Name   string
	Secret string
}

// DicomServerConfig defines server configurations
type DicomServerConfig struct {
	Host        string
	Port        int32
	AETitle     string
	StoragePath string
}

// DatabaseConf defines database configuration
type DatabaseConfig struct {
	Name     string
	User     string
	Password string
	Port     int32
}

type SecretKey struct {
	Secret string
}

// New configuration
// inputs filePath should contain path with file name and file extension e.g ./storage/config.yml
func New() (*Config, error) {
	//assumed cofiguration path
	configFile := "config.yml"
	confPath, err := os.Getwd()
	if util.IsError(err) {
		log.Error("error getting a working directory:%v", err)
		return nil, err
	}
	configPath := fmt.Sprintf("%s/%s", confPath, configFile)

	viper.SetConfigFile(configPath) //file name with extension
	viper.AutomaticEnv()            //enable reading environmental variables

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return nil, err
	}

	cfg := Config{}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
		return nil, err
	}
	return &cfg, nil
}

func (c *Config) GetSystemPrivateKey(systemName string) ([]byte, error) {
	for _, privKey := range c.PrivateKeys {
		if privKey.SystemName == systemName && privKey.KeyPath != "" {
			key, err := os.ReadFile(privKey.KeyPath)
			if util.IsError(err) {
				log.Errorf("error reading private key: %s", err)
				return nil, err
			}
			return key, nil
		}
	}
	return nil, errors.New("could not find the private key")
}

func (c *Config) GetSystemPublicKey(systemName string) ([]byte, error) {
	for _, pubKey := range c.PublicKeys {
		if pubKey.SystemName == systemName && pubKey.KeyPath != "" {
			key, err := os.ReadFile(pubKey.KeyPath)
			if util.IsError(err) {
				log.Errorf("error reading public key: %s", err)
				return nil, err
			}
			return key, nil
		}
	}
	return nil, errors.New("could not find the public key")
}

func (c *Config) GetDatabaseConnection() string {
	conn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%d", c.WebServer.Host, c.Database.Name, c.Database.User, c.Database.Password, c.Database.Port)
	return conn
}

func LoggerPath() string {
	path, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting file path %s\n", err)
	}
	return path + "/.logs"
}

// LogoPath returns logo path
func LogoPath() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		log.Error(err)
		return "", err
	}
	path = fmt.Sprintf("%s/webserver/public/images/dit-logo.png", path)
	return path, nil
}

// ReportDir returns .storage/reports path
func ReportDir() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		log.Error(err)
		return "", err
	}
	// path = fmt.Sprintf("%s\\.storage\\reports", path)
	path = fmt.Sprintf("%s/.storage/reports/", path)
	return path, nil
}

// UploadsDir returns .storage/uploads
func UploadsDir() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		log.Error(err)
		return "", err
	}
	path = fmt.Sprintf("%s/.storage/uploads/", path)
	return path, nil
}

// DownloadsDir returns .storage/downloads path
func DownloadsDir() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		log.Error(err)
		return "", err
	}
	path = fmt.Sprintf("%s/.storage/downloads/", path)
	return path, nil
}

func TamisemiDir() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		log.Error(err)
		return "", err
	}
	path = fmt.Sprintf("%s/.storage/tamisemi/", path)
	return path, nil
}

func TemplatePath() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		log.Error(err)
		return "", err
	}
	// path = fmt.Sprintf("%s\\.storage\\templates", path)
	path = fmt.Sprintf("%s/.storage/templates/", path)
	return path, nil
}