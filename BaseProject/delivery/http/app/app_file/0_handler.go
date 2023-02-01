package app_file

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/file_storage/domain/media"
)

type mediaHandler struct {
	mediaUC media.UseCase
}

func NewMediaHandler(mediaUC media.UseCase) media.HttpHandler {
	return &mediaHandler{mediaUC: mediaUC}
}
