package models

import (
	"github.com/francobottoni/client-api/internal/database"
	"github.com/francobottoni/client-api/internal/log"
)

type ClientStorageI interface {
	Add(cmd *CreateClientCMD) (*Client, error)
}

type ClientStorage struct {
	*database.MySqlClient
}

func (s *ClientStorage) Add(cmd *CreateClientCMD) (*Client, error) {
	tx, err := s.MySqlClient.Begin()

	if err != nil {
		log.Log().Error("cannot create transaction")
		return nil, err
	}

	res, err := tx.Exec(`insert into client (dni, name, price, country_origin)
	values (?, ?, ?, ?)`, cmd.DNI, cmd.Name, cmd.LastName, cmd.CountryOrigin)

	if err != nil {
		log.Log().Error("cannot execute statement")
		_ = tx.Rollback() //Use rollback in case error to close the transaction
		return nil, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		log.Log().Error("cannot fetch last id")
		_ = tx.Rollback()
		return nil, err
	}
	_ = tx.Commit() //Use commit in a happy case

	return &Client{
		ID:            id,
		DNI:           cmd.DNI,
		Name:          cmd.Name,
		LastName:      cmd.LastName,
		CountryOrigin: cmd.CountryOrigin,
	}, nil

}
