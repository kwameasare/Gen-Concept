package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type DbActionType int

const (
	Insert DbActionType = iota
	ReadAction
	UpdateAction
	DeleteAction
	Upsert
	BulkInsert
	BulkUpdate
	BulkDelete
	TransactionBegin
	TransactionCommit
	TransactionRollback
	StoredProcedureCall
	FunctionCall
	ExecuteRawQuery
	CreateTable
	AlterTable
	DropTable
	TruncateTable
	BackupDatabase
	RestoreDatabase
	MigrateDatabase
	IndexCreation
)

func (d DbActionType) String() string {
	names := [...]string{
		"INSERT",
		"READ",
		"UPDATE",
		"DELETE",
		"UPSERT",
		"BULK_INSERT",
		"BULK_UPDATE",
		"BULK_DELETE",
		"TRANSACTION_BEGIN",
		"TRANSACTION_COMMIT",
		"TRANSACTION_ROLLBACK",
		"STORED_PROCEDURE_CALL",
		"FUNCTION_CALL",
		"EXECUTE_RAW_QUERY",
		"CREATE_TABLE",
		"ALTER_TABLE",
		"DROP_TABLE",
		"TRUNCATE_TABLE",
		"BACKUP_DATABASE",
		"RESTORE_DATABASE",
		"MIGRATE_DATABASE",
		"INDEX_CREATION",
	}

	if d < Insert || int(d) >= len(names) {
		return "UNKNOWN"
	}
	return names[d]
}

func (d DbActionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *DbActionType) UnmarshalJSON(data []byte) error {
	var actionStr string
	if err := json.Unmarshal(data, &actionStr); err != nil {
		return err
	}

	switch actionStr {
	case "INSERT":
		*d = Insert
	case "READ":
		*d = ReadAction
	case "UPDATE":
		*d = UpdateAction
	case "DELETE":
		*d = DeleteAction
	case "UPSERT":
		*d = Upsert
	case "BULK_INSERT":
		*d = BulkInsert
	case "BULK_UPDATE":
		*d = BulkUpdate
	case "BULK_DELETE":
		*d = BulkDelete
	case "TRANSACTION_BEGIN":
		*d = TransactionBegin
	case "TRANSACTION_COMMIT":
		*d = TransactionCommit
	case "TRANSACTION_ROLLBACK":
		*d = TransactionRollback
	case "STORED_PROCEDURE_CALL":
		*d = StoredProcedureCall
	case "FUNCTION_CALL":
		*d = FunctionCall
	case "EXECUTE_RAW_QUERY":
		*d = ExecuteRawQuery
	case "CREATE_TABLE":
		*d = CreateTable
	case "ALTER_TABLE":
		*d = AlterTable
	case "DROP_TABLE":
		*d = DropTable
	case "TRUNCATE_TABLE":
		*d = TruncateTable
	case "BACKUP_DATABASE":
		*d = BackupDatabase
	case "RESTORE_DATABASE":
		*d = RestoreDatabase
	case "MIGRATE_DATABASE":
		*d = MigrateDatabase
	case "INDEX_CREATION":
		*d = IndexCreation
	default:
		return fmt.Errorf("invalid DbActionType: %s", actionStr)
	}

	return nil
}

func (d DbActionType) Value() (driver.Value, error) {
	return d.String(), nil
}

func (d *DbActionType) Scan(value interface{}) error {
	if value == nil {
		*d = Insert
		return nil
	}

	var actionStr string

	switch v := value.(type) {
	case string:
		actionStr = v
	case []byte:
		actionStr = string(v)
	default:
		return fmt.Errorf("unsupported Scan type for DbActionType: %T", value)
	}

	switch actionStr {
	case "INSERT":
		*d = Insert
	case "READ":
		*d = ReadAction
	case "UPDATE":
		*d = UpdateAction
	case "DELETE":
		*d = DeleteAction
	case "UPSERT":
		*d = Upsert
	case "BULK_INSERT":
		*d = BulkInsert
	case "BULK_UPDATE":
		*d = BulkUpdate
	case "BULK_DELETE":
		*d = BulkDelete
	case "TRANSACTION_BEGIN":
		*d = TransactionBegin
	case "TRANSACTION_COMMIT":
		*d = TransactionCommit
	case "TRANSACTION_ROLLBACK":
		*d = TransactionRollback
	case "STORED_PROCEDURE_CALL":
		*d = StoredProcedureCall
	case "FUNCTION_CALL":
		*d = FunctionCall
	case "EXECUTE_RAW_QUERY":
		*d = ExecuteRawQuery
	case "CREATE_TABLE":
		*d = CreateTable
	case "ALTER_TABLE":
		*d = AlterTable
	case "DROP_TABLE":
		*d = DropTable
	case "TRUNCATE_TABLE":
		*d = TruncateTable
	case "BACKUP_DATABASE":
		*d = BackupDatabase
	case "RESTORE_DATABASE":
		*d = RestoreDatabase
	case "MIGRATE_DATABASE":
		*d = MigrateDatabase
	case "INDEX_CREATION":
		*d = IndexCreation
	default:
		return fmt.Errorf("invalid DbActionType: %s", actionStr)
	}

	return nil
}
