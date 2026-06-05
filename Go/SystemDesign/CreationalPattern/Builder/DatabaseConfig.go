package builder


type DatabaseConfig struct{
	Host string
	Port int
	Username string
	Password string
	MaxConn int
}

type DatabaseConfigBuilder struct{
	config DatabaseConfig
}

func NewDatabaseConfigBuilder() *DatabaseConfigBuilder{

	return &DatabaseConfigBuilder{
		config: DatabaseConfig{
			Port: 5342,
		},
	}
}

func(c *DatabaseConfigBuilder) HostName(hostName string) *DatabaseConfigBuilder{
	c.config.Host = hostName;
	return c;
}


func(c *DatabaseConfigBuilder) Username(userName string) *DatabaseConfigBuilder{
	c.config.Username = userName
	return c
}

func(c *DatabaseConfigBuilder) Password(password string) *DatabaseConfigBuilder{
	c.config.Password = password
	return c
}

func(c *DatabaseConfigBuilder) MaxConn(maxconn int) *DatabaseConfigBuilder{
	c.config.MaxConn = maxconn
	return c
}

func (c *DatabaseConfigBuilder) Build() *DatabaseConfig {
	return &c.config
}


// func main(){

// 	config:= NewDatabaseConfigBuilder().HostName("localhost").Username("psql").Password("Password").MaxConn(100).Build()

// 	fmt.Println(config)


// }