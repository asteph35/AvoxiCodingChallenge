package service

import (
	"AvoxiCodingChallenge/models"
	"AvoxiCodingChallenge/repository"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_CheckIP_Handler(t *testing.T) {
	var (
		ip = "132.170.0.1"
		us = "US"
		jp = "JP"
		ua = "UA"
	)
	s := New(repository.New())

	body := models.NewIPCheckerReq(ip, us, jp, ua)
	byt, err := json.Marshal(body)
	require.NoError(t, err)

	req, err := http.NewRequest("GET", "/ip-check", bytes.NewReader(byt))
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.CheckIP)
	handler.ServeHTTP(rr, req)

	var res models.IPCheckerRes
	err = json.NewDecoder(rr.Body).Decode(&res)
	require.NoError(t, err)
	assert.Len(t, res.Countries, 3)
	assert.Equal(t, ip, res.IP)
	assert.Equal(t, us, res.Countries[0].Country)
	assert.True(t, res.Countries[0].Status)
	assert.Equal(t, jp, res.Countries[1].Country)
	assert.False(t, res.Countries[1].Status)
	assert.Equal(t, ua, res.Countries[2].Country)
	assert.False(t, res.Countries[2].Status)
}
