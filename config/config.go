package config

import (
  "github.com/gogf/gf/frame/g"
  "github.com/gogf/gf/os/gcfg"
  "github.com/guozhe001/ex-api-go/constant"
)

var config *gcfg.Config

func init() {
  config = g.Config()
  config.SetFileName(constant.ConfigApiKey)
}

func Get(table, key string) string {
  return config.GetMapStrStr(table)[key]
}
