package statement

import "time"

type Pagination struct {
	Page     int `json:"page" form:"page" binding:"required,numeric"`
	PageSize int `json:"page_size" form:"page_size" binding:"required,numeric"`
}

type VDate time.Time
type VDateTime time.Time

func (d *VDate) UnmarshalJSON(data []byte) (err error) {
	// 空值不进行解析
	if len(data) == 2 {
		*d = VDate(time.Time{})
		return
	}

	// 指定解析的格式
	now, err := time.Parse(time.DateOnly, string(data))
	*d = VDate(now)

	return
}

func (d VDate) MarshalJSON() (data []byte, err error) {
	b := make([]byte, 0, len(time.DateOnly)+2)
	b = append(b, '"')
	b = time.Time(d).AppendFormat(b, time.DateOnly)
	b = append(b, '"')
	return b, nil
}

func (d *VDateTime) UnmarshalJSON(data []byte) (err error) {
	return nil
}
