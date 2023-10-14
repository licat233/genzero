package internal

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/modules/api/conf"
	"github.com/licat233/genzero/modules/utils"
	"github.com/licat233/genzero/tools"
)

type Struct struct {
	Name    string
	TagType string
	Comment string
	Fields  StructFieldCollection
}

func NewStruct(name, TagType, comment string, fields StructFieldCollection) *Struct {
	return &Struct{
		Name:    name,
		TagType: TagType,
		Comment: comment,
		Fields:  fields,
	}
}

func (st *Struct) Copy() *Struct {
	return &Struct{
		Name:    st.Name,
		TagType: st.TagType,
		Comment: st.Comment,
		Fields:  st.Fields.Copy(),
	}
}

func (st *Struct) IgnoreStructFields(needIgnoreFields []string, more ...string) *Struct {
	all := append(needIgnoreFields, more...)
	filterFields := []*StructField{}
	for _, field := range st.Fields {
		if tools.SliceContain(all, field.Name) {
			continue
		}
		field.TagType = st.TagType
		filterFields = append(filterFields, field)
	}
	st.Fields = filterFields
	return st
}

func (st *Struct) GenCommonStructs() []*Struct {
	res := []*Struct{}
	//default
	defaultStruct := st.Copy().IgnoreStructFields(config.C.Api.IgnoreColumns)
	res = append(res, defaultStruct)

	//add req
	addReqStruct := st.Copy().IgnoreStructFields(config.C.Api.IgnoreColumns, conf.MoreIgnoreColumns...) //AddReqStruct
	addReqStruct.Name = "Add" + tools.ToCamel(st.Name) + "Req"
	addReqStruct.Comment = "添加" + st.Comment + "请求"
	res = append(res, addReqStruct)

	//put req
	putReqStruct := st.Copy().IgnoreStructFields(config.C.Api.IgnoreColumns) //PutReqStruct
	putReqStruct.Name = "Put" + tools.ToCamel(st.Name) + "Req"
	putReqStruct.Comment = "更新" + st.Comment + "请求"
	res = append(res, putReqStruct)

	//del req
	delReqStruct := st.Copy() //DelReqStruct
	delReqStruct.Name = "Del" + tools.ToCamel(st.Name) + "Req"
	delReqStruct.Comment = "删除" + st.Comment + "请求"
	delReqStruct.Fields = []*StructField{
		NewStructField("Id", "int64", "json", "id", "", st.Comment+" ID"),
	}
	res = append(res, delReqStruct)

	//get req
	getReqStruct := st.Copy() //GetReqStruct
	getReqStruct.Name = "Get" + tools.ToCamel(st.Name) + "Req"
	getReqStruct.Comment = "获取" + st.Comment + "请求"
	getReqStruct.Fields = []*StructField{
		NewStructField("Id", "int64", "form", "id", "", st.Comment+" ID"),
	}
	res = append(res, getReqStruct)

	//list req
	listReqStruct := st.Copy() //ListReqStruct
	listReqStruct.Name = "Get" + tools.ToCamel(st.Name) + "ListReq"
	listReqStruct.Comment = "获取" + st.Comment + "列表请求"
	listReqStruct.TagType = "form"
	for _, field := range listReqStruct.Fields {
		if tools.SliceContain(config.C.Api.IgnoreColumns, field.Name) {
			continue
		}
		if !strings.Contains(field.TagOpt, "optional") {
			field.TagOpt += ",optional"
			switch field.Typ {
			case "int64":
				if tools.IsTimeTypeField(field.Name) {
					field.TagOpt += ",default=0"
				} else {
					field.TagOpt += ",default=-1"
				}
			case "float64":
				field.TagOpt += ",default=-1"
			}
		}
	}
	listReqStruct.Fields = append(GenListReqFields(), listReqStruct.Fields...)
	listReqStruct.Fields.PutTagType("form")
	res = append(res, listReqStruct)

	//enums req
	enumsReqStruct := st.Copy() //EnumsReqStruct
	enumsReqStruct.Name = "Get" + tools.ToCamel(st.Name) + "EnumsReq"
	enumsReqStruct.Comment = "获取" + st.Comment + "枚举请求"
	enumsReqStruct.Fields = []*StructField{
		NewStructField("ParentId", "int64", "json", "parent_id", "optional,default=-1", "父级ID"),
	}
	res = append(res, enumsReqStruct)

	return res
}

