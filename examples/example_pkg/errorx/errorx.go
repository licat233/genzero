package errorx

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/licat233/genzero/examples/example_pkg/uniqueid"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/** Errorx 对接前端antpro框架的接口格式规范，用于实现error接口
 */
type Errorx struct {
	Status       bool   `json:"status"`       // 响应状态
	Success      bool   `json:"success"`      // 响应状态，用于对接umijs
	ErrorCode    int    `json:"errorCode"`    // 【选填】错误类型代码：400错误请求，401未授权，500服务器内部错误，200成功
	ErrorMessage string `json:"errorMessage"` // 【选填】向用户显示消息
	TraceMessage string `json:"traceMessage"` // 【选填】自己加上去的，调试错误信息，请勿在生产环境下使用，可有可无
	ShowType     int    `json:"showType"`     // 【选填】错误显示类型：0.不提示错误;1.警告信息提示；2.错误信息提示；4.通知提示；9.页面跳转
	TraceId      string `json:"traceId"`      // 【选填】方便后端故障排除：唯一的请求ID
	Host         string `json:"host"`         // 【选填】方便后端故障排除：当前访问服务器的主机
}

/** ResponseError 对接前端antpro框架的接口格式规范，用于实现展现错误
 */
type ResponseError struct {
	Status       bool   `json:"status"`       // 响应状态
	Success      bool   `json:"success"`      // 响应状态，用于对接umijs
	ErrorCode    int    `json:"errorCode"`    // 【选填】错误类型代码：400错误请求，401未授权，500服务器内部错误，200成功
	ErrorMessage string `json:"errorMessage"` // 【选填】向用户显示消息
	TraceMessage string `json:"traceMessage"` // 【选填】自己加上去的，调试错误信息，请勿在生产环境下使用，可有可无
	ShowType     int    `json:"showType"`     // 【选填】错误显示类型：0.不提示错误;1.警告信息提示；2.错误信息提示；4.通知提示；9.页面跳转
	TraceId      string `json:"traceId"`      // 【选填】方便后端故障排除：唯一的请求ID
	Host         string `json:"host"`         // 【选填】方便后端故障排除：当前访问服务器的主机
}

const (
	_ERROR_CODE_INTERNAL         = 500  //内部错误
	_ERROR_CODE_EXTERNAL         = 400  //外部错误（默认）会给与提示
	_ERROR_CODE_EXTERNAL_SECRECY = 4001 //外部错误，但错误原因保密
	_ERROR_CODE_AUTH             = 401  //外部错误（默认）
	_ERROR_CODE_ACCESS           = 403  //服务器禁止访问
	_SHOW_TYPE                   = 2    //信息错误
)

var debug = os.Getenv("SERVICE_MODE") == "dev"

// var mode = os.Getenv("SERVICE_MODE")
var ServerError = InternalError
var RequestError = ExternalError

// 消息提示
var ServerErrorMsg = "server error"
var RequestErrorMsg = "bad request"
var AuthErrorMsg = "permission denied"
var RequestSuccessMsg = "request successful"
var RequestFailedMsg = "request failed"

var Ip = "0.0.0.0"

// 获取本地服务器的ip
func LocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "0.0.0.0", err
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "0.0.0.0", nil
}

func init() {
	Ip, _ = LocalIP()
}

// Error 返回错误信息，主要用于实现原生error接口，
// 注意：打印errorx的时候，会调用该函数值，所以返回errorx，等于返回该Error()的结果
func (e *Errorx) Error() string {
	return e.ErrorMessage
}

// Data 返回errorx的值
// 如果需要返回errorx的值怎么办？
// 我们需要自己定义一个展现errorx结果的函数，使用ResponseError结构体
func (e *Errorx) Data() *ResponseError {
	return &ResponseError{
		Status:       e.Status,
		Success:      e.Status,
		ErrorCode:    e.ErrorCode,
		ErrorMessage: e.ErrorMessage,
		TraceMessage: e.TraceMessage,
		ShowType:     e.ShowType,
		TraceId:      e.TraceId,
		Host:         e.Host,
	}
}

/** InternalError 内部错误 500
 * traceErr error 调试信息
 * alertErr error 提示信息
 */
func InternalError(errs ...any) error {
	var traceErr error
	alertErr := errors.New(ServerErrorMsg)
	n := len(errs)
	if n > 0 && errs[0] != nil {
		traceErr = fmt.Errorf("%s", errs[0])
	}
	if n > 1 && errs[1] != nil {
		alertErr = fmt.Errorf("%s", errs[1])
	}
	return internalError(alertErr.Error(), traceErr)
}

/** ExternalError 外部错误 400
 * traceErr error 调试信息
 * alertErr error 提示信息
 */
