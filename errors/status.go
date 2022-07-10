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
	StatusOK                  Status = 200  //success
	StatusBadRequest          Status = 4000 //数据绑定错误
	StatusUnauthorized        Status = 4001 //账户类错误
	StatusForbidden           Status = 4003 //权限
	StatusNotFound            Status = 4004 //
	StatusRequestTimeout      Status = 4008 //
	StatusDomainDisable       Status = 4009 // domain 被禁用
	StatusInternalServerError Status = 5000 //服务器通用错误 前端不显示
	StatusBreakerServerError  Status = 5001 //服务器通用错误 前端不显示
	StatusBatchError          Status = 5040 //服务器批量通用错误 前端显示
	// 数据重复 52... DB 错误
	StatusDBInternalErr      Status = 5200 // 数据库内错误
	StatusDBDuplication      Status = 5201 // 数据重复
	StatusDBNotFound         Status = 5202 // 数据不存在
	StatusDBRowsAffectedZero Status = 5203 // 数据影响条数为0
	StatusPulishConfigError  Status = 5204 // 发布配置错误

)

var statusText = map[Status]string{
	StatusOK:                  "ok",
	StatusBadRequest:          "BadRequest",
	StatusUnauthorized:        "Unauthorized",
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
