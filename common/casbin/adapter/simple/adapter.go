package simple

import (
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
)

type Adapter struct {
	data [][]string
}

func NewAdapter(data [][]string) *Adapter {
	return &Adapter{data: data}
}

// LoadPolicy loads all policy rules from the storage.
func (mp *Adapter) LoadPolicy(model model.Model) error {

	// data := strings.Split(mp.data, "\r\n")
	for _, line := range mp.data {
		persist.LoadPolicyArray(line, model)
	}
	return nil
}

// SavePolicy saves all policy rules to the storage.
func (mp *Adapter) SavePolicy(model model.Model) error {
	panic("not implemented") // TODO: Implement
}

// AddPolicy adds a policy rule to the storage.
// This is part of the Auto-Save feature.
func (mp *Adapter) AddPolicy(sec string, ptype string, rule []string) error {
	panic("not implemented") // TODO: Implement
}

// RemovePolicy removes a policy rule from the storage.
// This is part of the Auto-Save feature.
func (mp *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
	panic("not implemented") // TODO: Implement
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
// This is part of the Auto-Save feature.
func (mp *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	panic("not implemented") // TODO: Implement
}
