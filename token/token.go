package token

// Define a Type for store types
// of each token and be different from
// string.
type Type string

type DataType string
type MetaData string

var DataTypes map[string]DataType = map[string]DataType{
	"INT":       INT,
	"INTEGER":   INTEGER,
	"TINYINT":   TINYINT,
	"SMALLINT":  SMALLINT,
	"MEDIUMINT": MEDIUMINT,
	"BIGINT":    BIGINT,
	"INT2":      INT2,
	"INT8":      INT8,
	"UNSIGNED":  UNSIGNED,
	"BIG":       BIG,
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
}

// Define token struct with two value.
// type & literal
type Token struct {
	Type    Type
	Literal string
}

/*
* Resource => https://sqlite.org/lang_keywords.html
 */
const (
	CREATE               = "CREATE"
	TABLE                = "TABLE"
	IF                   = "IF"
	EXISTS               = "EXISTS"
	NULL                 = "NULL"
	NOT                  = "NOT"
	UNIQUE               = "UNIQUE"
	SEMICOLON            = ";"
	COMMA                = ","
	KEY                  = "KEY"
	PRIMARY              = "PRIMARY"
	LPARENTHESIS         = "("
	RPARENTHESIS         = ")"
	INT                  = "INT"
	INTEGER              = "INTEGER"
	TINYINT              = "TINYINT"
	SMALLINT             = "SMALLINT"
	MEDIUMINT            = "MEDIUMINT"
	BIGINT               = "BIGINT"
	UNSIGNED             = "UNSIGNED"
	BIG                  = "BIG"
	INT2                 = "INT2"
	INT8                 = "INT8"
	CHARACTER            = "CHARACTER"
	VARCHAR              = "VARCHAR"
	VARYING              = "VARYING"
	NCHAR                = "NCHAR"
	NATIVE               = "NATIVE"
	NVARCHAR             = "NVARCHAR"
	TEXT                 = "TEXT"
	CLOB                 = "CLOB"
	BLOB                 = "BLOB"
	REAL                 = "REAL"
	DOUBLE               = "DOUBLE"
	PRECISION            = "PRECISION"
	FLOAT                = "FLOAT"
	NUMERIC              = "NUMERIC"
	DECIMAL              = "DECIMAL"
	BOOLEAN              = "BOOLEAN"
	DATE                 = "DATE"
	DATETIME             = "DATETIME"
	DEFAULT              = "DEFAULT"
	TRUE                 = "TRUE"
	FALSE                = "FALSE"
	NAME                 = "NAME"
	NUMBER               = "NUMBER"
	EOQ                  = "EOQ"
	ILLEGAL              = "ILLEGAL"
	STRING               = "\""
	BYTE                 = "'"
	ABORT                = "ABORT"
	ACTION               = "ACTION"
	ADD                  = "ADD"
	AFTER                = "AFTER"
	ALL                  = "ALL"
	ALTER                = "ALTER"
	ALWAYS               = "ALWAYS"
	ANALYZE              = "ANALYZE"
	AND                  = "AND"
	AS                   = "AS"
	ASC                  = "ASC"
	ATTACH               = "ATTACH"
	AUTOINCREMENT        = "AUTOINCREMENT"
	BEFORE               = "BEFORE"
	BEGIN                = "BEGIN"
	BETWEEN              = "BETWEEN"
	BY                   = "BY"
	CASCADE              = "CASCADE"
	CASE                 = "CASE"
	CAST                 = "CAST"
	CHECK                = "CHECK"
	COLLATE              = "COLLATE"
	COLUMN               = "COLUMN"
	COMMIT               = "COMMIT"
	CONFLICT             = "CONFLICT"
	CONSTRAINT           = "CONSTRAINT"
	CROSS                = "CROSS"
	CURRENT              = "CURRENT"
	CURRENT_DATE         = "CURRENT_DATE"
	CURRENT_TIME         = "CURRENT_TIME"
	CURRENT_TIMESTAMP    = "CURRENT_TIMESTAMP"
	DATABASE             = "DATABASE"
	DEFERRABLE           = "DEFERRABLE"
	DEFERRED             = "DEFERRED"
	DELETE               = "DELETE"
	DESC                 = "DESC"
	DETACH               = "DETACH"
	DISTINCT             = "DISTINCT"
	DO                   = "DO"
	DROP                 = "DROP"
	EACH                 = "EACH"
	ELSE                 = "ELSE"
	END                  = "END"
	ESCAPE               = "ESCAPE"
	EXCEPT               = "EXCEPT"
	EXCLUDE              = "EXCLUDE"
	EXCLUSIVE            = "EXCLUSIVE"
	EXPLAIN              = "EXPLAIN"
	FAIL                 = "FAIL"
	FILTER               = "FILTER"
	FIRST                = "FIRST"
	FOLLOWING            = "FOLLOWING"
	FOR                  = "FOR"
	FOREIGN              = "FOREIGN"
	FROM                 = "FROM"
	FULL                 = "FULL"
	GENERATED            = "GENERATED"
	GLOB                 = "GLOB"
	GROUP                = "GROUP"
	GROUPS               = "GROUPS"
	HAVING               = "HAVING"
	IGNORE               = "IGNORE"
	IMMEDIATE            = "IMMEDIATE"
	IN                   = "IN"
	INDEX                = "INDEX"
	INDEXED              = "INDEXED"
	INITIALLY            = "INITIALLY"
	INNER                = "INNER"
	INSERT               = "INSERT"
	INSTEAD              = "INSTEAD"
	INTERSECT            = "INTERSECT"
	INTO                 = "INTO"
	IS                   = "IS"
	ISNULL               = "ISNULL"
	JOIN                 = "JOIN"
	LAST                 = "LAST"
	LEFT                 = "LEFT"
	LIKE                 = "LIKE"
	LIMIT                = "LIMIT"
	MATCH                = "MATCH"
	MATERIALIZED         = "MATERIALIZED"
	NATURAL              = "NATURAL"
	NO                   = "NO"
	NOTHING              = "NOTHING"
	NOTNULL              = "NOTNULL"
	NULLS                = "NULLS"
	OF                   = "OF"
	OFFSET               = "OFFSET"
	ON                   = "ON"
	OR                   = "OR"
	ORDER                = "ORDER"
	OTHERS               = "OTHERS"
	OUTER                = "OUTER"
	OVER                 = "OVER"
	PARTITION            = "PARTITION"
	PLAN                 = "PLAN"
	PRAGMA               = "PRAGMA"
	PRECEDING            = "PRECEDING"
	QUERY                = "QUERY"
	RAISE                = "RAISE"
	RANGE                = "RANGE"
	RECURSIVE            = "RECURSIVE"
	REFERENCES           = "REFERENCES"
	REGEXP               = "REGEXP"
	REINDEX              = "REINDEX"
	RELEASE              = "RELEASE"
	RENAME               = "RENAME"
	REPLACE              = "REPLACE"
	RESTRICT             = "RESTRICT"
	RETURNING            = "RETURNING"
	RIGHT                = "RIGHT"
	ROLLBACK             = "ROLLBACK"
	ROW                  = "ROW"
	ROWS                 = "ROWS"
	SAVEPOINT            = "SAVEPOINT"
	SELECT               = "SELECT"
	SET                  = "SET"
	TEMP                 = "TEMP"
	TEMPORARY            = "TEMPORARY"
	THEN                 = "THEN"
	TIES                 = "TIES"
	TO                   = "TO"
	TRANSACTION          = "TRANSACTION"
	TRIGGER              = "TRIGGER"
	UNBOUNDED            = "UNBOUNDED"
	UNION                = "UNION"
	UPDATE               = "UPDATE"
	USING                = "USING"
	VACUUM               = "VACUUM"
	VALUES               = "VALUES"
	VIEW                 = "VIEW"
	VIRTUAL              = "VIRTUAL"
	WHEN                 = "WHEN"
	WHERE                = "WHERE"
	WINDOW               = "WINDOW"
	WITH                 = "WITH"
	WITHOUT              = "WITHOUT"
	STAR                 = "*"
	EQUALITY             = "="
	SINGLE_QUOTES_STRING = "'"
	LESSTHAN             = "<"
	GREATERTHAN          = ">"
)

