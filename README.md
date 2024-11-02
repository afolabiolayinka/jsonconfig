# jsonconfig

For the love of curly braces syntax, I present jsonconfig. This is an alternative to indexed, tabbed or spaced based configuration files in your go projects.
It is simple enough to use, just remember to add it to your gitignore file to avoid exposing your secret credentials.

## Usage
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
