module github.com/sworisbreathing/go-ibbq/v2/examples/datalogger

go 1.12

require (
	github.com/go-ble/ble v0.0.0-20230130210458-dd4b07d15402
	github.com/mgutz/logxi v0.0.0-20161027140823-aebf8a7d67ab
	github.com/sworisbreathing/go-ibbq/v2 v2.0.0
)

replace github.com/sworisbreathing/go-ibbq/v2 v2.0.0 => ../../
