package demo

import "fmt"

type BizEror struct {
	errorCode int
	errorMsg  string
}

func (de *BizEror) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    code: %d
    msg: %s
`
	return fmt.Sprintf(strFormat, de.errorCode, de.errorMsg)
}

func TestError() (code int, error string) {
	odata := BizEror{
		errorCode: 311,
		errorMsg:  "love",
	}
	odata.Error()
	println(odata.errorMsg)
	return 333, odata.errorMsg
}
