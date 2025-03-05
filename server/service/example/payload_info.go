
package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
    exampleReq "github.com/flipped-aurora/gin-vue-admin/server/model/example/request"
)

type PayloadInfoService struct {}
// CreatePayloadInfo 创建payloadInfo表记录
// Author [yourname](https://github.com/yourname)
func (payloadInfoService *PayloadInfoService) CreatePayloadInfo(payloadInfo *example.PayloadInfo) (err error) {
	err = global.GVA_DB.Create(payloadInfo).Error
	return err
}

// DeletePayloadInfo 删除payloadInfo表记录
// Author [yourname](https://github.com/yourname)
func (payloadInfoService *PayloadInfoService)DeletePayloadInfo(id string) (err error) {
	err = global.GVA_DB.Delete(&example.PayloadInfo{},"id = ?",id).Error
	return err
}

// DeletePayloadInfoByIds 批量删除payloadInfo表记录
// Author [yourname](https://github.com/yourname)
func (payloadInfoService *PayloadInfoService)DeletePayloadInfoByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]example.PayloadInfo{},"id in ?",ids).Error
	return err
}

// UpdatePayloadInfo 更新payloadInfo表记录
// Author [yourname](https://github.com/yourname)
func (payloadInfoService *PayloadInfoService)UpdatePayloadInfo(payloadInfo example.PayloadInfo) (err error) {
	err = global.GVA_DB.Model(&example.PayloadInfo{}).Where("id = ?",payloadInfo.Id).Updates(&payloadInfo).Error
	return err
}

// GetPayloadInfo 根据id获取payloadInfo表记录
// Author [yourname](https://github.com/yourname)
func (payloadInfoService *PayloadInfoService)GetPayloadInfo(id string) (payloadInfo example.PayloadInfo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&payloadInfo).Error
	return
}
// GetPayloadInfoInfoList 分页获取payloadInfo表记录
// Author [yourname](https://github.com/yourname)
func (payloadInfoService *PayloadInfoService)GetPayloadInfoInfoList(info exampleReq.PayloadInfoSearch) (list []example.PayloadInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&example.PayloadInfo{})
    var payloadInfos []example.PayloadInfo
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name LIKE ?","%"+*info.Name+"%")
    }
    if info.VulType != nil && *info.VulType != "" {
        db = db.Where("vulType LIKE ?","%"+*info.VulType+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["id"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&payloadInfos).Error
	return  payloadInfos, total, err
}
func (payloadInfoService *PayloadInfoService)GetPayloadInfoPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
