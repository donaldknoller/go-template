package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func InitDefault() {
	viper.SetEnvPrefix(ENV_PREFIX)
	viper.SetEnvKeyReplacer(strings.NewReplacer("_", "-"))
	viper.AutomaticEnv()
	prepareEnv()

	log.Info().Any("a", viper.AllSettings()).Msg("dump")
}
func populateDefaults() {
	viper.Set(SERVER_PORT, "9115")
}

func prepareEnv() {
	populateDefaults()
	underScoreReplacer := strings.NewReplacer("_", "-")
	viper.AutomaticEnv()
	// unfortunate workaround
	for _, envEntry := range os.Environ() {
		envVar := strings.Split(envEntry, "=")
		if strings.HasPrefix(envVar[0], ENV_PREFIX) {
			stripPrefix := fmt.Sprintf("%s_", ENV_PREFIX)
			newKey := strings.TrimPrefix(envVar[0], stripPrefix)
			newKey = underScoreReplacer.Replace(newKey)
			viper.Set(newKey, envVar[1])
		}
	}
}
