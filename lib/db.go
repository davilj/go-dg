package dg

import (
  //"fmt"
)


func Load(modules []XMLivyModule) DB {
  db := DB{make(map[string][]string),  make(map[string]XMLivyModule)}
  for _, module := range(modules) {
      info:=module.Info
      name:=info.Module
      rev:=info.Rev
      key:= name + "|" + rev
      db.DbInfo[key]=module
      ds:=module.Dependencies.Dependency
      dsArray := make([]string,0)
      for _, d := range(ds) {
          dName:=d.Name
          dRev:=d.Rev
          dKey:= dName + "|" + dRev
          dsArray=append(dsArray,dKey)
      }
      db.DbLink[key]=dsArray
  }
  return db
}
