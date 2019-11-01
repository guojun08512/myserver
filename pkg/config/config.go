package config

import (
	"bytes"
	"fmt"
	"html/template"
	"myserver/pkg/logger"
	"myserver/pkg/utils"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

// Paths is the list of directories used to search for a
// configuration file
var Paths = []string{
	".",
	"$HOME/Documents/keyayun.com/myserver",
	"$HOME/.myserver",
}

// Filename is the default configuration filename
// search for
const Filename = "my-config"

var config *Config

type db struct {
	UserName       string
	PassWord       string
	DBName         string
	Host           string
	DBIdleconnsMax int
	DBOpenconnsMax int
}

//Config 配置信息
type Config struct {
	DB      *db
	Version string
	Port    int
}

var log = logger.WithNamespace("config")

// FindConfigFile search in the Paths directories for the file with the given
// name. It returns an error if it cannot find it or if an error occurs while
// searching.
func FindConfigFile(name string) (string, error) {
	for _, cp := range Paths {
		filename := filepath.Join(utils.AbsPath(cp), name)
		ok, err := utils.FileExists(filename)
		if err != nil {
			return "", err
		}
		if ok {
			return filename, nil
		}
	}
	return "", fmt.Errorf("Could not find config file %q", name)
}

// UseViper default configure
func UseViper(v *viper.Viper) error {
	config = &Config{
		Port:    v.GetInt("port"),
		Version: v.GetString("version"),
		DB: &db{
			UserName:       v.GetString("db.username"),
			PassWord:       v.GetString("db.password"),
			DBName:         v.GetString("db.dbname"),
			Host:           v.GetString("db.host"),
			DBIdleconnsMax: v.GetInt("db.dbIdleconns_max"),
			DBOpenconnsMax: v.GetInt("db.dbOpenconns_max"),
		},
	}
	return nil
}

func envMap() map[string]string {
	env := make(map[string]string)
	for _, i := range os.Environ() {
		sep := strings.Index(i, "=")
		env[i[0:sep]] = i[sep+1:]
	}
	return env
}

// Setup set configure
func Setup(cfgFile string) (err error) {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix("my")
	viper.AutomaticEnv()

	if cfgFile == "" {
		for _, ext := range viper.SupportedExts {
			var file string
			file, err = FindConfigFile(Filename + "." + ext)
			if file != "" && err == nil {
				cfgFile = file
				break
			}
		}
	}

	if cfgFile == "" {
		return UseViper(viper.GetViper())
	}

	log.Debugf("Using config file: %s", cfgFile)
	tmpl := template.New(filepath.Base(cfgFile))
	tmpl = tmpl.Option("missingkey=zero")
	tmpl, err = tmpl.Funcs(numericFuncsMap).ParseFiles(cfgFile)
	if err != nil {
		return fmt.Errorf("Unable to open and parse configuration file "+
			"template %s: %s", cfgFile, err)
	}

	dest := new(bytes.Buffer)
	ctxt := &struct {
		Env    map[string]string
		NumCPU int
	}{
		Env:    envMap(),
		NumCPU: runtime.NumCPU(),
	}
	err = tmpl.ExecuteTemplate(dest, filepath.Base(cfgFile), ctxt)
	if err != nil {
		return fmt.Errorf("Template error for config file %s: %s", cfgFile, err)
	}

	if ext := filepath.Ext(cfgFile); len(ext) > 0 {
		viper.SetConfigType(ext[1:])
	}
	if err := viper.ReadConfig(dest); err != nil {
		if _, isParseErr := err.(viper.ConfigParseError); isParseErr {
			log.Errorf("Failed to read configurations from %s", cfgFile)
			log.Error(dest.String())
			return err
		}
	}

	return UseViper(viper.GetViper())
}

func GetConfig() *Config {
	return config
}
