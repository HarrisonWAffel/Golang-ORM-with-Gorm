package config

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

//To override default values use env vars with prefix LIFE_POD_LOAD_TEST_APP, i.e:
var (
	v             *viper.Viper
	propertyNames map[string]struct{}
	Dsn           string
)

func init() {
	v = viper.New()
}

func Read() error {
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	v.SetConfigName("service_config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("./config")
	v.AddConfigPath("../config")
	v.AddConfigPath("../../config")
	v.AddConfigPath("../../../config")

	err := v.ReadInConfig()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Error reading build file %s", v.ConfigFileUsed()))
	}

	propertyNames = make(map[string]struct{})
	for _, k := range v.AllKeys() {
		propertyNames[k] = struct{}{}
	}

	if Dsn == "" {
		Dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			GetString("database.host"), GetString("database.username"), GetString("database.password"),
			GetString("database.name"), GetString("database.port"))
	}

	return nil
}

func requireNotMissing(key string) {
	_, ok := propertyNames[key]
	if !ok {
		panic(fmt.Sprintf("configuration file %s is missing property %s", v.ConfigFileUsed(), key))
	}
}

func GetString(key string) string {
	requireNotMissing(key)
	value := v.GetString(key)
	return value
}

func GetSlice(key string) []string {
	requireNotMissing(key)
	return v.GetStringSlice(key)
}

func GetInt(key string) int {
	requireNotMissing(key)
	value := v.GetInt(key)
	return value
}

func GetBool(key string) bool {
	requireNotMissing(key)
	return v.GetBool(key)
}
