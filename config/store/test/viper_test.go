package test

import (
	"fmt"
	"github.com/liujunren93/share_utils/config"
	"github.com/liujunren93/share_utils/config/store"
	"testing"
)

func Test_sViper_GetConfig(t *testing.T) {
	var acm store.AcmOptions
	options := config.DataOptions{
		Path: "./",
		//FileType: "yaml",
		FileName: "init.yml",
	}
	v := store.NewViperStore(&options)
	err := config.GetConfig(v, &acm, nil)

fmt.Println(err,acm.NamespaceID)
}
