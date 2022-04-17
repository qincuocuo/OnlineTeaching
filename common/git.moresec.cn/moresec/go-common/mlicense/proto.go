/*
@Time : 2020-11-16 14:27
@Author : gaoxl@moresec.cn
@Description:
@Software: GoLand
*/
package mlicense

const (
	LICENSE_TYPE_VERIFY_TIME_V3     int32 = 1
	LICENSE_TYPE_VERIFY_COUNT_V3    int32 = 2
	LICENSE_TYPE_TARIT_V3           int32 = 3
	LICENSE_TYPE_UPDATE_V3          int32 = 4
	LICENSE_TYPE_LICENSE_CONTENT_V3 int32 = 5
	LICENSE_TYPE_JSON_DATA_V3       int32 = 6
	LICENSE_TYPE_VERSION_V3         int32 = 7
	LICENSE_TYPE_LICENSE_V3         int32 = 8
)

const (
	RET_CODE_OK    int32 = 0
	RET_CODE_ERROR int32 = 1
)

type LicenseRequest struct {
	ReqType    int32  `json:"req_type,omitempty"`
	Ip         string `json:"ip,omitempty"`
	Service    string `json:"service,omitempty"`
	License    string `json:"license,omitempty"`
	AssetCount int32  `json:"asset_count,omitempty"`
}

type LicenseResponse struct {
	Code      int32  `json:"code,omitempty"`
	Msg       string `json:"msg,omitempty"`
	StartTime uint64 `json:"start_time,omitempty"`
	EndTime   uint64 `json:"end_time,omitempty"`
	Count     int32  `json:"count,omitempty"`
	Data      []byte `json:"data,omitempty"`
}
