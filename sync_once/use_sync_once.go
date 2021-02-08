package sync_once

func UseLog() error {
	if err := InitInstance(); err != nil {
		return err
	}
	Info("logrus is been initialised")
	return nil
}
