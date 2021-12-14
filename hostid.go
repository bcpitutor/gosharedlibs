package gosharedlibs

import (
	"github.com/denisbrodbeck/machineid"
)

func GetMachineId() (string, error) {
	id, err := machineid.ID()
	if err != nil {
		return "", err
	}
	return id, nil
}

func GetAppKey(appName string) (string, error) {
	id, err := machineid.ProtectedID(appName)
	if err != nil {
		return "", err
	}

	return id, nil
}
