package drtm

import (
	"encoding/json"
	"io/ioutil"
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
	if _, err := os.Stat(c.Path + c.Name + "." + c.Type); os.IsNotExist(err) {
		c.Path, err = osext.ExecutableFolder()
		if err != nil {
			c.Path = "."
		}

	} else {
		c.Path, err = filepath.Abs(filepath.Dir(c.Path + c.Name + "." + c.Type))
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(dir)
	}

	//fmt.Println("dir >>>>", dir)
	v := viper.New()
	v.SetConfigType(c.Type)
	v.SetConfigName(c.Name)
	v.AddConfigPath(c.Path)
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

func (c Config) Save() error {
	var dir string
	if _, err := os.Stat(c.Path + c.Name + "." + c.Type); os.IsNotExist(err) {
		c.Path, err = osext.ExecutableFolder()
		if err != nil {
			c.Path = "."
		}

	} else {
		c.Path, err = filepath.Abs(filepath.Dir(c.Path + c.Name + "." + c.Type))
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(dir)
	}

	if _, err := os.Stat(c.Path + c.Name + "." + c.Type); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	file, _ := json.MarshalIndent(c.Config, "", " ")

	return ioutil.WriteFile(c.Path+c.Name+"."+c.Type, file, 0644)
}
