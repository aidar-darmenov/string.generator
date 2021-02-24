package service

import (
	"github.com/aidar-darmenov/string.generator/db"
	"github.com/aidar-darmenov/string.generator/model"
)

type Service struct {
	Configuration *model.Configuration
	Manager       *db.Manager
	ChannelString chan model.StringElement
	ChannelFiller chan model.HashedStringElement
}

func NewCoreManager(
	//dbm *db.Manager,
	config *model.Configuration,
) *Service {
	coreManager := Service{
		Configuration: config,
		//Manager:       dbm,
		ChannelString: make(chan model.StringElement),
		ChannelFiller: make(chan model.HashedStringElement),
	}
	return &coreManager
}
