package piaconf

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var _appConf *PiaAppConf = nil

type PiaAppConf struct {
	Applications []CatalogValue `yaml:"applications"`
}

type CatalogValue struct {
	Id       string   `yaml:"id"`
	Name     string   `yaml:"name"`
	Language string   `yaml:"prog_lang"`
	Command  string   `yaml:"entry_cmd"`
	AvroIn   []string `yaml:"avro_in"`
	AvroOut  []string `yaml:"avro_out"`
	InitCmd  string   `yaml:"init_cmd"`
	ExecCmd  string   `yaml:"exec_cmd"`
}

func (c *PiaAppConf) Load(path string) error {
	dat, err := ioutil.ReadFile("catalog.yml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(dat, &c)
	return err
}

func GetConfig() *PiaAppConf {
	if _appConf == nil {
		_appConf = new(PiaAppConf)
		_appConf.Load("catalog.yml")
	}
	return _appConf
}

func GetApp(appId string, appConf *CatalogValue) error {
	conf := GetConfig()
	for _, app := range conf.Applications {
		if app.Id == appId {
			*appConf = app
			return nil
		}
	}
	return errors.New("App not found")
}
