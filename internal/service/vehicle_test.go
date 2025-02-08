package service_test

import (
	"app/internal"
	"app/internal/service"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) FindAll() (v map[int]internal.Vehicle, err error) {
	args := r.Called()
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

func (r *RepositoryMock) FindByColorAndYear(color string, fabricationYear int) (v map[int]internal.Vehicle, err error) {
	args := r.Called(color, fabricationYear)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

func (r *RepositoryMock) FindByBrandAndYearRange(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error) {
	args := r.Called(brand, startYear, endYear)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

func (r *RepositoryMock) FindByBrand(brand string) (v map[int]internal.Vehicle, err error) {
	args := r.Called(brand)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

func (r *RepositoryMock) FindByWeightRange(fromWeight float64, toWeight float64) (v map[int]internal.Vehicle, err error) {
	args := r.Called(fromWeight, toWeight)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

func TestFindByColorAndYear(t *testing.T) {
	t.Run("success - returns vehicles match", func(t *testing.T) {
		repo := new(RepositoryMock)
		expected := map[int]internal.Vehicle{
			1: {
				Id: 1,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Pejou",
					Model:           "204",
					Registration:    "Sachi",
					Color:           "BLue",
					FabricationYear: 2023,
					Capacity:        3,
					MaxSpeed:        200,
					FuelType:        "gas",
					Transmission:    "auto",
					Weight:          2000,
					Dimensions: internal.Dimensions{
						Height: 140,
						Length: 154,
						Width:  172,
					},
				},
			},
		}
		repo.On("FindByColorAndYear", "BLue", 2023).Return(expected, nil)
		s := service.NewServiceVehicleDefault(repo)

		actual, err := s.FindByColorAndYear("BLue", 2023)

		require.Equal(t, expected, actual)
		require.NoError(t, err)
	})

	t.Run("fail - not found", func(t *testing.T) {
		repo := new(RepositoryMock)
		expected := map[int]internal.Vehicle{}
		repo.On("FindByColorAndYear", "Green", 2000).Return(expected, nil)
		s := service.NewServiceVehicleDefault(repo)

		actual, err := s.FindByColorAndYear("Green", 2000)

		require.Equal(t, expected, actual)
		require.NoError(t, err)
	})
}

func TestAverageMaxSpeedByBrand(t *testing.T) {
	t.Run("success - returns speed average", func(t *testing.T) {
		repo := new(RepositoryMock)
		expected := map[int]internal.Vehicle{
			1: {
				Id: 1,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Pejou",
					Model:           "204",
					Registration:    "Sachi",
					Color:           "BLue",
					FabricationYear: 2023,
					Capacity:        3,
					MaxSpeed:        200,
					FuelType:        "gas",
					Transmission:    "auto",
					Weight:          2000,
					Dimensions: internal.Dimensions{
						Height: 140,
						Length: 154,
						Width:  172,
					},
				},
			},
			2: {
				Id: 2,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Pejou",
					Model:           "3008",
					Registration:    "chassi",
					Color:           "Green",
					FabricationYear: 2024,
					Capacity:        6,
					MaxSpeed:        250,
					FuelType:        "gas",
					Transmission:    "manual",
					Weight:          2000,
					Dimensions: internal.Dimensions{
						Height: 140,
						Length: 154,
						Width:  172,
					},
				},
			},
		}
		expectedSpeed := 225
		repo.On("FindByBrand", "Pejou").Return(expected, nil)
		s := service.NewServiceVehicleDefault(repo)

		actual, err := s.AverageMaxSpeedByBrand("Pejou")

		require.Equal(t, float64(expectedSpeed), actual)
		require.NoError(t, err)
	})

	t.Run("fail - not found", func(t *testing.T) {
		repo := new(RepositoryMock)
		expected := map[int]internal.Vehicle{}
		expectedSpeed := 0
		repo.On("FindByBrand", "Pejou").Return(expected, nil)
		s := service.NewServiceVehicleDefault(repo)

		actual, err := s.AverageMaxSpeedByBrand("Pejou")

		require.Equal(t, float64(expectedSpeed), actual)
		require.ErrorIs(t, err, internal.ErrServiceNoVehicles)
	})
}

func TestFindByWeightRange(t *testing.T) {
	t.Run("success - returns vehicles match", func(t *testing.T) {
		repo := new(RepositoryMock)
		expected := map[int]internal.Vehicle{
			1: {
				Id: 1,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Pejou",
					Model:           "204",
					Registration:    "Sachi",
					Color:           "BLue",
					FabricationYear: 2023,
					Capacity:        3,
					MaxSpeed:        200,
					FuelType:        "gas",
					Transmission:    "auto",
					Weight:          275,
					Dimensions: internal.Dimensions{
						Height: 140,
						Length: 154,
						Width:  172,
					},
				},
			},
			2: {
				Id: 2,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Fiat",
					Model:           "Uno",
					Registration:    "Sachi",
					Color:           "BLue",
					FabricationYear: 2023,
					Capacity:        3,
					MaxSpeed:        200,
					FuelType:        "gas",
					Transmission:    "auto",
					Weight:          250,
					Dimensions: internal.Dimensions{
						Height: 140,
						Length: 154,
						Width:  172,
					},
				},
			},
			3: {
				Id: 3,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Volkswagem",
					Model:           "Golf",
					Registration:    "Toma",
					Color:           "Gurin",
					FabricationYear: 2024,
					Capacity:        3,
					MaxSpeed:        200,
					FuelType:        "gas",
					Transmission:    "auto",
					Weight:          200,
					Dimensions: internal.Dimensions{
						Height: 140,
						Length: 154,
						Width:  172,
					},
				},
			},
		}
		query := internal.SearchQuery{
			FromWeight: 200.0,
			ToWeight:   300.0,
		}
		repo.On("FindByWeightRange", 200.0, 300.0).Return(expected, nil)
		s := service.NewServiceVehicleDefault(repo)

		actual, err := s.SearchByWeightRange(query, true)

		require.Equal(t, expected, actual)
		require.NoError(t, err)
	})

	t.Run("success - without query, returns all", func(t *testing.T) {
		repo := new(RepositoryMock)
		expected := map[int]internal.Vehicle{
			1: {
				Id: 1,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Pejou",
					Model:           "204",
					Registration:    "Sachi",
					Color:           "BLue",
					FabricationYear: 2023,
					Capacity:        3,
					MaxSpeed:        200,
					FuelType:        "gas",
					Transmission:    "auto",
					Weight:          275,
					Dimensions: internal.Dimensions{
						Height: 140,
						Length: 154,
						Width:  172,
					},
				},
			},
			2: {
				Id: 2,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Fiat",
					Model:           "Uno",
					Registration:    "Sachi",
					Color:           "BLue",
					FabricationYear: 2023,
					Capacity:        3,
					MaxSpeed:        200,
					FuelType:        "gas",
					Transmission:    "auto",
					Weight:          250,
					Dimensions: internal.Dimensions{
						Height: 140,
						Length: 154,
						Width:  172,
					},
				},
			},
			3: {
				Id: 3,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Volkswagem",
					Model:           "Golf",
					Registration:    "Toma",
					Color:           "Gurin",
					FabricationYear: 2024,
					Capacity:        3,
					MaxSpeed:        200,
					FuelType:        "gas",
					Transmission:    "auto",
					Weight:          200,
					Dimensions: internal.Dimensions{
						Height: 140,
						Length: 154,
						Width:  172,
					},
				},
			},
		}
		repo.On("FindAll").Return(expected, nil)
		s := service.NewServiceVehicleDefault(repo)

		actual, err := s.SearchByWeightRange(internal.SearchQuery{}, false)

		require.Equal(t, expected, actual)
		require.NoError(t, err)
	})

	t.Run("fail - not found", func(t *testing.T) {
		repo := new(RepositoryMock)
		expected := map[int]internal.Vehicle{}
		query := internal.SearchQuery{
			FromWeight: 50.0,
			ToWeight:   100.0,
		}
		repo.On("FindByWeightRange", 50.0, 100.0).Return(expected, nil)
		s := service.NewServiceVehicleDefault(repo)

		actual, err := s.SearchByWeightRange(query, true)

		require.Equal(t, expected, actual)
		require.NoError(t, err)
	})
}

