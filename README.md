# jsonconfig

##Usage
```go
  configuration := helper.ReadConfigurationFile()

	dbConn := database.StartDatabaseClient(
		configuration.DatabaseHost,
		configuration.DatabasePort,
		configuration.DatabaseName,
		configuration.DatabaseUsername,
		configuration.DatabasePassword,
	)

	router.InitializeRouter(app, dbConn.Connection())
```
