package mlicense

type Licenser interface {
	//获取机器指纹
	GetMachineTrait() (string, error)
	//跟新license
	UpdateLicense(data string) (*LicenseData, error)
	//获取license 信息
	GetLicense() (*LicenseData, error)
	//获取license 文本
	GetLicenseContent() (string, error)
	//校验license时间有效性
	VerifyTime() error
	//校验资产个数有效性
	VerifyCount(count int32) error
	//获取license版本号
	GetLicenseVersion() (string, error)
	//获取自定义json数据
	GetLicenseJsonData() (string, error)
}

type LicenseData struct {
	StartTime int64 //开始时间
	EndTime   int64 //结束时间
	Count     int32 //支持资产个数
}

//	desc：获取license实例
//  inparam： strProductId 产品代码
//			  strVmIp      虚拟机ip
//			  isVm         是否是虚拟机
//  out param：
//				license实例
//				错误信息
func NewLicense(strProductId, strVmIp string, isVm bool) (Licenser, error) {
	if strProductId == "" {
		return nil, errSrvType
	}
	if isVm && strVmIp == "" {
		return nil, errIp
	}
	return &license{
		strSrvType: strProductId,
		strVmIp:    strVmIp,
		isVm:       isVm,
	}, nil
}

//	desc：获取license实例
//  inparam： productId 产品代码
//			  vmIp      虚拟机ip
//			  isVm         是否是虚拟机
//            addr     连接地址 为空默认走unix socke
//            network  连接地址 为空默认走unix socke
//
//  out param：
//				license实例
//

func NewLicenseV3(productId, vmIp string, addr, network string, isVm bool) (Licenser, error) {
	if productId == "" {
		return nil, errSrvType
	}

	if isVm && vmIp == "" {
		return nil, errIp
	}

	return &licenseV3{
		serviceType: productId,
		vmIp:        vmIp,
		addr:        addr,
		network:     network,
	}, nil
}
