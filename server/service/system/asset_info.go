package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

type AssetService struct{}

func (assetService *AssetService) GetAssetInfoList(info systemReq.GetAssetList) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//db := global.GVA_DB.Model(&system.SysUser{})
	db := global.GVA_DB.Model(&system.AssetInfo{})
	var assetList []system.AssetInfo

	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.SystemCode != "" {
		db = db.Where("systemCode LIKE ?", "%"+info.SystemCode+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	//err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&assetList).Error
	err = db.Debug().Limit(limit).Offset(offset).Find(&assetList).Error
	return assetList, total, err
}
