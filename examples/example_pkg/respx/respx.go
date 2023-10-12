package respx

import (
	"reflect"

	"github.com/licat233/genzero/examples/example_pkg/errorx"
)

var (
	//respx只会返回成功的请求
	SUCCESS_MSG = "request success"
)

// StatusData 状态响应结构体
type StatusData struct {
	Status  bool   `json:"status"`  // 响应状态
	Success bool   `json:"success"` // 响应状态，用于对接umijs
	Message string `json:"message"` // 响应消息
}

// SingleData 单数据响应结构体
type SingleData struct {
	Status  bool   `json:"status"`  // 响应状态
	Success bool   `json:"success"` // 响应状态，用于对接umijs
	Message string `json:"message"` // 响应消息
	Data    any    `json:"data"`    // 返回数据
}

// ListData 列表响应数据结构体
type ListData struct {
	Status    bool        `json:"status"`    // 响应状态
	Success   bool        `json:"success"`   // 响应状态，用于对接umijs
	Message   string      `json:"message"`   //【给予的提示信息】
	Data      interface{} `json:"data"`      //【选填】响应的业务数据
	Total     int64       `json:"total"`     //【选填】数据总个数
	PageSize  int64       `json:"pageSize"`  //【选填】单页数量
	Page      int64       `json:"page"`      //【选填】当前页码
	Current   int64       `json:"current"`   //【选填】当前页码，用于对接umijs
	TotalPage int64       `json:"totalPage"` //【自己加上去的】总共有多少页，根据前端的pageSize来计算
}

func NewData(data any) *SingleData {
	return &SingleData{
		Status:  true,
		Success: true,
		Message: SUCCESS_MSG,
		Data:    data,
	}
}

func NewListData(listData any, totalRecord, pageSize, Page int64) *ListData {
	data := listData
	if listData != nil {
		if reflect.TypeOf(listData).Kind().String() != "slice" {
			data = []any{listData}
		}
	} else {
		data = []any{}
	}
	if pageSize == 0 {
		pageSize = 1
	}
	return &ListData{
		Status:    true,
		Success:   true,
		Message:   SUCCESS_MSG,
		Data:      data,
		Total:     totalRecord,
		PageSize:  pageSize,
		Page:      Page,
		Current:   Page,
		TotalPage: (totalRecord + pageSize - 1) / pageSize,
	}
}

// StatusResp 返回响应状态
func StatusResp(alertMsg string, status bool) *StatusData {
	if alertMsg == "" {
		alertMsg = SUCCESS_MSG
	}
	return &StatusData{
		Status:  status,
		Success: status,
		Message: alertMsg,
	}
}

// DefaultStatusResp 默认返回响应状态，主要用于对接go-zero
func DefaultStatusResp(err error) (*StatusData, error) {
	if err != nil {
		return nil, errorx.Convert(err)
	}
	return StatusResp(SUCCESS_MSG, true), nil
}

// SingleResp 返回单数据响应
func SingleResp(alertMsg string, data any) *SingleData {
	if alertMsg == "" {
		alertMsg = SUCCESS_MSG
	}
	return &SingleData{
		Status:  true,
		Success: true,
		Message: alertMsg,
		Data:    data,
	}
}

// DefaultSingleResp 默认返回单数据响应，主要用于对接go-zero
func DefaultSingleResp(data any, err error) (*SingleData, error) {
	if err != nil {
		return nil, errorx.Convert(err)
	}
	return SingleResp(SUCCESS_MSG, data), nil
}

/** ListResp 返回列表数据响应，用于分页查询数据
* alertMsg string 提示信息
* data any 返回数据
* totalRecord in64 总数据量
* pageSize int64 单页数据量
* Current int64 当前页码
 */
func ListResp(alertMsg string, listData any, totalRecord, pageSize, Page int64) *ListData {
	// data := listData
	// if listData != nil {
	// 	if reflect.TypeOf(listData).Kind().String() != "slice" {
	// 		data = []any{listData}
	// 	}
	// } else {
	// 	data = []any{}
	// }
	if listData == nil {
		listData = make([]any, 0)
	}
	if pageSize == 0 {
		pageSize = 1
	}
	if alertMsg == "" {
		alertMsg = SUCCESS_MSG
	}
	return &ListData{
		Status:    true,
		Success:   true,
		Message:   alertMsg,
		Data:      listData,
		Total:     totalRecord,
		PageSize:  pageSize,
		Page:      Page,
		Current:   Page,
		TotalPage: (totalRecord + pageSize - 1) / pageSize,
	}
}

/** DefaultListResp 默认返回列表数据响应，用于分页查询数据，主要用于对接go-zero
* err error 提示信息
* data any 返回数据
* totalRecord in64 总数据量
* pageSize int64 单页数据量
* Current int64 当前页码
 */
func DefaultListResp(listData any, totalRecord, pageSize, Page int64, err error) (*ListData, error) {
	if err != nil {
		return nil, errorx.Convert(err)
	}
	return ListResp(SUCCESS_MSG, listData, totalRecord, pageSize, Page), nil
}

// Msg 根据error返回不同提示信息
func Msg(data any, err error, successMsg, failedMsg string) (any, error) {
	if err != nil {
		return nil, errorx.InternalError(err)
	}
	return SingleResp(successMsg, data), nil
}

// DefaultBoolResp 默认返回响应状态，主要用于对接go-zero
func DefaultBoolResp(ok bool, err error) (*StatusData, error) {
	if ok {
		return &StatusData{
			Status:  true,
			Success: true,
			Message: SUCCESS_MSG,
		}, nil
	}
	if err == nil {
		err = errorx.ExternalError("Request failed! But the server has no error prompt")
	}
	return nil, errorx.Convert(err)
}
