package model

import (
	"strconv"
	"strings"

	"github.com/Biubiubiuuuu/yuepuwebsite/db/mysql"
)

// PropertyInfo 物业信息
type PropertyInfo struct {
	Model
	IndustryID     int64           `json:"industry_id"`                                                                          // 经营业态ID
	Title          string          `gorm:"not null;size:255;" json:"title"`                                                      // 标题
	Nickname       string          `gorm:"size:50;" json:"nickname"`                                                             // 联系人
	Telephone      string          `gorm:"not null;size:30;" json:"telephone"`                                                   // 联系手机
	ShopName       string          `gorm:"size:100;" json:"shop_name"`                                                           // 店名（后台录入）
	Image          string          `gorm:"size:200;" json:"image"`                                                               // 图片
	Video          string          `json:"video"`                                                                                // 视频（后台录入）
	BusType        int64           `gorm:"size:1;" json:"bus_type"`                                                              // 业务类型（后台录入）0-商铺 ｜ 1-写字楼 ｜ 2-厂房仓库
	ModelType      int64           `gorm:"size:1;" json:"model_type"`                                                            // 模型类型（后台录入）0-转让 ｜ 1-出售 ｜ 3-出租 | 4-求租 ｜ 5-求购
	ProvinceCode   string          `gorm:"size:10;" json:"province_code"`                                                        // 省代码
	CityCode       string          `gorm:"size:10;" json:"city_code"`                                                            // 城市代码
	DistrictCode   string          `gorm:"size:10;" json:"district_code"`                                                        // 区代码
	StreetCode     string          `gorm:"size:10;" json:"street_code"`                                                          // 街道代码
	Address        string          `json:"address"`                                                                              // 详细地址
	StoreTypeID    int64           `json:"store_type_id"`                                                                        // 店铺类型ID
	AreaTypeID     int64           `json:"area_type_id"`                                                                         // 面积分类ID （后台录入或者自动判断）
	RentTypeID     int64           `json:"rent_type_id"`                                                                         // 租金分类ID （后台录入或者自动判断）
	MinArea        float64         `json:"min_area"`                                                                             // 最小面积（单位：平方米）
	MaxArea        float64         `json:"max_area"`                                                                             // 最大面积（单位：平方米）
	MinRent        float64         `json:"min_rent"`                                                                             // 最低租金（单位：元/月）
	MaxRent        float64         `json:"max_rent"`                                                                             // 最高租金（单位：元/月）
	Lots           []Lot           `gorm:"foreignkey:PropertyInfoID;association_foreignkey:ID" json:"lots,omitempty"`            // 考虑地段
	Idling         bool            `json:"idling"`                                                                               // 可否空转
	InOperation    int64           `json:"in_operation"`                                                                         // 是否营业中 0-新铺 ｜ 1-空置中 ｜ 2-营业中
	Area           float64         `json:"area"`                                                                                 // 面积（单位：平方米）
	Rent           float64         `json:"rent"`                                                                                 // 租金（单位：元/月）
	TransferFee    float64         `json:"transfer_fee"`                                                                         // 转让费用（单位：万元 不录入则前台显示为面议）
	IndustryRanges []IndustryRange `gorm:"foreignkey:PropertyInfoID;association_foreignkey:ID" json:"industry_ranges,omitempty"` // 适合经营范围
	Description    string          `json:"description"`                                                                          // 详细描述
	ExplicitTel    bool            `json:"explicit_tel"`                                                                         // 是否外显号码 true：客户号码 ｜ false：发布者号码
	Tel1           string          `gorm:"size:30;" json:"tel1"`                                                                 // 外显号码1
	Tel2           string          `gorm:"size:30;" json:"tel2"`                                                                 // 外显号码2
	Audit          bool            `json:"audit"`                                                                                // 是否审核 true：已审核 ｜ false：待审核 （后台录入）
	AuditID        int64           `json:"audit_id"`                                                                             // 审核人ID （后台录入）
	Protect        bool            `json:"protect"`                                                                              // 是否保护 true：已保护 ｜ false：未保护 （后台录入）
	QuotedPrice    float64         `json:"quoted_price"`                                                                         // 报价（后台录入，保护时显示）
	Pictures       []Picture       `gorm:"foreignkey:PropertyInfoID;association_foreignkey:ID" json:"pictures,omitempty"`        // 店图集（后台录入）
	Status         bool            `json:"status"`                                                                               // 是否成功 true：已成功 ｜ false：未成功 （后台录入）
	SourceID       int64           `json:"source_id"`                                                                            // 来源ID
	SourceInfo     string          `gorm:"size:200" json:"source_info"`                                                          // 来源描述
	Remake         string          `gorm:"size:200" json:"remake"`                                                               // 跟进备注
	Views          int64           `json:"views"`                                                                                // 浏览次数
}

