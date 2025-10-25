module github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus

go 1.23.0

retract v1.1.2 // Breaks customers in situations where close is slow/infinite.

require (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.19.1
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.13.0
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.11.2
	github.com/Azure/go-amqp v1.5.0
)

require (
	// used in tests only
	github.com/joho/godotenv v1.5.1

	// used in stress tests
	github.com/microsoft/ApplicationInsights-Go v0.4.4
	github.com/stretchr/testify v1.11.1

	// used in examples only
	nhooyr.io/websocket v1.8.17
)

require github.com/golang/mock v1.6.0

require (
	code.cloudfoundry.org/clock v1.1.0 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v1.5.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/golang-jwt/jwt/v5 v5.3.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/crypto v0.41.0 // indirect
	golang.org/x/net v0.43.0 // indirect
	golang.org/x/sys v0.35.0 // indirect
	golang.org/x/text v0.28.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
