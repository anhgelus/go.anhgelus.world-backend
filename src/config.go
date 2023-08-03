package src

import "fmt"

type Config struct {
	Redirections []Redirect
}

type Redirect struct {
	Id   string
	Path Path
}

type Path string

type Location struct {
	Origin string
	Uri    string
}

func (p *Path) generateOrigin() string {
	return fmt.Sprintf("https://%s.anhgelus.world/", *p)
}

func (l *Location) generateUrl() string {
	return l.Origin + l.Uri
}
