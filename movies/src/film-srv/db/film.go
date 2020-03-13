package db

import "hongbao/choujiang/models"


// 获取正在上映的电影
func SelectTickingFilims(status int64) ([]*models.Film, error) {
	films := []*models.Film{}
	err := db.Select(&films, "SELECT `rating_final`,`title_cn`,`is_3D`,`is_DMAX`,`is_IMAX`,`is_IMAX3D`,`img`,`movie_id`,`film_director`,`film_drama` FROM `film` WHERE `is_ticking` = ?", status)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return films, err
}

// 获取影片详情
func SelectFilmDetail(movieId int64) (*models.Film, error) {
	film := models.Film{}
	err := db.Get(&film, "SELECT * FROM `film` WHERE `movie_id` = ?", movieId)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &film, err
}