func ExternalError(errs ...any) error {
	var traceErr error
	alertErr := errors.New(RequestErrorMsg)
	n := len(errs)
	if n > 0 && errs[0] != nil {
		traceErr = fmt.Errorf("%s", errs[0])
	}
	if n > 1 && errs[1] != nil {
		alertErr = fmt.Errorf("%s", errs[1])
	}
	return externalError(alertErr.Error(), traceErr)
}

/** AccessError 外部错误 403
 * traceErr error 调试信息
 * alertErr error 提示信息
 */
func AccessError(errs ...any) error {
	var traceErr error
	alertErr := errors.New(AuthErrorMsg)
	n := len(errs)
	if n > 0 && errs[0] != nil {
		traceErr = fmt.Errorf("%s", errs[0])
	}
	if n > 1 && errs[1] != nil {
		alertErr = fmt.Errorf("%s", errs[1])
	}
	return accessError(alertErr.Error(), traceErr)
}

/** AuthError 身份校验失败 401
 * traceMsg string 调试信息
 */
func AuthError(traceErr any) error {
	return authError(AuthErrorMsg, traceErr)
}

// Alert 错误提示, 400
func Alert(msg any) error {
	var alertMsg string
	if msg != nil {
		alertMsg = strings.TrimSpace(fmt.Sprint(msg))
	}
	if alertMsg == "" {
		msg = RequestFailedMsg
	}
	return &Errorx{
		Status:       false,
		Success:      false,
		ErrorCode:    _ERROR_CODE_EXTERNAL,
		ErrorMessage: alertMsg,
		TraceMessage: "",
		ShowType:     _SHOW_TYPE,
		TraceId:      uniqueid.NewUUID(),
		Host:         Ip,
	}
}

// ----------实现--------
func internalError(alert any, trace any) *Errorx {
	errorx := newErrorx(alert, trace)
	if len(errorx.ErrorMessage) == 0 {
		errorx.ErrorMessage = ServerErrorMsg
	}
	errorx.ErrorCode = _ERROR_CODE_INTERNAL
	//非调试模式，统一报：服务器错误
	if !debug {
		errorx.ErrorMessage = ServerErrorMsg
		errorx.TraceMessage = "There is an error in the server, please contact the administrator licat233@gmail.com"
	}
	return errorx
}

func externalError(alert any, trace any) *Errorx {
	errorx := newErrorx(alert, trace)
	if len(errorx.ErrorMessage) == 0 {
		errorx.ErrorMessage = RequestErrorMsg
	}
	errorx.ErrorCode = _ERROR_CODE_EXTERNAL
	return errorx
}

func authError(alert any, trace any) *Errorx {
	errorx := newErrorx(alert, trace)
	if len(errorx.ErrorMessage) == 0 {
		errorx.ErrorMessage = AuthErrorMsg
	}
	errorx.ErrorCode = _ERROR_CODE_AUTH
	return errorx
}

func accessError(alert any, trace any) *Errorx {
	errorx := newErrorx(alert, trace)
	if len(errorx.ErrorMessage) == 0 {
		errorx.ErrorMessage = AuthErrorMsg
	}
	errorx.ErrorCode = _ERROR_CODE_ACCESS
	return errorx
}

// newErrorx 新建Errorx实例，默认为外部请求错误
func newErrorx(alert any, trace any) *Errorx {
	var alertMsg, traceMsg string
	if alert != nil {
		alertMsg = strings.TrimSpace(fmt.Sprint(alert))
	}
	if alertMsg == "" {
		alertMsg = RequestErrorMsg
	}
	//调试模式下，展现错误调试信息
	if debug && trace != nil {
		traceMsg = strings.TrimSpace(fmt.Sprint(trace))
	}
	errorx := &Errorx{
		Status:       false,
		Success:      false,
		ErrorCode:    _ERROR_CODE_EXTERNAL,
		ErrorMessage: alertMsg,
		TraceMessage: traceMsg,
		ShowType:     _SHOW_TYPE,
		TraceId:      uniqueid.NewUUID(),
		Host:         Ip,
	}
	return errorx
}

// 判断是否为*Errorx错误
func IsErrorx(err error) bool {
	_, ok := err.(*Errorx)
	return ok
}

// 将原生错误转化为errorx 400错误
func Convert(err error) (errorx *Errorx) {
	if err == nil {
		return nil
	}
	//类型断言看是否为 Errorx 类型
	if errx, ok := err.(*Errorx); ok {
		//无需处理，直接返回
		return errx
	}
	//类型断言看是否为status.Error 类型
	if se, ok := err.(interface {
		GRPCStatus() *status.Status
	}); ok {
		//如果是statusErr错误，则需要转化为Errorx错误
		return StatusErrToErrorx(se.GRPCStatus())
	}

	//原生错误和其它错误，统一做转化处理，归类于内部错误
	errorx = internalError("", err)
	// errorx = New(err.Error())
	return
}

