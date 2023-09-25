package infra

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog/log"
)

func NewMinioClient() *minio.Client {
	minioC, err := minio.New(MinIoEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(MinIoID, MinIoSecretKey, ""),
		Secure: MinIoSSL,
	})

	if err != nil {
		log.Fatal().Err(err).Msg("failed open connection minio")
	}

	return minioC
}
