package base

import (
	"errors"
	"gorm.io/gorm"
	"time"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/vars"
)

type ACToken struct {
	ID             int64     `json:"id"`
	MainUserId     int64     `json:"main_user_id"`
	UserId         int64     `json:"user_id"`
	IP             string    `json:"ip"`
	Scopes         string    `json:"scopes"`
	TokenType      string    `json:"token_type"`
	AccessToken    string    `json:"access_token"`
	AccessTokenMD5 string    `json:"access_token_md5"`
	RefreshToken   string    `json:"refresh_token"`
	ExpireTime     time.Time `json:"expire_time"`
	CreateTime     time.Time `json:"create_time"`
	UpdateTime     time.Time `json:"update_time"`
}

type ACTokenModel struct {
	dbName string
	db     *gorm.DB
}

func NewACTokenModel(connect string, connects *data.Data) *ACTokenModel {
	if connect == "" {
		connect = vars.DRCidTest
	}
	return &ACTokenModel{
		dbName: "access_token",
		db:     connects.DbConnects[connect],
	}
}

func (m *ACTokenModel) QueryByBuilder(builder data.QueryBuilder, fields []string) (list []*ACToken, err *errs.MyErr) {
	query := m.db.Table(m.dbName).Where("is_delete = 0")
	if len(fields) > 0 {
		query = query.Select(fields)
	}
	query = builder(query)
	e := query.Find(&list).Error
	if e != nil {
		return nil, errs.Err(errs.SysError, e)
	}
	return
}

func (m *ACTokenModel) GetOneByBuilder(builder data.QueryBuilder, fields []string) (one *ACToken, err *errs.MyErr) {
	query := m.db.Table(m.dbName)
	if len(fields) > 0 {
		query = query.Select(fields)
	}
	query = builder(query)
	e := query.Order("id desc").First(&one).Error
	if e != nil {
		if errors.Is(e, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errs.Err(errs.SysError, e)
	}
	return
}

func (m *ACTokenModel) Save(token *ACToken) (err *errs.MyErr) {
	e := m.db.Transaction(func(tx *gorm.DB) (e error) {
		// 一个号登一个IP
		e = tx.Exec("DELETE FROM access_token WHERE main_user_id = ? AND user_id = ? AND token_type = ?", token.MainUserId, token.UserId, token.TokenType).Error
		if e != nil {
			return
		}
		e = tx.Table(m.dbName).Create(token).Error
		if e != nil {
			return
		}
		return
	})
	if e != nil {
		err = errs.Err(errs.SysError, e)
	}
	return
}