var keywords map[string]Type = map[string]Type{
	"CREATE":            CREATE,
	"TABLE":             TABLE,
	"IF":                IF,
	"EXISTS":            EXISTS,
	"NULL":              NULL,
	"NOT":               NOT,
	"UNIQUE":            UNIQUE,
	"PRIMARY":           PRIMARY,
	"KEY":               KEY,
	"INT":               INT,
	"INTEGER":           INTEGER,
	"TINYINT":           TINYINT,
	"SMALLINT":          SMALLINT,
	"MEDIUMINT":         MEDIUMINT,
	"BIGINT":            BIGINT,
	"UNSIGNED":          UNSIGNED,
	"BIG":               BIG,
	"INT2":              INT2,
	"INT8":              INT8,
	"CHARACTER":         CHARACTER,
	"VARCHAR":           VARCHAR,
	"VARYING":           VARYING,
	"NCHAR":             NCHAR,
	"NATIVE":            NATIVE,
	"NVARCHAR":          NVARCHAR,
	"TEXT":              TEXT,
	"CLOB":              CLOB,
	"BLOB":              BLOB,
	"REAL":              REAL,
	"DOUBLE":            DOUBLE,
	"PRECISION":         PRECISION,
	"FLOAT":             FLOAT,
	"NUMERIC":           NUMERIC,
	"DECIMAL":           DECIMAL,
	"BOOLEAN":           BOOLEAN,
	"DATE":              DATE,
	"DATETIME":          DATETIME,
	"DEFAULT":           DEFAULT,
	"TRUE":              TRUE,
	"FALSE":             FALSE,
	"ABORT":             ABORT,
	"ACTION":            ACTION,
	"ADD":               ADD,
	"AFTER":             AFTER,
	"ALL":               ALL,
	"ALTER":             ALTER,
	"ALWAYS":            ALWAYS,
	"ANALYZE":           ANALYZE,
	"AND":               AND,
	"AS":                AS,
	"ASC":               ASC,
	"ATTACH":            ATTACH,
	"AUTOINCREMENT":     AUTOINCREMENT,
	"BEFORE":            BEFORE,
	"BEGIN":             BEGIN,
	"BETWEEN":           BETWEEN,
	"BY":                BY,
	"CASCADE":           CASCADE,
	"CASE":              CASE,
	"CAST":              CAST,
	"CHECK":             CHECK,
	"COLLATE":           COLLATE,
	"COLUMN":            COLUMN,
	"COMMIT":            COMMIT,
	"CONFLICT":          CONFLICT,
	"CONSTRAINT":        CONSTRAINT,
	"CROSS":             CROSS,
	"CURRENT":           CURRENT,
	"CURRENT_DATE":      CURRENT_DATE,
	"CURRENT_TIME":      CURRENT_TIME,
	"CURRENT_TIMESTAMP": CURRENT_TIMESTAMP,
	"DATABASE":          DATABASE,
	"DEFERRABLE":        DEFERRABLE,
	"DEFERRED":          DEFERRED,
	"DELETE":            DELETE,
	"DESC":              DESC,
	"DETACH":            DETACH,
	"DISTINCT":          DISTINCT,
	"DO":                DO,
	"DROP":              DROP,
	"EACH":              EACH,
	"ELSE":              ELSE,
	"END":               END,
	"ESCAPE":            ESCAPE,
	"EXCEPT":            EXCEPT,
	"EXCLUDE":           EXCLUDE,
	"EXCLUSIVE":         EXCLUSIVE,
	"EXPLAIN":           EXPLAIN,
	"FAIL":              FAIL,
	"FILTER":            FILTER,
	"FIRST":             FIRST,
	"FOLLOWING":         FOLLOWING,
	"FOR":               FOR,
	"FOREIGN":           FOREIGN,
	"FROM":              FROM,
	"FULL":              FULL,
	"GENERATED":         GENERATED,
	"GLOB":              GLOB,
	"GROUP":             GROUP,
	"GROUPS":            GROUPS,
	"HAVING":            HAVING,
	"IGNORE":            IGNORE,
	"IMMEDIATE":         IMMEDIATE,
	"IN":                IN,
	"INDEX":             INDEX,
	"INDEXED":           INDEXED,
	"INITIALLY":         INITIALLY,
	"INNER":             INNER,
	"INSERT":            INSERT,
	"INSTEAD":           INSTEAD,
	"INTERSECT":         INTERSECT,
	"INTO":              INTO,
	"IS":                IS,
	"ISNULL":            ISNULL,
	"JOIN":              JOIN,
	"LAST":              LAST,
	"LEFT":              LEFT,
	"LIKE":              LIKE,
	"LIMIT":             LIMIT,
	"MATCH":             MATCH,
	"MATERIALIZED":      MATERIALIZED,
	"NATURAL":           NATURAL,
	"NO":                NO,
	"NOTHING":           NOTHING,
	"NOTNULL":           NOTNULL,
	"NULLS":             NULLS,
	"OF":                OF,
	"OFFSET":            OFFSET,
	"ON":                ON,
	"OR":                OR,
	"ORDER":             ORDER,
	"OTHERS":            OTHERS,
	"OUTER":             OUTER,
	"OVER":              OVER,
	"PARTITION":         PARTITION,
	"PLAN":              PLAN,
	"PRAGMA":            PRAGMA,
	"PRECEDING":         PRECEDING,
	"QUERY":             QUERY,
	"RAISE":             RAISE,
	"RANGE":             RANGE,
	"RECURSIVE":         RECURSIVE,
	"REFERENCES":        REFERENCES,
	"REGEXP":            REGEXP,
	"REINDEX":           REINDEX,
	"RELEASE":           RELEASE,
	"RENAME":            RENAME,
	"REPLACE":           REPLACE,
	"RESTRICT":          RESTRICT,
	"RETURNING":         RETURNING,
	"RIGHT":             RIGHT,
	"ROLLBACK":          ROLLBACK,
	"ROW":               ROW,
	"ROWS":              ROWS,
	"SAVEPOINT":         SAVEPOINT,
	"SELECT":            SELECT,
	"SET":               SET,
	"TEMP":              TEMP,
	"TEMPORARY":         TEMPORARY,
	"THEN":              THEN,
	"TIES":              TIES,
	"TO":                TO,
	"TRANSACTION":       TRANSACTION,
	"TRIGGER":           TRIGGER,
	"UNBOUNDED":         UNBOUNDED,
	"UNION":             UNION,
	"UPDATE":            UPDATE,
	"USING":             USING,
	"VACUUM":            VACUUM,
	"VALUES":            VALUES,
	"VIEW":              VIEW,
	"VIRTUAL":           VIRTUAL,
	"WHEN":              WHEN,
	"WHERE":             WHERE,
	"WINDOW":            WINDOW,
	"WITH":              WITH,
	"WITHOUT":           WITHOUT,
	"*":                 STAR,
	"<":                 LESSTHAN,
	">":                 GREATERTHAN,
	"=":                 EQUALITY,
}

func Lookup(val string) Type {
	if token, ok := keywords[val]; ok {
		return token
	}

	return NAME
}
