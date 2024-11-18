package user

import (
	"gorm.io/gorm"
	"slices"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/internal/data/common"
	"xiaoniuds.com/cid/pkg/errs"
)

type AdminUserVersionService struct {
	DbConnect *data.Data
}

func (s *AdminUserVersionService) GetAdminUserVersionInfo(userId int64, productVersion int8, fields []string) (version *common.UserVersion, err *errs.MyErr) {
	builder := func(db *gorm.DB) *gorm.DB {
		db = db.Where("is_delete = 0").Where("user_id = ?", userId).Where("product_version = ?", productVersion)
		return db
	}
	return common.NewUserVersionModel("", s.DbConnect).GetAdminUserVersionInfoByBuilder(builder, fields)
}

// GetAdminUserVersionMappingByUserIds 通过用户ID查询获得用户版本列表
func (s *AdminUserVersionService) GetAdminUserVersionMappingByUserIds(userIds []int64, productVersion int8, fields []string) (result map[int64][]*common.UserVersion, err *errs.MyErr) {
	builder := func(query *gorm.DB) *gorm.DB {
		query = query.Where("is_delete = 0").Where("user_id in ?", userIds)

		if productVersion >= 0 && productVersion <= 3 {
			query = query.Where("product_version = ?", productVersion)
		}
		return query
	}
	if !slices.Contains(fields, "user_id") {
		fields = append(fields, "user_id")
	}
	dataList, err := common.NewUserVersionModel("", s.DbConnect).GetAdminUserVersionListByBuilder(builder, fields)
	if err != nil {
		return
	}
	result = make(map[int64][]*common.UserVersion)
	for i, version := range dataList {
		result[version.UserId] = append(result[version.UserId], dataList[i])
	}
	return
}
