package chubby

import (
	"testing"
)

func TestNew(t *testing.T) {
	rolls := New()
	if len(rolls) != 0 {
		t.Errorf("New Rolls object has non-zero length")
	}
}

func TestAdd(t *testing.T) {
	rolls := New()
	roll := rolls.NewRoll("Admin")
	if len(rolls) != 1 {
		t.Errorf("Unable to insert new roll")
	}
	if roll.Value != 1 {
		t.Errorf("New roll value isn't 1")
	}
}

func TestGet(t *testing.T) {
	rolls := New()
	_ = rolls.NewRoll("Admin")
	roll, err := rolls.Get("Admin")
	if err != nil {
		t.Errorf("Unable to get inserted roll: %s", err)
	}
	if roll.Value != 1 {
		t.Errorf("Inserted roll value is not 1")
	}
}

func TestHasRoll(t *testing.T) {
	rolls := New()
	adminRoll := rolls.NewRoll("Admin")
	roll, err := rolls.Get("Admin")
	if err != nil {
		t.Errorf("Unable to get inserted roll: %s", err)
	}
	has := HasRoll(roll.Value, adminRoll)
	if !has {
		t.Errorf("Failed to validate roll - should match")
	}
	userRoll := rolls.NewRoll("User")
	has = HasRoll(userRoll.Value, adminRoll)
	if has {
		t.Errorf("Failed to validate roll - shouldn't match")
	}
}

func TestCombineRolls(t *testing.T) {
	rolls := New()
	_ = rolls.NewRoll("Admin")
	_ = rolls.NewRoll("Manager")
	_ = rolls.NewRoll("Tester")
	_ = rolls.NewRoll("Anonymous")
	combinedRolls := rolls.Combine("Admin", "Tester")
	if combinedRolls != 5 {
		t.Errorf("Failed to combine rolls - got %d", combinedRolls)
	}
	combinedRolls = rolls.Combine("Admin", "Tester", "Anonymous")
	if combinedRolls != 13 {
		t.Errorf("Failed to combine rolls - got %d", combinedRolls)
	}
}

func TestHasCombinedRolls(t *testing.T) {
	rolls := New()
	adminRoll := rolls.NewRoll("Admin")
	managerRoll := rolls.NewRoll("Manager")
	_ = rolls.NewRoll("Tester")
	_ = rolls.NewRoll("Anonymous")

	combinedRolls := rolls.Combine("Admin", "Tester", "Anonymous")
	if combinedRolls != 13 {
		t.Errorf("Failed to combine rolls - got %d", combinedRolls)
	}

	has := HasRoll(combinedRolls, managerRoll)
	if has {
		t.Errorf("Failed to validate roll - shouldn't match")
	}

	has = HasRoll(combinedRolls, adminRoll)
	if !has {
		t.Errorf("Failed to validate roll - should match")
	}
}

func BenchmarkHasCombinedRolls(b *testing.B) {
	rolls := New()
	_ = rolls.NewRoll("Admin")
	managerRoll := rolls.NewRoll("Manager")
	_ = rolls.NewRoll("Tester")
	_ = rolls.NewRoll("Anonymous")

	combinedRolls := rolls.Combine("Admin", "Tester")

	for n := 0; n < b.N; n++ {
		HasRoll(combinedRolls, managerRoll)
	}
}
