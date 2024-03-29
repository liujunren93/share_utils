package main

import (
	"context"
	"flag"
	"io"
	"os"

	"github.com/liujunren93/share_utils/app"
)

var fileName string

func init() {

	flag.StringVar(&fileName, "f", "./conf/config.json", "data file")
	flag.Parse()

}

func main() {

	op, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer op.Close()

	data, err := io.ReadAll(op)
	ap := app.NewApp(nil, nil)

	_, err = ap.Cloud.PublishConfig(context.TODO(), ap.GetLocalConfig().GetLocalBase().ConfCenter.ConfName, ap.GetLocalConfig().GetLocalBase().ConfCenter.Group, string(data))
	if err != nil {
		panic(err)
	}
}
