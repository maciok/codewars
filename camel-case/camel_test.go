package sum_of_digits

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, "", ToCamelCase(""))
	assert.Equal(t, "TheStealthWarrior", ToCamelCase("The_Stealth_Warrior"))
	assert.Equal(t, "theStealthWarrior", ToCamelCase("the-Stealth-Warrior"))
	assert.Equal(t, "YouHaveChosenToTranslateThisKataForYourConvenienceWeHaveProvidedTheExistingTestCasesUsedForTheLanguageThatYouHaveAlreadyCompletedAsWellAsAllOfTheOtherRelatedFields", ToCamelCase("You_have_chosen_to_translate_this_kata_For_your_convenience_we_have_provided_the_existing_test_cases_used_for_the_language_that_you_have_already_completed_as_well_as_all_of_the_other_related_fields"))
}
