package config

type Config struct {
	Dir string `env:"MIGRATE_DIR"`

	Port       string `env:"DB_PORT"`
	Host       string `env:"DB_HOST"`
	User       string `env:"DB_USER"`
	Name       string `env:"DB_NAME"`
	SSLMode    string `env:"DB_SSLMODE"`
	Password   string `env:"DB_PASSWORD"`
	DriverName string `env:"DB_DRIVER"`

	SecretKey       string `env:"JWT_SECRET"`
	ValiditySeconds int    `env:"JWT_EXPIRATION"`
	Header          string `env:"JWT_HEADER"`

	HTTPPort string `env:"HTTP_PORT"`
}

//
//type DB struct {
//	Port       string `env:"DB_PORT"`
//	Host       string `env:"DB_HOST"`
//	User       string `env:"DB_USER"`
//	Name       string `env:"DB_NAME"`
//	SSLMode    string `env:"DB_SSLMODE"`
//	Password   string `env:"DB_PASSWORD"`
//	DriverName string `env:"DB_DRIVER"`
//}
//
////
////type Kafka struct {
////	Port string `evn:"KAFKA_PORT"`
////}
//
//type Migrate struct {
//	Dir string `env:"MIGRATE_DIR"`
//}
//
//type JWT struct {
//	SecretKey       string `env:"JWT_SECRET"`
//	ValiditySeconds int    `env:"JWT_EXPIRATION"`
//	Header          string `env:"JWT_HEADER"`
//}
