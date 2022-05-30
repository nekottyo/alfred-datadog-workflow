package dd

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"

	aw "github.com/deanishe/awgo"
)

type Service struct {
	svcs []serviceMap
	wf   *aw.Workflow
}

type serviceMap struct {
	Url   string `yaml:"url"`
	Title string `yaml:"title"`
}

func NewServices(path string, wf *aw.Workflow) (Service, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return Service{}, err
	}

	svcs := make([]serviceMap, 20)
	if err := yaml.Unmarshal(buf, &svcs); err != nil {
		return Service{}, err
	}

	return Service{
		svcs: svcs,
		wf:   wf,
	}, nil
}

func (s *Service) ListServices() error {
	for _, svc := range s.svcs {
		s.wf.NewItem(svc.Title).
			Subtitle(svc.Url).
			Arg(svc.Url).
			UID(svc.Title).
			Valid(true)
	}
	return nil
}
