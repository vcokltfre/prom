package impl

import (
	"errors"
	"fmt"
	"os"
	"path"
)

type Manager struct {
	Config Config
}

func GetManager() (*Manager, error) {
	manager := &Manager{}

	err := manager.Init()
	if err != nil {
		return nil, err
	}

	return manager, nil
}

func (m *Manager) Init() error {
	config, err := LoadConfig()
	if err != nil {
		return err
	}

	m.Config = config

	return nil
}

func (m *Manager) InitProject(name string, directory string) error {
	if _, found := m.Config.Projects[name]; found {
		return fmt.Errorf("Project %s already exists", name)
	}

	if _, err := os.Stat(directory); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(directory, 0755)
		if err != nil {
			return err
		}
	}

	m.Config.Projects[name] = directory

	err := m.Config.Save()
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Project %s initialised", name))

	return nil
}

func (m *Manager) CloseProject(name string) error {
	if _, found := m.Config.Projects[name]; !found {
		return fmt.Errorf("Project %s does not exist", name)
	}

	directory := m.Config.Projects[name]

	delete(m.Config.Projects, name)
	err := m.Config.Save()
	if err != nil {
		return err
	}

	if _, err := os.Stat(directory); errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("Project directory for %s does not exist (%s). Removing orphan project.", name, directory)
	}

	err = os.Rename(directory, path.Join(m.Config.StaleDir, name))
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Project %s closed", name))

	return nil
}
