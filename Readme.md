### How to use
```
config := braketini.Config{
  ProjectId:   1234,
  Key:         "blahblah",
  Environment: "development",
}
handler := martini.Classic()
handler.Use(braketini.Middleware(config))
```