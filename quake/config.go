/**------------------------------------------------------------**
 * @filename core/config.go
 * @author   jinycoo - admin@jinycoo.com
 * @version  1.0.0
 * @date     2021/8/17 16:35
 * @desc     go-quake - config setting
 **------------------------------------------------------------**/
package quake

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/jinycoo/go-quake/core/toml"
)

type Config struct {
	AppName string
	Mode    string
	Quake   *Quake
}

type Quake struct {
	BaseUrl string
	ApiKey  string
}

func NewConfig() *Config {
	cfg := &Config{}
	confPath := filepath.Join(getCurrentAbPath(), "config.toml")
	if _, err := toml.DecodeFile(confPath, &cfg); err != nil {
		log.Fatalln(err)

	}
	return cfg
}

func getCurrentAbPath() string {
	dir := getCurrentAbPathByExecutable()
	tmpDir, _ := filepath.EvalSymlinks(os.TempDir())
	if strings.Contains(dir, tmpDir) {
		return getCurrentAbPathByCaller()
	}
	return dir
}

// 获取当前执行文件绝对路径
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
		abPath = strings.Replace(abPath, "/quake", "/config.toml", 1)
	}
	return abPath
}
