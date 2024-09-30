package config

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

var Conf *viper.Viper

func Viper_Init(path string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigFile(path)
	err := conf.ReadInConfig()
	if err != nil {
		log.Fatalf("read conf err   " + err.Error())
	}
	conf.WatchConfig()
	conf.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("丸啦，谁把我配置改了。。。")
		err = conf.ReadInConfig()
		if err != nil {
			log.Fatalf("read conf err 2    " + err.Error())
		}
	})
	return conf
}

func Config_Init() {
	var env_config = flag.String("conf", "config/app.yml", "")
	flag.Parse()
	conf := Viper_Init(*env_config)
	Conf = conf
}
