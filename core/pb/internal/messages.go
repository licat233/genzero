package internal

import (
	"bytes"
	"fmt"

	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/core/pb/conf"
	"github.com/licat233/genzero/tools"
)

type Message struct {
	Name    string
	Comment string
	Fields  MessageFieldCollection
}

func NewMessage(name, comment string, fields MessageFieldCollection) *Message {
	return &Message{
		Name:    tools.ToCamel(name),
		Comment: comment,
		Fields:  fields,
	}
}

func (ms *Message) Copy() *Message {
	return &Message{
		Name:    ms.Name,
		Comment: ms.Comment,
		Fields:  ms.Fields.Copy(),
	}
}

func (ms *Message) IgnoreMessageFields(needIgnoreFields []string, more ...string) *Message {
	all := append(needIgnoreFields, more...)
	curFields := []*MessageField{}
	var filedTag int
	for _, field := range ms.Fields {
		if tools.HasInSlice(all, field.Name) {
			continue
		}
		filedTag++
		field.Tag = filedTag
		curFields = append(curFields, field)
	}
	ms.Fields = curFields
	return ms
}

func (ms *Message) GenCommonMessages() []*Message {
	//default
	defaultMessage := ms.Copy().IgnoreMessageFields(config.C.Pb.IgnoreColumns)
	//add req
	addReqMessage := ms.Copy().IgnoreMessageFields(config.C.Pb.IgnoreColumns, conf.MoreIgnoreColumns...) //AddReqMessage
	addReqMessage.Name = "Add" + tools.ToCamel(ms.Name) + "Req"
	addReqMessage.Comment = "添加" + ms.Comment + "请求"
	//add resp
	addRespMessage := ms.Copy() //AddRespMessage
	addRespMessage.Name = "Add" + tools.ToCamel(ms.Name) + "Resp"
	addRespMessage.Comment = "添加" + ms.Comment + "响应"
	addRespMessage.Fields = []*MessageField{
		NewMessageField(tools.ToCamel(ms.Name), ms.Name, 1, ms.Comment+"信息"),
	}
	//put req
	putReqMessage := ms.Copy().IgnoreMessageFields(config.C.Pb.IgnoreColumns) //PutReqMessage
	putReqMessage.Name = "Put" + tools.ToCamel(ms.Name) + "Req"
	putReqMessage.Comment = "更新" + ms.Comment + "请求"
	//put resp
	putRespMessage := ms.Copy() //PutRespMessage
	putRespMessage.Name = "Put" + tools.ToCamel(ms.Name) + "Resp"
	putRespMessage.Comment = "更新" + ms.Comment + "响应"
	putRespMessage.Fields = []*MessageField{}
	//del req
	delReqMessage := ms.Copy() //DelReqMessage
	delReqMessage.Name = "Del" + tools.ToCamel(ms.Name) + "Req"
	delReqMessage.Comment = "删除" + ms.Comment + "请求"
	delReqMessage.Fields = []*MessageField{
		NewMessageField("int64", "id", 1, ms.Comment+" ID"),
	}
	//del resp
	delRespMessage := ms.Copy() //DelRespMessage
	delRespMessage.Name = "Del" + tools.ToCamel(ms.Name) + "Resp"
	delRespMessage.Comment = "删除" + ms.Comment + "响应"
	delRespMessage.Fields = []*MessageField{}
	//get req
	getReqMessage := ms.Copy() //GetReqMessage
	getReqMessage.Name = "Get" + tools.ToCamel(ms.Name) + "Req"
	getReqMessage.Comment = "获取" + ms.Comment + "请求"
	getReqMessage.Fields = []*MessageField{
		NewMessageField("int64", "id", 1, ms.Comment+" ID"),
	}
	//get resp
	getRespMessage := ms.Copy() //GetRespMessage
	getRespMessage.Name = "Get" + tools.ToCamel(ms.Name) + "Resp"
	getRespMessage.Comment = "获取" + ms.Comment + "响应"
	getRespMessage.Fields = []*MessageField{
		NewMessageField(tools.ToCamel(ms.Name), ms.Name, 1, ms.Comment+" 信息"),
	}
	//list req
	listReqMessage := ms.Copy() //ListReqMessage
	listReqMessage.Name = "Get" + tools.ToCamel(ms.Name) + "ListReq"
	listReqMessage.Comment = "获取" + ms.Comment + "列表请求"
	listReqMessage.Fields = []*MessageField{
		NewMessageField("ListReq", "list_req", 1, "列表页码参数"),
		NewMessageField(tools.ToCamel(ms.Name), ms.Name, 2, ms.Comment+"参数"),
	}
	//list resp
	listRespMessage := ms.Copy() //ListRespMessage
	listRespMessage.Name = "Get" + tools.ToCamel(ms.Name) + "ListResp"
	listRespMessage.Comment = "获取" + ms.Comment + "列表响应"
	listRespMessage.Fields = []*MessageField{
		NewMessageField("repeated "+tools.ToCamel(ms.Name), tools.PluralizedName(ms.Name), 1, ms.Comment+"列表"),
		NewMessageField("int64", "total", 2, ms.Comment+"总数量"),
	}
	//enums req
	enumsReqMessage := ms.Copy() //EnumsReqMessage
	enumsReqMessage.Name = "Get" + tools.ToCamel(ms.Name) + "EnumsReq"
	enumsReqMessage.Comment = "获取" + ms.Comment + "枚举请求"
	enumsReqMessage.Fields = []*MessageField{
		NewMessageField("int64", "parent_id", 1, "父级ID"),
	}

	res := []*Message{
		defaultMessage,
		addReqMessage,
		addRespMessage,
		putReqMessage,
		putRespMessage,
		delReqMessage,
		delRespMessage,
		getReqMessage,
		getRespMessage,
		listReqMessage,
		listRespMessage,
		enumsReqMessage,
	}
	return res
}

