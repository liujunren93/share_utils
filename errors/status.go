package errors

type Status int32

func (s Status) GetCode() int32 {
	return int32(s)
}

func (s Status) GetMsg() (msg string) {

	return statusText[Status(s.GetCode())]
}

//database

const (
	StatusOK Status = 200 //success
	//数据相关
	StatusDBInternalErr      Status = 11001 // 数据库内错误
	StatusDBDuplication      Status = 11002 // 数据重复
	StatusDBNotFound         Status = 11003 // 数据不存在
	StatusDBRowsAffectedZero Status = 11004 // 数据影响条数为0
	// 系统相关错误
	StatusBadRequest          Status = 14000 //数据绑定错误
	StatusUnauthorized        Status = 14001 //账户类错误
	StatusTokenTimeout        Status = 14002 //token过期
	StatusForbidden           Status = 14003 //权限
	StatusNotFound            Status = 14004 //
	StatusRefreshTokenTimeout Status = 14005 //RefreshToken过期
	StatusMetadataNotFound    Status = 14006 // 从metadata中未获取到数据
	StatusRequestTimeout      Status = 14008 //
	StatusDomainDisable       Status = 14009 // domain 被禁用
	StatusInternalServerError Status = 15000 //服务器通用错误 前端不显示
	StatusBreakerServerError  Status = 15001 //服务器通用错误 前端不显示
	StatusPulishConfigError   Status = 15204 // 发布配置错误
	StatusBatchError          Status = 16501 //服务器批量通用错误

)

var statusText = map[Status]string{
	StatusOK:                  "ok",
	StatusBadRequest:          "BadRequest",
	StatusUnauthorized:        "Unauthorized",
	StatusTokenTimeout:        "Authorized Timeout",
	StatusRefreshTokenTimeout: "RefreshToken Timeout",
	StatusMetadataNotFound:    "Metadata Lost Val",
	StatusForbidden:           "Forbidden",
	StatusNotFound:            "NotFound",
	StatusRequestTimeout:      "Request Timeout",
	StatusDomainDisable:       "Domain Disable",
	StatusInternalServerError: "Internal Server Error",
	StatusDBInternalErr:       "DBInternalServerError",
	StatusDBDuplication:       "Duplication",
	StatusDBNotFound:          "DataNotFount",
	StatusDBRowsAffectedZero:  "NoDataToUpdate",
	StatusBreakerServerError:  "Status Breaker Server Error",
	StatusPulishConfigError:   "Pulish Config Error",
	StatusBatchError:          "Status Batch Error",
}
