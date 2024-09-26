// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"hotgo/internal/dao/internal"
)

// internalSysLogDao is internal type for wrapping internal DAO implements.
type internalSysLogDao = *internal.SysLogDao

// sysLogDao is the data access object for table hg_sys_log.
// You can define custom methods on it to extend its functionality as you wish.
type sysLogDao struct {
	internalSysLogDao
}

var (
	// SysLog is globally public accessible object for table hg_sys_log operations.
	SysLog = sysLogDao{
		internal.NewSysLogDao(),
	}
)

// Fill with you ideas below.
