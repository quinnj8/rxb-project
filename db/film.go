package db

import (
	"github.com/quinnj8/rxb-project/models"
)

var films = []models.Film{
	{Id: 1, Title: "WizardofOz", Rating: "8", Category: "Black and White"},
	{Id: 2, Title: "HarryPotter", Rating: "10", Category: "Adventure"},
	{Id: 3, Title: "LordoftheRings", Rating: "9", Category: "Adventure"},
	{Id: 4, Title: "KarateKid", Rating: "7", Category: "80s"},
	{Id: 5, Title: "Grinch", Rating: "5", Category: "Holiday"},
}

func GetAllFilms() []models.Film {
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

func GetFilmByTitle(title string) []models.Film {

	var result []models.Film

	for _, film := range films {
		if film.Title == title {
			result = append(result, film)
		}
	}
	return result
}

func GetFilmByRating(rating string) []models.Film {

	var result []models.Film

	for _, film := range films {
		if film.Rating == rating {
			result = append(result, film)
		}
	}
	return result
}

func GetFilmByCategory(category string) []models.Film {

	var result []models.Film

	for _, film := range films {
		if film.Category == category {
			result = append(result, film)
		}
	}
	return result
}
