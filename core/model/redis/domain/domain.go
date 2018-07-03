package domain

import (
	"github.com/go-redis/redis"
	"Stingray/core/model"
)

type Domain struct {
	data_connect *redis.Client
}

func DomainsService() (domain *Domain) {
	domainService := &Domain{}
	domainService.data_connect, _ = model.NewRedisConn().GetSiteListMasterConnect()
	return domainService
}
