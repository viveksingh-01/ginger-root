package restaurant

type Filter struct {
	Veg       *bool
	MinRating *float64
}

// NOTE:
// We use pointers because we need to distinguish three states, not two.
// For veg, those states are:
// 1. Not provided (no filter)
// 2. Provided = true
// 3. Provided = false
// A plain bool can only represent two states - true or false, but not nil

// Pointer = “optional value” in Go
// In Go, pointers are commonly used to mean: “This value may or may not be present”
