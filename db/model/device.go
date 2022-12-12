package model

type deviceUtil struct{}

var DeviceUtil = deviceUtil{}

func (deviceUtil) CreateOrUpdate(userID string, deviceID string, meta string, channel string, token string) error {
	return nil
}