func TestAverageCapacityByBrand(t *testing.T) {
	t.Run("success - returns capacity average", func(t *testing.T) {
		repo := new(RepositoryMock)
		expected := map[int]internal.Vehicle{
			1: {
				Id: 1,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Pejou",
					Model:           "204",
					Registration:    "Sachi",
					Color:           "BLue",
					FabricationYear: 2023,
					Capacity:        4,
					MaxSpeed:        200,
					FuelType:        "gas",
					Transmission:    "auto",
					Weight:          2000,
					Dimensions: internal.Dimensions{
						Height: 140,
						Length: 154,
						Width:  172,
					},
				},
			},
			2: {
				Id: 2,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Pejou",
					Model:           "3008",
					Registration:    "chassi",
					Color:           "Green",
					FabricationYear: 2024,
					Capacity:        6,
					MaxSpeed:        250,
					FuelType:        "gas",
					Transmission:    "manual",
					Weight:          2000,
					Dimensions: internal.Dimensions{
						Height: 140,
						Length: 154,
						Width:  172,
					},
				},
			},
		}
		expectedCapacity := 5
		repo.On("FindByBrand", "Pejou").Return(expected, nil)
		s := service.NewServiceVehicleDefault(repo)

		actual, err := s.AverageCapacityByBrand("Pejou")

		require.Equal(t, expectedCapacity, actual)
		require.NoError(t, err)
	})

	t.Run("fail - not found", func(t *testing.T) {
		repo := new(RepositoryMock)
		expected := map[int]internal.Vehicle{}
		expectedCapacity := 0
		repo.On("FindByBrand", "Pejou").Return(expected, nil)
		s := service.NewServiceVehicleDefault(repo)

		actual, err := s.AverageMaxSpeedByBrand("Pejou")

		require.Equal(t, float64(expectedCapacity), actual)
		require.ErrorIs(t, err, internal.ErrServiceNoVehicles)
	})
}