type StructCollection []*Struct

func (mc StructCollection) Len() int {
	return len(mc)
}

func (mc StructCollection) Less(i, j int) bool {
	return mc[i].Name < mc[j].Name
}

func (mc StructCollection) Swap(i, j int) {
	mc[i], mc[j] = mc[j], mc[i]
}

type StructField struct {
	Name    string
	Typ     string
	TagType string
	TagName string
	TagOpt  string
	// TagString string
	Comment string
}

func NewStructField(name, typ, tagType, tagName, tagOpt, comment string) *StructField {
	if tagName == "" {
		tagName = name
	}
	if tagType == "" {
		tagType = "json"
	}

	if typ == "time.Time" {
		typ = "int64"
	}

	if typ == "sql.NullString" {
		typ = "string"
	}

	tagName = utils.ConvertStringStyle(config.C.Api.JsonStyle, tagName)
	return &StructField{
		Name:    name,
		Typ:     typ,
		TagType: tagType,
		TagName: tagName,
		TagOpt:  tagOpt,
		// TagString: tagString,
		Comment: comment,
	}
}

func (sf *StructField) GetTagString() string {
	tagOptString := utils.HandleOptContent(sf.TagName, sf.TagOpt)
	return fmt.Sprintf("`%s:\"%s\"`", sf.TagType, tagOptString)
}

func (sf *StructField) Copy() *StructField {
	return NewStructField(sf.Name, sf.Typ, sf.TagType, sf.TagName, sf.TagOpt, sf.Comment)
}

type StructFieldCollection []*StructField

func (mfc StructFieldCollection) String() string {
	buf := new(bytes.Buffer)
	for _, f := range mfc {
		buf.WriteString(fmt.Sprint(f))
	}
	return buf.String()
}

func (mfc StructFieldCollection) Copy() StructFieldCollection {
	var data = make(StructFieldCollection, len(mfc))
	for k, v := range mfc {
		cp := *v
		data[k] = &cp
	}
	return data
}

func (sfc StructFieldCollection) PutTagType(tagType string) StructFieldCollection {
	for k := range sfc {
		sfc[k].TagType = tagType
		// sfc[k].TagString = sfc[k].GetTagString()
	}
	return sfc
}

func GenListReqFields() StructFieldCollection {
	return []*StructField{
		NewStructField("PageSize", "int64", "json", "pageSize", "optional,default=20", "页面容量，默认20，可选"),
		NewStructField("Page", "int64", "json", "page", "optional,default=1", "当前页码，默认1，可选"),
		NewStructField("Current", "int64", "json", "current", "optional,default=1", "当前页码，默认1，用于对接umijs，可选"),
		NewStructField("Keyword", "string", "json", "keyword", "optional", "关键词，可选"),
	}
}

