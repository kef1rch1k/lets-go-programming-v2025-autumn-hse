package wifi_test

import (
	"errors"
	"net"
	"testing"

	"github.com/stretchr/testify/require"

	mwifi "github.com/mdlayher/wifi"

	"github.com/kef1rch1k/task-6/internal/wifi"
)

var errIntf = errors.New("interfaces error")

func TestGetAddresses(t *testing.T) {
	t.Parallel()

	mockWiFi := NewWiFiHandle(t)

	mockWiFi.On("Interfaces").Return([]*mwifi.Interface{
		{
			Name:         "wlan0",
			HardwareAddr: net.HardwareAddr{0x00, 0x11},
		},
	}, nil)

	service := wifi.New(mockWiFi)

	addrs, err := service.GetAddresses()
	require.NoError(t, err)
	require.Len(t, addrs, 1)
}

func TestGetAddresses_Error(t *testing.T) {
	t.Parallel()

	mockWiFi := NewWiFiHandle(t)

	mockWiFi.On("Interfaces").
		Return(nil, errIntf)

	service := wifi.New(mockWiFi)

	_, err := service.GetAddresses()
	require.ErrorContains(t, err, "interfaces error")
}

func TestGetNames(t *testing.T) {
	t.Parallel()

	mockWiFi := NewWiFiHandle(t)

	mockWiFi.On("Interfaces").Return([]*mwifi.Interface{
		{Name: "wlan0"},
		{Name: "eth0"},
	}, nil)

	service := wifi.New(mockWiFi)

	names, err := service.GetNames()
	require.NoError(t, err)
	require.Len(t, names, 2)
}

func TestGetNames_Error(t *testing.T) {
	t.Parallel()

	mockWiFi := NewWiFiHandle(t)

	mockWiFi.On("Interfaces").
		Return(nil, errIntf)

	service := wifi.New(mockWiFi)

	_, err := service.GetNames()
	require.ErrorContains(t, err, "interfaces error")
}
