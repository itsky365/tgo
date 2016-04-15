package tgo

import (
	"sync"
)

var (
	appConfigMux sync.Mutex

	appConfig *ConfigApp
)

type ConfigApp struct {
	Configs map[string]interface{}
}

func configAppInit() {

	if appConfig == nil || len(appConfig.Configs) == 0 {

		appConfigMux.Lock()
		defer appConfigMux.Unlock()

		if appConfig == nil {
			appConfig = &ConfigApp{}

			defaultConfig := configAppGetDefault()

			configGet("app", appConfig, defaultConfig)
		}
	}
}
func configAppClear() {
	appConfigMux.Lock()
	defer appConfigMux.Unlock()

	appConfig = nil
}

func configAppGetDefault() *ConfigApp {
	return &ConfigApp{map[string]interface{}{"Env": "idc", "UrlUserLogin": "http://user.haiziwang.com/user/CheckLogin"}}
}
func ConfigAppGetString(key string, defaultConfig string) string {

	config := ConfigAppGet(key)

	configStr := config.(string)

	if UtilIsEmpty(configStr) {
		configStr = defaultConfig
	}
	return configStr
}

func ConfigAppGet(key string) interface{} {

	configAppInit()

	config, exists := appConfig.Configs[key]

	if !exists {
		return nil
	}
	return config
}

func ConfigEnvGet() string {
	strEnv := ConfigAppGet("Env")

	return strEnv.(string)
}

func ConfigEnvIsDev() bool {
	env := ConfigEnvGet()

	if env == "dev" {
		return true
	}
	return false
}

func ConfigEnvIsBeta() bool {
	env := ConfigEnvGet()

	if env == "beta" {
		return true
	}
	return false
}
