package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/spf13/viper"
	"vvvstore/internal/pkg/database"
)

var Enforcer *casbin.Enforcer

// 初始化
func InitCasbin() (*casbin.Enforcer, error) {
	m := model.NewModel()
	m.AddDef("r", "r", "sub, obj, act") // 定义请求参数的含义
	m.AddDef("p", "p", "sub, obj, act") // 定义策略规则
	m.AddDef("e", "e", "some(where (p.eft == allow))") // 定义生效范围
	//m.AddDef("m", "m", "r.sub == p.sub && r.obj == p.obj && r.act == p.act")
	m.AddDef("m", "m", `r.sub == p.sub && keyMatch(r.obj, p.obj) && (r.act == p.act || p.act == "*") || r.sub == "root"`) // 定义匹配规则

	a, err := adapter.NewAdapterByDBUsePrefix(database.GetDatabase(), viper.GetString("database.prefix"))
	if err != nil {
		return nil, err
	}

	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		return nil, err
	}

	e.LoadPolicy() // 加载规则

	Enforcer = e

	return e, nil
}

// 检测权限
func Enforce(e *casbin.Enforcer, sub, obj, act string) (bool, error) {
	return e.Enforce(sub, obj, act)
}
