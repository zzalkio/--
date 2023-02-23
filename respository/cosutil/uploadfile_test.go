/******
** @创建时间 : 2022/6/21 15:10
** @作者 : MUGUAGAI
******/
package cosutil

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestUploadFile(t *testing.T) {
	CosInit()
	f, _ := os.Open("/public/test")
	fileByte, _ := ioutil.ReadAll(f)
	UploadFile("test", fileByte)
}
