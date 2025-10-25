package client

import (
	"errors"
)

const (
	LocalConfig int = iota
	CloudConfig
)

func NewConfigClient(configType int) (ServerConfigClient, error) {
	switch configType {
	case LocalConfig:
		return newLocalServerConfigClient(), nil
	case CloudConfig:
		return newCloudServerConfigClient(), nil
	}
	return nil, errors.New("invalid config type")
}
