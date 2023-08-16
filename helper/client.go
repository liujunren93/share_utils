package helper

import (
	"context"
	"fmt"
	"strings"

	"github.com/liujunren93/share/client"
	"github.com/liujunren93/share_utils/errors"
)

func Invoke(client *client.Client, ctx context.Context, appName string, method string, req, res interface{}) errors.Error {
	names := strings.Split(appName, "_")
	cc, err := client.Client(appName)
	if err != nil {
		return errors.NewInternalError("Invoke init client error:" + err.Error())

	}

	err = client.Invoke(ctx, fmt.Sprintf("/%s.%s/%s", names[2], strings.ToUpper(names[2][:1])+names[2][1:], strings.ToUpper(method[:1])+method[1:]), req, res, cc)
	if err != nil {
		return errors.NewInternalError(err)
	}
	return nil
}
