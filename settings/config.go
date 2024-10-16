package settings

type Config struct {
	Port      int       `json:"port"`
	Databases Databases `json:"databases"`
}

type Databases struct {
	NeedMocks bool          `json:"need_mocks"`
	MongoDB   MongoSettings `json:"mongo"`
}

type MongoSettings struct {
	ConnectionString string `json:"connection"`
	Database         string `json:"database"`
	Collection       string `json:"collection"`
}
