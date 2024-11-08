package statement

import (
	"xiaoniuds.com/cid/pkg/auth_token"
)

type PddGoodsList struct {
	LoginData *auth_token.LoginData

	Keyword            string  `form:"keyword"`
	Note               string  `form:"note"`
	ReportStatus       int8    `form:"report_status"`
	IsPromotionLimited int8    `form:"is_promotion_limited"`
	OwnerNoteStatus    int8    `form:"owner_note_status"`
	MallIds            []int64 `form:"mall_ids"`
	OwnerUserId        int64   `form:"owner_user_id"`
	DeliveryStatus     int8    `form:"delivery_status"`
	TTAdvertiser       string  `form:"tt_advertiser"`
	KSAdvertiser       string  `form:"ks_advertiser"`
	GDTAdvertiser      string  `form:"gdt_advertiser"`
	SortField          string  `form:"sort_field"`
	SortDirection      string  `form:"sort_direction"`
	StartDate          string  `form:"start_date,date"`
	EndDate            string  `form:"end_date,date"`
	Page               int     `json:"page" form:"page" binding:"required,numeric"`
	PageSize           int     `json:"page_size" form:"page_size" binding:"required,numeric"`
}

type JdGoodsList struct {
	*Pagination
	LoginData *auth_token.LoginData
}

type TbGoodsList struct {
	*Pagination
	LoginData *auth_token.LoginData
}
