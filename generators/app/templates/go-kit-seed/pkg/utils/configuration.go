package utils

// Configuration - interface
type Configuration interface {
	GetTitle() string
	GetOrgName() string
	GetAppName() string
	GetDBDriver() string
	GetDBHost() string
	GetDBPort() string
	GetDBName() string
	GetDBUser() string
	GetDBPass() string
	IsSSLEnabled() bool
	GetSSLCertPath() string
	GetSSLKeyPath() string
	GetDomain() string
	GetPort() string
	GetLogLevel() string
	GetLogEnabled() string
	GetZipkinHost() string
	GetZipkinPort() string
	GetPrometheusMetricsPort() string
}

type configuration struct {
	title                 string
	orgName               string
	appName               string
	dbDriver              string
	dbHost                string
	dbPort                string
	dbName                string
	dbUser                string
	dbPass                string
	sslEnabled            bool
	sslCertPath           string
	sslKeyPath            string
	domain                string
	port                  string
	logLevel              string
	logEnabled            string
	zipkinHost            string
	zipkinPort            string
	prometheusMetricsPort string
}

// NewConfiguration - create a new configuration
func NewConfiguration(
	title string,
	orgName string,
	appName string,
	dbDriver string,
	dbHost string,
	dbPort string,
	dbName string,
	dbUser string,
	dbPass string,
	sslEnabled bool,
	sslCertPath string,
	sslKeyPath string,
	domain string,
	port string,
	logLevel string,
	logEnabled string,
	zipkinHost string,
	zipkinPort string,
	prometheusMetricsPort string,
) Configuration {
	return configuration{
		title,
		orgName,
		appName,
		dbDriver,
		dbHost,
		dbPort,
		dbName,
		dbUser,
		dbPass,
		sslEnabled,
		sslCertPath,
		sslKeyPath,
		domain,
		port,
		logLevel,
		logEnabled,
		zipkinHost,
		zipkinPort,
		prometheusMetricsPort,
	}
}

func (c configuration) GetTitle() string {
	return c.title
}

func (c configuration) GetOrgName() string {
	return c.orgName
}

func (c configuration) GetAppName() string {
	return c.appName
}

func (c configuration) GetDBDriver() string {
	return c.dbDriver
}

func (c configuration) GetDBHost() string {
	return c.dbHost
}

func (c configuration) GetDBPort() string {
	return c.dbPort
}

func (c configuration) GetDBName() string {
	return c.dbName
}

func (c configuration) GetDBUser() string {
	return c.dbUser
}

func (c configuration) GetDBPass() string {
	return c.dbPass
}

func (c configuration) IsSSLEnabled() bool {
	return c.sslEnabled
}

func (c configuration) GetSSLCertPath() string {
	return c.sslCertPath
}

func (c configuration) GetSSLKeyPath() string {
	return c.sslKeyPath
}

func (c configuration) GetDomain() string {
	return c.domain
}

func (c configuration) GetPort() string {
	return c.port
}

func (c configuration) GetLogLevel() string {
	return c.logLevel
}

func (c configuration) GetLogEnabled() string {
	return c.logEnabled
}

func (c configuration) GetZipkinHost() string {
	return c.zipkinHost
}

func (c configuration) GetZipkinPort() string {
	return c.zipkinPort
}

func (c configuration) GetPrometheusMetricsPort() string {
	return c.prometheusMetricsPort
}
