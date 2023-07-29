package psql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/dmytrodemianchuk/crud-app/internal/domain"
)

type Musics struct {
	db *sql.DB
}

func NewMusics(db *sql.DB) *Musics {
	return &Musics{db}
}

// create music
func (m *Musics) Create(ctx context.Context, music domain.Music) error {
	_, err := m.db.Exec("INSERT INTO musics (name, artist, album, genre, released_year) values ($1, $2, $3, $4, $5)",
		music.Name, music.Artist, music.Album, music.Genre, music.ReleasedYear)

	return err
}

// get music by id
func (m *Musics) GetByID(ctx context.Context, id int64) (domain.Music, error) {
	var music domain.Music
	err := m.db.QueryRow("SELECT id, name, artist, album, genre, released_year FROM musics WHERE id=$1", id).
		Scan(&music.ID, &music.Name, &music.Artist, &music.Album, &music.Genre, &music.ReleasedYear)
	if err == sql.ErrNoRows {
		return music, domain.ErrorMusicNotFound
	}

	return music, err
}

// get all musics
func (m *Musics) GetAll(ctx context.Context) ([]domain.Music, error) {
	rows, err := m.db.Query("SELECT id, name, artist, album, genre, released_year from musics")
	if err != nil {
		return nil, err
	}

	musics := make([]domain.Music, 0)
	for rows.Next() {
		var music domain.Music
		if err := rows.Scan(&music.ID, &music.Name, &music.Artist, &music.Album, &music.Genre, &music.ReleasedYear); err != nil {
			return nil, err
		}
		musics = append(musics, music)
	}

	return musics, rows.Err()
}

// delete music by id
func (m *Musics) Delete(ctx context.Context, id int64) error {
	_, err := m.db.Exec("DELETE FROM musics WHERE id=$1", id)

	return err
}

// update music
func (m *Musics) Update(ctx context.Context, id int64, inp domain.UpdateMusicInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	// updating element of name
	if inp.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *inp.Name)
		argId++
	}

	// updating element of artist
	if inp.Artist != nil {
		setValues = append(setValues, fmt.Sprintf("artist=$%d", argId))
		args = append(args, *inp.Artist)
		argId++
	}

	// updating element of album
	if inp.Album != nil {
		setValues = append(setValues, fmt.Sprintf("album=$%d", argId))
		args = append(args, *inp.Album)
		argId++
	}

	// updating element of genre
	if inp.Genre != nil {
		setValues = append(setValues, fmt.Sprintf("genre=$%d", argId))
		args = append(args, *inp.Genre)
		argId++
	}

	// updating element of released_year
	if inp.ReleasedYear != nil {
		setValues = append(setValues, fmt.Sprintf("released_year=$%d", argId))
		args = append(args, *inp.ReleasedYear)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE musics set %s WHERE id=$%d", setQuery, argId)
	args = append(args, id)

	_, err := m.db.Exec(query, args...)
	return err
}
