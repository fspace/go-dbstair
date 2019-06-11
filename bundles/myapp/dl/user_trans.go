package dl

import "dbstair/bundles/myapp/domain/users"

// UserTrans 实现PEE 中的事务脚本
type UserTrans struct {
	UsersGateway users.Gateway // 事务脚本 可能依赖多个db 网关    即一个事务可能跨多个表
}
