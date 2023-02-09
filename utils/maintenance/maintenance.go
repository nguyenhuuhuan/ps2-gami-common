package maintenance

import (
	"context"
	"time"

	"gitlab.id.vin/gami/gami-common/adapters/cache"
	"gitlab.id.vin/gami/gami-common/adapters/gami_service"
	"gitlab.id.vin/gami/gami-common/logger"
	"gitlab.id.vin/gami/gami-common/models"
)

const (
	MidTerm             = 100 * time.Millisecond
	MaintainCode        = 4000005
	ExpirationCacheTime = 1 * time.Minute
	cacheKeyPrefix      = "gms_maintenance:"
)

func CheckMaintenance(ctxReq context.Context, campaignType models.CampaignType, localCacheAdapter cache.LocalCacheAdapter,
	gamiAdapter gami_service.Adapter) bool {
	var (
		maintenanceStatus bool
		cacheKey          = cacheKeyPrefix + campaignType.String()
		ctx, cancel       = context.WithDeadline(ctxReq, time.Now().Add(MidTerm))
	)

	defer cancel()

	err := localCacheAdapter.Get(ctx, cacheKey, &maintenanceStatus)
	if err != nil {
		resp, err := gamiAdapter.MaintenanceStatusCampaign(ctx, gami_service.MaintenanceStatusCampaignRequest{
			CampaignType: campaignType.String(),
		})
		if err != nil {
			logger.Context(ctx).Errorf("[CheckMaintenance] call gami_service failed, err: %v", err)
			return false
		}

		if resp != nil {
			if resp.Meta.Code == MaintainCode {
				err = localCacheAdapter.Set(ctx, cacheKey, true, ExpirationCacheTime)
				if err != nil {
					logger.Context(ctx).Errorf("[CheckMaintenance] set local cache failed, code: %v, key: %v, err: %v",
						resp.Meta.Code, cacheKey, err)
				}

				return true
			} else {
				err = localCacheAdapter.Set(ctx, cacheKey, false, ExpirationCacheTime)
				if err != nil {
					logger.Context(ctx).Errorf("[CheckMaintenance] set local cache failed, code: %v, key: %v, err: %v",
						resp.Meta.Code, cacheKey, err)
				}

				return false
			}
		}
	}

	if maintenanceStatus {
		resp, err := gamiAdapter.MaintenanceStatusCampaign(ctx, gami_service.MaintenanceStatusCampaignRequest{
			CampaignType: campaignType.String(),
		})
		if err != nil {
			logger.Context(ctx).Errorf("[CheckMaintenance][2] call gami_service failed, err: %v", err)
		} else if resp != nil && resp.Meta.Code != MaintainCode {
			err = localCacheAdapter.Set(ctx, cacheKey, false, ExpirationCacheTime)
			if err != nil {
				logger.Context(ctx).Errorf("[CheckMaintenance][2] set local cache failed, code: %v, key: %v, err: %v",
					resp.Meta.Code, cacheKey, err)
			}

			return false
		}

		return true
	}

	return false
}
