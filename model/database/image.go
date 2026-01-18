package database

import (
	"server/global"
	"server/model/appTypes"
)

type Image struct {
	global.MODEL
	Name     string            `json:"name"`
	URL      string            `json:"url"`
	CateGory appTypes.Category `json:"cate_gory"`
	Storage  appTypes.Storage  `json:"storage"`
}
