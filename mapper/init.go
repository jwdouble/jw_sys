package mapper

import (
	"jw.lib/sqlx"
)

func init() {
	sqlx.Register(sqlx.Driver, sqlx.PGConfigMap)
}

func Register() {

}
