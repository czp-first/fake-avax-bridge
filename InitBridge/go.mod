module github.com/czp-first/fake-avax-bridge/InitBridge

go 1.18

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/mdlayher/socket v0.2.0 // indirect
	golang.org/x/net v0.0.0-20220607020251-c690dde0001d // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c // indirect
)

require (
	github.com/google/uuid v1.3.0
	github.com/joho/godotenv v1.4.0
	github.com/mdlayher/vsock v1.1.1
	github.com/sirupsen/logrus v1.8.1
	google.golang.org/grpc v1.48.0
	google.golang.org/protobuf v1.27.1 // indirect
)

require github.com/czp-first/fake-avax-bridge/WardenServer v0.0.0

replace github.com/czp-first/fake-avax-bridge/WardenServer => ../WardenServer
