package casbinpermit

import (
	dbEngine "blog-backend/pkg/db"
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

type Permit struct {
	*casbin.Enforcer
}

func NewPermit(db *gorm.DB) (*Permit, error) {
	adapter, err := gormadapter.NewAdapterByDBWithCustomTable(db, nil, dbEngine.GetCasbinRuleTableName())
	if err != nil {
		log.Printf("failed to create casbin adapter with err(%s)", err.Error())
		return nil, err
	}

	// 加载model配置
	casbinModel, err := model.NewModelFromString(CasbinConf)
	if err != nil {
		log.Printf("failed to create casbin model with err(%s)", err.Error())
		return nil, err
	}

	enforcer, err := casbin.NewEnforcer(casbinModel, adapter)
	if err != nil {
		log.Printf("failed to create casbin enforcer with err(%s)", err.Error())
		return nil, err
	}

	return &Permit{Enforcer: enforcer}, nil
}

func (p *Permit) CheckPermission(sub, obj, act string) (bool, error) {
	log.Printf("check user permission, sub=[%s] act=[%s] obj=[%s]", sub, act, obj)
	return p.Enforce(sub, obj, act)
}
