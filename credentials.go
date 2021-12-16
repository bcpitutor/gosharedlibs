package gosharedlibs

import (
	"fmt"

	"gopkg.in/ini.v1"
)

func UpdateAWSCredentials(sectionName string, vars map[string]string) error {

	home, err := HomeFolder()
	if err != nil {
		return err
	}

	fpath := fmt.Sprintf("%s/.aws/credentials", home)
	cfg, err := ini.Load(fpath)
	cfg.SaveTo(fpath + ".bak")

	if err != nil {
		return err
	}

	section, err := cfg.NewSection(sectionName)
	if err != nil {
		return err
	}

	for k, v := range vars {
		section.Key(k).SetValue(v)
	}

	cfg.SaveTo(fpath)

	return nil
}
