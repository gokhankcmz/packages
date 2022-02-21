package KvpConverter

import (
	"Packages/src/pkg/KVP/KVP"
	"Packages/src/pkg/KVP/Object"
)

type Settings struct {
	HierarchySeparator string
	Prefix             string
}

func New(HierarchySeparator, Prefix string) *Settings {
	return &Settings{
		HierarchySeparator: HierarchySeparator,
		Prefix:             Prefix,
	}
}

func (s Settings) GetKVP(Object interface{}) map[string]string {
	return KVP.GetKVPs(Object, s.HierarchySeparator, s.Prefix, map[string]string{})
}

func (s Settings) GetObject(Obj interface{}, Map map[string]string) {
	Object.GetObject(Obj, s.HierarchySeparator, s.Prefix, Map)
}
