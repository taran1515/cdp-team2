package services

import (
	"sync"

	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/internal/dao"
	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/util"
)

type AddRewardService interface {
	ProcessRequest()
}

var addRewardServiceStruct AddRewardService
var addRewardServiceOnce sync.Once

type addRewardService struct {
	config *util.RouterConfig
	dao    dao.DynamoDAO
}

func InitAddRewardService(config *util.RouterConfig, dao dao.DynamoDAO) AddRewardService {
	addRewardServiceOnce.Do(func() {
		addRewardServiceStruct = &addRewardService{
			config: config,
			dao:    dao,
		}
	})

	return addRewardServiceStruct
}

func GetAddRewardService() AddRewardService {
	if addRewardServiceStruct == nil {
		panic("add reward service not initialized")
	}

	return addRewardServiceStruct
}

func (s *addRewardService) ProcessRequest() {
}
