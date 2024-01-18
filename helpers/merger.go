package helpers

import (
	"io"
	"os"

	jsoniter "github.com/json-iterator/go"
)

type Merger struct {
	Swagger map[string]interface{}
}

func NewMerger() *Merger {
	merger := new(Merger)
	merger.Swagger = map[string]interface{}{}
	return merger
}

func (m *Merger) AddFile(file, title string) error {
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
	err = jsoniter.Unmarshal(content, &s1)
	if err != nil {
		return err
	}

	return m.merge(s1.(map[string]interface{}), title)
}

func (m *Merger) merge(f map[string]interface{}, title string) error {
	for key, item := range f {
		if i, ok := item.(map[string]interface{}); ok {
			for subKey, subitem := range i {
				if _, ok := m.Swagger[key]; !ok {
					m.Swagger[key] = map[string]interface{}{}
				}
				if subKey == "title" {
					subitem = title
				}

				m.Swagger[key].(map[string]interface{})[subKey] = subitem
			}
		} else {
			if key == "title" {
				item = title
			}
			m.Swagger[key] = item
		}
	}

	return nil
}

func (m *Merger) Save(fileName string) error {
	res, _ := jsoniter.MarshalIndent(m.Swagger, "", "    ")

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
