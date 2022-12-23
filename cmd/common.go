package cmd

import (
	"context"
	"fmt"

	"github.com/klaudiusz-czapla/my-cloud-home-go/mch"
)

type ContextKey string

const ContextProxyKey = ContextKey("proxy")

func GetProxyFromContext(context context.Context) (*mch.MchProxy, error) {
	contextProxyValue := context.Value(ContextProxyKey)

	if contextProxyValue != nil {
		proxy, ok := contextProxyValue.(*mch.MchProxy)
		if !ok {
			return nil, fmt.Errorf("invalid type: pointer to proxy object expected")
		}

		return proxy, nil
	}

	return nil, nil
}