// PropertyInfoScan 物业信息详细
type PropertyInfoScan struct {
	PropertyInfo
	IndustryName  string  `json:"industry_name"`   // 行业名称
	ProvinceName  string  `json:"province_name"`   // 省名称
	CityName      string  `json:"city_name"`       // 城市名称
	DistrictName  string  `json:"district_name"`   // 区名称
	StreetName    string  `json:"street_name"`     // 街道名称
	StoreTypeName string  `json:"store_type_name"` // 店铺类型名称
	AreaType_name string  `json:"area_type_name"`  // 面积分类名称
	RentTypeName  string  `json:"rent_type_name"`  // 租金分类名称
	AuditName     string  `json:"audit_name"`      // 审核人
	SourceName    string  `json:"source_name"`     // 来源人
	ActualAmount  float64 `json:"actual_amount"`   // 实收金额
}

// IndustryRange 适合经营范围
type IndustryRange struct {
	ID             int64
	IndustryID     int64  `json:"industry_id"`    // 行业ID
	IndustryName   string `json:"industry_name"`  // 行业名称
	PropertyInfoID int64  `gorm:"INDEX" json:"-"` // 物业信息ID
}

// 图片
type Picture struct {
	ID             int64
	Url            string `json:"url"`            // 店铺图
	PropertyInfoID int64  `gorm:"INDEX" json:"-"` // 物业信息ID
}

type Lot struct {
	ID             int64
	DistrictCode   string `json:"district_code"`  // 区代码
	DistrictName   string `json:"district_name"`  // 区名
	PropertyInfoID int64  `gorm:"INDEX" json:"-"` // 物业信息ID
}

// 添加物业信息
func (p *PropertyInfo) CreatePropertyInfo() error {
	db := mysql.GetMysqlDB()
	return db.Create(p).Error
}

// 查询物业信息 by id
func (p *PropertyInfoScan) QueryPropertyInfoByID() error {
	db := mysql.GetMysqlDB()
	query := db.Table("property_info").Preload("IndustryRanges").Preload("Lots").Preload("Pictures")
	query = query.Select("DISTINCT property_info.*,industry.name AS industry_name,province.name AS province_name,city.name AS city_name,district.name AS district_name,street.name AS street_name,store_type.name AS store_type_name,CONCAT( area_type.min_area, '~', area_type.max_area ) AS area_type_name,CONCAT( rent_type.min_rent, '~', rent_type.max_rent ) AS rent_type_name,user.nickname AS audit_name,user.nickname AS source_name,pay_info.actual_amount")
	query = query.Joins("LEFT JOIN industry ON industry.id = property_info.industry_id")
	query = query.Joins("LEFT JOIN province ON province.code = property_info.province_code")
	query = query.Joins("LEFT JOIN city ON city.code = property_info.city_code")
	query = query.Joins("LEFT JOIN district ON district.code = property_info.district_code")
	query = query.Joins("LEFT JOIN street ON street.code = property_info.street_code")
	query = query.Joins("LEFT JOIN store_type ON store_type.id = property_info.store_type_id")
	query = query.Joins("LEFT JOIN area_type ON area_type.id = property_info.area_type_id")
	query = query.Joins("LEFT JOIN rent_type ON rent_type.id = property_info.rent_type_id")
	query = query.Joins("LEFT JOIN user ON user.id = property_info.audit_id OR property_info.source_id = user.id")
	query = query.Joins("LEFT JOIN pay_info ON pay_info.pro_info_id = property_info.id")
	return query.First(&p).Error
}

