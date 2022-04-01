package utils

import "os"

// DataForPath DataForPath
func DataForPath(path string) ([]byte, error) {

	file, err := os.Open(path)
	if err != nil {
		return []byte(""), nil
	}
	defer file.Close()

	// 读取文件内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)

	return buf, nil
}
