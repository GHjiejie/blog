package db

import (
	"articleManage/pkg/config"
	"fmt"
	"strings"
)

type Handle interface {
}

func NewHandler(c *config.DBConfig) (Handle, error) {
	switch strings.ToLower(c.Type) {
	case "mysql":
		return NewSQLDB(c.SQLPara)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", c.Type)
	}

}
