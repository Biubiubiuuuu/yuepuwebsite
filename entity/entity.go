package entity

// 返回结果
type ResponseData struct {
	Status  bool                   `json:"status"`  // 成功失败标志；true：成功 、false：失败
	Data    map[string]interface{} `json:"data"`    // 返回数据
	Message string                 `json:"message"` // 提示信息
}

// 用户登录
type UserLogin struct {
	UserName string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

// 用户注册
type UserRegister struct {
	Telephone string `json:"telephone"` // 手机号码
	Password  string `json:"password"`  // 密码
	Code      string `json:"code"`      // 验证码
}

// 修改用户基础信息
type EditUser struct {
	Username       string `json:"username"`       // 用户名
	Telephone      string `json:"telephone"`      // 手机号码
	Nickname       string `json:"nickname"`       // 姓名
	Sex            string `json:"sex"`            // 性别 0:未知 ｜ 1:男 ｜ 2:女 （空或其他默认未知）
	Landlinenumber string `json:"landlinenumber"` // 座机号码
	QQ             string `json:"QQ"`             // QQ
	Email          string `json:"email"`          // 邮箱
}

// 修改用户密码
type EditUserPass struct {
	OldPass string `json:"old_pass"` // 旧密码
	NewPass string `json:"new_pass"` // 新密码
}

// 用户提交店铺转让请求
type UserStoretransferRequest struct {
	IndustryID     int64   `json:"industry_id"`     // 经营业态ID
	Title          string  `json:"title"`           // 标题
	Nickname       string  `json:"nickname"`        // 联系人
	Telephone      string  `json:"telephone"`       // 联系手机
	Image          string  `json:"image"`           // 图片
	ProvinceCode   string  `json:"province_code"`   // 省代码
	CityCode       string  `json:"city_code"`       // 城市代码
	DistrictCode   string  `json:"district_code"`   // 区代码
	StreetCode     string  `json:"street_code"`     // 街道代码
	Address        string  `json:"address"`         // 详细地址
	StoreTypeID    int64   `json:"store_type_id"`   // 店铺类型ID
	Idling         bool    `json:"idling"`          // 可否空转
	InOperation    string  `json:"in_operation"`    // 是否营业中 0-新铺 ｜ 1-空置中 ｜ 2-营业中
	Area           float64 `json:"area"`            // 面积（单位：平方米）
	Rent           float64 `json:"rent"`            // 租金（单位：元/月）
	TransferFee    float64 `json:"transfer_fee"`    // 转让费用（单位：万元 不录入则前台显示为面议）
	IndustryRanges []int64 `json:"industry_ranges"` // 适合经营范围id
	Description    string  `json:"description"`     // 详细描述
}

// 查询用户已发布物业信息条件
type UserPropertyInfoRequest struct {
	IndustryID   int64   `json:"industry_id"`   // 经营业态ID
	Title        string  `json:"title"`         // 标题
	Nickname     string  `json:"nickname"`      // 联系人
	ProvinceCode string  `json:"province_code"` // 省代码
	CityCode     string  `json:"city_code"`     // 城市代码
	DistrictCode string  `json:"district_code"` // 区代码
	StreetCode   string  `json:"street_code"`   // 街道代码
	StoreTypeID  int64   `json:"store_type_id"` // 店铺类型ID
	MinArea      float64 `json:"min_area"`      // 最小面积（单位：平方米）
	MaxArea      float64 `json:"max_area"`      // 最大面积（单位：平方米）
	MinRent      float64 `json:"min_rent"`      // 最低租金（单位：元/月）
	MaxRent      float64 `json:"max_rent"`      // 最高租金（单位：元/月）
}

// 用户提交我要找铺请求
type UserFindStoreRequest struct {
	IndustryID  int64    `json:"industry_id"`   // 经营业态ID
	Title       string   `json:"title"`         // 标题
	Nickname    string   `json:"nickname"`      // 联系人
	Telephone   string   `json:"telephone"`     // 联系手机
	StoreTypeID int64    `json:"store_type_id"` // 店铺类型ID
	Lots        []string `json:"lots"`          // 考虑地段
	Description string   `json:"description"`   // 详细描述
	MinArea     float64  `json:"min_area"`      // 最小面积（单位：平方米）
	MaxArea     float64  `json:"max_area"`      // 最大面积（单位：平方米）
	MinRent     float64  `json:"min_rent"`      // 最低租金（单位：元/月）
	MaxRent     float64  `json:"max_rent"`      // 最高租金（单位：元/月）
}