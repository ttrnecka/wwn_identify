package entity

import (
	cdb "github.com/ttrnecka/agent_poc/common/db"
	"github.com/ttrnecka/wwn_identity/webapi/db"
)

type User struct {
	cdb.BaseModel `bson:",inline"`
	Username      string `bson:"username"`
	Email         string `bson:"email"`
	Password      string `bson:"password,omitempty"`
}

func Users() *cdb.CRUD[User] {
	return cdb.NewCRUD[User](db.Database(), "users")
}
