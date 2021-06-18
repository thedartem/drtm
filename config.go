package drtm

import (
	"log"
	"os"
	"path/filepath"

	"github.com/kardianos/osext"
	"github.com/spf13/viper"
)

type Config struct {
	Config interface{}
	Path   string
	Name   string
	Type   string
}

func (c Config) Read() error {
	var dir string
	if _, err := os.Stat(c.Path + c.Name + "." + c.Type); os.IsNotExist(err) {
		dir, err := osext.ExecutableFolder()
		if err != nil {
			dir = "."
		} else {
			dir = dir + "/data"
		}

	} else {
		dir, err = filepath.Abs(filepath.Dir(c.Path + c.Name + "." + c.Type))
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(dir)
	}

	//fmt.Println("dir >>>>", dir)
	v := viper.New()
	v.SetConfigType(c.Type)
	v.SetConfigName(c.Name)
	v.AddConfigPath(dir + "/")
	//v.AddConfigPath("./data/")
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	err := v.Unmarshal(&c.Config)
	if err != nil {
		return err
	}

	return nil
}
