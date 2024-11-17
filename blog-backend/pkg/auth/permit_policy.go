package auth

import log "github.com/sirupsen/logrus"

// 为用户添加权限
func (a *Auth) AddSubRoles(subRoles []string, username string) error {
	for _, r := range subRoles {
		if _, err := a.casbinPermit.AddGroupingPolicy(username, r); err != nil {
			log.Errorf("add group policy(username=%s, role=%s) failed with err(%s)",
				username, r, err.Error())
			return err
		}
	}
	return nil
}

// 为用户删除权限
func (a *Auth) DelSubRoles(username string) error {
	roles, err := a.casbinPermit.GetRolesForUser(username)
	if err != nil {
		log.Errorf("get roles for user(%s) failed with err(%s)", username, err.Error())
		return err
	}

	for _, r := range roles {
		if _, err := a.casbinPermit.RemoveGroupingPolicy(username, r); err != nil {
			log.Errorf("remove group policy(username=%s, role=%s) failed with err(%s)",
				username, r, err.Error())
			return err
		}
	}

	return nil
}
