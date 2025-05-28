package utils

import (
	"strings"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}
