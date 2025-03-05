
// 自动生成模板PayloadInfo
package example
import (
	"time"
)

// payloadInfo表 结构体  PayloadInfo
type PayloadInfo struct {
    Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:10;"`  //id字段
    Name  *string `json:"name" form:"name" gorm:"column:name;comment:插件名称;" binding:"required"`  //插件名称
    Code  *string `json:"code" form:"code" gorm:"column:code;comment:编码;"`  //编码
    Description  *string `json:"description" form:"description" gorm:"column:description;comment:插件描述;"`  //插件描述
    VulType  *string `json:"vulType" form:"vulType" gorm:"column:vulType;comment:漏洞类型;"`  //漏洞类型
    ProType  *string `json:"proType" form:"proType" gorm:"column:proType;comment:产品类型;"`  //产品类型
    Objectives  *string `json:"objectives" form:"objectives" gorm:"column:objectives;comment:测试目标;"`  //测试目标
    Expectations  *string `json:"expectations" form:"expectations" gorm:"column:expectations;comment:测试预期;"`  //测试预期
    Procedure  *string `json:"procedure" form:"procedure" gorm:"column:procedure;comment:测试步骤;"`  //测试步骤
    Hazard  *string `json:"hazard" form:"hazard" gorm:"column:hazard;comment:危害分析;"`  //危害分析
    Level  *int `json:"level" form:"level" gorm:"column:level;comment:危害等级（0低级1中级2高级3超高级）;size:10;" binding:"required"`  //危害等级（0低级1中级2高级3超高级）
    Advise  *string `json:"advise" form:"advise" gorm:"column:advise;comment:修复建议;"`  //修复建议
    UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"column:updateTime;comment:修改时间;"`  //修改时间
}


// TableName payloadInfo表 PayloadInfo自定义表名 payload_info
func (PayloadInfo) TableName() string {
    return "payload_info"
}





