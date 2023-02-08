package configs

import (
	"fmt"
)

// MySQL represents configuration of MySQL database.
type MySQL struct {
	Username          string `default:"vin_id" envconfig:"MYSQL_USER"`
	Password          string `default:"vin_id" envconfig:"MYSQL_PASS"`
	Host              string `default:"127.0.0.1" envconfig:"MYSQL_HOST"`
	Port              int    `default:"3306" envconfig:"MYSQL_PORT"`
	Database          string `default:"gamezone" envconfig:"MYSQL_DB"`
	MaxOpenConnection int    `default:"10" envconfig:"MYSQL_MAX_OPEN"`
	MaxIdleConnection int    `default:"10" envconfig:"MYSQL_MAX_IDLE"`
	MaxLifeTime       int    `default:"24" envconfig:"MYSQL_MAX_LIFETIME"`
}

// ConnectionString returns connection string of MySQL database.
func (c *MySQL) ConnectionString() string {
	format := "%v:%v@tcp(%v:%v)/%v?parseTime=true&charset=utf8"
	return fmt.Sprintf(format, c.Username, c.Password, c.Host, c.Port, c.Database) + "&loc=Asia%2FHo_Chi_Minh"
}
