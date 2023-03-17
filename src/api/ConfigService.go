package api

import (
	"github.com/AstrayChao/gifnoc-go/src/protocal"
)

type ConfigService interface {
	GetConfig() *protocal.Config
	AddConfig(config *protocal.Config)
	UpdateConfig(config *protocal.Config)
	DeleteConfig(configId int)
}
