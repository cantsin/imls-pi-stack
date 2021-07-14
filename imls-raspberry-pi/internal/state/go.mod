module gsa.gov/18f/state

go 1.16

replace gsa.gov/18f/wifi-hardware-search v0.0.0 => ../../internal/wifi-hardware-search

replace gsa.gov/18f/version v0.0.0 => ../../internal/version

replace gsa.gov/18f/config v0.0.0 => ../../internal/config

replace gsa.gov/18f/http v0.0.0 => ../../internal/http

replace gsa.gov/18f/cryptopasta v0.0.0 => ../../internal/cryptopasta

replace gsa.gov/18f/analysis v0.0.0 => ../../internal/analysis

replace gsa.gov/18f/logwrapper v0.0.0 => ../../internal/logwrapper

replace gsa.gov/18f/structs v0.0.0 => ../../internal/structs

require (
	github.com/jmoiron/sqlx v1.3.4
	github.com/mattn/go-sqlite3 v1.14.7
	github.com/newrelic/go-agent/v3 v3.14.0 // indirect
	github.com/newrelic/go-agent/v3/integrations/nrlogrus v1.0.1 // indirect
	github.com/stretchr/testify v1.2.2
	golang.org/x/text v0.3.6 // indirect
	gsa.gov/18f/config v0.0.0
	gsa.gov/18f/logwrapper v0.0.0
	gsa.gov/18f/structs v0.0.0
)