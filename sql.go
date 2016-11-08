//line sql.y:6
package sqlparser

import __yyfmt__ "fmt"

//line sql.y:6
import "bytes"

func SetParseTree(yylex interface{}, stmt Statement) {
	yylex.(*Tokenizer).ParseTree = stmt
}

func SetAllowComments(yylex interface{}, allow bool) {
	yylex.(*Tokenizer).AllowComments = allow
}

func ForceEOF(yylex interface{}) {
	yylex.(*Tokenizer).ForceEOF = true
}

var (
	SHARE        = []byte("share")
	MODE         = []byte("mode")
	IF_BYTES     = []byte("if")
	VALUES_BYTES = []byte("values")
)

//line sql.y:31
type yySymType struct {
	yys         int
	empty       struct{}
	statement   Statement
	selStmt     SelectStatement
	byt         byte
	bytes       []byte
	bytes2      [][]byte
	str         string
	selectExprs SelectExprs
	selectExpr  SelectExpr
	columns     Columns
	colName     *ColName
	tableExprs  TableExprs
	tableExpr   TableExpr
	smTableExpr SimpleTableExpr
	tableName   *TableName
	indexHints  *IndexHints
	expr        Expr
	boolExpr    BoolExpr
	valExpr     ValExpr
	colTuple    ColTuple
	valExprs    ValExprs
	values      Values
	rowTuple    RowTuple
	subquery    *Subquery
	caseExpr    *CaseExpr
	whens       []*When
	when        *When
	orderBy     OrderBy
	order       *Order
	limit       *Limit
	insRows     InsertRows
	updateExprs UpdateExprs
	updateExpr  *UpdateExpr

	/*
	   for CreateTable
	*/
	createTableStmt   CreateTable
	columnDefinition  *ColumnDefinition
	columnDefinitions ColumnDefinitions
	columnAtts        ColumnAtts
}

const LEX_ERROR = 57346
const SELECT = 57347
const INSERT = 57348
const UPDATE = 57349
const DELETE = 57350
const FROM = 57351
const WHERE = 57352
const GROUP = 57353
const HAVING = 57354
const ORDER = 57355
const BY = 57356
const LIMIT = 57357
const FOR = 57358
const ALL = 57359
const DISTINCT = 57360
const AS = 57361
const EXISTS = 57362
const IN = 57363
const IS = 57364
const LIKE = 57365
const REGEXP = 57366
const BETWEEN = 57367
const NULL = 57368
const ASC = 57369
const DESC = 57370
const VALUES = 57371
const INTO = 57372
const DUPLICATE = 57373
const KEY = 57374
const DEFAULT = 57375
const SET = 57376
const LOCK = 57377
const ID = 57378
const STRING = 57379
const NUMBER = 57380
const VALUE_ARG = 57381
const LIST_ARG = 57382
const COMMENT = 57383
const LE = 57384
const GE = 57385
const NE = 57386
const NULL_SAFE_EQUAL = 57387
const PRIMARY = 57388
const UNIQUE = 57389
const UNION = 57390
const MINUS = 57391
const EXCEPT = 57392
const INTERSECT = 57393
const JOIN = 57394
const STRAIGHT_JOIN = 57395
const LEFT = 57396
const RIGHT = 57397
const INNER = 57398
const OUTER = 57399
const CROSS = 57400
const NATURAL = 57401
const USE = 57402
const FORCE = 57403
const ON = 57404
const OR = 57405
const AND = 57406
const NOT = 57407
const UNARY = 57408
const CASE = 57409
const WHEN = 57410
const THEN = 57411
const ELSE = 57412
const END = 57413
const CREATE = 57414
const ALTER = 57415
const DROP = 57416
const RENAME = 57417
const ANALYZE = 57418
const TABLE = 57419
const INDEX = 57420
const VIEW = 57421
const TO = 57422
const IGNORE = 57423
const IF = 57424
const USING = 57425
const SHOW = 57426
const DESCRIBE = 57427
const EXPLAIN = 57428
const BIT = 57429
const TINYINT = 57430
const SMALLINT = 57431
const MEDIUMINT = 57432
const INT = 57433
const INTEGER = 57434
const BIGINT = 57435
const REAL = 57436
const DOUBLE = 57437
const FLOAT = 57438
const UNSIGNED = 57439
const ZEROFILL = 57440
const DECIMAL = 57441
const NUMERIC = 57442
const DATE = 57443
const TIME = 57444
const TIMESTAMP = 57445
const DATETIME = 57446
const YEAR = 57447
const TEXT = 57448
const CHAR = 57449
const VARCHAR = 57450
const NULLX = 57451
const AUTO_INCREMENT = 57452
const BOOL = 57453
const APPROXNUM = 57454
const INTNUM = 57455

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LEX_ERROR",
	"SELECT",
	"INSERT",
	"UPDATE",
	"DELETE",
	"FROM",
	"WHERE",
	"GROUP",
	"HAVING",
	"ORDER",
	"BY",
	"LIMIT",
	"FOR",
	"ALL",
	"DISTINCT",
	"AS",
	"EXISTS",
	"IN",
	"IS",
	"LIKE",
	"REGEXP",
	"BETWEEN",
	"NULL",
	"ASC",
	"DESC",
	"VALUES",
	"INTO",
	"DUPLICATE",
	"KEY",
	"DEFAULT",
	"SET",
	"LOCK",
	"ID",
	"STRING",
	"NUMBER",
	"VALUE_ARG",
	"LIST_ARG",
	"COMMENT",
	"LE",
	"GE",
	"NE",
	"NULL_SAFE_EQUAL",
	"'('",
	"'='",
	"'<'",
	"'>'",
	"'~'",
	"PRIMARY",
	"UNIQUE",
	"UNION",
	"MINUS",
	"EXCEPT",
	"INTERSECT",
	"','",
	"JOIN",
	"STRAIGHT_JOIN",
	"LEFT",
	"RIGHT",
	"INNER",
	"OUTER",
	"CROSS",
	"NATURAL",
	"USE",
	"FORCE",
	"ON",
	"OR",
	"AND",
	"NOT",
	"'&'",
	"'|'",
	"'^'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'.'",
	"UNARY",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"END",
	"CREATE",
	"ALTER",
	"DROP",
	"RENAME",
	"ANALYZE",
	"TABLE",
	"INDEX",
	"VIEW",
	"TO",
	"IGNORE",
	"IF",
	"USING",
	"SHOW",
	"DESCRIBE",
	"EXPLAIN",
	"BIT",
	"TINYINT",
	"SMALLINT",
	"MEDIUMINT",
	"INT",
	"INTEGER",
	"BIGINT",
	"REAL",
	"DOUBLE",
	"FLOAT",
	"UNSIGNED",
	"ZEROFILL",
	"DECIMAL",
	"NUMERIC",
	"DATE",
	"TIME",
	"TIMESTAMP",
	"DATETIME",
	"YEAR",
	"TEXT",
	"CHAR",
	"VARCHAR",
	"NULLX",
	"AUTO_INCREMENT",
	"BOOL",
	"APPROXNUM",
	"INTNUM",
	"')'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 256
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 709

