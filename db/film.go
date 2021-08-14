package db

import "github.com/quinnj8/rxb-project/models"

func GetAllFilms() []models.Film {
	var films []models.Film

	var film models.Film
	film.Id = 123
	film.Title = "Tron Legacy"
	films = append(films, film)
	//rows, err := db.Conn.Query("SELECT * FROM film ORDER BY film_id DESC")
	//if err != nil {
	//	return films, err
	//}
	//for rows.Next() {
	//	var film models.Film
	//	err := rows.Scan(film.Id, film.Title)
	//	if err != nil {
	//		return films, err
	//	}
	//	films.Films = append(films.Films, film)
	//}
	return films
}
