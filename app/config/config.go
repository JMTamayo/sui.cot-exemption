package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

func getFromEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("Environment variable %s not found", key))
	}

	return value
}

type Logger struct {
	Logger log.Logger
}

func (l *Logger) Info(msg string) {
	l.Logger.Log("msg", msg)
}

func (l *Logger) Debug(msg string) {
	l.Logger.Log("msg", msg)
}

func (l *Logger) Error(msg string) {
	l.Logger.Log("msg", msg)
}

func buildLogger(logLevel string) *Logger {
	var logLevelOptions level.Option
	switch strings.ToUpper(logLevel) {
	case "DEBUG":
		logLevelOptions = level.AllowDebug()
	case "INFO":
		logLevelOptions = level.AllowInfo()
	case "ERROR":
		logLevelOptions = level.AllowError()
	default:
		logLevelOptions = level.AllowInfo()
	}

	iow := log.NewSyncWriter(os.Stdout)

	logger := log.NewLogfmtLogger(iow)
	logger = log.With(logger, "timestamp", log.DefaultTimestampUTC)
	logger = level.NewFilter(logger, logLevelOptions)

	return &Logger{
		Logger: logger,
	}
}

type Config struct {
	serviceName        string
	serviceVersion     string
	serviceDescription string
	servicePort        int

	logLevel string

	allowedOrigins []string

	tviSuperserviciosVerificarExentosHost string
}

func (c *Config) GetServiceName() string {
	return c.serviceName
}

func (c *Config) GetServiceVersion() string {
	return c.serviceVersion
}

func (c *Config) GetServiceDescription() string {
	return c.serviceDescription
}

func (c *Config) GetServiceAddress() string {
	return fmt.Sprintf(":%d", c.servicePort)
}

func (c *Config) GetLogLevel() string {
	return c.logLevel
}

func (c *Config) GetAllowedOrigins() []string {
	return c.allowedOrigins
}

func (c *Config) GetTviSuperserviciosVerificarExentosHost() string {
	return c.tviSuperserviciosVerificarExentosHost
}

func buildConfig() *Config {
	serviceName := getFromEnv("SERVICE_NAME")

	serviceVersion := getFromEnv("SERVICE_VERSION")

	serviceDescription := getFromEnv("SERVICE_DESCRIPTION")

	servicePort, err := strconv.Atoi(getFromEnv("SERVICE_PORT"))
	if err != nil {
		panic(fmt.Sprintf("SERVICE_PORT is not a valid number: %s", err))
	}

	logLevel := getFromEnv("LOG_LEVEL")

	allowedOrigins := strings.Split(getFromEnv("ALLOWED_ORIGINS"), ",")

	tviSuperserviciosVerificarExentosHost := getFromEnv("TVI_SUPERSERVICIOS_VERIFICAR_EXENTOS_HOST")

	return &Config{
		serviceName:                           serviceName,
		serviceVersion:                        serviceVersion,
		serviceDescription:                    serviceDescription,
		servicePort:                           servicePort,
		logLevel:                              logLevel,
		allowedOrigins:                        allowedOrigins,
		tviSuperserviciosVerificarExentosHost: tviSuperserviciosVerificarExentosHost,
	}
}

var Conf *Config = buildConfig()

var Log *Logger = buildLogger(Conf.GetLogLevel())
