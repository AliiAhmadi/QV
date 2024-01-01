package token

// Define a Type for store types
// of each token and be different from
// string.
type Type string

type DataType string
type MetaData string

// Define token struct with two value.
// type & literal
type Token struct {
	Type    Type
	Literal string
}

const (
	CREATE       = "CREATE"
	TABLE        = "TABLE"
	IF           = "IF"
	EXISTS       = "EXISTS"
	NULL         = "NULL"
	NOT          = "NOT"
	UNIQUE       = "UNIQUE"
	SEMICOLON    = ";"
	COMMA        = ","
	KEY          = "KEY"
	PRIMARY      = "PRIMARY"
	LPARENTHESIS = "("
	RPARENTHESIS = ")"
	INT          = "INT"
	INTEGER      = "INTEGER"
	TINYINT      = "TINYINT"
	SMALLINT     = "SMALLINT"
	MEDIUMINT    = "MEDIUMINT"
	BIGINT       = "BIGINT"
	UNSIGNED     = "UNSIGNED"
	BIG          = "BIG"
	INT2         = "INT2"
	INT8         = "INT8"
	CHARACTER    = "CHARACTER"
	VARCHAR      = "VARCHAR"
	VARYING      = "VARYING"
	NCHAR        = "NCHAR"
	NATIVE       = "NATIVE"
	NVARCHAR     = "NVARCHAR"
	TEXT         = "TEXT"
	CLOB         = "CLOB"
	BLOB         = "BLOB"
	REAL         = "REAL"
	DOUBLE       = "DOUBLE"
	PRECISION    = "PRECISION"
	FLOAT        = "FLOAT"
	NUMERIC      = "NUMERIC"
	DECIMAL      = "DECIMAL"
	BOOLEAN      = "BOOLEAN"
	DATE         = "DATE"
	DATETIME     = "DATETIME"
	DEFAULT      = "DEFAULT"
	TRUE         = "TRUE"
	FALSE        = "FALSE"
	NAME         = "NAME"
	NUMBER       = "NUMBER"
	EOQ          = "EOQ"
	ILLEGAL      = "ILLEGAL"
	STRING       = "\""
	BYTE         = "'"
)

var keywords map[string]Type = map[string]Type{
	"CREATE":    CREATE,
	"TABLE":     TABLE,
	"IF":        IF,
	"EXISTS":    EXISTS,
	"NULL":      NULL,
	"NOT":       NOT,
	"UNIQUE":    UNIQUE,
	"PRIMARY":   PRIMARY,
	"KEY":       KEY,
	"INT":       INT,
	"INTEGER":   INTEGER,
	"TINYINT":   TINYINT,
	"SMALLINT":  SMALLINT,
	"MEDIUMINT": MEDIUMINT,
	"BIGINT":    BIGINT,
	"UNSIGNED":  UNSIGNED,
	"BIG":       BIG,
	"INT2":      INT2,
	"INT8":      INT8,
	"CHARACTER": CHARACTER,
	"VARCHAR":   VARCHAR,
	"VARYING":   VARYING,
	"NCHAR":     NCHAR,
	"NATIVE":    NATIVE,
	"NVARCHAR":  NVARCHAR,
	"TEXT":      TEXT,
	"CLOB":      CLOB,
	"BLOB":      BLOB,
	"REAL":      REAL,
	"DOUBLE":    DOUBLE,
	"PRECISION": PRECISION,
	"FLOAT":     FLOAT,
	"NUMERIC":   NUMERIC,
	"DECIMAL":   DECIMAL,
	"BOOLEAN":   BOOLEAN,
	"DATE":      DATE,
	"DATETIME":  DATETIME,
	"DEFAULT":   DEFAULT,
	"TRUE":      TRUE,
	"FALSE":     FALSE,
}

func Lookup(val string) Type {
	if token, ok := keywords[val]; ok {
		return token
	}

	return NAME
}
