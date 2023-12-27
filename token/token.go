package token

// Define a Type for store types
// of each token and be different from
// string.
type Type string

// Define token struct with two value.
// type & literal
type Token struct {
	Type    Type
	Literal string
}

const (
	CREATE            = "CREATE"
	TABLE             = "TABLE"
	INE               = "IF NOT EXISTS"
	NULL              = "NULL"
	NOT               = "NOT"
	UNIQUE            = "UNIQUE"
	SEMICOLON         = ";"
	COMMA             = ","
	PK                = "PRIMARY KEY"
	LPARENTHESIS      = "("
	RPARENTHESIS      = ")"
	INT               = "INT"
	INTEGER           = "INTEGER"
	TINYINT           = "TINYINT"
	SMALLINT          = "SMALLINT"
	MEDIUMINT         = "MEDIUMINT"
	BIGINT            = "BIGINT"
	UNSIGNED_BIG_INT  = "UNSIGNED BIG INT"
	INT2              = "INT2"
	INT8              = "INT8"
	CHARACTER         = "CHARACTER"
	VARCHAR           = "VARCHAR"
	VARYING_CHARACTER = "VARYING CHARACTER"
	NCHAR             = "NCHAR"
	NATIVE_CHARACTER  = "NATIVE CHARACTER"
	NVARCHAR          = "NVARCHAR"
	TEXT              = "TEXT"
	CLOB              = "CLOB"
	BLOB              = "BLOB"
	REAL              = "REAL"
	DOUBLE            = "DOUBLE"
	DOUBLE_PRECISION  = "DOUBLE PRECISION"
	FLOAT             = "FLOAT"
	NUMERIC           = "NUMERIC"
	DECIMAL           = "DECIMAL"
	BOOLEAN           = "BOOLEAN"
	DATE              = "DATE"
	DATETIME          = "DATETIME"
	DEFAULT           = "DEFAULT"
	TRUE              = "TRUE"
	FALSE             = "FALSE"
	NAME              = "NAME"
	NUMBER            = "NUMBER"
)

var keywords map[string]Type = map[string]Type{
	"CREATE":            CREATE,
	"TABLE":             TABLE,
	"IF NOT EXISTS":     INE,
	"NULL":              NULL,
	"NOT":               NOT,
	"UNIQUE":            UNIQUE,
	"PRIMARY KEY":       PK,
	"INT":               INT,
	"INTEGER":           INTEGER,
	"TINYINT":           TINYINT,
	"SMALLINT":          SMALLINT,
	"MEDIUMINT":         MEDIUMINT,
	"BIGINT":            BIGINT,
	"UNSIGNED BIG INT":  UNSIGNED_BIG_INT,
	"INT2":              INT2,
	"INT8":              INT8,
	"CHARACTER":         CHARACTER,
	"VARCHAR":           VARCHAR,
	"VARYING CHARACTER": VARYING_CHARACTER,
	"NCHAR":             NCHAR,
	"NATIVE CHARACTER":  NATIVE_CHARACTER,
	"NVARCHAR":          NVARCHAR,
	"TEXT":              TEXT,
	"CLOB":              CLOB,
	"BLOB":              BLOB,
	"REAL":              REAL,
	"DOUBLE":            DOUBLE,
	"DOUBLE PRECISION":  DOUBLE_PRECISION,
	"FLOAT":             FLOAT,
	"NUMERIC":           NUMERIC,
	"DECIMAL":           DECIMAL,
	"BOOLEAN":           BOOLEAN,
	"DATE":              DATE,
	"DATETIME":          DATETIME,
	"DEFAULT":           DEFAULT,
	"TRUE":              TRUE,
	"FALSE":             FALSE,
}

func Lookup(val string) Type {
	if token, ok := keywords[val]; ok {
		return token
	}

	return NAME
}
