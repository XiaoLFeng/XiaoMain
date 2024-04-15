/*
 * --------------------------------------------------------------------------------
 * Copyright (c) 2016-NOW(至今) 筱锋
 * Author: 筱锋(https://www.x-lf.com)
 *
 * 本文件包含 XiaoMain 的源代码，该项目的所有源代码均遵循MIT开源许可证协议。
 * --------------------------------------------------------------------------------
 * 许可证声明：
 *
 * 版权所有 (c) 2016-2024 筱锋。保留所有权利。
 *
 * 本软件是“按原样”提供的，没有任何形式的明示或暗示的保证，包括但不限于
 * 对适销性、特定用途的适用性和非侵权性的暗示保证。在任何情况下，
 * 作者或版权持有人均不承担因软件或软件的使用或其他交易而产生的、
 * 由此引起的或以任何方式与此软件有关的任何索赔、损害或其他责任。
 *
 * 使用本软件即表示您了解此声明并同意其条款。
 *
 * 有关MIT许可证的更多信息，请查看项目根目录下的LICENSE文件或访问：
 * https://opensource.org/licenses/MIT
 * --------------------------------------------------------------------------------
 * 免责声明：
 *
 * 使用本软件的风险由用户自担。作者或版权持有人在法律允许的最大范围内，
 * 对因使用本软件内容而导致的任何直接或间接的损失不承担任何责任。
 * --------------------------------------------------------------------------------
 */

package user

import (
	"context"
	"github.com/gogf/gf/v2/os/glog"
	"xiaoMain/internal/dao"
	"xiaoMain/internal/model/do"
	"xiaoMain/internal/model/entity"
	"xiaoMain/internal/service"
)

type sUserMailLogic struct{}

func init() {
	service.RegisterUserMailLogic(New())
}

func New() *sUserMailLogic {
	return &sUserMailLogic{}
}

// CheckUserMail
// 检查用户输入的邮箱是否与数据库存储的邮箱保持正确，若保持争取的信息将会返回布尔值正确，否则返回错误
func (s *sUserMailLogic) CheckUserMail(ctx context.Context, email string) (checkMail bool, info string) {
	glog.Info(ctx, "[LOGIC] 执行 UserMailLogic:CheckUserMail 服务层")
	// 从数据库获取指定信息
	var getAdminEmail entity.XfIndex
	if dao.XfIndex.Ctx(ctx).Where(do.XfIndex{Key: "email"}).Scan(&getAdminEmail) != nil {
		return false, "未查询到邮箱"
	}
	// 对邮箱进行匹配
	if getAdminEmail.Value == email {
		// 返回正确信息
		return true, "邮箱匹配"
	} else {
		return false, "邮箱不匹配"
	}
}