var yyAct = [...]int{

	95, 298, 160, 433, 92, 362, 249, 93, 63, 163,
	372, 253, 368, 198, 290, 209, 240, 178, 443, 86,
	81, 162, 3, 136, 135, 446, 443, 103, 263, 264,
	265, 266, 267, 51, 268, 269, 82, 430, 411, 66,
	367, 186, 71, 65, 64, 74, 343, 345, 77, 78,
	54, 29, 30, 31, 32, 69, 443, 431, 130, 87,
	52, 53, 401, 259, 232, 301, 296, 130, 408, 402,
	130, 120, 124, 44, 232, 45, 344, 47, 48, 49,
	128, 390, 42, 230, 389, 132, 388, 407, 409, 70,
	445, 136, 135, 164, 73, 50, 46, 165, 444, 347,
	353, 241, 241, 288, 274, 121, 355, 400, 123, 159,
	161, 134, 172, 66, 117, 113, 66, 65, 182, 181,
	65, 176, 39, 220, 41, 119, 169, 231, 442, 430,
	352, 72, 135, 87, 204, 182, 349, 302, 295, 285,
	208, 385, 283, 217, 218, 219, 233, 222, 223, 224,
	225, 226, 227, 228, 229, 206, 207, 202, 398, 399,
	180, 403, 196, 149, 150, 151, 211, 203, 221, 234,
	87, 87, 205, 192, 115, 66, 66, 136, 135, 65,
	247, 291, 255, 245, 127, 291, 337, 256, 236, 238,
	387, 338, 335, 190, 248, 244, 193, 336, 386, 251,
	341, 340, 339, 144, 145, 146, 147, 148, 149, 150,
	151, 60, 179, 234, 273, 260, 115, 277, 278, 279,
	275, 257, 147, 148, 149, 150, 151, 232, 202, 276,
	263, 264, 265, 266, 267, 282, 268, 269, 179, 129,
	87, 211, 29, 30, 31, 32, 76, 395, 189, 191,
	188, 293, 431, 396, 357, 297, 287, 116, 284, 261,
	294, 289, 144, 145, 146, 147, 148, 149, 150, 151,
	201, 111, 212, 14, 114, 333, 334, 421, 210, 420,
	200, 174, 419, 351, 166, 115, 378, 130, 373, 202,
	202, 354, 369, 175, 183, 170, 79, 66, 237, 359,
	98, 358, 360, 363, 201, 168, 102, 167, 110, 108,
	426, 427, 438, 364, 200, 414, 85, 99, 100, 101,
	413, 412, 370, 371, 72, 67, 90, 254, 348, 346,
	106, 330, 329, 272, 195, 133, 374, 375, 376, 379,
	377, 380, 144, 145, 146, 147, 148, 149, 150, 151,
	271, 89, 72, 391, 194, 104, 105, 83, 392, 177,
	125, 122, 109, 381, 350, 394, 144, 145, 146, 147,
	148, 149, 150, 151, 440, 118, 280, 107, 144, 145,
	146, 147, 148, 149, 150, 151, 59, 234, 61, 415,
	80, 75, 429, 441, 417, 112, 428, 393, 423, 363,
	356, 425, 424, 416, 14, 418, 281, 448, 184, 235,
	14, 15, 16, 17, 126, 57, 384, 432, 55, 33,
	434, 434, 434, 66, 435, 436, 299, 65, 243, 300,
	213, 437, 214, 215, 216, 35, 36, 37, 38, 18,
	250, 383, 332, 179, 449, 62, 447, 422, 14, 450,
	34, 451, 317, 318, 319, 320, 321, 322, 323, 324,
	325, 326, 406, 405, 327, 328, 312, 313, 314, 315,
	316, 311, 309, 310, 365, 306, 308, 307, 404, 410,
	366, 98, 304, 305, 19, 252, 303, 102, 185, 40,
	108, 258, 20, 21, 23, 22, 24, 85, 99, 100,
	101, 187, 43, 68, 25, 26, 27, 90, 246, 173,
	439, 106, 397, 361, 382, 331, 286, 171, 14, 239,
	97, 94, 96, 292, 91, 242, 137, 88, 342, 199,
	262, 197, 89, 98, 84, 270, 104, 105, 83, 102,
	131, 56, 108, 109, 28, 58, 14, 13, 12, 67,
	99, 100, 101, 11, 10, 9, 98, 8, 107, 90,
	7, 6, 102, 106, 5, 108, 4, 102, 2, 1,
	108, 0, 67, 99, 100, 101, 0, 67, 99, 100,
	101, 0, 90, 0, 89, 0, 106, 166, 104, 105,
	0, 106, 0, 0, 0, 109, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 89, 0, 0,
	107, 104, 105, 0, 0, 0, 104, 105, 109, 102,
	0, 0, 108, 109, 0, 0, 0, 0, 0, 67,
	99, 100, 101, 107, 0, 0, 0, 0, 107, 166,
	0, 0, 0, 106, 0, 0, 0, 0, 0, 0,
	138, 143, 140, 141, 142, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 104, 105,
	0, 155, 156, 157, 158, 109, 152, 153, 154, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	107, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	139, 144, 145, 146, 147, 148, 149, 150, 151,
}
var yyPact = [...]int{

	405, -1000, -1000, 189, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	30, -21, 4, -15, 3, -1000, -1000, -1000, 443, 401,
	-1000, -1000, -1000, 397, -1000, 356, 352, 436, 289, -42,
	-4, 288, -1000, 2, 288, -1000, 355, -49, 288, -49,
	354, -1000, -1000, -1000, -1000, -1000, 461, -1000, 267, 352,
	361, 35, 352, 159, -1000, 210, -1000, 34, 339, 54,
	288, -1000, -1000, 325, -1000, -23, 324, 394, 116, 288,
	-1000, 230, -1000, -1000, 316, 31, 108, 629, -1000, 536,
	513, -1000, -1000, -1000, 593, 261, 259, -1000, 249, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 593,
	-1000, 247, 289, 323, 433, 289, 593, 288, 248, 388,
	-57, -1000, 160, -1000, 318, -1000, -1000, 298, -1000, 234,
	461, -1000, -1000, 288, 95, 536, 536, 593, 232, 409,
	593, 593, 593, 97, 593, 593, 593, 593, 593, 593,
	593, 593, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	629, -46, -2, 17, 629, -1000, 541, 280, 461, -1000,
	443, 19, 270, 399, 289, 289, 228, -1000, 427, 536,
	-1000, 270, -1000, 291, -1000, 114, 288, -1000, -32, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 202, 172, 314,
	268, 24, -1000, -1000, -1000, -1000, -1000, 62, 270, -1000,
	541, -1000, -1000, 232, 593, 593, 593, 270, 270, 306,
	-1000, 380, 147, 147, 147, 86, 86, -1000, -1000, -1000,
	-1000, -1000, 593, -1000, 270, -1000, 13, 461, 10, 18,
	-1000, 536, 113, 238, 189, 117, 9, -1000, 427, 411,
	415, 108, 8, -1000, 350, 296, -1000, -1000, 295, -1000,
	431, 234, 234, -1000, -1000, 134, 128, 144, 143, 142,
	-20, -1000, 293, -30, 292, 7, -1000, 270, 270, 294,
	593, -1000, 270, -1000, 1, -1000, 14, -1000, 593, 22,
	-1000, 369, 197, -1000, -1000, -1000, 289, 411, -1000, 593,
	593, 291, -1000, -1000, -72, -1000, -1000, 246, -1000, 246,
	246, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 242, 242, 242, 240, 240, -1000,
	-1000, 429, 402, 172, 73, -1000, 140, -1000, 132, -1000,
	-1000, -1000, -1000, -7, -9, -12, -1000, -1000, -1000, -1000,
	593, 270, -1000, -1000, 270, 593, 365, 238, -1000, -1000,
	190, 196, -1000, 131, -1000, 36, -75, -1000, -1000, 283,
	-1000, -1000, -1000, 282, -1000, -1000, -1000, -1000, 277, -1000,
	-1000, -1000, 427, 536, 593, 536, -1000, -1000, 236, 233,
	231, 270, 270, 440, -1000, 593, 593, -1000, -1000, -1000,
	375, -1000, 273, -1000, -1000, -1000, -1000, 364, -1000, 360,
	-1000, -1000, -92, 195, 0, 411, 108, 170, 108, 288,
	288, 288, 289, 270, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 274, 358, -1, -1000, -31, -39, 159, -104, -1000,
	439, 386, -1000, 288, -1000, -1000, -1000, -1000, 288, -1000,
	288, -1000,
}
var yyPgo = [...]int{

	0, 569, 568, 21, 566, 564, 561, 560, 557, 555,
	554, 553, 548, 547, 419, 545, 544, 541, 20, 36,
	540, 535, 534, 531, 13, 530, 529, 211, 528, 3,
	17, 19, 527, 526, 525, 524, 2, 15, 9, 523,
	7, 522, 27, 521, 4, 520, 519, 16, 517, 516,
	515, 514, 6, 513, 5, 512, 1, 510, 509, 508,
	14, 8, 44, 246, 503, 502, 501, 491, 489, 488,
	0, 33, 486, 11, 485, 484, 12, 483, 482, 480,
	479, 478, 477, 476, 10, 475, 474, 463, 462, 450,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 3, 3, 4, 4, 5, 6, 7,
	80, 80, 72, 72, 72, 85, 85, 85, 85, 85,
	77, 77, 77, 78, 78, 82, 82, 82, 82, 82,
	82, 82, 83, 83, 83, 83, 83, 83, 83, 84,
	84, 76, 76, 79, 79, 86, 86, 86, 86, 86,
	86, 86, 81, 81, 87, 87, 88, 88, 73, 74,
	74, 75, 8, 8, 8, 9, 9, 9, 10, 11,
	11, 11, 12, 13, 13, 13, 89, 14, 15, 15,
	16, 16, 16, 16, 16, 17, 17, 18, 18, 19,
	19, 19, 22, 22, 20, 20, 20, 23, 23, 24,
	24, 24, 24, 21, 21, 21, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 26, 26, 26, 27, 27,
	28, 28, 28, 28, 29, 29, 30, 30, 31, 31,
	31, 31, 31, 32, 32, 32, 32, 32, 32, 32,
	32, 32, 32, 32, 32, 33, 33, 33, 33, 33,
	33, 33, 37, 37, 37, 42, 38, 38, 36, 36,
	36, 36, 36, 36, 36, 36, 36, 36, 36, 36,
	36, 36, 36, 36, 36, 41, 41, 43, 43, 43,
	45, 48, 48, 46, 46, 47, 49, 49, 44, 44,
	35, 35, 35, 35, 50, 50, 51, 51, 52, 52,
	53, 53, 54, 55, 55, 55, 56, 56, 56, 57,
	57, 57, 58, 58, 59, 59, 60, 60, 34, 34,
	39, 39, 40, 40, 61, 61, 62, 63, 63, 64,
	64, 65, 65, 66, 66, 66, 66, 66, 67, 67,
	68, 68, 69, 69, 70, 71,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 12, 3, 7, 7, 8, 7, 3,
	0, 1, 3, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 1, 2, 1, 1, 1, 1, 1, 1,
	1, 1, 2, 2, 2, 2, 2, 2, 2, 0,
	5, 0, 3, 0, 1, 0, 3, 2, 3, 3,
	2, 2, 1, 1, 2, 1, 1, 2, 3, 1,
	3, 7, 1, 8, 4, 6, 7, 4, 5, 4,
	5, 5, 3, 2, 2, 2, 0, 2, 0, 2,
	1, 2, 1, 1, 1, 0, 1, 1, 3, 1,
	2, 3, 1, 1, 0, 1, 2, 1, 3, 3,
	3, 3, 5, 0, 1, 2, 1, 1, 2, 3,
	2, 3, 2, 2, 2, 1, 3, 1, 1, 3,
	0, 5, 5, 5, 1, 3, 0, 2, 1, 3,
	3, 2, 3, 3, 3, 4, 3, 4, 3, 4,
	5, 6, 3, 4, 2, 1, 1, 1, 1, 1,
	1, 1, 3, 1, 1, 3, 1, 3, 1, 1,
	1, 3, 3, 3, 3, 3, 3, 3, 3, 2,
	3, 4, 5, 4, 1, 1, 1, 1, 1, 1,
	5, 0, 1, 1, 2, 4, 0, 2, 1, 3,
	1, 1, 1, 1, 0, 3, 0, 2, 0, 3,
	1, 3, 2, 0, 1, 1, 0, 2, 4, 0,
	2, 4, 0, 3, 1, 3, 0, 5, 2, 1,
	1, 3, 3, 1, 1, 3, 3, 0, 2, 0,
	3, 0, 1, 1, 1, 1, 1, 1, 0, 1,
	0, 1, 0, 2, 1, 0,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -12, -13, 5, 6, 7, 8, 34, -75,
	87, 88, 90, 89, 91, 99, 100, 101, -16, 53,
	54, 55, 56, -14, -89, -14, -14, -14, -14, 92,
	-68, 94, 52, -65, 94, 96, 92, 92, 93, 94,
	92, -71, -71, -71, -3, 17, -17, 18, -15, 30,
	-27, 36, 9, -61, -62, -44, -70, 36, -64, 97,
	93, -70, 36, 92, -70, 36, -63, 97, -70, -63,
	36, -18, -19, 77, -22, 36, -31, -36, -32, 71,
	46, -35, -44, -40, -43, -70, -41, -45, 20, 37,
	38, 39, 26, -42, 75, 76, 50, 97, 29, 82,
	41, -27, 34, 80, -27, 57, 47, 80, 36, 71,
	-70, -71, 36, -71, 95, 36, 20, 68, -70, 9,
	57, -20, -70, 19, 80, 70, 69, -33, 21, 71,
	23, 24, 25, 22, 72, 73, 74, 75, 76, 77,
	78, 79, 47, 48, 49, 42, 43, 44, 45, -31,
	-36, -31, -3, -38, -36, -36, 46, 46, 46, -42,
	46, -48, -36, -58, 34, 46, -61, 36, -30, 10,
	-62, -36, -70, 46, 20, -69, 98, -66, 90, 88,
	33, 89, 13, 36, 36, 36, -71, -23, -24, -26,
	46, 36, -42, -19, -70, 77, -31, -31, -36, -37,
	46, -42, 40, 21, 23, 24, 25, -36, -36, -36,
	26, 71, -36, -36, -36, -36, -36, -36, -36, -36,
	129, 129, 57, 129, -36, 129, -18, 18, -18, -46,
	-47, 83, -34, 29, -3, -61, -59, -44, -30, -52,
	13, -31, -74, -73, 36, 68, -70, -71, -67, 95,
	-30, 57, -25, 58, 59, 60, 61, 62, 64, 65,
	-21, 36, 19, -24, 80, -38, -37, -36, -36, -36,
	70, 26, -36, 129, -18, 129, -49, -47, 85, -31,
	-60, 68, -39, -40, -60, 129, 57, -52, -56, 15,
	14, 57, 129, -72, -78, -77, -85, -82, -83, 122,
	123, 121, 116, 117, 118, 119, 120, 102, 103, 104,
	105, 106, 107, 108, 109, 110, 111, 114, 115, 36,
	36, -50, 11, -24, -24, 58, 63, 58, 63, 58,
	58, 58, -28, 66, 96, 67, 36, 129, 36, 129,
	70, -36, 129, 86, -36, 84, 31, 57, -44, -56,
	-36, -53, -54, -36, -73, -86, -79, 112, -76, 46,
	-76, -76, -84, 46, -84, -84, -84, -76, 46, -84,
	-76, -71, -51, 12, 14, 68, 58, 58, 93, 93,
	93, -36, -36, 32, -40, 57, 57, -55, 27, 28,
	71, 26, 33, 125, -81, -87, -88, 51, 32, 52,
	-80, 113, 38, 38, 38, -52, -31, -38, -31, 46,
	46, 46, 7, -36, -54, 26, 37, 38, 32, 32,
	129, 57, -56, -29, -70, -29, -29, -61, 38, -57,
	16, 35, 129, 57, 129, 129, 129, 7, 21, -70,
	-70, -70,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 86, 86, 86, 86, 86, 72,
	250, 241, 0, 0, 0, 255, 255, 255, 0, 90,
	92, 93, 94, 95, 88, 0, 0, 0, 0, 239,
	0, 0, 251, 0, 0, 242, 0, 237, 0, 237,
	0, 83, 84, 85, 14, 91, 0, 96, 87, 0,
	0, 128, 0, 19, 234, 0, 198, 254, 0, 0,
	0, 255, 254, 0, 255, 0, 0, 0, 0, 0,
	82, 0, 97, 99, 104, 254, 102, 103, 138, 0,
	0, 168, 169, 170, 0, 198, 0, 184, 0, 200,
	201, 202, 203, 233, 187, 188, 189, 185, 186, 191,
	89, 222, 0, 0, 136, 0, 0, 0, 0, 0,
	252, 74, 0, 77, 0, 79, 238, 0, 255, 0,
	0, 100, 105, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 155, 156, 157, 158, 159, 160, 161, 141,
	0, 0, 0, 0, 166, 179, 0, 0, 0, 154,
	0, 0, 192, 0, 0, 0, 136, 129, 208, 0,
	235, 236, 199, 0, 240, 0, 0, 255, 248, 243,
	244, 245, 246, 247, 78, 80, 81, 136, 107, 113,
	0, 125, 127, 98, 106, 101, 139, 140, 143, 144,
	0, 163, 164, 0, 0, 0, 0, 146, 148, 0,
	152, 0, 171, 172, 173, 174, 175, 176, 177, 178,
	142, 165, 0, 232, 166, 180, 0, 0, 0, 196,
	193, 0, 226, 0, 229, 226, 0, 224, 208, 216,
	0, 137, 0, 69, 0, 0, 253, 75, 0, 249,
	204, 0, 0, 116, 117, 0, 0, 0, 0, 0,
	130, 114, 0, 0, 0, 0, 145, 147, 149, 0,
	0, 153, 167, 181, 0, 183, 0, 194, 0, 0,
	15, 0, 228, 230, 16, 223, 0, 216, 18, 0,
	0, 0, 71, 55, 53, 23, 24, 51, 34, 51,
	51, 32, 25, 26, 27, 28, 29, 35, 36, 37,
	38, 39, 40, 41, 49, 49, 49, 49, 49, 255,
	76, 206, 0, 108, 111, 118, 0, 120, 0, 122,
	123, 124, 109, 0, 0, 0, 115, 110, 126, 162,
	0, 150, 182, 190, 197, 0, 0, 0, 225, 17,
	217, 209, 210, 213, 70, 68, 20, 54, 33, 0,
	30, 31, 42, 0, 43, 44, 45, 46, 0, 47,
	48, 73, 208, 0, 0, 0, 119, 121, 0, 0,
	0, 151, 195, 0, 231, 0, 0, 212, 214, 215,
	0, 57, 0, 60, 61, 62, 63, 0, 65, 66,
	22, 21, 0, 0, 0, 216, 207, 205, 112, 0,
	0, 0, 0, 218, 211, 56, 58, 59, 64, 67,
	52, 0, 219, 0, 134, 0, 0, 227, 0, 13,
	0, 0, 131, 0, 132, 133, 50, 220, 0, 135,
	0, 221,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 79, 72, 3,
	46, 129, 77, 75, 57, 76, 80, 78, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	48, 47, 49, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 74, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 73, 3, 50,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 51, 52, 53, 54, 55, 56,
	58, 59, 60, 61, 62, 63, 64, 65, 66, 67,
	68, 69, 70, 71, 81, 82, 83, 84, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 101, 102, 103, 104, 105, 106,
	107, 108, 109, 110, 111, 112, 113, 114, 115, 116,
	117, 118, 119, 120, 121, 122, 123, 124, 125, 126,
	127, 128,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:181
		{
			SetParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:187
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 13:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line sql.y:203
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(AST_WHERE, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(AST_HAVING, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:207
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 15:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:213
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 16:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:217
		{
			cols := make(Columns, 0, len(yyDollar[6].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[6].updateExprs))
			for _, col := range yyDollar[6].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 17:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:229
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(AST_WHERE, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 18:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:235
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(AST_WHERE, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:241
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 20:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:246
		{
			yyVAL.str = ""
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:250
		{
			yyVAL.str = AST_ZEROFILL
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:255
		{
			yyVAL.str = yyDollar[1].str
			if yyDollar[2].str != "" {
				yyVAL.str += " " + yyDollar[2].str
			}
			if yyDollar[3].str != "" {
				yyVAL.str += " " + yyDollar[3].str
			}
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:269
		{
			yyVAL.str = AST_DATE
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:273
		{
			yyVAL.str = AST_TIME
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:277
		{
			yyVAL.str = AST_TIMESTAMP
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:281
		{
			yyVAL.str = AST_DATETIME
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:285
		{
			yyVAL.str = AST_YEAR
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:291
		{
			if yyDollar[2].str == "" {
				yyVAL.str = AST_CHAR
			} else {
				yyVAL.str = AST_CHAR + yyDollar[2].str
			}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:299
		{
			if yyDollar[2].str == "" {
				yyVAL.str = AST_VARCHAR
			} else {
				yyVAL.str = AST_VARCHAR + yyDollar[2].str
			}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:307
		{
			yyVAL.str = AST_TEXT
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:313
		{
			yyVAL.str = yyDollar[1].str + yyDollar[2].str
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:317
		{
			yyVAL.str = yyDollar[1].str
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:323
		{
			yyVAL.str = AST_BIT
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:327
		{
			yyVAL.str = AST_TINYINT
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:331
		{
			yyVAL.str = AST_SMALLINT
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:335
		{
			yyVAL.str = AST_MEDIUMINT
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:339
		{
			yyVAL.str = AST_INT
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:343
		{
			yyVAL.str = AST_INTEGER
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:347
		{
			yyVAL.str = AST_BIGINT
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:353
		{
			yyVAL.str = AST_REAL + yyDollar[2].str
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:357
		{
			yyVAL.str = AST_DOUBLE + yyDollar[2].str
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:361
		{
			yyVAL.str = AST_FLOAT + yyDollar[2].str
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:365
		{
			yyVAL.str = AST_DECIMAL + yyDollar[2].str
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:369
		{
			yyVAL.str = AST_DECIMAL + yyDollar[2].str
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:373
		{
			yyVAL.str = AST_NUMERIC + yyDollar[2].str
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:377
		{
			yyVAL.str = AST_NUMERIC + yyDollar[2].str
		}
	case 49:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:382
		{
			yyVAL.str = ""
		}
	case 50:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:386
		{
			yyVAL.str = "(" + string(yyDollar[2].bytes) + ", " + string(yyDollar[4].bytes) + ")"
		}
	case 51:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:391
		{
			yyVAL.str = ""
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:395
		{
			yyVAL.str = "(" + string(yyDollar[2].bytes) + ")"
		}
	case 53:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:400
		{
			yyVAL.str = ""
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:404
		{
			yyVAL.str = AST_UNSIGNED
		}
	case 55:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:409
		{
			yyVAL.columnAtts = ColumnAtts{}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:413
		{
			yyVAL.columnAtts = append(yyVAL.columnAtts, AST_NOT_NULL)
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:419
		{
			node := StrVal(yyDollar[3].bytes)
			yyVAL.columnAtts = append(yyVAL.columnAtts, "default "+String(node))
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:424
		{
			node := NumVal(yyDollar[3].bytes)
			yyVAL.columnAtts = append(yyVAL.columnAtts, "default "+String(node))
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:429
		{
			yyVAL.columnAtts = append(yyVAL.columnAtts, AST_AUTO_INCREMENT)
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:433
		{
			yyVAL.columnAtts = append(yyVAL.columnAtts, yyDollar[2].str)
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:439
		{
			yyVAL.str = AST_PRIMARY_KEY
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:443
		{
			yyVAL.str = AST_UNIQUE_KEY
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:457
		{
			yyVAL.columnDefinition = &ColumnDefinition{ColName: string(yyDollar[1].bytes), ColType: yyDollar[2].str, ColumnAtts: yyDollar[3].columnAtts}
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:463
		{
			yyVAL.columnDefinitions = ColumnDefinitions{yyDollar[1].columnDefinition}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:467
		{
			yyVAL.columnDefinitions = append(yyVAL.columnDefinitions, yyDollar[3].columnDefinition)
		}
	case 71:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:473
		{
			yyVAL.statement = &CreateTable{Name: yyDollar[4].bytes, ColumnDefinitions: yyDollar[6].columnDefinitions}
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:479
		{
			yyVAL.statement = yyDollar[1].statement
		}
	case 73:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:483
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 74:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:488
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 75:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:494
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 76:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:498
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 77:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:503
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 78:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:509
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 79:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:515
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 80:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:519
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 81:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:524
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:530
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:536
		{
			yyVAL.statement = &Other{}
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:540
		{
			yyVAL.statement = &Other{}
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:544
		{
			yyVAL.statement = &Other{}
		}
	case 86:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:549
		{
			SetAllowComments(yylex, true)
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:553
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 88:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:559
		{
			yyVAL.bytes2 = nil
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:563
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:569
		{
			yyVAL.str = AST_UNION
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:573
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:577
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:581
		{
			yyVAL.str = AST_EXCEPT
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:585
		{
			yyVAL.str = AST_INTERSECT
		}
	case 95:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:590
		{
			yyVAL.str = ""
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:594
		{
			yyVAL.str = AST_DISTINCT
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:600
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:604
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:610
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:614
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:618
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:624
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:628
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 104:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:633
		{
			yyVAL.bytes = nil
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:637
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:641
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:647
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:651
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:657
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:661
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:665
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 112:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:669
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 113:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:674
		{
			yyVAL.bytes = nil
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:678
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 115:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:682
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:688
		{
			yyVAL.str = AST_JOIN
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:692
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 118:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:696
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:700
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 120:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:704
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:708
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:712
		{
			yyVAL.str = AST_JOIN
		}
	case 123:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:716
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:720
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:726
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:730
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:734
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:740
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:744
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 130:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:749
		{
			yyVAL.indexHints = nil
		}
	case 131:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:753
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 132:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:757
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 133:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:761
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:767
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:771
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 136:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:776
		{
			yyVAL.boolExpr = nil
		}
	case 137:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:780
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:787
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:791
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 141:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:795
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:799
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:805
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:809
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].colTuple}
		}
	case 145:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:813
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].colTuple}
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:817
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 147:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:821
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:825
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_REGEXP, Right: yyDollar[3].valExpr}
		}
	case 149:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:829
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_REGEXP, Right: yyDollar[4].valExpr}
		}
	case 150:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:833
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 151:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:837
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:841
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 153:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:845
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 154:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:849
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 155:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:855
		{
			yyVAL.str = AST_EQ
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:859
		{
			yyVAL.str = AST_LT
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:863
		{
			yyVAL.str = AST_GT
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:867
		{
			yyVAL.str = AST_LE
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:871
		{
			yyVAL.str = AST_GE
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:875
		{
			yyVAL.str = AST_NE
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:879
		{
			yyVAL.str = AST_NSE
		}
	case 162:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:885
		{
			yyVAL.colTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:889
		{
			yyVAL.colTuple = yyDollar[1].subquery
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:893
		{
			yyVAL.colTuple = ListArg(yyDollar[1].bytes)
		}
	case 165:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:899
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:905
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 167:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:909
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:915
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:919
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:923
		{
			yyVAL.valExpr = yyDollar[1].rowTuple
		}
	case 171:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:927
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:931
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 173:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:935
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 174:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:939
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 175:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:943
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:947
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 177:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:951
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 178:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:955
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 179:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:959
		{
			if num, ok := yyDollar[2].valExpr.(NumVal); ok {
				switch yyDollar[1].byt {
				case '-':
					yyVAL.valExpr = append(NumVal("-"), num...)
				case '+':
					yyVAL.valExpr = num
				default:
					yyVAL.valExpr = &UnaryExpr{Operator: yyDollar[1].byt, Expr: yyDollar[2].valExpr}
				}
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: yyDollar[1].byt, Expr: yyDollar[2].valExpr}
			}
		}
	case 180:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:974
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 181:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:978
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 182:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:982
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 183:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:986
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 184:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:990
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 185:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:996
		{
			yyVAL.bytes = IF_BYTES
		}
	case 186:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1000
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 187:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1006
		{
			yyVAL.byt = AST_UPLUS
		}
	case 188:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1010
		{
			yyVAL.byt = AST_UMINUS
		}
	case 189:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1014
		{
			yyVAL.byt = AST_TILDA
		}
	case 190:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1020
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 191:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1025
		{
			yyVAL.valExpr = nil
		}
	case 192:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1029
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1035
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 194:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1039
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 195:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1045
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 196:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1050
		{
			yyVAL.valExpr = nil
		}
	case 197:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1054
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1060
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 199:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1064
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1070
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1074
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 202:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1078
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1082
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 204:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1087
		{
			yyVAL.valExprs = nil
		}
	case 205:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1091
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 206:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1096
		{
			yyVAL.boolExpr = nil
		}
	case 207:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1100
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 208:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1105
		{
			yyVAL.orderBy = nil
		}
	case 209:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1109
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1115
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 211:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1119
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 212:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1125
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 213:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1130
		{
			yyVAL.str = AST_ASC
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1134
		{
			yyVAL.str = AST_ASC
		}
	case 215:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1138
		{
			yyVAL.str = AST_DESC
		}
	case 216:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1143
		{
			yyVAL.limit = nil
		}
	case 217:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1147
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 218:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1151
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 219:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1156
		{
			yyVAL.str = ""
		}
	case 220:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1160
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 221:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1164
		{
			if !bytes.Equal(yyDollar[3].bytes, SHARE) {
				yylex.Error("expecting share")
				return 1
			}
			if !bytes.Equal(yyDollar[4].bytes, MODE) {
				yylex.Error("expecting mode")
				return 1
			}
			yyVAL.str = AST_SHARE_MODE
		}
	case 222:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1177
		{
			yyVAL.columns = nil
		}
	case 223:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1181
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 224:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1187
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 225:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1191
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 226:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1196
		{
			yyVAL.updateExprs = nil
		}
	case 227:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1200
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 228:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1206
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 229:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1210
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 230:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1216
		{
			yyVAL.values = Values{yyDollar[1].rowTuple}
		}
	case 231:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1220
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].rowTuple)
		}
	case 232:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1226
		{
			yyVAL.rowTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 233:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1230
		{
			yyVAL.rowTuple = yyDollar[1].subquery
		}
	case 234:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1236
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 235:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1240
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 236:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1246
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 237:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1251
		{
			yyVAL.empty = struct{}{}
		}
	case 238:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1253
		{
			yyVAL.empty = struct{}{}
		}
	case 239:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1256
		{
			yyVAL.empty = struct{}{}
		}
	case 240:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1258
		{
			yyVAL.empty = struct{}{}
		}
	case 241:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1261
		{
			yyVAL.empty = struct{}{}
		}
	case 242:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1263
		{
			yyVAL.empty = struct{}{}
		}
	case 243:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1267
		{
			yyVAL.empty = struct{}{}
		}
	case 244:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1269
		{
			yyVAL.empty = struct{}{}
		}
	case 245:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1271
		{
			yyVAL.empty = struct{}{}
		}
	case 246:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1273
		{
			yyVAL.empty = struct{}{}
		}
	case 247:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1275
		{
			yyVAL.empty = struct{}{}
		}
	case 248:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1278
		{
			yyVAL.empty = struct{}{}
		}
	case 249:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1280
		{
			yyVAL.empty = struct{}{}
		}
	case 250:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1283
		{
			yyVAL.empty = struct{}{}
		}
	case 251:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1285
		{
			yyVAL.empty = struct{}{}
		}
	case 252:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1288
		{
			yyVAL.empty = struct{}{}
		}
	case 253:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1290
		{
			yyVAL.empty = struct{}{}
		}
	case 254:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1294
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 255:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1299
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
