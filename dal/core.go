package dal

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"github.com/pburtchaell/api.segment.social/utils"
)

type SQLAdapter interface {
	Conn() *gorm.DB
}

type PostgresAdapter struct {
	conn *gorm.DB
}

func NewPostgresAdapter(conn string) *PostgresAdapter {
	db, err := gorm.Open("postgres", fmt.Sprint(conn))
	utils.CheckErr(err, "sql.Open failed")

	err = db.DB().Ping()
	utils.CheckErr(err, "db.Ping failed")

	return &PostgresAdapter{conn: &db}
}

func (ph *PostgresAdapter) Conn() *gorm.DB {
	return ph.conn
}
