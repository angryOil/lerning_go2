package main

import (
	"context"
	"net/http"
)

func main() {

}

type logger interface {
	Log(context.Context, string)
}

type requestDeco func(r *http.Request) *http.Request

type businessLogic struct {
	RequestDeco requestDeco
	Logger      logger
	Remote      string
}

func (bl businessLogic) businessLogic(ctx context.Context, user string, data string) (string, error) {
	bl.Logger.Log(ctx, "tarting businessLogic for"+user+"with "+data)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, bl.Remote+"?query="+data, nil)
	if err != nil {
		bl.Logger.Log(ctx, "error building remote request: "+err.Error())
		return "", err
	}
	req = bl.RequestDeco(req)
	return "", nil
}
