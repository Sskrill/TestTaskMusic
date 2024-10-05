package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Sskrill/TestTaskMusic/internal/domain"
	"github.com/Sskrill/TestTaskMusic/pkg/customLogger"
	"strings"
	"time"
)

var (
	ErrorSongNotFound = errors.New("Song Not Found")
)

type Repo struct {
	db     *sql.DB
	logger customLogger.CustomLogger
}

func NewRepo(db *sql.DB, logger customLogger.CustomLogger) *Repo {
	return &Repo{db: db, logger: logger}
}

func (r *Repo) GetDetailsSong(nameSong, performerName string) (domain.Song, error) {
	go r.logger.PrintInfo("(DataBase) Finding song")
	var song domain.Song
	err := r.db.QueryRow("SELECT id,song_name,performer_name,link,song_text,release_date,created_at FROM song WHERE song_name=$1 AND performer_name=$2",
		nameSong, performerName).Scan(&song.Id, &song.Name, &song.PerformerName, &song.Link, &song.Text, &song.ReleaseDate, &song.CreatedAt)
	if err != nil {
		go r.logger.PrintError("(DataBase) Error in finding song | error:%s", ErrorSongNotFound.Error())
		return song, ErrorSongNotFound
	}
	go r.logger.PrintInfo("Find song succes")
	return song, nil
}
func (r *Repo) DeleteSong(id int) error {
	go r.logger.PrintInfo("(DataBase) Deleting song")
	_, err := r.db.Exec("DELETE FROM song WHERE id=$1", id)
	if err != nil {
		go r.logger.PrintError("(DataBase) Error in deleting song | error:%s", err.Error())
		return err
	}
	go r.logger.PrintInfo("(DataBase) Deleted song succes")
	return nil
}

func (r *Repo) CreateSong(song domain.Song) error {
	go r.logger.PrintInfo("(DataBase) Creating song")
	_, err := r.db.Exec("INSERT INTO song (song_name,performer_name,link,song_text,release_date,created_at) VALUES($1,$2,$3,$4,$5,$6)",
		song.Name, song.PerformerName, song.Link, song.Text, time.Now(), time.Now())
	if err != nil {
		go r.logger.PrintError("(DataBase) Error in creating song | error:%s", err.Error())
		return err
	}
	return nil
}

func (r *Repo) UpdateSong(id int, song domain.UpdateSong) error {
	go r.logger.PrintInfo("(DataBase) Updating song")

	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM song WHERE id=$1", id).Scan(&count)
	if err != nil || count == 0 {
		go r.logger.PrintError("(DataBase) Error in updating song | error:%s", ErrorSongNotFound.Error())
		return ErrorSongNotFound
	}

	query := "UPDATE song SET "
	params := []interface{}{}
	paramIdx := 1

	if song.Name != "" {
		query += fmt.Sprintf("song_name=$%d, ", paramIdx)
		params = append(params, song.Name)
		paramIdx++
	}
	if song.PerformerName != "" {
		query += fmt.Sprintf("performer_name=$%d, ", paramIdx)
		params = append(params, song.PerformerName)
		paramIdx++
	}
	if song.Link != "" {
		query += fmt.Sprintf("link=$%d, ", paramIdx)
		params = append(params, song.Link)
		paramIdx++
	}
	if song.Text != "" {
		query += fmt.Sprintf("song_text=$%d, ", paramIdx)
		params = append(params, song.Text)
		paramIdx++
	}

	if len(params) == 0 {
		go r.logger.PrintInfo("(DataBase) No updates were needed ")
		return nil
	}

	query = query[:len(query)-2] // убираем лишнюю запятую
	query += fmt.Sprintf(" WHERE id=$%d", paramIdx)
	params = append(params, id)

	_, err = r.db.Exec(query, params...)
	if err != nil {
		go r.logger.PrintError("(DataBase) Error in updating song | error:%s", err.Error())
		return err
	}

	go r.logger.PrintInfo("(DataBase) Updated song succes")
	return nil
}
func (r *Repo) GetSongText(nameSong, performerName string) ([]string, error) {
	go r.logger.PrintInfo("(DataBase) Getting song text ")

	var songText string
	err := r.db.QueryRow("SELECT song_text FROM song WHERE song_name=$1 AND performer_name=$2", nameSong, performerName).Scan(&songText)
	if err != nil {
		go r.logger.PrintError("(DataBase) Song not found | error:%s", err.Error())
		return []string{}, ErrorSongNotFound
	}

	verses := strings.Split(songText, "\n\n")
	go r.logger.PrintInfo("(DataBase) Getting song text succes")
	return verses, nil
}

func (r *Repo) GetSongsWithFilter(songFilters *domain.FiltersForSong) ([]*domain.Song, error) {
	go r.logger.PrintInfo("(DataBase) Getting songs with filter and pagination")

	query := `SELECT id, song_name, performer_name, link, song_text, release_date, created_at FROM song WHERE 1=1`
	var params []interface{}
	paramIdx := 1

	if songFilters.Name != nil {
		query += fmt.Sprintf(" AND song_name ILIKE $%d", paramIdx)
		params = append(params, "%"+*songFilters.Name+"%")
		paramIdx++
	}
	if songFilters.PerformerName != nil {
		query += fmt.Sprintf(" AND performer_name ILIKE $%d", paramIdx)
		params = append(params, "%"+*songFilters.PerformerName+"%")
		paramIdx++
	}
	if songFilters.ReleaseDate != nil {
		query += fmt.Sprintf(" AND release_date = $%d", paramIdx)
		params = append(params, *songFilters.ReleaseDate)
		paramIdx++
	}
	if songFilters.Text != nil {
		query += fmt.Sprintf(" AND song_text ILIKE $%d", paramIdx)
		params = append(params, "%"+*songFilters.Text+"%")
		paramIdx++
	}
	if songFilters.Link != nil {
		query += fmt.Sprintf(" AND link ILIKE $%d", paramIdx)
		params = append(params, "%"+*songFilters.Link+"%")
		paramIdx++
	}

	if songFilters.Limit != nil {
		query += fmt.Sprintf(" LIMIT %d", *songFilters.Limit)
	}
	if songFilters.Offset != nil {
		query += fmt.Sprintf(" OFFSET %d", *songFilters.Offset)
	}
	rows, err := r.db.Query(query, params...)
	if err != nil {
		go r.logger.PrintError("(DataBase) Error in getting songs | error: %v", err)
		return nil, err
	}
	defer rows.Close()
	var songs []*domain.Song
	for rows.Next() {
		var song domain.Song
		if err := rows.Scan(&song.Id, &song.Name, &song.PerformerName, &song.Link, &song.Text, &song.ReleaseDate, &song.CreatedAt); err != nil {
			go r.logger.PrintError("(DataBase) Error in scanning song | error: %v", err)
			return nil, err
		}
		songs = append(songs, &song)
	}

	if err := rows.Err(); err != nil {
		go r.logger.PrintError("(DataBase) Error during row iteration | error: %v", err)
		return nil, err
	}
	go r.logger.PrintInfo("(DataBase) Successfully retrieved %d songs", len(songs))
	return songs, nil
}
