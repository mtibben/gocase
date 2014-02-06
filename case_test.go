package gocase

import "testing"

type caseTest struct {
	testFunc func(string) string
	expected string
}

func TestCase(t *testing.T) {

	tests := map[string][]caseTest{
		"May the force be with you": {
			caseTest{ToSnake, "may_the_force_be_with_you"},
			caseTest{ToUpperSnake, "MAY_THE_FORCE_BE_WITH_YOU"},
			caseTest{ToSpinal, "may-the-force-be-with-you"},
			caseTest{ToTrain, "May-The-Force-Be-With-You"},
			caseTest{ToCamel, "MayTheForceBeWithYou"},
			caseTest{ToLowerCamel, "mayTheForceBeWithYou"},
		},
		"Toto, I've got a feeling we're not in Kansas anymore.": {
			caseTest{ToSnake, "toto_ive_got_a_feeling_were_not_in_kansas_anymore"},
			caseTest{ToUpperSnake, "TOTO_IVE_GOT_A_FEELING_WERE_NOT_IN_KANSAS_ANYMORE"},
			caseTest{ToSpinal, "toto-ive-got-a-feeling-were-not-in-kansas-anymore"},
			caseTest{ToTrain, "Toto-Ive-Got-A-Feeling-Were-Not-In-Kansas-Anymore"},
			caseTest{ToCamel, "TotoIveGotAFeelingWereNotInKansasAnymore"},
			caseTest{ToLowerCamel, "totoIveGotAFeelingWereNotInKansasAnymore"},
		},
		"E.T. phone home.": {
			caseTest{ToSnake, "et_phone_home"},
			caseTest{ToUpperSnake, "ET_PHONE_HOME"},
			caseTest{ToSpinal, "et-phone-home"},
			caseTest{ToTrain, "ET-Phone-Home"},
			caseTest{ToCamel, "ETPhoneHome"},
			caseTest{ToLowerCamel, "eTPhoneHome"},
		},
		"\twhitespace  lots \n of   whitespace  ": {
			caseTest{ToSnake, "whitespace_lots_of_whitespace"},
			caseTest{ToUpperSnake, "WHITESPACE_LOTS_OF_WHITESPACE"},
			caseTest{ToSpinal, "whitespace-lots-of-whitespace"},
			caseTest{ToTrain, "Whitespace-Lots-Of-Whitespace"},
			caseTest{ToCamel, "WhitespaceLotsOfWhitespace"},
			caseTest{ToLowerCamel, "whitespaceLotsOfWhitespace"},
		},
		"": {
			caseTest{ToSnake, ""},
			caseTest{ToUpperSnake, ""},
			caseTest{ToSpinal, ""},
			caseTest{ToTrain, ""},
			caseTest{ToCamel, ""},
			caseTest{ToLowerCamel, ""},
		},
		"----{}{{}[[]]][./';';.. <><><><>.. --- --": {
			caseTest{ToSnake, ""},
			caseTest{ToUpperSnake, ""},
			caseTest{ToSpinal, ""},
			caseTest{ToTrain, ""},
			caseTest{ToCamel, ""},
			caseTest{ToLowerCamel, ""},
		},
		"君子坦蕩蕩，小人長戚戚": {
			caseTest{ToSnake, "君子坦蕩蕩小人長戚戚"},
			caseTest{ToUpperSnake, "君子坦蕩蕩小人長戚戚"},
			caseTest{ToSpinal, "君子坦蕩蕩小人長戚戚"},
			caseTest{ToTrain, "君子坦蕩蕩小人長戚戚"},
			caseTest{ToCamel, "君子坦蕩蕩小人長戚戚"},
			caseTest{ToLowerCamel, "君子坦蕩蕩小人長戚戚"},
		},
	}

	for testString, testlist := range tests {
		for _, test := range testlist {
			result := test.testFunc(testString)
			if result != test.expected {
				t.Errorf(`Error in "%s"`+"\n"+`  Expected: "%s"`+"\n"+`  Got:      "%s"`, testString, test.expected, result)
			}
		}
	}
}
