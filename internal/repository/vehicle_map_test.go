package repository_test

import (
	"app/internal"
	"app/internal/repository"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindAll(t *testing.T) {
	db := map[int]internal.Vehicle{
		1: {
			Id: 1,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           "sitroein",
				Model:           "c3",
				Registration:    "gfsdas",
				Color:           "black",
				FabricationYear: 2031,
				Capacity:        7,
				MaxSpeed:        120,
				FuelType:        "gas",
				Transmission:    "manual",
				Weight:          600,
				Dimensions: internal.Dimensions{
					Height: 170,
					Length: 200,
					Width:  160,
				},
			},
		},
		2: {
			Id: 2,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           "sitroein",
				Model:           "cactus",
				Registration:    "gfsdas",
				Color:           "green",
				FabricationYear: 2032,
				Capacity:        4,
				MaxSpeed:        190,
				FuelType:        "etanol",
				Transmission:    "auto",
				Weight:          800,
				Dimensions: internal.Dimensions{
					Height: 270,
					Length: 400,
					Width:  660,
				},
			},
		},
	}

	repo := repository.NewRepositoryReadVehicleMap(db)

	t.Run("success - returns vehicles match", func(t *testing.T) {
		expected := db
		actual, err := repo.FindAll()

		require.Equal(t, expected, actual)
		require.NoError(t, err)
	})
	t.Run("fail - not found", func(t *testing.T) {
		rp := repository.NewRepositoryReadVehicleMap(map[int]internal.Vehicle{})
		actual, err := rp.FindAll()
		expected := map[int]internal.Vehicle{}

		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})
}

func TestFindByColorAndYear(t *testing.T) {
	db := map[int]internal.Vehicle{
		1: {
			Id: 1,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           "sitroein",
				Model:           "c3",
				Registration:    "gfsdas",
				Color:           "black",
				FabricationYear: 2031,
				Capacity:        7,
				MaxSpeed:        120,
				FuelType:        "gas",
				Transmission:    "manual",
				Weight:          600,
				Dimensions: internal.Dimensions{
					Height: 170,
					Length: 200,
					Width:  160,
				},
			},
		},
		2: {
			Id: 2,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           "sitroein",
				Model:           "cactus",
				Registration:    "gfsdas",
				Color:           "green",
				FabricationYear: 2032,
				Capacity:        4,
				MaxSpeed:        190,
				FuelType:        "etanol",
				Transmission:    "auto",
				Weight:          800,
				Dimensions: internal.Dimensions{
					Height: 270,
					Length: 400,
					Width:  660,
				},
			},
		},
	}

	repo := repository.NewRepositoryReadVehicleMap(db)

	t.Run("success - returns vehicles match", func(t *testing.T) {
		expected := map[int]internal.Vehicle{2: db[2]}
		actual, err := repo.FindByColorAndYear("green", 2032)

		require.Equal(t, expected, actual)
		require.NoError(t, err)
	})
	t.Run("fail - not found", func(t *testing.T) {
		actual, err := repo.FindByColorAndYear("azu", 1990)
		expected := map[int]internal.Vehicle{}

		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})
}

func TestFindByBrandAndYearRange(t *testing.T) {
	db := map[int]internal.Vehicle{
		1: {
			Id: 1,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           "sitroein",
				Model:           "c3",
				Registration:    "gfsdas",
				Color:           "black",
				FabricationYear: 2031,
				Capacity:        7,
				MaxSpeed:        120,
				FuelType:        "gas",
				Transmission:    "manual",
				Weight:          600,
				Dimensions: internal.Dimensions{
					Height: 170,
					Length: 200,
					Width:  160,
				},
			},
		},
		2: {
			Id: 2,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           "sitroein",
				Model:           "cactus",
				Registration:    "gfsdas",
				Color:           "green",
				FabricationYear: 2032,
				Capacity:        4,
				MaxSpeed:        190,
				FuelType:        "etanol",
				Transmission:    "auto",
				Weight:          800,
				Dimensions: internal.Dimensions{
					Height: 270,
					Length: 400,
					Width:  660,
				},
			},
		},
	}

	repo := repository.NewRepositoryReadVehicleMap(db)

	t.Run("success - returns vehicles match", func(t *testing.T) {
		expected := map[int]internal.Vehicle{2: db[2]}
		actual, err := repo.FindByBrandAndYearRange("sitroein", 2032, 2035)

		require.Equal(t, expected, actual)
		require.NoError(t, err)
	})
	t.Run("fail - not found", func(t *testing.T) {
		actual, err := repo.FindByBrandAndYearRange("pejou", 1990, 2000)
		expected := map[int]internal.Vehicle{}

		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})
}

func TestFindByBrand(t *testing.T) {
	db := map[int]internal.Vehicle{
		1: {
			Id: 1,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           "sitroein",
				Model:           "c3",
				Registration:    "gfsdas",
				Color:           "black",
				FabricationYear: 2031,
				Capacity:        7,
				MaxSpeed:        120,
				FuelType:        "gas",
				Transmission:    "manual",
				Weight:          600,
				Dimensions: internal.Dimensions{
					Height: 170,
					Length: 200,
					Width:  160,
				},
			},
		},
		2: {
			Id: 2,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           "sitroein",
				Model:           "cactus",
				Registration:    "gfsdas",
				Color:           "green",
				FabricationYear: 2032,
				Capacity:        4,
				MaxSpeed:        190,
				FuelType:        "etanol",
				Transmission:    "auto",
				Weight:          800,
				Dimensions: internal.Dimensions{
					Height: 270,
					Length: 400,
					Width:  660,
				},
			},
		},
	}

	repo := repository.NewRepositoryReadVehicleMap(db)

	t.Run("success - returns vehicles match", func(t *testing.T) {
		expected := db
		actual, err := repo.FindByBrand("sitroein")

		require.Equal(t, expected, actual)
		require.NoError(t, err)
	})
	t.Run("fail - not found", func(t *testing.T) {
		actual, err := repo.FindByBrand("pejou")
		expected := map[int]internal.Vehicle{}

		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})
}

func TestFindByWeightRange(t *testing.T) {
	db := map[int]internal.Vehicle{
		1: {
			Id: 1,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           "sitroein",
				Model:           "c3",
				Registration:    "gfsdas",
				Color:           "black",
				FabricationYear: 2031,
				Capacity:        7,
				MaxSpeed:        120,
				FuelType:        "gas",
				Transmission:    "manual",
				Weight:          600,
				Dimensions: internal.Dimensions{
					Height: 170,
					Length: 200,
					Width:  160,
				},
			},
		},
		2: {
			Id: 2,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           "sitroein",
				Model:           "cactus",
				Registration:    "gfsdas",
				Color:           "green",
				FabricationYear: 2032,
				Capacity:        4,
				MaxSpeed:        190,
				FuelType:        "etanol",
				Transmission:    "auto",
				Weight:          800,
				Dimensions: internal.Dimensions{
					Height: 270,
					Length: 400,
					Width:  660,
				},
			},
		},
	}

	repo := repository.NewRepositoryReadVehicleMap(db)

	t.Run("success - returns vehicles match", func(t *testing.T) {
		expected := map[int]internal.Vehicle{2: db[2]}
		actual, err := repo.FindByWeightRange(700, 900)

		require.Equal(t, expected, actual)
		require.NoError(t, err)
	})
	t.Run("fail - not found", func(t *testing.T) {
		actual, err := repo.FindByWeightRange(100, 150)
		expected := map[int]internal.Vehicle{}

		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})
}
