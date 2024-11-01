package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type PreferredDB int

const (
	Postgres PreferredDB = iota
	Mysql
	MongoDB
	Mssql
	MariaDB
	Oracle
	SQLite
	Redis
	Cassandra
	DynamoDB
	CouchDB
	Couchbase
	Elasticsearch
	Neo4j
	Firestore
	InfluxDB
	HBase
	Hive
	DB2
	Sybase
	Teradata
	Informix
	SAPHANA
	Snowflake
	Firebird
	Derby
	VoltDB
	NuoDB
	AmazonAurora
	GoogleBigQuery
	ApacheDrill
	ClickHouse
	Memcached
	LevelDB
	Riak
	OpenEdge
	ParAccel
	NoPreference
)

// String method for pretty printing
func (p PreferredDB) String() string {
	names := [...]string{
		"Postgres",
		"MySQL",
		"MongoDB",
		"MS SQL",
		"MariaDB",
		"Oracle",
		"SQLite",
		"Redis",
		"Cassandra",
		"DynamoDB",
		"CouchDB",
		"Couchbase",
		"Elasticsearch",
		"Neo4j",
		"Firestore",
		"InfluxDB",
		"HBase",
		"Hive",
		"DB2",
		"Sybase",
		"Teradata",
		"Informix",
		"SAP HANA",
		"Snowflake",
		"Firebird",
		"Derby",
		"VoltDB",
		"NuoDB",
		"Amazon Aurora",
		"Google BigQuery",
		"Apache Drill",
		"ClickHouse",
		"Memcached",
		"LevelDB",
		"Riak",
		"OpenEdge",
		"ParAccel",
		"N/A",
	}
	if p < Postgres || int(p) >= len(names) {
		return "UNKNOWN"
	}
	return names[p]
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
	case "MySQL":
		*p = Mysql
	case "MongoDB":
		*p = MongoDB
	case "MS SQL":
		*p = Mssql
	case "MariaDB":
		*p = MariaDB
	case "Oracle":
		*p = Oracle
	case "SQLite":
		*p = SQLite
	case "Redis":
		*p = Redis
	case "Cassandra":
		*p = Cassandra
	case "DynamoDB":
		*p = DynamoDB
	case "CouchDB":
		*p = CouchDB
	case "Couchbase":
		*p = Couchbase
	case "Elasticsearch":
		*p = Elasticsearch
	case "Neo4j":
		*p = Neo4j
	case "Firestore":
		*p = Firestore
	case "InfluxDB":
		*p = InfluxDB
	case "HBase":
		*p = HBase
	case "Hive":
		*p = Hive
	case "DB2":
		*p = DB2
	case "Sybase":
		*p = Sybase
	case "Teradata":
		*p = Teradata
	case "Informix":
		*p = Informix
	case "SAP HANA":
		*p = SAPHANA
	case "Snowflake":
		*p = Snowflake
	case "Firebird":
		*p = Firebird
	case "Derby":
		*p = Derby
	case "VoltDB":
		*p = VoltDB
	case "NuoDB":
		*p = NuoDB
	case "Amazon Aurora":
		*p = AmazonAurora
	case "Google BigQuery":
		*p = GoogleBigQuery
	case "Apache Drill":
		*p = ApacheDrill
	case "ClickHouse":
		*p = ClickHouse
	case "Memcached":
		*p = Memcached
	case "LevelDB":
		*p = LevelDB
	case "Riak":
		*p = Riak
	case "OpenEdge":
		*p = OpenEdge
	case "ParAccel":
		*p = ParAccel
	case "N/A":
		*p = NoPreference
	default:
		return fmt.Errorf("invalid PreferredDB: %s", preferredDBStr)
	}

	return nil
}

// Implement the driver.Valuer interface
func (p PreferredDB) Value() (driver.Value, error) {
	return p.String(), nil
}

// Implement the sql.Scanner interface
func (p *PreferredDB) Scan(value interface{}) error {
	if value == nil {
		*p = NoPreference
		return nil
	}

	var preferredDBStr string

	switch v := value.(type) {
	case string:
		preferredDBStr = v
	case []byte:
		preferredDBStr = string(v)
	default:
		return fmt.Errorf("unsupported Scan type for PreferredDB: %T", value)
	}

	switch preferredDBStr {
	case "Postgres":
		*p = Postgres
	case "MySQL":
		*p = Mysql
	case "MongoDB":
		*p = MongoDB
	case "MS SQL":
		*p = Mssql
	case "MariaDB":
		*p = MariaDB
	case "Oracle":
		*p = Oracle
	case "SQLite":
		*p = SQLite
	case "Redis":
		*p = Redis
	case "Cassandra":
		*p = Cassandra
	case "DynamoDB":
		*p = DynamoDB
	case "CouchDB":
		*p = CouchDB
	case "Couchbase":
		*p = Couchbase
	case "Elasticsearch":
		*p = Elasticsearch
	case "Neo4j":
		*p = Neo4j
	case "Firestore":
		*p = Firestore
	case "InfluxDB":
		*p = InfluxDB
	case "HBase":
		*p = HBase
	case "Hive":
		*p = Hive
	case "DB2":
		*p = DB2
	case "Sybase":
		*p = Sybase
	case "Teradata":
		*p = Teradata
	case "Informix":
		*p = Informix
	case "SAP HANA":
		*p = SAPHANA
	case "Snowflake":
		*p = Snowflake
	case "Firebird":
		*p = Firebird
	case "Derby":
		*p = Derby
	case "VoltDB":
		*p = VoltDB
	case "NuoDB":
		*p = NuoDB
	case "Amazon Aurora":
		*p = AmazonAurora
	case "Google BigQuery":
		*p = GoogleBigQuery
	case "Apache Drill":
		*p = ApacheDrill
	case "ClickHouse":
		*p = ClickHouse
	case "Memcached":
		*p = Memcached
	case "LevelDB":
		*p = LevelDB
	case "Riak":
		*p = Riak
	case "OpenEdge":
		*p = OpenEdge
	case "ParAccel":
		*p = ParAccel
	case "N/A":
		*p = NoPreference
	default:
		return fmt.Errorf("invalid PreferredDB: %s", preferredDBStr)
	}

	return nil
}
