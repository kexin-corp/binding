# binding [![wercker status](https://app.wercker.com/status/f00480949f8a4e4130557f802c5b5b6b "wercker status")](https://app.wercker.com/project/bykey/f00480949f8a4e4130557f802c5b5b6b)

Request data binding and validation for Martini.

[![GoDoc](https://godoc.org/github.com/martini-contrib/binding?status.svg)](https://godoc.org/github.com/martini-contrib/binding)


## Usage

为了适应具有中国特色的参数格式，已改成使用[validator](https://github.com/kexin-corp/validator)进行数据校验

#### 你的request可以这样写

```go
type User struct {
	Name    string `json:"name" validate:"nonone"`
	Email   string `json:"email" validate:"email"`
	Message string `json:"message"`
}
```


