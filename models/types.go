package models

import (
	"database/sql"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}

func (r NullString) MarshalJSON() ([]byte, error) {
	if r.Valid {
		return json.Marshal(r.String)
	}
	return json.Marshal(nil)
}
