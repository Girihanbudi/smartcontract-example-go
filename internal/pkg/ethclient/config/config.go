package config

type Config struct {
	Host            string `mapstructure:"host"`
	DefaultFund     int64  `mapstructure:"defaultfund"`
	DefaultGasLimit uint64 `mapstructure:"defaultgaslimit"`
	DefaultGasPrice int64  `mapstructure:"defaultgasprice"`
}