type MessageCollection []*Message

func (mc MessageCollection) Len() int {
	return len(mc)
}

func (mc MessageCollection) Less(i, j int) bool {
	return mc[i].Name < mc[j].Name
}

func (mc MessageCollection) Swap(i, j int) {
	mc[i], mc[j] = mc[j], mc[i]
}

type MessageField struct {
	Typ     string
	Name    string
	Tag     int
	Comment string
}

func NewMessageField(typ, name string, tag int, comment string) *MessageField {
	return &MessageField{
		Typ:     typ,
		Name:    tools.ToSnake(name),
		Tag:     tag,
		Comment: comment,
	}
}

func (f MessageField) String() string {
	return fmt.Sprintf("%s %s = %d; //%s\n", f.Typ, tools.ToSnake(f.Name), f.Tag, f.Comment)
}

type MessageFieldCollection []*MessageField

func (mfc MessageFieldCollection) String() string {
	buf := new(bytes.Buffer)
	for _, f := range mfc {
		buf.WriteString(fmt.Sprint(f))
	}
	return buf.String()
}

func (mfc MessageFieldCollection) Copy() MessageFieldCollection {
	var data = make(MessageFieldCollection, len(mfc))
	for k, v := range mfc {
		cp := *v
		data[k] = &cp
	}
	return data
}

func ListRespFields(dataType string) MessageFieldCollection {
	name := tools.PluralizedName(dataType)
	return MessageFieldCollection{
		NewMessageField("repeated "+dataType, name, 1, "数据列表"),
		NewMessageField("int64", "total", 2, "总数量"),
	}
}

func GenBaseMessages() MessageCollection {
	return []*Message{
		NewMessage("Enum", "枚举", MessageFieldCollection{
			NewMessageField("string", "label", 1, "标签"),
			NewMessageField("int64", "value", 2, "值"),
		}),
		NewMessage("Enums", "枚举列表", MessageFieldCollection{
			NewMessageField("repeated Enum", "enums", 1, "枚举列表数据"),
		}),
		NewMessage("Option", "选项", MessageFieldCollection{
			NewMessageField("string", "title", 1, "标题"),
			NewMessageField("int64", "value", 2, "值"),
		}),
		NewMessage("Options", "选项列表", MessageFieldCollection{
			NewMessageField("repeated Option", "options", 1, "选项列表数据"),
		}),
		NewMessage("TreeOption", "树形选项", MessageFieldCollection{
			NewMessageField("string", "title", 1, "标题"),
			NewMessageField("int64", "value", 2, "值"),
			NewMessageField("repeated TreeOption", "children", 3, "子集"),
		}),
		NewMessage("TreeOptions", "树形选项列表", MessageFieldCollection{
			NewMessageField("repeated TreeOption", "tree_options", 1, "树形选项列表数据"),
		}),
		NewMessage("StatusResp", "状态响应", MessageFieldCollection{
			NewMessageField("bool", "status", 1, "状态"),
		}),
		NewMessage("ListReq", "列表数据请求", MessageFieldCollection{
			NewMessageField("int64", "page_size", 1, "页容量"),
			NewMessageField("int64", "page", 2, "页码"),
			NewMessageField("string", "keyword", 3, "关键词"),
		}),
		NewMessage("ByIdReq", "通过ID请求", MessageFieldCollection{
			NewMessageField("int64", "id", 1, "主键"),
		}),
		NewMessage("NilReq", "空请求", nil),
		NewMessage("NilResp", "空响应", nil),
	}

}
