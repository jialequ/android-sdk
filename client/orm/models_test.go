// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package orm

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jialequ/android-sdk/client/orm/internal/models"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// A slice string field.
type SliceStringField []string

func (e SliceStringField) Value() []string {
	return []string(e)
}

func (e *SliceStringField) Set(d []string) {
	*e = SliceStringField(d)
}

func (e *SliceStringField) Add(v string) {
	*e = append(*e, v)
}

func (e *SliceStringField) String() string {
	return strings.Join(e.Value(), ",")
}

func (e *SliceStringField) FieldType() int {
	return TypeVarCharField
}

func (e *SliceStringField) SetRaw(value interface{}) error {
	f := func(str string) {
		if len(str) > 0 {
			parts := strings.Split(str, ",")
			v := make([]string, 0, len(parts))
			for _, p := range parts {
				v = append(v, strings.TrimSpace(p))
			}
			e.Set(v)
		}
	}

	switch d := value.(type) {
	case []string:
		e.Set(d)
	case string:
		f(d)
	case []byte:
		f(string(d))
	default:
		return fmt.Errorf("<SliceStringField.SetRaw> unknown value `%v`", value)
	}
	return nil
}

func (e *SliceStringField) RawValue() interface{} {
	return e.String()
}

var _ models.Fielder = new(SliceStringField)

// A json field.
type JSONFieldTest struct {
	Name string
	Data string
}

func (e *JSONFieldTest) String() string {
	data, _ := json.Marshal(e)
	return string(data)
}

func (e *JSONFieldTest) FieldType() int {
	return TypeTextField
}

func (e *JSONFieldTest) SetRaw(value interface{}) error {
	switch d := value.(type) {
	case string:
		return json.Unmarshal([]byte(d), e)
	case []byte:
		return json.Unmarshal(d, e)
	default:
		return fmt.Errorf("<JSONField.SetRaw> unknown value `%v`", value)
	}
}

func (e *JSONFieldTest) RawValue() interface{} {
	return e.String()
}

var _ models.Fielder = new(JSONFieldTest)

type Data struct {
	ID       int `orm:"column(id)"`
	Boolean  bool
	Byte     byte
	Int8     int8
	Uint8    uint8
	Rune     rune
	Char     string    `orm:"size(50)"`
	Text     string    `orm:"type(text)"`
	JSON     string    `orm:"type(json);default({\"name\":\"json\"})"`
	Jsonb    string    `orm:"type(jsonb)"`
	Time     time.Time `orm:"type(time)"`
	Date     time.Time `orm:"type(date)"`
	DateTime time.Time `orm:"column(datetime)"`
	Int      int
	Uint     uint
	Int16    int16
	Uint16   uint16
	Int32    int32
	Int64    int64
	Uint32   uint32
	Float32  float32
	Uint64   uint64
	Float64  float64
	Decimal  float64 `orm:"digits(8);decimals(4)"`
}

type DataNull struct {
	ID                int             `orm:"column(id)"`
	Char              string          `orm:"null;size(50)"`
	Text              string          `orm:"null;type(text)"`
	JSON              string          `orm:"type(json);null"`
	Jsonb             string          `orm:"type(jsonb);null"`
	Time              time.Time       `orm:"null;type(time)"`
	Date              time.Time       `orm:"null;type(date)"`
	DateTime          time.Time       `orm:"null;column(datetime)"`
	DateTimePrecision time.Time       `orm:"null;type(datetime);precision(4)"`
	Boolean           bool            `orm:"null"`
	Byte              byte            `orm:"null"`
	Int8              int8            `orm:"null"`
	Uint8             uint8           `orm:"null"`
	Rune              rune            `orm:"null"`
	Int               int             `orm:"null"`
	Uint              uint            `orm:"null"`
	Int16             int16           `orm:"null"`
	Uint16            uint16          `orm:"null"`
	Int32             int32           `orm:"null"`
	Int64             int64           `orm:"null"`
	Uint32            uint32          `orm:"null"`
	Float32           float32         `orm:"null"`
	Uint64            uint64          `orm:"null"`
	Float64           float64         `orm:"null"`
	Decimal           float64         `orm:"digits(8);decimals(4);null"`
	NullString        sql.NullString  `orm:"null"`
	NullBool          sql.NullBool    `orm:"null"`
	NullFloat64       sql.NullFloat64 `orm:"null"`
	NullInt64         sql.NullInt64   `orm:"null"`
	BooleanPtr        *bool           `orm:"null"`
	CharPtr           *string         `orm:"null;size(50)"`
	TextPtr           *string         `orm:"null;type(text)"`
	BytePtr           *byte           `orm:"null"`
	RunePtr           *rune           `orm:"null"`
	IntPtr            *int            `orm:"null"`
	Int8Ptr           *int8           `orm:"null"`
	Int16Ptr          *int16          `orm:"null"`
	Int32Ptr          *int32          `orm:"null"`
	Int64Ptr          *int64          `orm:"null"`
	UintPtr           *uint           `orm:"null"`
	Uint8Ptr          *uint8          `orm:"null"`
	Uint16Ptr         *uint16         `orm:"null"`
	Uint32Ptr         *uint32         `orm:"null"`
	Uint64Ptr         *uint64         `orm:"null"`
	Float32Ptr        *float32        `orm:"null"`
	Float64Ptr        *float64        `orm:"null"`
	DecimalPtr        *float64        `orm:"digits(8);decimals(4);null"`
	TimePtr           *time.Time      `orm:"null;type(time)"`
	DatePtr           *time.Time      `orm:"null;type(date)"`
	DateTimePtr       *time.Time      `orm:"null"`
}

