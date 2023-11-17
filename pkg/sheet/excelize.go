package sheet

import (
	"github.com/xuri/excelize/v2"
)

func Excelize(opts ...excelize.Options) *excelize.File {
	return excelize.NewFile(opts...)
}
