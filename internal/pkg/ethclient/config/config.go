package config

type Config struct {
	Host            string `mapstructure:"host"`
	SmHolder        string `mapstructure:"smartcontractholder"`
	SmAddr          string `mapstructure:"smartcontractaddress"`
	DefaultFund     int64  `mapstructure:"defaultfund"`
	DefaultGasLimit uint64 `mapstructure:"defaultgaslimit"`
	DefaultGasPrice int64  `mapstructure:"defaultgasprice"`
}
