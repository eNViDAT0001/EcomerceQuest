package wrap_viper

import (
	"github.com/spf13/viper"
	"time"
)

// WrapViper ..Wrapper of *viper.Viper
type WrapViper struct {
	*viper.Viper
}

var vp *WrapViper

// GetViper ..get singleton WrapViper instance, default is merge with Environment variable
func GetViper() *WrapViper {
	if vp != nil {
		return vp
	}
	vp = &WrapViper{viper.New()}
	vp.AutomaticEnv()
	vp.setDefaultValue()
	return vp
}

// LoadConfigFile ..Viper auto realize the media type in yaml|toml|json
func (v *WrapViper) LoadConfigFile(path string, filenameWithoutExtension string) {
	v.SetConfigName(filenameWithoutExtension)
	v.AddConfigPath(path)
	err := v.MergeInConfig()
	if err != nil {
		panic(err)
	}
}
func (v *WrapViper) setDefaultValue() {
	v.SetDefault("PORT", DEFAULT_PORT)
	v.SetDefault("VERSION", "0.0.0")
}

const (
	PRODUCTION  string = "production"
	DEVELOPMENT string = "development"
	STAGING     string = "staging"

	DEFAULT_PORT int = 8082

	// DEFAULT_LOG_LEVEL https://github.com/sirupsen/logrus/blob/f8bf7650dccb756cea26edaf9217aab85500fe07/logrus.go#L93
	//		0 PanicLevel
	//		1 FatalLevel
	//		2 ErrorLevel
	//		3 WarnLevel
	//		4 InfoLevel
	//		5 DebugLevel
	//		6 TraceLevel
	DEFAULT_LOG_LEVEL            int = 4
	DEFAULT_LOG_TIMESTAMP_FORMAT     = time.RFC3339
)
