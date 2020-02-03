package constant

const LOCAL = "local"
const DEV = "dev"
const PROD = "prod"
var Host = map[string]string{"local": "127.0.0.1", "dev": "127.0.0.1"}
var Port = map[string]string{"local": "3306", "dev": "3306"}
var UserName = map[string]string{"local": "medici", "dev": "medici"}
var Password = map[string]string{"local": "medici_local", "dev": "medici_dev"}
const DB_MASTER = "medici_master"
const DB_USER = "medici_user"