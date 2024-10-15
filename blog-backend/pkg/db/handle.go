package db

import "gorm.io/gorm"

type Handler interface {
	GetORMDB() *gorm.DB
}

func NEWHandler() (Handler, error) {
	return NEWSQLDB()
}
