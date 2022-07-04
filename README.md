# Go-Social_Network

Applying Clean Architecture in Go

cmd: Contains the main file for the delivery mechansim

deployment: Deployment scripts

internal: The most important folder, contains the implementation of all the alyers and components

pkg: Project indepoendent packages with the intention of extending the language. Each one of these packages could be released as a standalone library

tests: Contains integration tests and other test assets

web: Web related stuff like templates and static files.
 
# The go module commands

go mod init - initializes new moduole in current directory
go mod tidy - adds missing and removes unused modules
go mod download - downloads modules to local cache
go mod vendor - makes vendored copy of dependecies
go mod graph - prints module requirement graph
go mod verify - verifies dependencies have expected content
go mod why - explains why packages or modules are needed


https://dev.to/techschoolguru/generate-crud-golang-code-from-sql-and-compare-db-sql-gorm-sqlx-sqlc-560j
https://docs.sqlc.dev/en/latest/overview/install.html