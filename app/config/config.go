package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

// getFromEnv gets a value from the environment variables. It panics if the variable is not found.
//
// Arguments:
//   - key: The key of the environment variable.
//
// Returns:
//   - The value of the environment variable.
func getFromEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("Environment variable %s not found", key))
	}

	return value
}

// Logger represents the logger for the application.
type Logger struct {
	Logger log.Logger
}

// Info logs an info message.
//
// Arguments:
//   - msg: The message to log.
//
// Returns:
//   - None.
func (l *Logger) Info(msg string) {
	l.Logger.Log("msg", msg)
}

// Debug logs a debug message.
//
// Arguments:
//   - msg: The message to log.
//
// Returns:
//   - None.
func (l *Logger) Debug(msg string) {
	l.Logger.Log("msg", msg)
}

// Error logs an error message.
//
// Arguments:
//   - msg: The message to log.
//
// Returns:
//   - None.
func (l *Logger) Error(msg string) {
	l.Logger.Log("msg", msg)
}

// buildLogger builds the logger for the application.
//
// Arguments:
//   - logLevel: The level of the logger.
//
// Returns:
//   - The logger instance.
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

// Config represents the configuration for the application.
type Config struct {
	serviceName        string
	serviceVersion     string
	serviceDescription string
	servicePort        int

	logLevel string

	allowedOrigins []string

	tviSuperserviciosVerificarExentosHost string

	vpnAccessToken string
	vpnServer      string
}

// GetServiceName gets the name of the service.
//
// Arguments:
//   - None.
//
// Returns:
//   - The name of the service.
func (c *Config) GetServiceName() string {
	return c.serviceName
}

// GetServiceVersion gets the version of the service.
//
// Arguments:
//   - None.
//
// Returns:
//   - The version of the service.
func (c *Config) GetServiceVersion() string {
	return c.serviceVersion
}

// GetServiceDescription gets the description of the service.
//
// Arguments:
//   - None.
//
// Returns:
//   - The description of the service.
func (c *Config) GetServiceDescription() string {
	return c.serviceDescription
}

// GetServiceAddress gets the address of the service.
//
// Arguments:
//   - None.
//
// Returns:
//   - The address of the service.
func (c *Config) GetServiceAddress() string {
	return fmt.Sprintf(":%d", c.servicePort)
}

// GetLogLevel gets the log level of the application.
//
// Arguments:
//   - None.
//
// Returns:
//   - The log level of the application.
func (c *Config) GetLogLevel() string {
	return c.logLevel
}

// GetAllowedOrigins gets the allowed origins of the application.
//
// Arguments:
//   - None.
//
// Returns:
//   - The allowed origins of the application.
func (c *Config) GetAllowedOrigins() []string {
	return c.allowedOrigins
}

// GetTviSuperserviciosVerificarExentosHost gets the host of the TVI Superservicios service.
//
// Arguments:
//   - None.
//
// Returns:
//   - The host of the TVI Superservicios service.
func (c *Config) GetTviSuperserviciosVerificarExentosHost() string {
	return c.tviSuperserviciosVerificarExentosHost
}

// GetVpnAccessToken gets the access token of the VPN.
//
// Arguments:
//   - None.
//
// Returns:
//   - The access token of the VPN.
func (c *Config) GetVpnAccessToken() string {
	return c.vpnAccessToken
}

// GetVpnServer gets the server of the VPN.
//
// Arguments:
//   - None.
//
// Returns:
//   - The server of the VPN.
func (c *Config) GetVpnServer() string {
	return c.vpnServer
}

// buildConfig builds the configuration for the application.
//
// Arguments:
//   - None.
//
// Returns:
//   - The configuration for the application.
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

	vpnAccessToken := getFromEnv("VPN_ACCESS_TOKEN")

	vpnServer := getFromEnv("VPN_SERVER")

	return &Config{
		serviceName:                           serviceName,
		serviceVersion:                        serviceVersion,
		serviceDescription:                    serviceDescription,
		servicePort:                           servicePort,
		logLevel:                              logLevel,
		allowedOrigins:                        allowedOrigins,
		tviSuperserviciosVerificarExentosHost: tviSuperserviciosVerificarExentosHost,
		vpnAccessToken:                        vpnAccessToken,
		vpnServer:                             vpnServer,
	}
}

// Conf represents the configuration for the application. It must be used to get the configuration values across the application.
var Conf *Config = buildConfig()

// Log represents the logger for the application. It must be used to log messages across the application.
var Log *Logger = buildLogger(Conf.GetLogLevel())
