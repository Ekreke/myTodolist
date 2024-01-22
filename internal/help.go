package mytodolist

const (
	// recommendedHomeDir 定义放置服务配置的默认目录 maybe linux？
	recommendedHomeDir = ".myTodolist"
	// defaultConfigName 指定了默认配置文件名称
	defaultConfigName = "myTodolist.yaml"
)

// initconfig 设置需要读取的配置文件名、环境变量、并读取配置文件到viper中
// init store 读取db配置，创建gorm.DB 实例，并初始化mytodolist store层
// func initStore() error {
// 	dbOptions := &db.MySQLOptions{}
// }

func ConvertStringToUnixTime() {

}
