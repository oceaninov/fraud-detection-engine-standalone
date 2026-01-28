package environments

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type ApplicationEnvs struct {
	ApplicationName                string `envconfig:"APPLICATION_NAME"`
	ApplicationVersion             string `envconfig:"APPLICATION_VERSION"`
	ApplicationPort                string `envconfig:"APPLICATION_PORT"`
	ApplicationBO                  string `envconfig:"APPLICATION_BO"`
	ApplicationAzureSSORedirectUri string `envconfig:"APPLICATION_AZURE_SSO_REDIRECT_URI"`
}

type LoggingEnvs struct {
	LogLevel     string `envconfig:"LOG_LEVEL"`
	LogFilePath  string `envconfig:"LOG_FILE_PATH"`
	LogFormatter string `envconfig:"LOG_FORMATTER"`
	LogMaxSize   int    `envconfig:"LOG_MAX_SIZE"`
	LogMaxBackup int    `envconfig:"LOG_MAX_BACKUP"`
	LogMaxAge    int    `envconfig:"LOG_MAX_AGE"`
	LogCompress  bool   `envconfig:"LOG_COMPRESS"`
}

type DatabaseEnvs struct {
	DatabaseName           string `envconfig:"DATABASE_NAME"`
	DatabaseUser           string `envconfig:"DATABASE_USER"`
	DatabasePass           string `envconfig:"DATABASE_PASS"`
	DatabaseHost           string `envconfig:"DATABASE_HOST"`
	DatabasePort           string `envconfig:"DATABASE_PORT"`
	DatabaseSchema         string `envconfig:"DATABASE_SCHEMA"`
	DatabaseMaxIdleConn    int    `envconfig:"DATABASE_MAX_IDLE_CONNECTION"`
	DatabaseMaxOpenConn    int    `envconfig:"DATABASE_MAX_OPEN_CONNECTION"`
	DatabaseConMaxLifetime string `envconfig:"DATABASE_CON_MAX_LIFETIME"`
	DatabaseLowSqlQuery    bool   `envconfig:"DATABASE_LOW_SQL_QUERY"`
	DatabaseDriver         string `envconfig:"DATABASE_DRIVER"`
	DatabasePathMigration  string `envconfig:"DATABASE_PATH_MIGRATION"`
}

type JwtEnvs struct {
	JWTSecret               string `envconfig:"JWT_SECRET"`
	JWTAccessTokenDuration  string `envconfig:"JWT_ACCESS_TOKEN_DURATION"`
	JWTRefreshTokenDuration string `envconfig:"JWT_REFRESH_TOKEN_DURATION"`
}

type SettingsEnv struct {
	SettingsFDSGovDetectionDefaultAct string `envconfig:"SETTINGS_FDS_GOV_DETECTION_DEFAULT_ACT"`
	SettingsFDSBasicAuthUser          string `envconfig:"SETTINGS_FDS_BASIC_AUTH_USER"`
	SettingsFDSBasicAuthPassword      string `envconfig:"SETTINGS_FDS_BASIC_AUTH_PASSWORD"`
}

type OTP struct {
	SMTPHost     string `envconfig:"SMTP_HOST"`
	SMTPPort     string `envconfig:"SMTP_PORT"`
	SMTPAuthPass string `envconfig:"SMTP_AUTH_PASSWORD"`
	SMTPAuth     string `envconfig:"SMTP_AUTH"`
	SMTPSender   string `envconfig:"SMTP_SENDER"`
}

type Externals struct {
	// Proxy Adaptor
	ProxyAdapterHost string `envconfig:"PROXY_ADAPTER_HOST"`
}

type Envs struct {
	SetMode string `envconfig:"SET_MODE"`
	JwtEnvs
	ApplicationEnvs
	LoggingEnvs
	DatabaseEnvs
	SettingsEnv
	OTP
	Externals
}

func NewEnvs() (*Envs, error) {
	var environments = Envs{}
	filename := ".env"

	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		err := envconfig.Process("", &environments)
		if err != nil {
			return nil, err
		}
	}

	err = godotenv.Load(filename)
	if err != nil {
		return nil, err
	}

	err = envconfig.Process("", &environments)
	if err != nil {
		return nil, err
	}

	return &environments, nil
}