// 查询物业信息 by source_id
func (p *PropertyInfo) QueryPropertyInfoByUserID() (propertyInfoScans []PropertyInfoScan) {
	db := mysql.GetMysqlDB()
	query := db.Table("property_info").Preload("IndustryRanges").Preload("Lots").Preload("Pictures")
	query = query.Select("property_info.*,industry.name AS industry_name,province.name AS province_name,city.name AS city_name,district.name AS district_name,street.name AS street_name,store_type.name AS store_type_name,CONCAT( area_type.min_area, '~', area_type.max_area ) AS area_type_name,CONCAT( rent_type.min_rent, '~', rent_type.max_rent ) AS rent_type_name,user.nickname AS audit_name,user.nickname AS source_name,pay_info.actual_amount")
	query = query.Joins("LEFT JOIN industry ON industry.id = property_info.industry_id")
	query = query.Joins("LEFT JOIN province ON province.code = property_info.province_code")
	query = query.Joins("LEFT JOIN city ON city.code = property_info.city_code")
	query = query.Joins("LEFT JOIN district ON district.code = property_info.district_code")
	query = query.Joins("LEFT JOIN street ON street.code = property_info.street_code")
	query = query.Joins("LEFT JOIN store_type ON store_type.id = property_info.store_type_id")
	query = query.Joins("LEFT JOIN area_type ON area_type.id = property_info.area_type_id")
	query = query.Joins("LEFT JOIN rent_type ON rent_type.id = property_info.rent_type_id")
	query = query.Joins("LEFT JOIN user ON user.id = property_info.audit_id OR property_info.source_id = user.id")
	query = query.Joins("LEFT JOIN pay_info ON pay_info.pro_info_id = property_info.id")
	query = query.Where("property_info.source_id = ?", p.SourceID)
	query.Find(&propertyInfoScans)
	return
}

// 修改物业信息 by id
func (p *PropertyInfo) EditPropertyInfoByID(args map[string]interface{}) error {
	db := mysql.GetMysqlDB()
	query := db.Model(p)
	if _, ok := args["industry_ranges"]; ok {
		query.Association("IndustryRanges").Replace("IndustryRanges", args["industry_ranges"])
		delete(args, "industry_ranges")
	}
	if _, ok := args["pictures"]; ok {
		query.Association("Pictures").Replace(p.Pictures, args["pictures"])
		delete(args, "pictures")
	}
	if _, ok := args["lots"]; ok {
		query.Association("Lots").Replace(p.Lots, args["lots"])
		delete(args, "lots")
	}
	return query.Update(args).Error
}

