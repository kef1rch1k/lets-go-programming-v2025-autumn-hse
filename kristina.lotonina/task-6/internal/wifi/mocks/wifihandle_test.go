package mocks_test

import (
	"errors"
	"testing"

	mwifi "github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"

	"github.com/kef1rch1k/task-6/internal/wifi/mocks"
)

var errBoom = errors.New("boom")

func TestWiFiHandle_Interfaces_PanicWhenNoReturn(t *testing.T) {
	t.Parallel()

	h := mocks.NewWiFiHandle(t)

	h.Mock.On("Interfaces").Return()

	require.PanicsWithValue(t, "no return value specified for Interfaces", func() {
		_, _ = h.Interfaces()
	})
}

func TestWiFiHandle_Interfaces_FuncBoth(t *testing.T) {
	t.Parallel()

	h := mocks.NewWiFiHandle(t)

	h.Mock.On("Interfaces").Return(func() ([]*mwifi.Interface, error) {
		return []*mwifi.Interface{
			{Name: "wlan0"},
		}, nil
	}, nil)

	ifaces, err := h.Interfaces()
	require.NoError(t, err)
	require.Len(t, ifaces, 1)
	require.Equal(t, "wlan0", ifaces[0].Name)
}

func TestWiFiHandle_Interfaces_FuncSlice(t *testing.T) {
	t.Parallel()

	h := mocks.NewWiFiHandle(t)

	h.Mock.On("Interfaces").Return(func() []*mwifi.Interface {
		return []*mwifi.Interface{
			{Name: "wlan0"},
		}
	}, nil)

	ifaces, err := h.Interfaces()
	require.NoError(t, err)
	require.Len(t, ifaces, 1)
}

func TestWiFiHandle_Interfaces_FuncError(t *testing.T) {
	t.Parallel()

	h := mocks.NewWiFiHandle(t)

	h.Mock.On("Interfaces").Return([]*mwifi.Interface{}, func() error {
		return errBoom
	})

	_, err := h.Interfaces()
	require.EqualError(t, err, "boom")
}
