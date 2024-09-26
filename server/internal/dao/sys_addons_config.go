// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"hotgo/internal/dao/internal"
)

// internalSysAddonsConfigDao is internal type for wrapping internal DAO implements.
type internalSysAddonsConfigDao = *internal.SysAddonsConfigDao

// sysAddonsConfigDao is the data access object for table hg_sys_addons_config.
// You can define custom methods on it to extend its functionality as you wish.
type sysAddonsConfigDao struct {
	internalSysAddonsConfigDao
}

var (
	// SysAddonsConfig is globally public accessible object for table hg_sys_addons_config operations.
	SysAddonsConfig = sysAddonsConfigDao{
		internal.NewSysAddonsConfigDao(),
	}
)

// Fill with you ideas below.
