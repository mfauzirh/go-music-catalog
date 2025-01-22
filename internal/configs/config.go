package configs

import "github.com/spf13/viper"

// config is global variable that holds the loaded configurations.
var config *Config

// option is a struct that holds configuration options such as folders, file name, and file type.
//
// Fields:
//
// - configFolders is a list of directories to search for the configuration file.
//
// - configFile is the name of the configuration file.
//
// - configType is the type of the configuration file (e.g., yaml, json).
type option struct {
	configFolders []string
	configFile    string
	configType    string
}

// Init initializes the configuration using the provided options.
// It reads the configuration file and unmarshals it into the global config object.
// It accepts variadic Option functions to customize the configuration path, file name, and file type.
//
// Example usage:
//
//	err := Init(
//	    WithConfigFolder([]string{"./configs", "./custom_configs"}),
//	    WithConfigFile("app_config"),
//	    WithConfigType("yaml"),
//	)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// Returns an error if the configuration file cannot be read or unmarshalled.
func Init(opts ...Option) error {
	opt := &option{
		configFolders: getDefaultConfigFolder(),
		configFile:    getDefaultConfigFile(),
		configType:    getDefaultConfigType(),
	}

	// Apply options
	for _, optFunc := range opts {
		optFunc(opt)
	}

	// Add config folder to viper, set config file name and type
	for _, configFolder := range opt.configFolders {
		viper.AddConfigPath(configFolder)
	}
	viper.SetConfigName(opt.configFile)
	viper.SetConfigType(opt.configType)

	// Automaticly bind environment variable
	viper.AutomaticEnv()

	config = new(Config)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.Unmarshal(&config)
}

// Option is a function that modifies the option struct. It is used to configure Init.
type Option func(*option)

// getDefaultConfigFolder returns the default folder for configuration files.
func getDefaultConfigFolder() []string {
	return []string{"./configs"}
}

// getDefaultConfigFile returns the default name of the configuration file.
func getDefaultConfigFile() string {
	return "config"
}

// getDefaultConfigType returns the default type of the configuration file (e.g., yaml).
func getDefaultConfigType() string {
	return "yaml"
}

// WithConfigFolder allows setting custom folders to search for configuration files.
// The option can be passed to Init.
func WithConfigFolder(configFolders []string) Option {
	return func(opt *option) {
		opt.configFolders = configFolders
	}
}

// WithConfigFile allows setting a custom configuration file name.
// The option can be passed to Init.
func WithConfigFile(configFile string) Option {
	return func(opt *option) {
		opt.configFile = configFile
	}
}

// WithConfigType allows setting a custom configuration file type (e.g., json, yaml).
// The option can be passed to Init.
func WithConfigType(configType string) Option {
	return func(opt *option) {
		opt.configType = configType
	}
}

// Get returns the global config object. If the config is not initialized, it initializes it.
// This function is used to retrieve the loaded configuration in other parts of the application.
func Get() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}
