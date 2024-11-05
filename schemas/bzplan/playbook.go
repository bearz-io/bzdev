package playbook

import (
	"github.com/bearz-io/bzdev/lib/errors"
	"github.com/bearz-io/bzdev/schemas/primitives"
	"gopkg.in/yaml.v3"
)

type Playbook struct {
	Id    string               `json:"id"`
	Name  *string              `json:"name"`
	Tasks []primitives.Task    `json:"tasks"`
	Env   *primitives.ExprsMap `json:"env"`
}

func (pb *Playbook) UnmarshalYAML(node *yaml.Node) error {
	if node.Kind != yaml.MappingNode {
		return errors.NewYamlError(node, "playbook", "playbook must be a mapping node")
	}

	var playbook struct {
		Id    string               `json:"id"`
		Name  *string              `json:"name"`
		Tasks []primitives.Task    `json:"tasks"`
		Env   *primitives.ExprsMap `json:"env"`
	}

	if err := node.Decode(&playbook); err != nil {
		return err
	}

	pb.Id = playbook.Id
	pb.Name = playbook.Name
	pb.Tasks = playbook.Tasks
	pb.Env = playbook.Env
	if pb.Env == nil {
		pb.Env = &primitives.ExprsMap{}
	}

	return nil
}
