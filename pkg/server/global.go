package server

import (
	"log"

	"github.com/bwmarrin/snowflake"
)

func init() {
	var err error
	IdGenerator, err = snowflake.NewNode(1)
	if err != nil {
		log.Fatalln(err)
	}
}

var IdGenerator *snowflake.Node
