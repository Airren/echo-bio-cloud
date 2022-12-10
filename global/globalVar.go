package global

import (
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
)

var MinioClient *minio.Client
var Logger *zap.Logger
