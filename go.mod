module logging-service

go 1.15

require (
	cloud.google.com/go/pubsub v1.12.0
	github.com/RackSec/srslog v0.0.0-20180709174129-a4725f04ec91
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20210208195552-ff826a37aa15 // indirect
	github.com/cloudfoundry-incubator/candiedyaml v0.0.0-20170901234223-a41693b7b7af
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.6.1
	google.golang.org/api v0.50.0
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
)
