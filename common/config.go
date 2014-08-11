package common

func GetRedisHostAndPort(args ...int) (host string, port string) {
	var index int = 0
	if len(args) > 0 {
		if _, ok := redis[args[0]]; ok {
			index = args[0]
		}
	}
	return redis[index]["host"], redis[index]["port"]
}
