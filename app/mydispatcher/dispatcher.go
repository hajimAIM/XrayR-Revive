// Package mydispatcher Package dispatcher implement the rate limiter and the online device counter
package mydispatcher

//go:generate go run github.com/xtls/xray-core/common/errors/errorgen

import "github.com/xtls/xray-core/features/routing"

// Type returns the feature type token for the custom dispatcher feature.
func Type() interface{} {
	return routing.DispatcherType()
}
