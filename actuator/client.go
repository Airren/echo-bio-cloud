package actuator

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/airren/echo-bio-backend/utils"
)

const (
	BasePath  = "/public/analysis/wuqy/tmp/echo-bio-cloud"
	SshClient = "-p 5868 wuqy@210.22.109.162"
)

var (
	// <sshClient> perl <cmd> -i <inputFile> -o <outPutFile>
	paint      = "%s perl %s -i %s -o %s"
	createPath = "%s mkdir -p %s"
	squashFile = "%s tar -C %s -czvf %s %s"
	upload     = "-P 5868 %s wuqy@210.22.109.162:%s"
	download   = "-P 5868 wuqy@210.22.109.162:%s %s"
)

var commandMap = map[string]string{
	"pie": "/public/analysis/wuqy/Scripts/16S/pie_v2.pl",
}

func CreateDirectory(userId string) (err error) {
	args := fmt.Sprintf(createPath, SshClient, path.Join(BasePath, userId))
	return utils.Exec("ssh", os.Stdout, splitArgs(args)...)

}

func UpLoadFile(userId, fileName string) (remoteFilePath string, err error) {
	_, file := path.Split(fileName)
	remoteFilePath = path.Join(BasePath, userId, file)
	args := fmt.Sprintf(upload, fileName, path.Join(BasePath, userId))
	return remoteFilePath, utils.Exec("scp", os.Stdout, splitArgs(args)...)
}

func DownLoadFile(fileName, localPath string) error {
	args := fmt.Sprintf(download, fileName, localPath)
	return utils.Exec("scp", os.Stdout, splitArgs(args)...)

}

func Paint(cmd, inPath, outPath string) (err error) {
	cmdPath, ok := commandMap[cmd]
	if !ok {
		return errors.New("command not found")
	}
	args := fmt.Sprintf(paint, SshClient, cmdPath, inPath, outPath)
	return utils.Exec("ssh", os.Stdout, splitArgs(args)...)
}

func SquashFile(baseDir, fileName string) (target string, err error) {
	target = path.Join(baseDir, fileName+".tar.gz")
	args := fmt.Sprintf(squashFile, SshClient, baseDir, target, fileName)
	err = utils.Exec("ssh", os.Stdout, splitArgs(args)...)
	return target, err
}

func splitArgs(argStr string) (args []string) {
	return strings.Split(argStr, " ")
}
