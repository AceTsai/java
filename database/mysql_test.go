package database

import (
	"fmt"
	"testing"
)

func TestMysql(t *testing.T) {
	Mysql()
	fmt.Println(Db.Name())
}
