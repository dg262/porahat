package servicesinitializer

import (
	"context"
	"flower-management/api/rest"
	"flower-management/internal/core/config"
	"flower-management/internal/core/servicecore"
	persistency "flower-management/internal/persistency/contracts"
	"flower-management/internal/persistency/dal"
	"flower-management/internal/persistency/mock"
)

func Execute(envFilename string) {
	configSet, err := config.LoadConfig(envFilename)
	if err != nil {
		panic(err)
	}

	var dalInstance persistency.DalInterface
	if configSet.Mocks.DalMocked {
		dalInstance = mock.NewDalMock()
	} else {
		dalInstance, err = dal.NewDal(context.Background(), configSet.DalConfig)
		if err != nil {
			panic(err)
		}
	}

	servicecore := servicecore.NewServiceCore(dalInstance)

	restServer := rest.NewRestServer(configSet.RestServerConfig, servicecore)
	if err := restServer.Start(); err != nil {
		panic(err)
	}

	select {}
}