type (
	String  string
	Boolean bool
	Byte    byte
	Rune    rune
	Int     int
	Int8    int8
	Int16   int16
	Int32   int32
	Int64   int64
	Uint    uint
	Uint8   uint8
	Uint16  uint16
	Uint32  uint32
	Uint64  uint64
	Float32 float64
	Float64 float64
)

type DataCustom struct {
	ID      int `orm:"column(id)"`
	Boolean Boolean
	Byte    Byte
	Int8    Int8
	Uint8   Uint8
	Rune    Rune
	Char    string `orm:"size(50)"`
	Text    string `orm:"type(text)"`
	Int     Int
	Uint    Uint
	Int16   Int16
	Uint16  Uint16
	Int32   Int32
	Int64   Int64
	Uint32  Uint32
	Float32 Float32
	Uint64  Uint64
	Float64 Float64
	Decimal Float64 `orm:"digits(8);decimals(4)"`
}

// only for mysql
type UserBig struct {
	ID   uint64 `orm:"column(id)"`
	Name string
}

type TM struct {
	ID           int       `orm:"column(id)"`
	TMPrecision1 time.Time `orm:"type(datetime);precision(3)"`
	TMPrecision2 time.Time `orm:"auto_now_add;type(datetime);precision(4)"`
}

func (t *TM) TableName() string {
	return "tm"
}

func NewTM() *TM {
	obj := new(TM)
	return obj
}

type DeptInfo struct {
	ID           int       `orm:"column(id)"`
	Created      time.Time `orm:"auto_now_add"`
	DeptName     string
	EmployeeName string
	Salary       int
}

type UnregisterModel struct {
	ID           int       `orm:"column(id)"`
	Created      time.Time `orm:"auto_now_add"`
	DeptName     string
	EmployeeName string
	Salary       int
}

type User struct {
	ID             int    `orm:"column(id)"`
	UserName       string `orm:"size(30);unique"`
	Email          string `orm:"size(100)"`
	Password       string `orm:"size(100)"`
	Status         int16  `orm:"column(Status)"`
	IsStaff        bool
	IsActive       bool `orm:"default(true)"`
	Unexported     bool `orm:"-"`
	UnexportedBool bool
	Created        time.Time `orm:"auto_now_add;type(date)"`
	Updated        time.Time `orm:"auto_now"`
	Profile        *Profile  `orm:"null;rel(one);on_delete(set_null)"`
	Posts          []*Post   `orm:"reverse(many)" json:"-"`
	ShouldSkip     string    `orm:"-"`
	Nums           int
	Langs          SliceStringField `orm:"size(100)"`
	Extra          JSONFieldTest    `orm:"type(text)"`
}

func (u *User) TableIndex() [][]string {
	return [][]string{
		{"Id", "UserName"},
		{"Id", "Created"},
	}
}

func (u *User) TableUnique() [][]string {
	return [][]string{
		{"UserName", "Email"},
	}
}

func NewUser() *User {
	obj := new(User)
	return obj
}

type Profile struct {
	ID       int `orm:"column(id)"`
	Age      int16
	Money    float64
	User     *User `orm:"reverse(one)" json:"-"`
	BestPost *Post `orm:"rel(one);null"`
}

func (u *Profile) TableName() string {
	return "user_profile"
}

func NewProfile() *Profile {
	obj := new(Profile)
	return obj
}

type Post struct {
	ID               int       `orm:"column(id)"`
	User             *User     `orm:"rel(fk)"`
	Title            string    `orm:"size(60)"`
	Content          string    `orm:"type(text)"`
	Created          time.Time `orm:"auto_now_add"`
	Updated          time.Time `orm:"auto_now"`
	UpdatedPrecision time.Time `orm:"auto_now;type(datetime);precision(4)"`
	Tags             []*Tag    `orm:"rel(m2m);rel_through(github.com/jialequ/android-sdk/client/orm.PostTags)"`
}

func (u *Post) TableIndex() [][]string {
	return [][]string{
		{"Id", "Created"},
	}
}

func NewPost() *Post {
	obj := new(Post)
	return obj
}

type NullValue struct {
	ID    int    `orm:"column(id)"`
	Value string `orm:"size(30);null"`
}

type Tag struct {
	ID       int     `orm:"column(id)"`
	Name     string  `orm:"size(30)"`
	BestPost *Post   `orm:"rel(one);null"`
	Posts    []*Post `orm:"reverse(many)" json:"-"`
}

