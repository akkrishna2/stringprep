package stringprep_test

import (
	"fmt"

	"github.com/xdg/stringprep"
)

func Example_customProfile() {

	myProfile := stringprep.Profile{
		Mappings: []stringprep.Mapping{
			stringprep.TableB1,
			stringprep.TableB2,
		},
		Normalize: true,
		Prohibits: []stringprep.Set{
			stringprep.TableC1_1,
			stringprep.TableC1_2,
		},
		CheckBiDi: true,
	}

	prepped, err := myProfile.Prepare("TrustNô1")
	if err != nil {
		panic("stringprep failed")
	}

	fmt.Print(prepped)
	// Output: trustnô1
}

func Example_saslprep() {

	prepped, err := stringprep.SASLprep.Prepare("TrustNô1")
	if err != nil {
		panic("SASLprep failed")
	}
	fmt.Print(prepped)
	// Output: TrustNô1
}
