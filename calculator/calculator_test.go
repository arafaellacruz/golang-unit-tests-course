package calculator

import "testing"

type DiscountRepositoryMock struct{}

func (drm *DiscountRepositoryMock) FindCurrentDiscount() int {
	return 20
}

func TestDiscountAppliedCalculator(t *testing.T) {
	type testCase struct {
		name            string
		minimumPurchase int
		discount        int
		purchaseAmount  int
		expectedAmount  int
	}

	testCases := []testCase{
		{name: "should apply 20", minimumPurchase: 100, discount: 20, purchaseAmount: 150, expectedAmount: 130},
		{name: "should apply 40", minimumPurchase: 100, discount: 20, purchaseAmount: 200, expectedAmount: 160},
		{name: "should apply 60", minimumPurchase: 100, discount: 20, purchaseAmount: 350, expectedAmount: 290},
		{name: "should not apply", minimumPurchase: 100, discount: 20, purchaseAmount: 50, expectedAmount: 50},

	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			DiscountRepositoryMock := DiscountRepositoryMock{}

			calculator, err := NewDiscountCalculator(tc.minimumPurchase, tc.discount)
			if err != nil {
				// Fail + Log = Errof
				// t.Errorf("could not instantiate the calculator %s", err.Error())
				// FailNow will stop the test execution
				// t.FailNow()
				// Fatalf = FailNow + Log
				t.Fatalf("could not instantiate the calculator %s", err.Error())
			}
			amount := calculator.Calculate(tc.purchaseAmount)

			if amount != tc.expectedAmount {
				t.Errorf("expected %v, got %v, failed because the discount was not expected to be applied", tc.expectedAmount, amount) // Error = Log + Fail
			}
		})

	}
}

func TestDiscountCalculatorShouldFailWithZeroMinimumAmount(t *testing.T) {
			_, err := NewDiscountCalculator(0, 20)
			if err == nil {
				t.Fatalf("should not create the calculator with zero purchase amount")
			}
	
		}


