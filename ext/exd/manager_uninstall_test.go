package exd_test

import (
	"context"
	"errors"

	tMock "github.com/stretchr/testify/mock"

	"github.com/odpf/optimus/ext/exd"
	"github.com/odpf/optimus/mock"
)

func (m *ManagerTestSuite) TestUninstall() {
	ctx := context.Background()
	httpDoer := &mock.HTTPDoer{}
	verbose := true

	m.Run("should return error if one or more required fields are empty", func() {
		manager := &exd.Manager{}
		commandName := "valor"
		tagName := "v1.0"

		actualErr := manager.Uninstall(commandName, tagName)

		m.Error(actualErr)
	})

	m.Run("should return error if command name is empty", func() {
		manifester := &mock.Manifester{}
		assetOperator := &mock.AssetOperator{}

		manager, err := exd.NewManager(ctx, httpDoer, manifester, assetOperator, verbose)
		if err != nil {
			panic(err)
		}

		var commandName string
		tagName := "v1.0"

		actualErr := manager.Uninstall(commandName, tagName)

		m.Error(actualErr)
	})

	m.Run("should return error if error loading manifest", func() {
		assetOperator := &mock.AssetOperator{}

		manifester := &mock.Manifester{}
		manifester.On("Load", tMock.Anything).Return(nil, errors.New("random error"))

		manager, err := exd.NewManager(ctx, httpDoer, manifester, assetOperator, verbose)
		if err != nil {
			panic(err)
		}

		commandName := "valor"
		tagName := "v1.0"

		actualErr := manager.Uninstall(commandName, tagName)

		m.Error(actualErr)
	})

	m.Run("should return error if command name is not found", func() {
		assetOperator := &mock.AssetOperator{}

		release := &exd.RepositoryRelease{
			TagName: "v1.0",
		}
		project := &exd.RepositoryProject{
			CommandName:   "valor",
			ActiveTagName: "v1.0",
			Releases:      []*exd.RepositoryRelease{release},
		}
		release.Project = project
		owner := &exd.RepositoryOwner{
			Projects: []*exd.RepositoryProject{project},
		}
		project.Owner = owner
		manifest := &exd.Manifest{
			RepositoryOwners: []*exd.RepositoryOwner{owner},
		}
		manifester := &mock.Manifester{}
		manifester.On("Load", tMock.Anything).Return(manifest, nil)

		manager, err := exd.NewManager(ctx, httpDoer, manifester, assetOperator, verbose)
		if err != nil {
			panic(err)
		}

		commandName := "valor2"
		tagName := "v1.0"

		actualErr := manager.Uninstall(commandName, tagName)

		m.Error(actualErr)
	})

	m.Run("should return error if tag name is not found", func() {
		assetOperator := &mock.AssetOperator{}

		release := &exd.RepositoryRelease{
			TagName: "v1.0",
		}
		project := &exd.RepositoryProject{
			CommandName:   "valor",
			ActiveTagName: "v1.0",
			Releases:      []*exd.RepositoryRelease{release},
		}
		release.Project = project
		owner := &exd.RepositoryOwner{
			Projects: []*exd.RepositoryProject{project},
		}
		project.Owner = owner
		manifest := &exd.Manifest{
			RepositoryOwners: []*exd.RepositoryOwner{owner},
		}
		manifester := &mock.Manifester{}
		manifester.On("Load", tMock.Anything).Return(manifest, nil)

		manager, err := exd.NewManager(ctx, httpDoer, manifester, assetOperator, verbose)
		if err != nil {
			panic(err)
		}

		commandName := "valor"
		tagName := "v1.1"

		actualErr := manager.Uninstall(commandName, tagName)

		m.Error(actualErr)
	})

	m.Run("should return error if error encountered during preparation", func() {
		release := &exd.RepositoryRelease{
			TagName: "v1.0",
		}
		project := &exd.RepositoryProject{
			CommandName:   "valor",
			ActiveTagName: "v1.0",
			Releases:      []*exd.RepositoryRelease{release},
		}
		release.Project = project
		owner := &exd.RepositoryOwner{
			Projects: []*exd.RepositoryProject{project},
		}
		project.Owner = owner
		manifest := &exd.Manifest{
			RepositoryOwners: []*exd.RepositoryOwner{owner},
		}
		manifester := &mock.Manifester{}
		manifester.On("Load", tMock.Anything).Return(manifest, nil)

		assetOperator := &mock.AssetOperator{}
		assetOperator.On("Prepare", tMock.Anything).Return(errors.New("random error"))

		manager, err := exd.NewManager(ctx, httpDoer, manifester, assetOperator, verbose)
		if err != nil {
			panic(err)
		}

		commandName := "valor"
		tagName := "v1.0"

		actualErr := manager.Uninstall(commandName, tagName)

		m.Error(actualErr)
	})

	m.Run("should return error if error encountered during uninstallation", func() {
		release := &exd.RepositoryRelease{
			TagName: "v1.0",
		}
		project := &exd.RepositoryProject{
			CommandName:   "valor",
			ActiveTagName: "v1.0",
			Releases:      []*exd.RepositoryRelease{release},
		}
		release.Project = project
		owner := &exd.RepositoryOwner{
			Projects: []*exd.RepositoryProject{project},
		}
		project.Owner = owner
		manifest := &exd.Manifest{
			RepositoryOwners: []*exd.RepositoryOwner{owner},
		}
		manifester := &mock.Manifester{}
		manifester.On("Load", tMock.Anything).Return(manifest, nil)

		assetOperator := &mock.AssetOperator{}
		assetOperator.On("Prepare", tMock.Anything).Return(nil)
		assetOperator.On("Uninstall", tMock.Anything).Return(errors.New("random error"))

		manager, err := exd.NewManager(ctx, httpDoer, manifester, assetOperator, verbose)
		if err != nil {
			panic(err)
		}

		commandName := "valor"
		tagName := "v1.0"

		actualErr := manager.Uninstall(commandName, tagName)

		m.Error(actualErr)
	})

	m.Run("should return error if error encountered during updating manifest", func() {
		release := &exd.RepositoryRelease{
			TagName: "v1.0",
		}
		project := &exd.RepositoryProject{
			CommandName:   "valor",
			ActiveTagName: "v1.0",
			Releases:      []*exd.RepositoryRelease{release},
		}
		release.Project = project
		owner := &exd.RepositoryOwner{
			Projects: []*exd.RepositoryProject{project},
		}
		project.Owner = owner
		manifest := &exd.Manifest{
			RepositoryOwners: []*exd.RepositoryOwner{owner},
		}
		manifester := &mock.Manifester{}
		manifester.On("Load", tMock.Anything).Return(manifest, nil)
		manifester.On("Flush", tMock.Anything, tMock.Anything).Return(errors.New("random error"))

		assetOperator := &mock.AssetOperator{}
		assetOperator.On("Prepare", tMock.Anything).Return(nil)
		assetOperator.On("Uninstall", tMock.Anything).Return(nil)

		manager, err := exd.NewManager(ctx, httpDoer, manifester, assetOperator, verbose)
		if err != nil {
			panic(err)
		}

		commandName := "valor"
		tagName := "v1.0"

		actualErr := manager.Uninstall(commandName, tagName)

		m.Error(actualErr)
	})

	m.Run("should return nil if no error encountered during the whole process", func() {
		release := &exd.RepositoryRelease{
			TagName: "v1.0",
		}
		project := &exd.RepositoryProject{
			CommandName:   "valor",
			ActiveTagName: "v1.0",
			Releases:      []*exd.RepositoryRelease{release},
		}
		release.Project = project
		owner := &exd.RepositoryOwner{
			Projects: []*exd.RepositoryProject{project},
		}
		project.Owner = owner
		manifest := &exd.Manifest{
			RepositoryOwners: []*exd.RepositoryOwner{owner},
		}
		manifester := &mock.Manifester{}
		manifester.On("Load", tMock.Anything).Return(manifest, nil)
		manifester.On("Flush", tMock.Anything, tMock.Anything).Return(nil)

		assetOperator := &mock.AssetOperator{}
		assetOperator.On("Prepare", tMock.Anything).Return(nil)
		assetOperator.On("Uninstall", tMock.Anything).Return(nil)

		manager, err := exd.NewManager(ctx, httpDoer, manifester, assetOperator, verbose)
		if err != nil {
			panic(err)
		}

		commandName := "valor"
		tagName := "v1.0"

		actualErr := manager.Uninstall(commandName, tagName)

		m.NoError(actualErr)
	})
}
