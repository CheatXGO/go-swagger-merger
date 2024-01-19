package helpers

import (
	"encoding/json"
	"io"
	"os"
)

type Merger struct {
	Swagger map[string]interface{}
	Title   string
	Version string
}

func NewMerger(title, version string) *Merger {
	return &Merger{
		Swagger: make(map[string]interface{}),
		Title:   title,
		Version: version,
	}
}

func (m *Merger) AddFile(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	var s1 interface{}
	err = json.Unmarshal(content, &s1)
	if err != nil {
		return err
	}

	return m.merge(s1.(map[string]interface{}))
}

func (m *Merger) merge(f map[string]interface{}) error {
	for key, item := range f {
		if i, ok := item.(map[string]interface{}); ok {
			for subKey, subitem := range i {
				if _, ok := m.Swagger[key]; !ok {
					m.Swagger[key] = map[string]interface{}{}
				}

				m.Swagger[key].(map[string]interface{})[subKey] = m.checkBaseHeaders(subKey, subitem)
			}
		} else {
			m.Swagger[key] = m.checkBaseHeaders(key, item)
		}
	}

	return nil
}

func (m *Merger) checkBaseHeaders(header string, item interface{}) interface{} {
	switch header {
	case "title":
		return m.Title
	case "version":
		return m.Version
	default:
		return item
	}
}

func (m *Merger) Save(fileName string) error {
	res, _ := json.MarshalIndent(m.Swagger, "", "    ")

	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write(res)
	if err != nil {
		return err
	}

	return nil
}
