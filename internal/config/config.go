package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App       AppConfig       `mapstructure:"app"`
	Database  DatabaseConfig  `mapstructure:"database"`
	Redis     RedisConfig     `mapstructure:"redis"`
	JWT       JWTConfig       `mapstructure:"jwt"`
	RSA       RSAConfig       `mapstructure:"rsa"`
	OSS       OSSConfig       `mapstructure:"oss"`
	CORS      CORSConfig      `mapstructure:"cors"`
	RateLimit RateLimitConfig `mapstructure:"ratelimit"`
}

type AppConfig struct {
	Env      string `mapstructure:"env"`
	Port     int    `mapstructure:"port"`
	LogLevel string `mapstructure:"log_level"`
}

type DatabaseConfig struct {
	Host                   string `mapstructure:"host"`
	Port                   int    `mapstructure:"port"`
	User                   string `mapstructure:"user"`
	Password               string `mapstructure:"password"`
	Name                   string `mapstructure:"name"`
	SSLMode                string `mapstructure:"sslmode"`
	TimeZone               string `mapstructure:"timezone"`
	MaxOpenConns           int    `mapstructure:"max_open_conns"`
	MaxIdleConns           int    `mapstructure:"max_idle_conns"`
	ConnMaxLifetimeMinutes int    `mapstructure:"conn_max_lifetime_minutes"`
}

func (d DatabaseConfig) DSN() string {
	parts := []string{
		fmt.Sprintf("host=%s", d.Host),
		fmt.Sprintf("port=%d", d.Port),
		fmt.Sprintf("user=%s", d.User),
	}
	// pgx/lib pq parse the space-separated DSN positionally; an empty
	// `password=` value swallows the next key/value pair, so omit it entirely
	// when no password is configured.
	if d.Password != "" {
		parts = append(parts, fmt.Sprintf("password=%s", d.Password))
	}
	parts = append(parts,
		fmt.Sprintf("dbname=%s", d.Name),
		fmt.Sprintf("sslmode=%s", d.SSLMode),
		fmt.Sprintf("TimeZone=%s", d.TimeZone),
	)
	return strings.Join(parts, " ")
}

func (d DatabaseConfig) ConnMaxLifetime() time.Duration {
	if d.ConnMaxLifetimeMinutes <= 0 {
		return time.Hour
	}
	return time.Duration(d.ConnMaxLifetimeMinutes) * time.Minute
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type JWTConfig struct {
	Secret     string `mapstructure:"secret"`
	TTLSeconds int    `mapstructure:"ttl_seconds"`
	Issuer     string `mapstructure:"issuer"`
}

func (j JWTConfig) TTL() time.Duration {
	if j.TTLSeconds <= 0 {
		return 24 * time.Hour
	}
	return time.Duration(j.TTLSeconds) * time.Second
}

type RSAConfig struct {
	PrivateKeyPath string `mapstructure:"private_key_path"`
	PublicKeyPath  string `mapstructure:"public_key_path"`
}

type OSSConfig struct {
	Endpoint        string `mapstructure:"endpoint"`
	AccessKeyID     string `mapstructure:"access_key_id"`
	AccessKeySecret string `mapstructure:"access_key_secret"`
	Bucket          string `mapstructure:"bucket"`
	PublicURL       string `mapstructure:"public_url"`
}

type CORSConfig struct {
	AllowedOrigins []string `mapstructure:"allowed_origins"`
}

type RateLimitConfig struct {
	LoginPerMinute int `mapstructure:"login_per_minute"`
	APIPerMinute   int `mapstructure:"api_per_minute"`
	GuestPerMinute int `mapstructure:"guest_per_minute"`
}

// Load reads configuration from configs/config.yaml (preferred) and overrides
// with environment variables. Missing config file is not fatal — env vars alone
// can fully drive the app.
func Load(configPath string) (*Config, error) {
	v := viper.New()
	v.SetConfigType("yaml")
	if configPath != "" {
		v.SetConfigFile(configPath)
	} else {
		v.SetConfigName("config")
		v.AddConfigPath("./configs")
		v.AddConfigPath(".")
	}

	v.SetEnvPrefix("")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	bindEnvAliases(v)
	setDefaults(v)

	if err := v.ReadInConfig(); err != nil {
		// Tolerate missing file when env vars supply enough.
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("read config: %w", err)
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}
	return &cfg, nil
}

func bindEnvAliases(v *viper.Viper) {
	pairs := map[string]string{
		"app.env":                          "APP_ENV",
		"app.port":                         "APP_PORT",
		"app.log_level":                    "APP_LOG_LEVEL",
		"database.host":                    "DB_HOST",
		"database.port":                    "DB_PORT",
		"database.user":                    "DB_USER",
		"database.password":                "DB_PASSWORD",
		"database.name":                    "DB_NAME",
		"database.sslmode":                 "DB_SSLMODE",
		"database.timezone":                "DB_TIMEZONE",
		"redis.addr":                       "REDIS_ADDR",
		"redis.password":                   "REDIS_PASSWORD",
		"redis.db":                         "REDIS_DB",
		"jwt.secret":                       "JWT_SECRET",
		"jwt.ttl_seconds":                  "JWT_TTL_SECONDS",
		"jwt.issuer":                       "JWT_ISSUER",
		"rsa.private_key_path":             "RSA_PRIVATE_KEY_PATH",
		"rsa.public_key_path":              "RSA_PUBLIC_KEY_PATH",
		"oss.endpoint":                     "OSS_ENDPOINT",
		"oss.access_key_id":                "OSS_ACCESS_KEY_ID",
		"oss.access_key_secret":            "OSS_ACCESS_KEY_SECRET",
		"oss.bucket":                       "OSS_BUCKET",
		"oss.public_url":                   "OSS_PUBLIC_URL",
	}
	for key, env := range pairs {
		_ = v.BindEnv(key, env)
	}
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("app.env", "local")
	v.SetDefault("app.port", 8080)
	v.SetDefault("app.log_level", "info")
	v.SetDefault("database.port", 5432)
	v.SetDefault("database.sslmode", "disable")
	v.SetDefault("database.timezone", "Asia/Shanghai")
	v.SetDefault("database.max_open_conns", 50)
	v.SetDefault("database.max_idle_conns", 10)
	v.SetDefault("database.conn_max_lifetime_minutes", 60)
	v.SetDefault("redis.db", 0)
	v.SetDefault("jwt.ttl_seconds", 86400)
	v.SetDefault("jwt.issuer", "ai_hs_go")
	v.SetDefault("ratelimit.login_per_minute", 10)
	v.SetDefault("ratelimit.api_per_minute", 600)
	v.SetDefault("ratelimit.guest_per_minute", 120)
	v.SetDefault("cors.allowed_origins", []string{"http://localhost:5173"})
}
