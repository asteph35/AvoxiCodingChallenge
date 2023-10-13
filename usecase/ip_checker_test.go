package usecase_test

import (
	"AvoxiCodingChallenge/repository"
	"AvoxiCodingChallenge/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIP(t *testing.T) {
	var (
		us = "US"
		jp = "JP"
		ua = "UA"
	)
	ipm := usecase.NewIPManager(ipRepo())

	t.Run("us ipv4 address with three allowable countries", func(t *testing.T) {
		var ip = "132.170.0.1"

		countryMap, err := ipm.IPChecker(ip, us, jp, ua)
		require.NoError(t, err)
		assert.Len(t, countryMap, 3)
		assert.Equal(t, true, countryMap[us])
		assert.Equal(t, false, countryMap[jp])
		assert.Equal(t, false, countryMap[ua])
	})
	t.Run("jp ipv4 address with three allowable countries", func(t *testing.T) {
		var ip = "101.110.128.122"

		countryMap, err := ipm.IPChecker(ip, us, jp, ua)
		require.NoError(t, err)
		assert.Len(t, countryMap, 3)
		assert.Equal(t, false, countryMap[us])
		assert.Equal(t, true, countryMap[jp])
		assert.Equal(t, false, countryMap[ua])
	})
	t.Run("invalid ip", func(t *testing.T) {
		var ip = "not an ip"
		countryMap, err := ipm.IPChecker(ip, us, jp, ua)
		require.Equal(t, err, usecase.InvalidIP)
		require.Nil(t, countryMap)
	})
}

func ipRepo() usecase.IPRepo {
	ipRepo := repository.New()
	return ipRepo
}