// 根据条件查看物业信息
func QueryPropertyInfo(pageSize int, page int, args map[string]interface{}) (propertyInfoScans []PropertyInfoScan, count int) {
	db := mysql.GetMysqlDB()
	query := db.Table("property_info").Preload("IndustryRanges").Preload("Lots").Preload("Pictures")
	query = query.Select("DISTINCT property_info.*,industry.name AS industry_name,province.name AS province_name,city.name AS city_name,district.name AS district_name,street.name AS street_name,store_type.name AS store_type_name,CONCAT( area_type.min_area, '~', area_type.max_area ) AS area_type_name,CONCAT( rent_type.min_rent, '~', rent_type.max_rent ) AS rent_type_name,user.nickname AS audit_name,user.nickname AS source_name,pay_info.actual_amount")
	query = query.Joins("LEFT JOIN industry ON industry.id = property_info.industry_id")
	query = query.Joins("LEFT JOIN province ON province.code = property_info.province_code")
	query = query.Joins("LEFT JOIN city ON city.code = property_info.city_code")
	query = query.Joins("LEFT JOIN district ON district.code = property_info.district_code")
	query = query.Joins("LEFT JOIN street ON street.code = property_info.street_code")
	query = query.Joins("LEFT JOIN store_type ON store_type.id = property_info.store_type_id")
	query = query.Joins("LEFT JOIN area_type ON area_type.id = property_info.area_type_id")
	query = query.Joins("LEFT JOIN rent_type ON rent_type.id = property_info.rent_type_id")
	query = query.Joins("LEFT JOIN user ON user.id = property_info.audit_id OR property_info.source_id = user.id")
	query = query.Joins("LEFT JOIN lot ON lot.property_info_id = property_info.id")
	query = query.Joins("LEFT JOIN pay_info ON pay_info.pro_info_id = property_info.id")
	if v, ok := args["telephone"]; ok && v.(string) != "" {
		var buf strings.Builder
		buf.WriteString("%")
		buf.WriteString(v.(string))
		buf.WriteString("%")
		query = query.Where("property_info.telephone like ?", buf.String())
	}
	if v, ok := args["nickname"]; ok && v.(string) != "" {
		var buf strings.Builder
		buf.WriteString("%")
		buf.WriteString(v.(string))
		buf.WriteString("%")
		query = query.Where("property_info.nickname like ?", buf.String())
	}
	if v, ok := args["title"]; ok && v.(string) != "" {
		var buf strings.Builder
		buf.WriteString("%")
		buf.WriteString(v.(string))
		buf.WriteString("%")
		query = query.Where("property_info.title like ?", buf.String())
	}
	if v, ok := args["province_code"]; ok && v.(string) != "" {
		query = query.Where("province.code = ?", v.(string))
	}
	if v, ok := args["show_status"]; ok {
		query = query.Where("property_info.status = ?", v.(bool))
	}
	if v, ok := args["district_code"]; ok && v.(string) != "" {
		query = query.Where("district.code = ?", v.(string))
	}
	if v, ok := args["city_code"]; ok && v.(string) != "" {
		query = query.Where("city.code = ?", v.(string))
	}
	if v, ok := args["street_code"]; ok && v.(string) != "" {
		query = query.Where("street.code = ?", v.(string))
	}
	if v, ok := args["audit"]; ok && v.(string) != "" {
		query = query.Where("property_info.audit = ?", v.(string))
	}
	if v, ok := args["industry_id"]; ok {
		if arr, ok2 := v.([]int64); ok2 {
			query = query.Where("property_info.industry_id in (?)", arr)
		} else if item, ok3 := v.(string); ok3 && item != "" {
			id, _ := strconv.ParseInt(item, 10, 64)
			query = query.Where("property_info.industry_id = ?", id)
		}
	}
	if v, ok := args["area_type_id"]; ok && v.(string) != "" {
		query = query.Where("property_info.area_type_id = ?", v.(string))
	}
	if v, ok := args["rent_type_id"]; ok && v.(string) != "" {
		query = query.Where("property_info.rent_type_id = ?", v.(string))
	}
	v1, ok1 := args["min_area"]
	v2, ok2 := args["max_area"]
	if ok1 && ok2 && v1.(string) != "" && v2.(string) != "" {
		query = query.Where("property_info.area BETWEEN ? AND ?", v1.(string), v2.(string))
	} else if ok1 && v1.(string) != "" {
		query = query.Where("property_info.area >= ?", v1.(string))
	} else if ok2 && v2.(string) != "" {
		query = query.Where("property_info.area <= ?", v2.(string))
	}
	v3, ok3 := args["min_rent"]
	v4, ok4 := args["max_rent"]
	if ok3 && ok4 && v3.(string) != "" && v4.(string) != "" {
		query = query.Where("property_info.rent BETWEEN ? AND ?", v3.(string), v4.(string))
	} else if ok3 && v3.(string) != "" {
		query = query.Where("property_info.rent >= ?", v3.(string))
	} else if ok4 && v4.(string) != "" {
		query = query.Where("property_info.rent <= ?", v4.(string))
	}
	if v, ok := args["store_type_id"]; ok && v.(string) != "" {
		query = query.Where("property_info.store_type_id = ?", v.(string))
	}
	if v, ok := args["bus_type"]; ok && v.(string) != "" {
		query = query.Where("property_info.bus_type = ?", v.(string))
	}
	if v, ok := args["model_type"]; ok && v.(string) != "" {
		var arr []int64
		strArr := strings.Split(v.(string), ",")
		for _, item := range strArr {
			id, _ := strconv.ParseInt(item, 10, 64)
			arr = append(arr, id)
		}
		query = query.Where("property_info.model_type in (?)", arr)
		if v2, ok2 := args["district_code"]; ok2 && v2.(string) != "" && (v.(string) == "3" || v.(string) == "4") {
			query = query.Where("lot.district_code = ?", v2.(string))
		} else {
			if v3, ok3 := args["district_code"]; ok3 && v3.(string) != "" {
				query = query.Where("district.code = ?", v3.(string))
			}
		}
	}
	if v, ok := args["source_id"]; ok && v.(int64) > 0 {
		query = query.Where("property_info.source_id = ?", v.(int64))
	}
	if v, ok := args["status"]; ok && v.(string) != "" {
		query = query.Where("property_info.status = ?", v.(string))
	}
	if v, ok := args["protect"]; ok && v.(string) != "" {
		query = query.Where("property_info.protect = ?", v.(string))
	}
	if v, ok := args["id"]; ok {
		ids, _ := v.([]int64)
		query = query.Where("property_info.id in (?)", ids)
	}
	query.Count(&count)
	if v, ok := args["sort_condition"]; ok && v.(string) != "" {
		query = query.Order("property_info." + v.(string) + " desc")
	}
	query.Limit(pageSize).Offset((page - 1) * pageSize).Find(&propertyInfoScans)
	return
}

