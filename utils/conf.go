package utils

import (
	"os"
	"path/filepath"
	"github.com/Unknwon/goconfig"
)

var Cfg *goconfig.ConfigFile
func init(){
	dir :="./conf"
	Cfg = new(goconfig.ConfigFile)

	filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f == nil  {return err}
		if f.IsDir() {return nil}
		err = Cfg.AppendFiles(path)
		if err != nil {
			panic("load config error: " + err.Error())
		}
		return nil
	})

}