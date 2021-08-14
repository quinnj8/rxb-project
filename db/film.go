package db

import "github.com/quinnj8/rxb-project/Models"

func (db Database) GetAllFilms() (Models.Films, error) {
	films := Models.Films{}
	rows, err := db.Conn.Query("SELECT * FROM film ORDER BY film_id DESC")
	if err != nil {
		return films, err
	}
	for rows.Next() {
		var film Models.Film
		err := rows.Scan(film.Id, film.Title)
		if err != nil {
			return films, err
		}
		films.Films = append(films.Films, film)
	}
	return films, nil
}
