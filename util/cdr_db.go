package util

//CdrDB this struct represet the Database of each CDR table
type CdrDB struct {
	ProviderName string
	IPAddress    string
	Port         int
	Username     string
	Password     string
}
