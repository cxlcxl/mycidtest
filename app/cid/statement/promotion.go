package statement

import (
	"time"
	"xiaoniuds.com/cid/pkg/auth_token"
)

type PddGoodsList struct {
	*Pagination
	LoginData *auth_token.LoginData

	Keyword            string    `form:"keyword"`
	Note               string    `form:"note"`
	ListType           int64     `form:"list_type"`
	ReportStatus       uint8     `form:"report_status"`
	IsPromotionLimited uint8     `form:"is_promotion_limited"`
	OwnerNoteStatus    uint8     `form:"owner_note_status"`
	MallIds            []int64   `form:"mall_ids"`
	OwnerUserId        int64     `form:"owner_user_id"`
	CreateTime         time.Time `form:"create_time"`
}

type JdGoodsList struct {
	*Pagination
	LoginData *auth_token.LoginData
}

type TbGoodsList struct {
	*Pagination
	LoginData *auth_token.LoginData
}
