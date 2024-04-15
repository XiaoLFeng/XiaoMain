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

// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// XfVerificationCodeDao is the data access object for table xf_verification_code.
type XfVerificationCodeDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns XfVerificationCodeColumns // columns contains all the column names of Table for convenient usage.
}

// XfVerificationCodeColumns defines and stores column names for table xf_verification_code.
type XfVerificationCodeColumns struct {
	Id        string // 主键
	Type      string // 是否为邮箱（邮箱为 true，手机号为 false）
	Contact   string // 联系方式
	Code      string // 验证码
	Scenes    string // 使用场景
	CreatedAt string // 创建时间
	ExpiredAt string // 过期时间
}

// xfVerificationCodeColumns holds the columns for table xf_verification_code.
var xfVerificationCodeColumns = XfVerificationCodeColumns{
	Id:        "id",
	Type:      "type",
	Contact:   "contact",
	Code:      "code",
	Scenes:    "scenes",
	CreatedAt: "created_at",
	ExpiredAt: "expired_at",
}

// NewXfVerificationCodeDao creates and returns a new DAO object for table data access.
func NewXfVerificationCodeDao() *XfVerificationCodeDao {
	return &XfVerificationCodeDao{
		group:   "default",
		table:   "xf_verification_code",
		columns: xfVerificationCodeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *XfVerificationCodeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *XfVerificationCodeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *XfVerificationCodeDao) Columns() XfVerificationCodeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *XfVerificationCodeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *XfVerificationCodeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *XfVerificationCodeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
