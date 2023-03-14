package internal

import (
	"bytes"
	"fmt"

	"github.com/licat233/genzero/tools"
)

type Service struct {
	Name    string
	Comment string
	Rpcs    Rpcs
}

func NewService(name, comment string) *Service {
	s := &Service{
		Name:    tools.ToCamel(name),
		Comment: comment,
		Rpcs:    nil,
	}
	s.initBaseServiceRpcs()
	return s
}

func (s *Service) initBaseServiceRpcs() {
	name := s.Name
	s.Rpcs = []*Rpc{
		NewRpc("BaseAdd"+name, "Add"+name+"Req", "Add"+name+"Resp", "添加"+s.Comment),
		NewRpc("BasePut"+name, "Put"+name+"Req", "Put"+name+"Resp", "更新"+s.Comment),
		NewRpc("BaseGet"+name, "Get"+name+"Req", "Get"+name+"Resp", "获取"+s.Comment),
		NewRpc("BaseDel"+name, "Del"+name+"Req", "Del"+name+"Resp", "删除"+s.Comment),
		NewRpc("BaseGet"+name+"List", "Get"+name+"ListReq", "Get"+name+"ListResp", "获取"+s.Comment+"列表"),
		NewRpc("BaseGet"+name+"Enums", "Get"+name+"EnumsReq", "Enums", "获取"+s.Comment+"枚举列表"),
	}
}

type ServiceCollection []*Service

func (sc ServiceCollection) Len() int {
	return len(sc)
}

func (sc ServiceCollection) Less(i, j int) bool {
	return sc[i].Name < sc[j].Name
}

func (sc ServiceCollection) Swap(i, j int) {
	sc[i], sc[j] = sc[j], sc[i]
}

type Rpc struct {
	Name    string
	Req     string
	Resp    string
	Comment string
}

func NewRpc(name, req, resp, comment string) *Rpc {
	return &Rpc{
		Name:    name,
		Req:     req,
		Resp:    resp,
		Comment: comment,
	}
}

func (r *Rpc) String() string {
	comment := fmt.Sprintf("\n	//%s", r.Comment)
	rpcContent := fmt.Sprintf("\n	rpc %s(%s) returns (%s);", r.Name, r.Req, r.Resp)
	return comment + rpcContent
}

type Rpcs []*Rpc

func (rc Rpcs) Len() int {
	return len(rc)
}

func (rc Rpcs) Less(i, j int) bool {
	return rc[i].Name < rc[j].Name
}

func (rc Rpcs) Swap(i, j int) {
	rc[i], rc[j] = rc[j], rc[i]
}

func (rc Rpcs) String() string {
	var buf = new(bytes.Buffer)
	for _, rpc := range rc {
		buf.WriteString(fmt.Sprint(rpc))
	}
	return buf.String()
}
