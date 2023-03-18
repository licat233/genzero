package internal

import (
	"path"

	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/tools"
)

type Service struct {
	Name    string
	Comment string
	Apis    ApiCollection
	Server  *Server
}

func NewService(name, comment string) *Service {
	s := &Service{
		Name:    name,
		Comment: comment,
		Apis:    ApiCollection{},
		Server:  NewServer(name, config.C.ApiConfig.Jwt, tools.ToLowerCamel(name), config.C.ApiConfig.Middleware, path.Join(config.C.ApiConfig.Prefix, tools.ToLowerCamel(name))),
	}
	s.initBaseApiServiceItems()
	return s
}

func (s *Service) initBaseApiServiceItems() {
	name := tools.ToCamel(s.Name)
	s.Apis = []*Api{
		NewApi("post", "/", "Add"+name, "Add"+name+"Req", "BaseResp", "添加"+s.Comment+" 基础API"),
		NewApi("put", "/", "Put"+name, "Put"+name+"Req", "BaseResp", "更新"+s.Comment+" 基础API"),
		NewApi("get", "/", "Get"+name, "Get"+name+"Req", "BaseResp", "获取"+s.Comment+" 基础API"),
		NewApi("delete", "/", "Del"+name, "Del"+name+"Req", "BaseResp", "删除"+s.Comment+" 基础API"),
		NewApi("get", "/list", "Get"+name+"List", "Get"+name+"ListReq", "BaseResp", "获取"+s.Comment+"列表"+" 基础API"),
		NewApi("get", "/enums", "Get"+name+"Enums", "Get"+name+"EnumsReq", "BaseResp", "获取"+s.Comment+"枚举列表"+" 基础API"),
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

type Api struct {
	Method  string
	Path    string
	Handler string
	Req     string
	Resp    string
	Comment string
}

func NewApi(method, path, handler, req, resp, comment string) *Api {
	if path == "" {
		path = "/"
	}
	return &Api{
		Method:  method,
		Path:    path,
		Handler: handler,
		Req:     req,
		Resp:    resp,
		Comment: comment,
	}
}

type ApiCollection []*Api

func (ac ApiCollection) Len() int {
	return len(ac)
}

func (ac ApiCollection) Less(i, j int) bool {
	return ac[i].Path < ac[j].Path
}

func (ac ApiCollection) Swap(i, j int) {
	ac[i], ac[j] = ac[j], ac[i]
}
