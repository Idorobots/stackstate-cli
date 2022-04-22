package printer

import (
	"net/http"
)

type MockPrinter struct {
	PrintStructCalls  *[]interface{}
	PrintErrCalls     *[]error
	StartSpinnerCalls *[]LoadingMsg
	StopSpinnerCalls  *int
	UseColor          bool
	outputType        OutputType
	SuccessCalls      *[]string
	PrintWarnCalls    *[]string
	TableCalls        *[]TableData
	PrintLnCalls      *[]string
}

type PrintErrResponseCall struct {
	Err  error
	Resp *http.Response
}

func NewMockPrinter() MockPrinter {
	printStructCalls := make([]interface{}, 0)
	printErrCalls := make([]error, 0)
	startSpinnerCalls := make([]LoadingMsg, 0)
	successCalls := make([]string, 0)
	printWarnCalls := make([]string, 0)
	printTableCalls := make([]TableData, 0)
	printLnCalls := make([]string, 0)
	return MockPrinter{
		PrintStructCalls:  &printStructCalls,
		PrintErrCalls:     &printErrCalls,
		StartSpinnerCalls: &startSpinnerCalls,
		SuccessCalls:      &successCalls,
		PrintWarnCalls:    &printWarnCalls,
		TableCalls:        &printTableCalls,
		PrintLnCalls:      &printLnCalls,
		outputType:        Auto,
	}
}

func (p *MockPrinter) PrintStruct(s interface{}) {
	*p.PrintStructCalls = append(*p.PrintStructCalls, s)
}

func (p *MockPrinter) PrintErr(err error) {
	*p.PrintErrCalls = append(*p.PrintErrCalls, err)
}

func (p *MockPrinter) StartSpinner(loadingMsg LoadingMsg) StopPrinterFn {
	*p.StartSpinnerCalls = append(*p.StartSpinnerCalls, loadingMsg)
	return func() {}
}

func (p *MockPrinter) SetUseColor(useColor bool) {
	p.UseColor = useColor
}

func (p *MockPrinter) GetUseColor() bool {
	return p.UseColor
}

func (p *MockPrinter) SetOutputType(outputType OutputType) {
	p.outputType = outputType
}

func (p *MockPrinter) GetOutputType() OutputType {
	return p.outputType
}

func (p *MockPrinter) Success(msg string) {
	*p.SuccessCalls = append(*p.SuccessCalls, msg)
}

func (p *MockPrinter) PrintWarn(msg string) {
	*p.PrintWarnCalls = append(*p.PrintWarnCalls, msg)
}

func (p *MockPrinter) Table(t TableData) {
	*p.TableCalls = append(*p.TableCalls, t)
}

func (p *MockPrinter) PrintLn(text string) {
	*p.PrintLnCalls = append(*p.PrintLnCalls, text)
}
