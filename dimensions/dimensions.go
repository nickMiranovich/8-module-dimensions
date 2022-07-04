package dimensions

type Phone interface {
	Brand() string
	Model() string
	Type() string
}

type StationPhone interface {
	ButtonsCount() int
}

type Smartphone interface {
	OS() string
}

type applePhone struct {
	brand     string
	model     string
	typePhone string
	os        string
}

type androidPhone struct {
	brand     string
	model     string
	typePhone string
	os        string
}
type stantionPhone struct {
	brand       string
	model       string
	typePhone   string
	buttonCount int
}

// конструкторы
func NewApplePhone(model string) applePhone {
	return applePhone{
		brand:     "Apple",
		model:     model,
		typePhone: "smartphone",
		os:        "MacOS",
	}
}

func NewAndroidPhone(brand, model string) androidPhone {
	return androidPhone{
		brand:     brand,
		model:     model,
		typePhone: "smartphone",
		os:        "Android",
	}

}

func NewStantionPhone(brand, model string, buttonCount int) stantionPhone {
	return stantionPhone{
		brand:       brand,
		model:       model,
		typePhone:   "stantion",
		buttonCount: buttonCount,
	}
}

//методы для ApplePhone
func (a applePhone) Brand() string {
	return a.brand
}

func (a applePhone) Model() string {
	return a.model
}

func (a applePhone) Type() string {
	return a.typePhone
}

func (a applePhone) OS() string {
	return a.os
}

//методы для AndroidPhone
func (a androidPhone) Brand() string {
	return a.brand
}

func (a androidPhone) Model() string {
	return a.model
}

func (a androidPhone) Type() string {
	return a.typePhone
}

func (a androidPhone) OS() string {
	return a.os
}

//методы для StantionPhone
func (s stantionPhone) Brand() string {
	return s.brand
}

func (s stantionPhone) Model() string {
	return s.model
}

func (s stantionPhone) Type() string {
	return s.typePhone
}

func (s stantionPhone) ButtonsCount() int {
	return s.buttonCount
}

