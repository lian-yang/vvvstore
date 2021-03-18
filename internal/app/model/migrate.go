package model

import "errors"

// 迁移
type Migrate struct {

}

// 安装
func (m *Migrate) Install() error {
	return nil
}

// 升级
func (m *Migrate) Upgrade(CurVersionID int) error {
	if CurVersionID == 100 {
		return errors.New("v1.0.0版本不支持升级")
	}
	return nil
}

// 升级到1.0.1版本
func (m *Migrate) upgradeFor101() error {
	return nil
}