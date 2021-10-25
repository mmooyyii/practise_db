package service

import (
	"exercise_db/internal/configure"
	"exercise_db/internal/network"
)

type Db struct {
	Config  *configure.Config
	Network *network.Tcp
}

func New(c *configure.Config, tcp *network.Tcp) *Db {
	db := new(Db)
	db.Network = tcp
	db.Config = c
	return db
}

func (db *Db) Start() {

}
