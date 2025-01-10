package xmysql

const invalidMysqlPoolConfigProvided = "Invalid MySQL pool config provided"

type InvalidMySQLPoolConfigProvided struct {
	extraItems map[string]interface{}
}

func (u InvalidMySQLPoolConfigProvided) Error() string {
	return invalidMysqlPoolConfigProvided
}

func (u InvalidMySQLPoolConfigProvided) ExtraItems() map[string]interface{} {
	return u.extraItems
}

func NewInvalidMysqlPoolConfigProvided() *InvalidMySQLPoolConfigProvided {
	return &InvalidMySQLPoolConfigProvided{extraItems: map[string]interface{}{}}
}
