package service

import (
	"go_gin_pragmatic/entity"
	"go_gin_pragmatic/repository"
)

type VideoService interface {
	// FindAll() []entity.Video
	// Save(entity.Video) entity.Video
	Save(entity.Video) error
	Update(entity.Video) error
	Delete(entity.Video) error
	FindAll() []entity.Video
}

type videoService struct {
	// videos []entity.Video
	repository repository.VideoRepository
}

func New(videoRepository repository.VideoRepository) VideoService {
	return &videoService{
		repository: videoRepository,
	}
}

func (service *videoService) FindAll() []entity.Video {
	// return service.videos
	return service.repository.FindAll()
}

func (service *videoService) Save(video entity.Video) error {
	// service.videos = append(service.videos, video)
	service.repository.Save(video)
	return nil
}

func (service *videoService) Update(video entity.Video) error {
	service.repository.Update(video)
	return nil
}

func (service *videoService) Delete(video entity.Video) error {
	service.repository.Delete(video)
	return nil
}
