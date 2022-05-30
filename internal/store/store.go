package store

import (
	"database/sql"
	logger "github.com/bumblebizoom/bumblogger"
	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	config *Config
	db     *sql.DB
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("mysql", s.config.DataBaseUrl)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	} else {
		logger.DebugMsg("DB was successfully connected")
	}

	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}