func NewTag() *Tag {
	obj := new(Tag)
	return obj
}

type PostTags struct {
	ID   int   `orm:"column(id)"`
	Post *Post `orm:"rel(fk)"`
	Tag  *Tag  `orm:"rel(fk)"`
}

func (m *PostTags) TableName() string {
	return "prefix_post_tags"
}

type Comment struct {
	ID      int       `orm:"column(id)"`
	Post    *Post     `orm:"rel(fk);column(post)"`
	Content string    `orm:"type(text)"`
	Parent  *Comment  `orm:"null;rel(fk)"`
	Created time.Time `orm:"auto_now_add"`
}

func NewComment() *Comment {
	obj := new(Comment)
	return obj
}

type Group struct {
	ID          int `orm:"column(gid);size(32)"`
	Name        string
	Permissions []*Permission `orm:"reverse(many)" json:"-"`
}

type Permission struct {
	ID     int `orm:"column(id)"`
	Name   string
	Groups []*Group `orm:"rel(m2m);rel_through(github.com/jialequ/android-sdk/client/orm.GroupPermissions)"`
}

type GroupPermissions struct {
	ID         int         `orm:"column(id)"`
	Group      *Group      `orm:"rel(fk)"`
	Permission *Permission `orm:"rel(fk)"`
}

type ModelID struct {
	ID int64
}

type ModelBase struct {
	ModelID

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

type InLine struct {
	// Common Fields
	ModelBase

	// Other Fields
	Name  string `orm:"unique"`
	Email string
}

type Index struct {
	// Common Fields
	Id int `orm:"column(id)"`

	// Other Fields
	F1 int `orm:"column(f1);index"`
	F2 int `orm:"column(f2);index"`
}

func NewInLine() *InLine {
	return new(InLine)
}

type InLineOneToOne struct {
	// Common Fields
	ModelBase

	Note   string
	InLine *InLine `orm:"rel(fk);column(inline)"`
}

func NewInLineOneToOne() *InLineOneToOne {
	return new(InLineOneToOne)
}

type IntegerPk struct {
	ID    int64 `orm:"pk"`
	Value string
}

type UintPk struct {
	ID   uint32 `orm:"pk"`
	Name string
}

type PtrPk struct {
	ID       *IntegerPk `orm:"pk;rel(one)"`
	Positive bool
}

type StrPk struct {
	Id    string `orm:"column(id);size(64);pk"`
	Value string
}

var DBARGS = struct {
	Driver string
	Source string
	Debug  string
}{
	os.Getenv("ORM_DRIVER"),
	os.Getenv("ORM_SOURCE"),
	os.Getenv("ORM_DEBUG"),
}

var (
	IsMysql    = DBARGS.Driver == "mysql"
	IsSqlite   = DBARGS.Driver == "sqlite3"
	IsPostgres = DBARGS.Driver == "postgres"
	IsTidb     = DBARGS.Driver == "tidb"
)

var (
	dORM     Ormer
	dDbBaser dbBaser
)

var helpinfo = `need driver and source!

	Default DB Drivers.

	  driver: url
	   mysql: https://github.com/go-sql-driver/mysql
	 sqlite3: https://github.com/mattn/go-sqlite3
	postgres: https://github.com/lib/pq
	tidb: https://github.com/pingcap/tidb

	usage:

	go Get -u github.com/jialequ/android-sdk/client/orm
	go Get -u github.com/go-sql-driver/mysql
	go Get -u github.com/mattn/go-sqlite3
	go Get -u github.com/lib/pq
	go Get -u github.com/pingcap/tidb

	#### MySQL
	mysql -u root -e 'create database orm_test;'
	export ORM_DRIVER=mysql
	export ORM_SOURCE="root:@/orm_test?charset=utf8"
	go test -v github.com/jialequ/android-sdk/client/orm


	#### Sqlite3
	export ORM_DRIVER=sqlite3
	export ORM_SOURCE='file:memory_test?mode=memory'
	go test -v github.com/jialequ/android-sdk/client/orm


	#### PostgreSQL
	psql -c 'create database orm_test;' -U postgres
	export ORM_DRIVER=postgres
	export ORM_SOURCE="user=postgres dbname=orm_test sslmode=disable"
	go test -v github.com/jialequ/android-sdk/client/orm

	#### TiDB
	export ORM_DRIVER=tidb
	export ORM_SOURCE='memory://test/test'
	go test -v github.com/jialequ/android-sdk/pgk/orm

	`

func init() {
	// Debug, _ = StrTo(DBARGS.Debug).Bool()
	Debug = true

	if DBARGS.Driver == "" || DBARGS.Source == "" {
		fmt.Println(helpinfo)
		os.Exit(2)
	}

	err := RegisterDataBase("default", DBARGS.Driver, DBARGS.Source, MaxIdleConnections(20))
	if err != nil {
		panic(fmt.Sprintf("can not Register database: %v", err))
	}

	alias := getDbAlias("default")
	if alias.Driver == DRMySQL {
		alias.Engine = "INNODB"
	}
}
