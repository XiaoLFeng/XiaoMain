// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"develop/internal/dao/internal"
)

// internalXfIndexDao is internal type for wrapping internal DAO implements.
type internalXfIndexDao = *internal.XfIndexDao

// xfIndexDao is the data access object for table xf_index.
// You can define custom methods on it to extend its functionality as you wish.
type xfIndexDao struct {
	internalXfIndexDao
}

var (
	// XfIndex is globally public accessible object for table xf_index operations.
	XfIndex = xfIndexDao{
		internal.NewXfIndexDao(),
	}
)

// Fill with you ideas below.
