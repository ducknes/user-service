package settings

type Config struct {
	Port      int      `json:"port"`
	Databases Database `json:"databases"`
}

type Database struct {
	MongoDB MongoSettings `json:"mongo"`
}

type MongoSettings struct {
	ConnectionString string `json:"connection"`
	Database         string `json:"database"`
	Collection       string `json:"collection"`
}
