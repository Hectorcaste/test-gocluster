package database

import "os"

func NewConnection() *LealDBAdapter {
	conn, err := ClientMariaGormDTB(
		os.Getenv("DB_USER_SASS"),
		os.Getenv("DB_PASSWORD_SASS"),
		os.Getenv("DB_HOST_SASS"),
		os.Getenv("DB_NAME_SASS"),
	)
	if err != nil {
		return nil
	}

	return NewLealDBAdapter(conn)
}