func GenBaseStructCollection() StructCollection {
	return []*Struct{
		NewStruct("Enum", "json", "枚举", StructFieldCollection{
			NewStructField("Label", "interface{}", "json", "label", "", "名"),
			NewStructField("Value", "interface{}", "json", "value", "", "值"),
		}),
		NewStruct("Enums", "json", "枚举列表", StructFieldCollection{
			NewStructField("List", "[]Enum", "json", "list", "", "枚举列表数据"),
		}),
		NewStruct("Option", "json", "选项", StructFieldCollection{
			NewStructField("Title", "string", "json", "title", "", "名"),
			NewStructField("Value", "int64", "json", "value", "", "值"),
		}),
		NewStruct("Options", "json", "选项列表", StructFieldCollection{
			NewStructField("List", "[]Option", "json", "list", "", "选项列表数据"),
		}),
		NewStruct("TreeOption", "json", "树形选项", StructFieldCollection{
			NewStructField("Title", "string", "json", "title", "", "名"),
			NewStructField("Value", "int64", "json", "value", "", "值"),
			NewStructField("Children", "[]TreeOption", "json", "children", "optional", "子集"),
		}),
		NewStruct("TreeOptions", "json", "树形选项列表", StructFieldCollection{
			NewStructField("List", "[]TreeOption", "json", "list", "", "树形选项列表数据"),
		}),
		NewStruct("JwtToken", "json", "jwt token", StructFieldCollection{
			NewStructField("AccessToken", "string", "json", "accessToken", "", "token"),
			NewStructField("AccessExpire", "int64", "json", "accessExpire", "", "expire"),
			NewStructField("RefreshAfter", "int64", "json", "refreshAfter", "", "refresh at"),
		}),
		NewStruct("ListReq", "form", "列表数据请求", GenListReqFields().PutTagType("form")),
		NewStruct("ByIdReq", "form", "通过ID请求", StructFieldCollection{
			NewStructField("Id", "int64", "form", "id", "", "主键"),
		}),
		NewStruct("NilReq", "json", "空请求", nil),
		NewStruct("NilResp", "json", "空响应", nil),
		NewStruct("Resp", "json", "通用数据响应", StructFieldCollection{
			NewStructField("Body", "interface{}", "json", "body", "", "响应数据"),
		}),
		NewStruct("CaptchaResp", "json", "验证码响应", StructFieldCollection{
			NewStructField("CaptchaId", "string", "json", "captchaId", "", "captcha id"),
			NewStructField("ExpiresAt", "int64", "json", "expiresAt", "", "expires time"),
		}),
		NewStruct("BaseResp", "json", "规范响应体", StructFieldCollection{
			NewStructField("Status", "bool", "json", "status", "", "响应状态"),
			NewStructField("Success", "bool", "json", "success", "", "响应状态，用于对接umijs"),
			NewStructField("Message", "string", "json", "message", "optional,omitempty", "给予的提示信息"),
			NewStructField("Data", "interface{}", "json", "data", "optional,omitempty", "【选填】响应的业务数据"),
			NewStructField("Total", "int64", "json", "total", "optional,omitempty", "【选填】数据总个数"),
			NewStructField("PageSize", "int64", "json", "pageSize", "optional,omitempty", "【选填】单页数量"),
			NewStructField("Current", "int64", "json", "current", "optional,omitempty", "【选填】当前页码，用于对接umijs"),
			NewStructField("Page", "int64", "json", "page", "optional,omitempty", "【选填】当前页码"),
			NewStructField("TotalPage", "int64", "json", "totalPage", "optional,omitempty", "【选填】自增项，总共有多少页，根据前端的pageSize来计算"),
			NewStructField("ErrorCode", "int64", "json", "errorCode", "optional,omitempty", "【选填】错误类型代码：400错误请求，401未授权，500服务器内部错误，200成功"),
			NewStructField("ErrorMessage", "string", "json", "errorMessage", "optional,omitempty", "【选填】向用户显示消息"),
			NewStructField("TraceMessage", "string", "json", "traceMessage", "optional,omitempty", "【选填】调试错误信息，请勿在生产环境下使用，可有可无"),
			NewStructField("ShowType", "int64", "json", "showType", "optional,omitempty", "【选填】错误显示类型：0.不提示错误;1.警告信息提示；2.错误信息提示；4.通知提示；9.页面跳转"),
			NewStructField("TraceId", "string", "json", "traceId", "optional,omitempty", "【选填】方便后端故障排除：唯一的请求ID"),
			NewStructField("Host", "string", "json", "host", "optional,omitempty", "【选填】方便后端故障排除：当前访问服务器的主机"),
		}),
	}
}
