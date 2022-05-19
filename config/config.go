package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	App      `mapstructure:"app"`
	Auth     `mapstructure:"auth"`
	Log      `mapstructure:"log"`
	DataBase `mapstructure:"database"`
}

type App struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Port      string `mapstructure:"port"`
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
}

type Auth struct {
	JwtExpire time.Duration `mapstructure:"jwt_expire"`
}

type Log struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type DataBase struct {
	Dbtype       string `mapstructure:"dbtype"`
	Hostname     string `mapstructure:"hostname"`
	Port         int    `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Base         string `mapstructure:"base"`
	PingAttempts int    `mapstructure:"ping_attempts"`
}

var Conf *Config = &Config{}

func Init() (err error) {
	//方式1：直接指定配置文件路径（相对路径或绝对路径）
	//相对路径： 相对于可执行文件的相对路径
	//viper.SetConfigFile("./config.yaml") // 指定配置文件路径
	//绝对路径：系统中实际的文件路径
	//viper.SetConfigFile("/usr/local/src/config.yaml") // 指定配置文件路径

	//方式2：只能配置文件名和配置文件的位置，viper自动查找可用的配置文件
	//配置文件名不需要带后缀
	//配置文件位置可以配置多个
	viper.SetConfigName("config") // 配置文件名称(无扩展名)
	viper.AddConfigPath(".")      // 还可以在工作目录中查找配置
	viper.AddConfigPath("./conf") // 还可以在工作目录中查找配置

	//j基本上是配合远程配置中心使用的，告诉viper当前的数据使用什么格式去解析
	//viper.SetConfigType("yaml")   // 如果配置文件的名称中没有扩展名，则需要配置此项

	err = viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {            // 处理读取配置文件的错误
		fmt.Printf("viper.ReadInConfig() failed, err: %v\n", err)
		return
	}

	//把配置反序列化到Conf结构体变量
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal() failed, err: %v\n", err)
		return
	}

	//配置文件热加载
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
		//如果配置文件发生更改 就在反序列化配置到Conf结构体变量
		if err = viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal() failed, err: %v\n", err)
			return
		}
	})
	return

}