// 查询物业信息是否关联面积分类
func QueryPropertyInfoRelationAreaType(ids []int64) bool {
	db := mysql.GetMysqlDB()
	var propertyInfos []PropertyInfo
	if count := db.Where("area_type_id in (?)", ids).Find(&propertyInfos).RowsAffected; count > 0 {
		return true
	}
	return false
}

// 查询物业信息是否关联租金分类
func QueryPropertyInfoRelationRentType(ids []int64) bool {
	db := mysql.GetMysqlDB()
	var propertyInfos []PropertyInfo
	if count := db.Where("rent_type_id in (?)", ids).Find(&propertyInfos).RowsAffected; count > 0 {
		return true
	}
	return false
}

// 查询物业信息是否关联行业
func QueryPropertyInfoRelationIndustry(ids []int64) bool {
	db := mysql.GetMysqlDB()
	var propertyInfos []PropertyInfo
	if count := db.Where("industry_id in (?)", ids).Find(&propertyInfos).RowsAffected; count > 0 {
		return true
	}
	return false
}

// 查询物业信息是否关联店铺类型
func QueryPropertyInfoRelationStoreTypeID(ids []int64) bool {
	db := mysql.GetMysqlDB()
	var propertyInfos []PropertyInfo
	if count := db.Where("store_type_id in (?)", ids).Find(&propertyInfos).RowsAffected; count > 0 {
		return true
	}
	return false
}

// 删除图集图片
func (p *Picture) DelPicturre() error {
	db := mysql.GetMysqlDB()
	return db.Where("property_info_id = ?", p.PropertyInfoID).Unscoped().Delete(&p, p.ID).Error
}

// 上传图片
func (p *Picture) AddPicture() error {
	db := mysql.GetMysqlDB()
	return db.Create(&p).Error
}

// 最新动态
func QueryProInfoDynamic() (count1 int, count2 int, count3 int) {
	//"zzzp": cu2,
	//"zzzd": cu3,
	//"zcjl": cu1,
	db := mysql.GetMysqlDB()
	query := db.Table("property_info")

	query1 := query.Where("status = true")
	query1.Count(&count1)

	query2 := query.Where("model_type IN (0,1,3) AND `status` IS NOT TRUE")
	query2.Count(&count3)

	query3 := query.Where("model_type IN (4,5) AND `status` IS NOT TRUE")
	query3.Count(&count2)
	return
}

// 删除物业信息 By ID
func (p *PropertyInfo) DelProInfo() error {
	db := mysql.GetMysqlDB()
	return db.Where("id = ?", p.ID).Delete(&PropertyInfo{}).Error
}
