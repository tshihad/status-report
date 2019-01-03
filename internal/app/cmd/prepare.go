package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func mustPrepareApp() error {
	// Do initial task which should invoke before
	// App starts rather than using init function.
	// Implement database connection and preparation in
	// this file

	if err := mustReadConfig(); err != nil {
		return errors.Wrap(err, "Failed to read config file")
	}
	return nil
}

func mustReadConfig() error {

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
