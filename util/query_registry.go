package util

import (
	"github.com/go-kid/gok/registry"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"log"
)

func QueryInRegistry(name string) (string, error) {
	bytes := registry.Registry
	src := gjson.GetBytes(bytes, name)
	if src.Exists() {
		return src.String(), nil
	}
	var registries = []string{
		"https://raw.githubusercontent.com/go-kid/go-kid-ctl/refs/heads/main/registry/default.json",
	}
	for _, reg := range registries {
		res, err := resty.New().R().Get(reg)
		if err != nil {
			log.Printf("find %s in %s error: %v", name, reg, err)
			continue
		}
		if res.IsError() {
			log.Printf("find %s in %s failed: %s", name, reg, res.String())
			continue
		}
		src := gjson.GetBytes(res.Body(), name)
		if src.Exists() {
			return src.String(), nil
		}
	}
	return "", errors.Errorf("%s not found", name)
}
