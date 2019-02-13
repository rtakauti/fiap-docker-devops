package shared

// Modelo de conf de conexao com o db
type DatabaseConf struct {
	DBDriver string
	DBUser   string
	DBPass   string
	DBName   string
}
