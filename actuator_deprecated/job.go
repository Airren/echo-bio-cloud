package actuator_deprecated

import (
	"log"
	"path"

	"github.com/airren/echo-bio-backend/model"
)

const LocalDataBase = "./static"

func StartPaint(job model.AnalysisJob) (outFile string, err error) {

	userId := job.AccountId
	localDataPath := path.Join(LocalDataBase, userId)
	// create directory for user on remote work server
	err = CreateDirectory(userId)
	if err != nil {
		log.Fatalf("create directory on remote server for user: %s failed", userId)
		return
	}

	//upload input file
	fileLocalPath := path.Join(localDataPath, job.Name)
	remoteFilePath, err := UpLoadFile(userId, fileLocalPath)
	if err != nil {
		log.Fatalf("upload input file failed, user: %s, file: %s", userId, job.Name)
		return
	}

	// execute the command to generate the chart
	remoteOutPath := remoteFilePath + "-out"
	err = Paint(job.Algorithm, remoteFilePath, remoteOutPath)
	if err != nil {
		log.Fatalf("exec job %v failed %s", job.Id, err)
		return
	}

	remoteSquashFile, err := SquashFile(path.Join(BasePath, userId), job.Name+"-out")
	if err != nil {
		log.Fatalf("job %v squash failed %v", job.Id, err)
	}
	err = DownLoadFile(remoteSquashFile, localDataPath)
	if err != nil {
		log.Fatalf("job  %v download failed %v", job.Id, err)
		return
	}

	// todo refactor
	outFile = job.Name + "-out.tar.gz"

	return
}
