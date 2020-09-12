package config

import (
	"fmt"
	"github.com/micro/go-micro/v3/config/source/file"
	"github.com/liujunren93/share_utils/config"
	"github.com/liujunren93/share_utils/config/store"
	"testing"
)

func TestGetConfig(t *testing.T) {
	var AcmOption *store.AcmOptions
	newSource := file.NewSource(
		file.WithPath("./init.yml"),
	)
	microStore, err := store.NewMicroStore(newSource)
	if err != nil {
		panic(err)
	}
	err = config.GetConfig(microStore, &AcmOption,&config.DataOptions{
		Path:    "",
	})
	fmt.Println(AcmOption)
}
