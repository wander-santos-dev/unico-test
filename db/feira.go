package db

import (
	"database/sql"

	"github.com/teste-unico/models"
)

func (db Database) GetAllFeiras() (*models.FeiraList, error) {
	list := &models.FeiraList{}
	rows, err := db.Conn.Query("SELECT * FROM feiras ORDER BY ID DESC")

	if err != nil {
		return list, err
	}

	for rows.Next() {
		var feira models.Feira
		err := rows.Scan(&feira.Long, &feira.Lat, &feira.Setcens, &feira.Areap, &feira.Coddist, &feira.Distrito, &feira.Codsubpref, &feira.Subprefe, &feira.Regiao5, &feira.Regiao8, &feira.Nome_feira, &feira.Registro, &feira.Logradouro, &feira.Numero, &feira.Bairro, &feira.Referencia, &feira.Created_at)

		if err != nil {
			return list, err
		}

		list.Feiras = append(list.Feiras, feira)
	}

	return list, nil
}

func (db Database) AddFeira(feira models.Feira) error {
	var id int
	var createdAt string

	query := `INSERT INTO feiras (
		long,
		lat,
		setcens,
		areap,
		coddist,
		distrito,
		codsubpref,
		subprefe,
		regiao5,
		regiao8,
		nome_feira,
		registro,
		logradouro,
		numero,
		bairro,
		referencia,
		created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17) RETURNING id, created_at`

	err := db.Conn.QueryRow(query, feira.Long, feira.Lat, feira.Setcens, feira.Areap, feira.Coddist, feira.Distrito, feira.Codsubpref, feira.Subprefe, feira.Regiao5, feira.Regiao8, feira.Nome_feira, feira.Registro, feira.Logradouro, feira.Numero, feira.Bairro, feira.Referencia).Scan(&id, &createdAt)

	if err != nil {
		return err
	}

	feira.ID = id
	feira.Created_at = createdAt
	return nil
}

func (db Database) GetFeiraById(feiraId int) (models.Feira, error) {
	feira := models.Feira{}
	query := `SELECT * FROM feiras WHERE id = $1;`
	row := db.Conn.QueryRow(query, feiraId)

	switch err := row.Scan(&feira.Long, &feira.Lat, &feira.Setcens, &feira.Areap, &feira.Coddist, &feira.Distrito, &feira.Codsubpref, &feira.Subprefe, &feira.Regiao5, &feira.Regiao8, &feira.Nome_feira, &feira.Registro, &feira.Logradouro, &feira.Numero, &feira.Bairro, &feira.Referencia, &feira.Created_at); err {
	case sql.ErrNoRows:
		return feira, ErrNoMatch
	default:
		return feira, err
	}
}

func (db Database) DeleteFeira(feiraId int) error {
	query := `DELETE FROM feiras WHERE id = $1;`
	_, err := db.Conn.Exec(query, feiraId)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}

func (db Database) UpdateFeira(feiraId int, feiraData models.Feira) (models.Feira, error) {
	feira := models.Feira{}
	query := `UPDATE feiras SET
		long=$1,
		lat=$2,
		setcens=$3,
		areap=$4,
		coddist=$5,
		distrito=$6,
		codsubpref=$7,
		subprefe=$8,
		regiao5=$9,
		regiao8=$10,
		nome_feira=$11,
		registro=$12,
		logradouro=$13,
		numero=$14,
		bairro=$15,
		referencia=$16  WHERE id = $17 RETURNING
		id,
		long,
		lat,
		setcens,
		areap,
		coddist,
		distrito,
		codsubpref,
		subprefe,
		regiao5,
		regiao8,
		nome_feira,
		registro,
		logradouro,
		numero,
		bairro,
		referencia,
		created_at;`

	err := db.Conn.QueryRow(query, feiraData.Long, feiraData.Lat, feiraData.Setcens, feiraData.Areap, feiraData.Coddist, feiraData.Distrito, feiraData.Codsubpref, feiraData.Subprefe, feiraData.Regiao5, feiraData.Regiao8, feiraData.Nome_feira, feiraData.Registro, feiraData.Logradouro, feiraData.Numero, feiraData.Bairro, feiraData.Referencia, feiraId).Scan(
		&feira.ID, &feira.Long, &feira.Lat, &feira.Setcens, &feira.Areap, &feira.Coddist, &feira.Distrito, &feira.Codsubpref, &feira.Subprefe, &feira.Regiao5, &feira.Regiao8, &feira.Nome_feira, &feira.Registro, &feira.Logradouro, &feira.Numero, &feira.Bairro, &feira.Referencia, &feira.Created_at,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return feira, ErrNoMatch
		}
		return feira, err
	}

	return feira, nil
}
