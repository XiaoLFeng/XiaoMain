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

// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"xiaoMain/internal/model/vo"
)

type (
	IInfoLogic interface {
		// GetMainInfo 获取主要信息
		// 用于获取主要信息，如果返回成功则返回具体的信息，若某些情况下无法获取则获取的内容为空
		// 接口的返回都会有结果，如果返回错误将会返回空值
		//
		// 参数：
		// ctx: 请求的上下文，用于管理超时和取消信号。
		//
		// 返回：
		// getMainInfo: 如果获取成功，返回具体的信息；否则返回空值。
		GetMainInfo(ctx context.Context) *vo.MainVO
		GetBloggerInfo(ctx context.Context) *vo.BloggerVO
	}
)

var (
	localInfoLogic IInfoLogic
)

func InfoLogic() IInfoLogic {
	if localInfoLogic == nil {
		panic("implement not found for interface IInfoLogic, forgot register?")
	}
	return localInfoLogic
}

func RegisterInfoLogic(i IInfoLogic) {
	localInfoLogic = i
}
