package config

//configStruct 文件中配置结构
type configStruct struct {
	Port    int
	Debug   bool
	Logpath string

	Db *dbStruct
}

//dbStruct db配置
type dbStruct struct {
	Mysqlurl     string
	Mysqlidle    int
	Mysqlmaxopen int
}
