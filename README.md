# binding [![wercker status](https://app.wercker.com/status/f00480949f8a4e4130557f802c5b5b6b "wercker status")](https://app.wercker.com/project/bykey/f00480949f8a4e4130557f802c5b5b6b)

Request data binding and validation for Martini.

[![GoDoc](https://godoc.org/github.com/martini-contrib/binding?status.svg)](https://godoc.org/github.com/martini-contrib/binding)


## Usage

现在已改成使用[validator](https://github.com/kexin-corp/validator)进行数据校验

#### Getting form data from a request

```go
type User struct {
	Name    string `json:"name" validate:"nonone"`
	Email   string `json:"email" validate:"email"`
	Message string `json:"message"`
}
```


