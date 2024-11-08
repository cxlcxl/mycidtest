package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"slices"
	"xiaoniuds.com/cid/app/cid/statement"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/internal/data/base"
	"xiaoniuds.com/cid/internal/data/common"
	"xiaoniuds.com/cid/pkg/auth_token"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/pkg/util"
)

type Service struct {
	C         *config.Config
	DbConnect *data.Data
}

func (s *Service) ZoneDomain(params statement.ZoneDomain) (zone int, err *errs.MyErr) {
	userZone, err := common.NewUserZoneModel("", s.DbConnect).FindUserZoneByEmail(params.Email)
	if err != nil {
		return
	}
	if userZone == nil {
		return -1, errs.Err(errs.LoginEmailNotExist)
	}
	zone = userZone.ZoneIndex
	return
}

func (s *Service) Login(params statement.LoginData) (builder *auth_token.WebToken, err *errs.MyErr) {
	user, err := base.NewUserModel("", s.DbConnect).
		FindUserByLogin(params.Email, util.Password(params.Password, false))
	if err != nil {
		return
	}
	if user.IsLock != 0 {
		return nil, errs.Err(errs.LoginUserExpireError)
	}

	// 判断是否过期
	if user.ParentId == 0 {

	}

	builder = &auth_token.WebToken{
		User: &auth_token.LoginData{
			UserId:         user.UserId,
			ProjectId:      user.ProjectId,
			GroupId:        user.GroupId,
			GroupName:      "",
			Email:          user.Email,
			UserName:       user.UserName,
			UserFullName:   user.UserFullName,
			Mobile:         user.Mobile,
			DataRange:      user.DataRange,
			LatestNews:     user.LatestNews,
			CompanyType:    user.CompanyType,
			Industry:       user.Industry,
			MediaLaunch:    user.MediaLaunch,
			CreateLevel:    user.CreateLevel,
			UsedStatus:     user.UsedStatus,
			MainUserId:     user.MainUserId,
			ContractType:   user.ContractType,
			ProductVersion: params.ProductVersion,
		},
	}
	err = auth_token.CreateJwtToken(builder, s.C.Auth.Login, s.DbConnect)
	return
}

func (s *Service) GetMyAuthorizedUsers(module, moduleRangeType string, userId int64, productVersion, isLock int8) (users []*base.User, err *errs.MyErr) {
	user, err := base.NewUserModel("", s.DbConnect).FindUserById(userId)
	if err != nil {
		return
	}
	// 用户的 GroupId 需要根据版本查询
	adminUserVersionService := AdminUserVersionService{C: s.C, DbConnect: s.DbConnect}
	version, err := adminUserVersionService.GetAdminUserVersionInfo(user.UserId, productVersion, []string{})
	if err != nil {
		return
	}
	user.GroupId = version.GroupId

	dataRange, err := s.GetModuleRange(module, moduleRangeType, &UserModuleRange{UserId: userId}, productVersion)
	if err != nil {
		return
	}
	builder := func(query *gorm.DB) *gorm.DB {
		query = query.Where("is_delete = 0")
		if isLock >= 0 {
			query = query.Where("is_lock = ?", isLock)
		}
		if dataRange == "self" || dataRange == "media_project" { //本人或者项目
			query = query.Where("user_id = ?", user.UserId)
		} else if dataRange == "project" { // 部门所有人
			projectIds := s.GetProjectAllLeafNodeData(nil, user.ProjectId, user.MainUserId, true)
			query = query.Where("project_id in ?", projectIds)
		} else if dataRange == "company" { //公司所有人
			query = query.Where("parent_id = ? or user_id = ?", user.MainUserId, user.MainUserId)
		} else if dataRange == "platform" {
			// 平台所有人无需加其他条件
		} else {
			query = query.Where("1 = 0") //不应该出现这个分支
		}

		return query.Order("is_lock asc, create_time asc")
	}
	users, err = base.NewUserModel("", s.DbConnect).FindUserByQuery(builder, []string{})
	userIds := make([]int64, len(users))
	for i, u := range users {
		userIds[i] = u.UserId
	}
	// 用户的 GroupId 需要根据版本查询
	userVersionInfo, err := adminUserVersionService.GetAdminUserVersionMappingByUserIds(userIds, productVersion, []string{"user_id", "group_id"})
	if err != nil {
		return
	}
	for i, u := range users {
		users[i].GroupId = userVersionInfo[u.UserId][0].GroupId
	}

	return
}

