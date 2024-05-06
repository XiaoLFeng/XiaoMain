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

// XfLinkRssDao is the data access object for table xf_link_rss.
type XfLinkRssDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns XfLinkRssColumns // columns contains all the column names of Table for convenient usage.
}

// XfLinkRssColumns defines and stores column names for table xf_link_rss.
type XfLinkRssColumns struct {
	LinkId    string // 链接 id
	RssJson   string // Rss 内容解析
	CheckTime string // 检查时间
}

// xfLinkRssColumns holds the columns for table xf_link_rss.
var xfLinkRssColumns = XfLinkRssColumns{
	LinkId:    "link_id",
	RssJson:   "rss_json",
	CheckTime: "check_time",
}

// NewXfLinkRssDao creates and returns a new DAO object for table data access.
func NewXfLinkRssDao() *XfLinkRssDao {
	return &XfLinkRssDao{
		group:   "default",
		table:   "xf_link_rss",
		columns: xfLinkRssColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *XfLinkRssDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *XfLinkRssDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *XfLinkRssDao) Columns() XfLinkRssColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *XfLinkRssDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *XfLinkRssDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *XfLinkRssDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
