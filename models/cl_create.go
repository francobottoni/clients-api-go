package models

import (
	"github.com/francobottoni/client-api/internal/database"
)

type ClientS struct {
	ClientStorageI
}

type CreateClientI interface {
	Create(cmd *CreateClientCMD) (*Client, error)
}

func (s *ClientS) Create(cmd *CreateClientCMD) (*Client, error) {
	return s.Add(cmd)
}

func NewClientCreate(c *database.MySqlClient) CreateClientI {
	return &ClientS{&ClientStorage{c}}
}
