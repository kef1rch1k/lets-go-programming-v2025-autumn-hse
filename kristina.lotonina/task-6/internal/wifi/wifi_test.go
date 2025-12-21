package wifi_test

import (
	"errors"
	"net"
	"testing"

	"github.com/kef1rch1k/task-6/internal/wifi"
	mwifi "github.com/mdlayher/wifi"
)

var (
	errIntf = errors.New("interfaces error")
)

func TestGetAddresses(t *testing.T) {
	mockWiFi := NewWiFiHandle(t)

	mockWiFi.On("Interfaces").Return([]*mwifi.Interface{
		{
			Name:         "wlan0",
			HardwareAddr: net.HardwareAddr{0x00, 0x11},
		},
	}, nil)

	service := wifi.New(mockWiFi)

	addrs, err := service.GetAddresses()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(addrs) != 1 {
		t.Fatalf("expected 1 address, got %d", len(addrs))
	}
}

func TestGetAddresses_Error(t *testing.T) {
	mockWiFi := NewWiFiHandle(t)

	mockWiFi.On("Interfaces").
		Return(nil, errIntf)

	service := wifi.New(mockWiFi)

	_, err := service.GetAddresses()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestGetNames(t *testing.T) {
	mockWiFi := NewWiFiHandle(t)

	mockWiFi.On("Interfaces").Return([]*mwifi.Interface{
		{Name: "wlan0"},
		{Name: "eth0"},
	}, nil)

	service := wifi.New(mockWiFi)

	names, err := service.GetNames()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(names) != 2 {
		t.Fatalf("expected 2 names, got %d", len(names))
	}
}

func TestGetNames_Error(t *testing.T) {
	mockWiFi := NewWiFiHandle(t)

	mockWiFi.On("Interfaces").
		Return(nil, errIntf)

	service := wifi.New(mockWiFi)

	_, err := service.GetNames()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
