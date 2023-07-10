package actuator_deprecated

import (
	"path"
	"testing"
)

var (
	userId = "89757"
)

func TestCreateDirectory(t *testing.T) {
	err := CreateDirectory(userId)
	if err != nil {
		t.Fatal("create file failed")
	}
}
func TestUpLoadInputFile(t *testing.T) {
	fileName := "../data/pie-data.xls"
	remotePath, err := UpLoadFile(userId, fileName)
	if err != nil {
		t.Fatal("upload file failed")
	}
	t.Logf("remote file path: %s ", remotePath)
}

func TestPaint(t *testing.T) {
	cmd := "pie"
	fileName := "pie-data.xls"
	inPath := path.Join(BasePath, userId, fileName)
	outPath := path.Join(BasePath, userId, fileName+"-out")

	err := Paint(cmd, inPath, outPath)
	if err != nil {
		t.Fatal("paint failed")
	}
}

func TestSquashFile(t *testing.T) {

	baseDir := "/public/analysis/wuqy/tmp/89757/"
	fileName := "pie-data.xls-out"

	file, err := SquashFile(baseDir, fileName)
	if err != nil {
		t.Fatal("squash failed")
	}
	t.Logf("new file is %s\n", file)
}

func TestDownLoadFile(t *testing.T) {
	fileName := "/public/analysis/wuqy/tmp/89757/pie-data.xls-out.tar.gz"
	localPath := "../data"
	err := DownLoadFile(fileName, localPath)
	if err != nil {
		t.Fatal("download file failed")
	}
}
