package comm

type Config struct {
	Server struct {
		Host     string `yaml:"host"` //外网访问地址
		RunLimit int    `yaml:"runLimit"`
		HbtpHost string `yaml:"hbtpHost"`
		Secret   string `yaml:"secret"`
	} `yaml:"server"`
	Database struct {
		Driver string `yaml:"driver"`
		Url    string `yaml:"url"`
	} `yaml:"database"`
}
