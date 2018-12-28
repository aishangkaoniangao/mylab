# 开发相关
```
./native/consts.go
const (
    MYSQL_TYPE_DECIMAL     = 0x00
    MYSQL_TYPE_TINY        = 0x01 // int8, uint8, bool
    MYSQL_TYPE_SHORT       = 0x02 // int16, uint16
    MYSQL_TYPE_LONG        = 0x03 // int32, uint32
    MYSQL_TYPE_FLOAT       = 0x04 // float32
    MYSQL_TYPE_DOUBLE      = 0x05 // float64
    MYSQL_TYPE_NULL        = 0x06 // nil
    MYSQL_TYPE_TIMESTAMP   = 0x07 // Timestamp
    MYSQL_TYPE_LONGLONG    = 0x08 // int64, uint64
    MYSQL_TYPE_INT24       = 0x09
    MYSQL_TYPE_DATE        = 0x0a // Date
    MYSQL_TYPE_TIME        = 0x0b // Time
    MYSQL_TYPE_DATETIME    = 0x0c // time.Time
    MYSQL_TYPE_YEAR        = 0x0d
    MYSQL_TYPE_NEWDATE     = 0x0e
    MYSQL_TYPE_VARCHAR     = 0x0f
    MYSQL_TYPE_BIT         = 0x10
    MYSQL_TYPE_NEWDECIMAL  = 0xf6
    MYSQL_TYPE_ENUM        = 0xf7
    MYSQL_TYPE_SET         = 0xf8
    MYSQL_TYPE_TINY_BLOB   = 0xf9
    MYSQL_TYPE_MEDIUM_BLOB = 0xfa
    MYSQL_TYPE_LONG_BLOB   = 0xfb
    MYSQL_TYPE_BLOB        = 0xfc // Blob
    MYSQL_TYPE_VAR_STRING  = 0xfd // []byte
    MYSQL_TYPE_STRING      = 0xfe // string
    MYSQL_TYPE_GEOMETRY    = 0xff

    MYSQL_UNSIGNED_MASK = uint16(1 << 15) 
)

rows, res, err := db.Query("select * from user where id>=%d", 1)
res.Fields()返回的是Mysql.Field结构
Mysql.Field结构参考：https://godoc.org/github.com/ziutek/mymysql/mysql#Field
Filed结构中有个字段是Type，它对应的是这个字段的类型是什么，这个类型字典见如上的字典
字典链接：https://godoc.org/github.com/ziutek/mymysql/native#pkg-constants
```

# TODO

+ 1, get后的key-value组织这里并没有做得很自动;
+ 2, 还未曾尝试transaction操作；