type UserModuleRange struct {
	UserId      int64
	ModuleRange string
}

// GetModuleRange 根据模块名获取权限
func (s *Service) GetModuleRange(module, moduleRangeType string, user *UserModuleRange, productVersion int8) (ret string, err *errs.MyErr) {
	if user == nil {
		return "", errs.Err(errs.ErrMissUserInfo)
	}
	var e error
	if user.ModuleRange == "" {
		if productVersion == -1 {
			return "", errs.Err(errs.SysError, errors.New("获取权限配置失败: 版本信息缺失"))
		}
		var infoAdminUser *common.UserVersion
		infoAdminUser, err = (&AdminUserVersionService{C: s.C, DbConnect: s.DbConnect}).GetAdminUserVersionInfo(user.UserId, productVersion, []string{})
		if err != nil {
			return "", errs.Err(errs.SysError, err)
		}
		if infoAdminUser == nil {
			return "", errs.Err(errs.SysError, errors.New("获取权限配置失败: 用户不存在"))
		}

		var infoAdminGroup *base.UserCustom
		if slices.Contains([]int64{2, 24, 4033, 4100}, infoAdminUser.GroupId) {
			infoAdminGroup, e = common.NewUserCustomModel("", s.DbConnect).FindByGroupId(infoAdminUser.GroupId, []string{"group_id", "module_range"})
			if e != nil {
				return "", errs.Err(errs.SysError, e)
			}
		} else {
			infoAdminGroup, e = base.NewUserCustomModel("", s.DbConnect).FindByGroupId(infoAdminUser.GroupId, []string{"group_id", "module_range"})
			if e != nil {
				return "", errs.Err(errs.SysError, e)
			}
		}
		if infoAdminGroup == nil {
			err = errs.Err(errs.SysError, fmt.Errorf("权限组不存在, user_id: %d, group_id: %d", infoAdminUser.UserId, infoAdminUser.GroupId))
			return
		}
		user.ModuleRange = infoAdminGroup.ModuleRange
	}
	var moduleRanges map[string]*base.ModuleRange
	e = json.Unmarshal([]byte(user.ModuleRange), &moduleRanges)
	if e != nil {
		return "", errs.Err(errs.ErrJsonUnmarshal, e)
	}

	resultModuleRange, ok := moduleRanges[module]
	if !ok {
		resultModuleRange = moduleRanges["default"]
	}

	// 做一个特殊处理，以下账号可以看到企业数据
	// 2024-02-21 请在完成权限控制后，删除该代码
	if slices.Contains([]int64{12000021719, 12000021734, 12000022307, 12000023524}, user.UserId) {
		return "company", nil
	}

	if moduleRangeType == "info_range" {
		return resultModuleRange.InfoRange, nil
	} else {
		return resultModuleRange.DataRange, nil
	}
}

func (s *Service) GetProjectAllLeafNodeData(projects []*base.AdminProject, projectId, mainUserId int64, isNeedParent bool) (projectIds []int64) {
	projectIds = []int64{}
	// 如果没有传入data数据，则自查部门数据
	if len(projects) == 0 {
		// 查询部门的数据
		var err *errs.MyErr
		projects, err = base.NewAdminProjectModel("", s.DbConnect).FindByOwnerUserId(mainUserId, []string{})
		if err != nil {
			return
		}
	}
	projects = s.projectLeafNodeRecursion(projects, projectId)

	// 是否需要把父节点也进行合并
	if isNeedParent {
		projectIds = []int64{projectId}
	}
	for _, row := range projects {
		if row.ProjectId > 0 {
			projectIds = append(projectIds, row.ProjectId)
		}
	}
	return
}

func (s *Service) projectLeafNodeRecursion(projects []*base.AdminProject, projectId int64) []*base.AdminProject {
	var subs []*base.AdminProject
	for _, item := range projects {
		if item.ParentId == projectId {
			subs = append(subs, item)
			subs = append(subs, s.projectLeafNodeRecursion(projects, item.ProjectId)...)
		}
	}
	return subs
}
