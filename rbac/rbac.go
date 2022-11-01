package rbac

import "github.com/casbin/casbin/v2"

type RBACService struct {
	*casbin.Enforcer
}

func New(enforcer *casbin.Enforcer) (*RBACService, error) {
	return &RBACService{enforcer}, nil
}