/** New 新建基本的errorx错误 400
 * alertErr string 提示信息
 * traceErr string 调试信息
 */
func New(errs ...any) error {
	alertErr := RequestErrorMsg
	n := len(errs)
	if n > 0 {
		if errs[0] != nil {
			alertErr = strings.TrimSpace(fmt.Sprint(errs[0]))
		}
	}
	var traceErr string
	if n > 1 {
		if errs[1] != nil {
			traceErr = strings.TrimSpace(fmt.Sprint(errs[1]))
		}
	}
	return externalError(alertErr, errors.New(traceErr))
}

/** ResponseErrorHandler 对接go-zero的错误处理函数，将原生错误转化为 Errorx
 * err 传入的错误
 */
func ResponseErrorHandler(err error) (int, interface{}) {
	if err == nil {
		//没有错误，应该正常返回
		return http.StatusOK, &struct {
			Status  bool   `json:"status"`
			Success bool   `json:"success"`
			Message string `json:"message"`
		}{
			Status:  true,
			Success: true,
			Message: RequestSuccessMsg,
		}
	}
	//类型断言看是否为 Errorx 类型
	if errorx, ok := err.(*Errorx); ok {
		//无需处理，直接返回
		return http.StatusOK, errorx.Data()
	}

	//类型断言看是否为 status.Error 类型
	if se, ok := err.(interface {
		GRPCStatus() *status.Status
	}); ok {
		//如果是statusErr错误，则需要转化为Errorx错误
		return http.StatusOK, StatusErrToErrorx(se.GRPCStatus()).Data()
	}

	/** 20220826修正：
	 * ①httpx.Parse(r, &req)返回原生error，为request错误，默认
	 * ②业务代码panic返回的error不经过错误处理，为internal server错误
	 */

	return http.StatusOK, externalError(RequestErrorMsg, err).Data()
}

func ResponseErrorHandlerCtx(ctx context.Context, err error) (int, interface{}) {
	return ResponseErrorHandler(err)
}

/**需要注意的是：rpc中返回的status.Error Internael错误没有traceMessage，只有一个message**/

// 将error转化成External RPC status.Status，主要用于rpc服务中
func ExtRpcErr(err error) (statusError error) {
	statusError = status.Error(_ERROR_CODE_EXTERNAL, err.Error())
	return
}

// 将error转化成Internal RPC status.Status，主要用于rpc服务中
func IntRpcErr(err error) (statusError error) {
	statusError = status.Error(_ERROR_CODE_INTERNAL, err.Error())
	return
}

// 将error转化成Internal RPC status.Status，主要用于rpc服务中，外部错误，但错误信息保密
func SecRpcErr(err error) (statusError error) {
	statusError = status.Error(_ERROR_CODE_EXTERNAL_SECRECY, err.Error())
	return
}

// 将status.Status error转化为Errorx，主要用于api服务处理rpc返回的错误
func FromRpcErr(statusErr error) (errorx *Errorx) {
	s := status.Convert(statusErr)
	return StatusErrToErrorx(s)
}

// status.Error转化为errorx
func StatusErrToErrorx(statusErr *status.Status) (errorx *Errorx) {
	switch statusErr.Code() {
	case _ERROR_CODE_INTERNAL:
		errorx = internalError(ServerErrorMsg, errors.New(statusErr.Message()))
	case _ERROR_CODE_EXTERNAL:
		errorx = externalError(statusErr.Message(), nil)
	case _ERROR_CODE_EXTERNAL_SECRECY:
		errorx = externalError(RequestErrorMsg, errors.New(statusErr.Message()))
	case _ERROR_CODE_AUTH:
		errorx = authError(statusErr.Message(), nil)
	case _ERROR_CODE_ACCESS:
		errorx = accessError(RequestFailedMsg, errors.New(statusErr.Message()))
	default:
		//未知错误，可能是panic引起的，内部错误，不抛给前端
		//记录一下
		err := fmt.Errorf("未知错误: %+v", statusErr.Message())
		logx.WithContext(context.Background()).Error(err)
		errorx = internalError(ServerErrorMsg, err)
	}
	return
}

// 将errorx转化为status.Status error
func ToStatusErr(errorx error) (statusErr error) {
	if errx, ok := errorx.(*Errorx); ok {
		return status.Error(codes.Code(errx.ErrorCode), errx.ErrorMessage)
	}
	//原生错误，我们默认转化为500错误
	return status.Error(_ERROR_CODE_INTERNAL, errorx.Error())
}

// 转化为status.Error
func (errx *Errorx) ToStatusErr() status.Status {
	if errx.TraceMessage != "" {
		return *status.New(codes.Code(errx.ErrorCode), errx.TraceMessage)
	}
	return *status.New(codes.Code(errx.ErrorCode), errx.ErrorMessage)
}

// 转化为error接口
func (errx *Errorx) ToError() error {
	return errx
}
