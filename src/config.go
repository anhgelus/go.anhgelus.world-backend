package src

import (
	"bytes"
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"os"
)

var cfg Config

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
	return fmt.Sprintf("%s%s", l.Origin, l.Uri)
}

func LoadConfig(path string) {
	val, err := os.ReadFile(path)
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	} else if os.IsNotExist(err) {
		createFile(path)
		return
	}
	err = toml.Unmarshal(val, &cfg)
	if err != nil {
		panic(err)
	}
}

func createFile(path string) {
	var def Config
	var red []Redirect
	red = append(red, Redirect{"lmdln", "lemondedelanuit"})
	red = append(red, Redirect{"b", "blog"})
	red = append(red, Redirect{"c", "cloud"})
	def.Redirections = red
	cfg = def
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bytes.Buffer{}
	enc := toml.NewEncoder(&buf)
	enc.SetIndentTables(true)
	err = enc.Encode(def)
	if err != nil {
		panic(err)
	}
	_, err = f.Write(buf.Bytes())
}
