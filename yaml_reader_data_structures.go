package main

type TestDescriber struct {
	Test_name string
	Comment string
	Command_sequence []Command
}

type Command struct {
	Order_id int
	Comment string
	Type string
	Method string
	Url string
	Data string
	Expect Expect
	Headers map[string]string
	Repeat_times int
	Waiting_time int
}

type Expect struct {
	Value string
	Respones_code int
	Type string
}