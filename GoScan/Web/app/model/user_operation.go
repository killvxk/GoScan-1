// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"github.com/CTF-MissFeng/GoScan/Web/app/model/internal"
)

// UserOperation is the golang structure for table user_operation.
type UserOperation internal.UserOperation

// Fill with you ideas below.

// 用户登录日志管理 模糊分页查询返回数据所需信息
type UserOperationRespLogins struct{
	Code int `json:"code"`
	Msg string `json:"msg"`
	Count int64 `json:"count"`
	Data [] *UserLog `json:"data"`
}

// 用户操作日志管理 模糊分页查询返回数据所需信息
type UserOperationRespLogs struct{
	Code int `json:"code"`
	Msg string `json:"msg"`
	Count int64 `json:"count"`
	Data [] *UserOperation `json:"data"`
}