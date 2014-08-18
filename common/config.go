package common

import "go/build"
import "path/filepath"

var prefix = ""

func SetPrefix(prfx string) string {
	prefix = prfx
	return prefix
}
func Prefix() string {
	return prefix
}

func GetRedisHostAndPort(args ...int) (host string, port string) {
	var index int = 0
	if len(args) > 0 {
		if args[0] == -1 {
			return redis_test[0]["host"], redis_test[0]["port"]
		}
		if _, ok := redis[args[0]]; ok {
			index = args[0]
		}
	}
	return redis[index]["host"], redis[index]["port"]
}

func GetPushHostAndCertFilesPath(typ PushType) (host, certPath, keyPath string) {
	projectPath := filepath.Join(build.Default.GOPATH, "src", "github.com/otiai10/push-kcwidget")
	// 以下Apn用とりあえずハードコーディング
	host = "gateway.sandbox.push.apple.com:2195"
	certPath = filepath.Join(projectPath, "service/certfiles", "apple", Prefix()+"cert.pem")
	keyPath = filepath.Join(projectPath, "service/certfiles", "apple", Prefix()+"key.pem")
	return
}
