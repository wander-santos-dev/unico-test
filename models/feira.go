package models

import (
	"fmt"
	"net/http"
)

type Feira struct {
	ID         int    `json:"id"`
	Long       string `json:"long"`
	Lat        string `json:"lat"`
	Setcens    string `json:"setcens"`
	Areap      string `json:"areap"`
	Coddist    string `json:"coddist"`
	Distrito   string `json:"distrito"`
	Codsubpref string `json:"codsubpref"`
	Subprefe   string `json:"subprefe"`
	Regiao5    string `json:"regiao5"`
	Regiao8    string `json:"regiao8"`
	Nome_feira string `json:"nome_feira"`
	Registro   string `json:"registro"`
	Logradouro string `json:"logradouro"`
	Numero     string `json:"numero"`
	Bairro     string `json:"bairro"`
	Referencia string `json:"referencia"`
	Created_at string `json:"created_at"`
}

type FeiraList struct {
	Feiras []Feira `json:"feiras"`
}

func (i *Feira) Bind(r *http.Request) error {
	if i.Nome_feira == "" {
		return fmt.Errorf("Nome feira is a required field!")
	}

	return nil
}

func (*FeiraList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Feira) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
