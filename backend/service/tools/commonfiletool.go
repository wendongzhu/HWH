package tool

import (
	"fmt"
	"github.com/wendongzhu/hwh/backend/service/param"
	"io/ioutil"
	"os"
	"strings"
)

// PathExists 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		// 创建文件夹
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			//Logger("sys").Error().Msg("Mkdir file path failed! err: " + err.Error())
		} else {
			//Logger("sys").Info().Msg("Mkdir file path <" + path + "> success!")
			return true, nil
		}
	}

	return false, err
}

func ReadFileName(filePath, fileType string) (res param.FileNameListParam) {
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		fmt.Println("ReadFileName Error:", err)
		return
	}
	for _, file := range files {
		fileName := strings.Fields(strings.Replace(file.Name(), ".", " ", -1))
		if fileName[1] == fileType {
			res.Data = append(res.Data, param.FileNameParam{Value: file.Name(), Label: fileName[0]})
		}
	}
	return res
}
