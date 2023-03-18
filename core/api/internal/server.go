package internal

import (
	"bytes"
)

type Server struct {
	Name       string
	Jwt        string
	Group      string
	Middleware string
	Prefix     string
}

func NewServer(name, jwt, group, middleware, prefix string) *Server {
	return &Server{
		Name:       name,
		Jwt:        jwt,
		Group:      group,
		Middleware: middleware,
		Prefix:     prefix,
	}
}

func (s *Server) String() string {
	var exists bool
	var buf bytes.Buffer
	buf.WriteString("@server(\n")
	if s.Jwt != "" {
		exists = true
		buf.WriteString("  jwt: " + s.Jwt + "\n")
	}
	if s.Group != "" {
		exists = true
		buf.WriteString("  group: " + s.Group + "\n")
	}
	if s.Middleware != "" {
		exists = true
		buf.WriteString("  middleware: " + s.Middleware + "\n")
	}
	if s.Prefix != "" {
		exists = true
		buf.WriteString("  prefix: " + s.Prefix + "\n")
	}
	buf.WriteString(")")
	if exists {
		return buf.String()
	}
	return ""
}