func TestFindByBrandAndYearRange(t *testing.T) {
	t.Run("success - returns vehicles match", func(t *testing.T) {
		repo := new(RepositoryMock)
		expected := map[int]internal.Vehicle{
			1: {
				Id: 1,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Pejou",
					Model:           "204",
					Registration:    "Sachi",
					Color:           "BLue",
					FabricationYear: 2023,
					Capacity:        3,
					MaxSpeed:        200,
					FuelType:        "gas",
					Transmission:    "auto",
					Weight:          2000,
					Dimensions: internal.Dimensions{
						Height: 140,
						Length: 154,
						Width:  172,
					},
				},
			},
		}
		repo.On("FindByBrandAndYearRange", "BLue", 2000, 2025).Return(expected, nil)
		s := service.NewServiceVehicleDefault(repo)

		actual, err := s.FindByBrandAndYearRange("BLue", 2000, 2025)

		require.Equal(t, expected, actual)
		require.NoError(t, err)
	})

	t.Run("fail - not found", func(t *testing.T) {
		repo := new(RepositoryMock)
		expected := map[int]internal.Vehicle{}
		repo.On("FindByBrandAndYearRange", "BLue", 2000, 2025).Return(expected, nil)
		s := service.NewServiceVehicleDefault(repo)

		actual, err := s.FindByBrandAndYearRange("BLue", 2000, 2025)

		require.Equal(t, expected, actual)
		require.NoError(t, err)
	})
}