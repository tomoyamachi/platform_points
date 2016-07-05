package model

import (
	"github.com/gocraft/dbr"
)

type MPoint struct {
	Id         int64  `json:"id"`
	Code       string `json:"code"`
	PointLabel string `json:"point_label"`
	UnitLabel  string `json:"unit_label"`
	Default    string `json:"default"`
	Max        string `json:"max"`
}

func (m *MPoint) Load(tx *dbr.Tx, id int64) error {
	return tx.Select("*").
		From("m_point").
		Where("id = ?", id).
		LoadStruct(m)
}

type MPoints []MPoint

func (m *MPoints) Load(tx *dbr.Tx, code string) error {
	var condition dbr.Condition
	if code != "" {
		condition = dbr.Eq("code", code)
	}

	return tx.Select("*").
		From("m_point").
		Where(condition).
		LoadStruct(m)
}
