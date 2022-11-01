package rbac

import (
	"fmt"

	"github.com/casbin/casbin/v2"
)

type RBACService struct {
	*casbin.Enforcer
}

func New(enforcer *casbin.Enforcer) (*RBACService, error) {
	return &RBACService{enforcer}, nil
}

func (r *RBACService) HasAccessToBlock(user, block, metric string) (bool, error) {
	if !r.HasAccessToMetric(user, metric) {
		return false, fmt.Errorf("you don't have access to underlying metric")
	}

	result, err := r.Enforce(user, block, "read")
	return (err == nil && result), err
}

func (r *RBACService) HasAccessToMetric(user, metric string) bool {
	result, err := r.Enforce(user, metric, "read")
	return err == nil && result
}
