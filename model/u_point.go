package model

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gocraft/dbr"
)

type UPoint struct {
	Id        int64  `json:"id"`
	AccountId int64  `json:"account_id"`
	MPointId  int64  `json:"m_point_id"`
	Value     int64  `json:"value"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// 指定したユーザーにポイント情報がない場合に呼び出す
// 各m_pointデータの初期値を入れる
func NewUPoint(account_id int64, mpoint *MPoint) *UPoint {
	return &UPoint{
		AccountId: account_id,
		MPointId:  mpoint.Id,
		Value:     mpoint.Default,
		CreatedAt: time.Now().Format("2000-01-01 00:00:00"),
		UpdatedAt: time.Now().Format("2000-01-01 00:00:00"),
	}
}

// データを保存する
func (p *UPoint) Save(tx *dbr.Tx) error {
	_, err := tx.InsertInto("u_point").
		Columns("account_id", "m_point_id", "value", "created_at", "updated_at").
		Record(p).
		Exec()

	return err
}

type UPoints []UPoint

func (p *UPoints) LoadTargetAccountPoints(tx *dbr.Tx, account_id int64) error {
	logrus.Debug(account_id)
	return tx.Select("*").
		From("u_point").
		Where("account_id = ?", account_id).
		LoadStruct(p)
}

func (p *UPoints) Load(tx *dbr.Tx) error {
	return tx.Select("*").
		From("u_point").
		LoadStruct(p)
}
