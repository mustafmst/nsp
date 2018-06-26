package nsp

// ConfigurationBuilder interface
type ConfigurationBuilder interface {
	ConfigRoutes(l Logger)
}

// Logger interface
type Logger interface {
	LogInfo(string)
	LogDebug(string)
	LogError(string)
}

// AppInterface - main application interface
type AppInterface interface {
	DebugMode(option bool) AppInterface
	UseBuilder(b ConfigurationBuilder) AppInterface
	UseLogger(l Logger) AppInterface
	Run() AppInterface
}
