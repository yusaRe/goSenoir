package config

type Config struct {
    App struct {
        Name string
        Port string
    }
    Database struct {
        Host     string
        Port     string
        User     string
        Password string
        Name     string
    }
}

var AppConfig *Config

func InitConfig() {
    viper.SerConfigName("config")
    viper.AddConfigPath("yml")
    viper.A
}