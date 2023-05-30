package internal

import (
	"bytes"
	"strings"
)

type Server struct {
	Name       string
	Jwt        string
	Group      string
	Middleware []string
	Prefix     string
}

func NewServer(name string, jwt string, group string, middleware []string, prefix string) *Server {
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
	if len(s.Middleware) != 0 {
		exists = true
		buf.WriteString("  middleware: " + strings.Join(s.Middleware, ",") + "\n")
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
