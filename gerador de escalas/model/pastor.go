package model

type Pastor struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	IdCreator    int    `json:"id_creator"`
	CreationDate int    `json:"creation_date"`
}
