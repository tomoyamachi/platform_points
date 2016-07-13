package model

import (
	"github.com/gocraft/dbr"
)

type MPoint struct {
	Id         int64  `json:"id"`
	Code       string `json:"code"`
	PointLabel string `json:"point_label"`
	UnitLabel  string `json:"unit_label"`
	Default    int64  `json:"default"`
	Max        int64  `json:"max"`
}

func (m *MPoint) Load(tx *dbr.Tx, id int64) error {
	return tx.Select("*").
		From("m_point").
		Where("id = ?", id).
		LoadStruct(m)
}

type MPoints []MPoint

func (m *MPoints) Load(tx *dbr.Tx) error {
	return tx.Select("*").
		From("m_point").
		LoadStruct(m)
}
