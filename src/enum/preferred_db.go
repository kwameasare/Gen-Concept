package enum

import (
	"encoding/json"
)

type PreferredDB int

const (
	Postgres PreferredDB = iota
	Mysql
	Mongo
	Mssql
	Maria
	Oracle
	NA
)

// String method for pretty printing
func (p PreferredDB) String() string {
	return [...]string{"Postgres", "Mysql", "Mongo", "Mssql", "Maria", "Oracle","N/A"}[p]
}

// MarshalJSON for custom JSON encoding
func (p PreferredDB) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

// UnmarshalJSON for custom JSON decoding
func (p *PreferredDB) UnmarshalJSON(data []byte) error {
	var preferredDBStr string
	if err := json.Unmarshal(data, &preferredDBStr); err != nil {
		return err
	}

	switch preferredDBStr {
	case "Postgres":
		*p = Postgres
	case "Mysql":
		*p = Mysql
	case "Mongo":
		*p = Mongo
	case "Mssql":
		*p = Mssql
	case "Maria":
		*p = Maria
	case "Oracle":
		*p = Oracle
	default:
		*p = NA
	}

	return nil
}
