package isbn

import "testing"

type TestScenario struct {
	name     string
	args     interface{}
	expected interface{}
}

func TestValidate(t *testing.T) {
	validator := Validator{}

	testScenarios := []TestScenario{
		{
			name:     "Valid ISBN with 10 digits should return true",
			args:     "0716703440",
			expected: true,
		},
		{
			name:     "ISBN with less than 10 digits should return false",
			args:     "889213",
			expected: false,
		},
		{
			name:     "Valid ISBN with 13 digits should return true",
			args:     "9780716703440",
			expected: true,
		},
		{
			name:     "ISBN with 13 digits not starting with 978 should return false",
			args:     "9780716703440",
			expected: false,
		},
	}

	for _, test := range testScenarios {
		t.Run(test.name, func(t *testing.T) {
			isbnNumber := test.args.(string)

			if result := validator.Validate(isbnNumber); result != test.expected {
				t.Errorf(" Validate() returns %#v but expected %#v", result, test.expected)
			}
		})
	}
}
func TestGet10Digits(t *testing.T){
	testScenarios := []TestScenario{
		{
			name: "Should return 187",
			args: []string{
				"0","7","1","6","7","0","3","4","4","0",
			},
			expected: 187,
		},

	for _, test := range testScenarios{
		t.Run(test.name, func(t *testing.T){
			isBnNumber := test.args.([]string)

			result, err := get10Digits(isbnNumber)
			if err != nil {
				t.Error(err)
			}

		if result != test.expected {
			t.Errorf(" Validate() returns %#v but expected %#v", result, test.expected)
		},
	}
}
