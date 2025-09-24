// service/minio_service.go
package service

const (
	MinioEndpoint = "http://localhost:9000"
	BucketName    = "cardsandromeda"
)

type MinioService struct{}

func NewMinioService() *MinioService {
	return &MinioService{}
}

func (s *MinioService) GetImageURL(imageName string) string {
	// Генерируем URL к изображению в Minio
	return MinioEndpoint + "/" + BucketName + "/" + imageName + ".jpg"
}
