package repository_test

import (
	"AvoxiCodingChallenge/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var repo = repository.New()

func Test_GetNetworks(t *testing.T) {
	t.Run("GetNetworks by country name", func(t *testing.T) {
		countryName := "United States"
		networks, err := repo.GetNetworks(countryName)
		require.NoError(t, err)
		assert.NotEmpty(t, networks)
	})
	t.Run("GetNetworks by country name lowercase", func(t *testing.T) {
		countryName := "united states"
		networks, err := repo.GetNetworks(countryName)
		require.NoError(t, err)
		assert.NotEmpty(t, networks)
	})
	t.Run("GetNetworks by country iso code", func(t *testing.T) {
		countryName := "US"
		networks, err := repo.GetNetworks(countryName)
		require.NoError(t, err)
		assert.NotEmpty(t, networks)
	})
	t.Run("GetNetworks by country iso code lowercase", func(t *testing.T) {
		countryName := "us"
		networks, err := repo.GetNetworks(countryName)
		require.NoError(t, err)
		assert.NotEmpty(t, networks)
	})
	t.Run("GetNetworks country not found", func(t *testing.T) {
		countryName := "fake country"
		networks, err := repo.GetNetworks(countryName)
		require.NoError(t, err)
		assert.Empty(t, networks)
	})
}
