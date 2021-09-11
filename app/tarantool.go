package app

import (
	"github.com/tarantool/go-tarantool"
	"log"
)

type TTool struct {
	Conn *tarantool.Connection
	ConnectionString string
}

func NewTarantool(connectionString string) *TTool {

	conn, err := tarantool.Connect(connectionString, tarantool.Opts{
		User: "itlbroker",
		Pass: "superpassword",
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	return &TTool{
		Conn: conn,
		ConnectionString: connectionString,
	}
}

func (t *TTool) Reconnect()  {
	var err error
	t.Conn, err = tarantool.Connect(t.ConnectionString, tarantool.Opts{
		User: "itlbroker",
		Pass: "superpassword",
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
}


