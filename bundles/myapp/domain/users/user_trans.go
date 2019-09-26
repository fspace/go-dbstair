package users

// UserTrans 实现PEE 中的事务脚本  封装领域逻辑  此处事务并非专指sql中的事务
// 事务脚本 类似yii中的controller 还有一种实现是 每类型一事务（相当于很多action 类)
// 控制器有点偏web功能 针对传统mvc结构此处确实是类似控制器  但对于java系 此处类似 application-service层
// 如果采用“每类型一事务"风格 目录结构变为 domain/users/transactions/{register_user.go, update_user.go}
type UserTrans struct {
	UsersGateway users.Gateway // 事务脚本 可能依赖多个db 网关    即一个事务可能跨多个表
}
