package domain

import (
	"github.com/go-redis/redis"
	"Stingray/core/model"
)

type Domain struct {
	data_connect *redis.Client
}

func DomainsService() (domain *Domain) {
	return &Domain{}
}

func (d *Domain) Init() {
	d.data_connect, _ = model.NewRedisConn().GetSiteListMasterConnect()
}
