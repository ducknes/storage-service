package settings

type Config struct {
	Port      int       `json:"port"`
	Databases Databases `json:"databases"`
}

type Databases struct {
	NeedMocks bool          `json:"need_mocks"`
	MongoDB   MongoSettings `json:"mongo"`
	Redis     RedisSettings `json:"redis"`
}

type MongoSettings struct {
	ConnectionString string `json:"connection"`
	Database         string `json:"database"`
	Collection       string `json:"collection"`
}

type RedisSettings struct {
	Address  string `json:"address"`
	Password string `json:"password"`
	Database int    `json:"database"`
}
