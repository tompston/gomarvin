# Gomarvin changelog

<!--

- controllers.gen.go added to gitignore
- controllers don't call the  `database.GetDbConn()` function anymore. Instead, use the settings.DB
- renamed .env db variables
- moved typescript fetch functions to a public dir, so that i  could be hosted and shared by frontend easier.
- 

 -->

### v0.2.0

- Support for echo framework
- refactored typescript functions to fetch endpoints
- automatic formatting after codegen using gofmt

### v0.1.0

- init release
