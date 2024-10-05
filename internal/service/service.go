package service

import (
	"errors"
	"github.com/Sskrill/TestTaskMusic/internal/domain"
	"github.com/Sskrill/TestTaskMusic/pkg/customLogger"
)

var ErrorNullValueOfParams = errors.New("Null Value Of Params")

type MusicRepository interface {
	GetDetailsSong(nameSong, performerName string) (domain.Song, error)
	DeleteSong(id int) error
	UpdateSong(id int, song domain.UpdateSong) error
	CreateSong(song domain.Song) error
	GetSongText(nameSong, performerName string) ([]string, error)
	GetSongsWithFilter(songFilters *domain.FiltersForSong) ([]*domain.Song, error)
}
type Service struct {
	logger    customLogger.CustomLogger
	MusicRepo MusicRepository
}

func NewService(repository MusicRepository, logger customLogger.CustomLogger) *Service {
	return &Service{MusicRepo: repository, logger: logger}
}

func (s *Service) AddSong(song domain.Song) error {
	go s.logger.PrintInfo("(Service) Adding song")
	err := s.MusicRepo.CreateSong(song)
	if err != nil {
		go s.logger.PrintError("(Service) Adding song | error:%s", err.Error())
		return err
	}
	go s.logger.PrintInfo("(Service) Added song succes")
	return nil
}
func (s *Service) EditSong(id int, song domain.UpdateSong) error {
	go s.logger.PrintInfo("(Service) Updating song")
	err := s.MusicRepo.UpdateSong(id, song)
	if err != nil {
		go s.logger.PrintError("(Service) Updating song | error:%s", err.Error())
		return err
	}
	go s.logger.PrintInfo("(Service) Updated song succes")
	return nil
}
func (s *Service) GetSongDetails(songName, performerName string) (domain.Song, error) {
	go s.logger.PrintInfo("(Service) Getting song details")
	if songName == "" || performerName == "" {
		go s.logger.PrintError("(Service) Null Parameters error:%s", ErrorNullValueOfParams.Error())
		return domain.Song{}, ErrorNullValueOfParams
	}
	song, err := s.MusicRepo.GetDetailsSong(songName, performerName)
	if err != nil {
		go s.logger.PrintError("(Service) Getting song details | error:%s", err.Error())
		return song, err
	}
	go s.logger.PrintInfo("(Service) Getting song details succes")
	return song, nil
}
func (s *Service) DeleteSong(id int) error {
	go s.logger.PrintInfo("(Service) Deleting song")
	err := s.MusicRepo.DeleteSong(id)
	if err != nil {
		go s.logger.PrintError("(Service) Deleting song | error:%s", err.Error())
		return err
	}
	go s.logger.PrintInfo("(Service) Deleted song succes")
	return nil
}

func (s *Service) GetSongText(nameSong, performerName string) ([]string, error) {
	go s.logger.PrintInfo("(Service) Getting song text")
	songText, err := s.MusicRepo.GetSongText(nameSong, performerName)
	if err != nil {
		go s.logger.PrintError("(Service) Getting song text | error:%s", err.Error())
		return nil, err
	}
	go s.logger.PrintInfo("(Service) Getting song text succes")
	return songText, nil
}
func (s *Service) GetSongsByFilters(songFilters *domain.FiltersForSong) ([]*domain.Song, error) {
	go s.logger.PrintInfo("(Service) Getting songs by filters")
	songs, err := s.MusicRepo.GetSongsWithFilter(songFilters)
	if err != nil {
		go s.logger.PrintError("(Service) Getting songs by filters | error:%s", err.Error())
		return nil, err
	}
	return songs, nil
}
