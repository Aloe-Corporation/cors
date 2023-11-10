# cors

This project is a middlewares module to manage Cross-origin resource sharing for Gin

## Example to use

```go
func Init() {
    Router := gin.New()
    Router.Use(cors.Middleware(cors.DefaultConf()))
}
```
Or 
```go
func Init() {
    Router := gin.New()
    Router.Use(cors.Middleware(cors.Conf{
        ...
    }))
}
```