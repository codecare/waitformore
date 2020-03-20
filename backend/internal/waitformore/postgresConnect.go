package waitformore

type DbConnect struct {
	Hostname string
	Port     string
	User     string
	Password string
	Database string
}

func MakeDbConnect() DbConnect {

	hostname := ReadFromEnv("db_hostname", "localhost", false)
	port := ReadFromEnv("db_port", "5432", false)
	user := ReadFromEnv("db_user", "wait4more", false)
	password := ReadFromEnv("db_password", "zCpzQ7VF7%GFmTM", true)
	database := ReadFromEnv("db_database", "wait4more", false)

	return DbConnect{
		Hostname: hostname,
		Port:     port,
		User:     user,
		Password: password,
		Database: database,
	}
}
