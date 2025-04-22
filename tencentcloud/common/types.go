package common

// IntPtr returns a pointer to the int value passed.
func IntPtr(v int) *int {
	return &v
}

// Int64Ptr returns a pointer to the int64 value passed.
func Int64Ptr(v int64) *int64 {
	return &v
}

// UintPtr returns a pointer to the uint value passed.
func UintPtr(v uint) *uint {
	return &v
}

// Uint64Ptr returns a pointer to the uint64 value passed.
func Uint64Ptr(v uint64) *uint64 {
	return &v
}

// Float64Ptr returns a pointer to the float64 value passed.
func Float64Ptr(v float64) *float64 {
	return &v
}

// BoolPtr returns a pointer to the bool value passed.
func BoolPtr(v bool) *bool {
	return &v
}

// StringPtr returns a pointer to the string value passed.
func StringPtr(v string) *string {
	return &v
}

// StringValues converts a slice of string pointers to a slice of strings.
// Nil pointers are converted to empty strings.
func StringValues(ptrs []*string) []string {
	values := make([]string, len(ptrs))
	for i := 0; i < len(ptrs); i++ {
		if ptrs[i] != nil {
			values[i] = *ptrs[i]
		}
	}
	return values
}

// IntPtrs converts a slice of ints to a slice of int pointers.
func IntPtrs(vals []int) []*int {
	ptrs := make([]*int, len(vals))
	for i := 0; i < len(vals); i++ {
		ptrs[i] = &vals[i]
	}
	return ptrs
}

// Int64Ptrs converts a slice of int64s to a slice of int64 pointers.
func Int64Ptrs(vals []int64) []*int64 {
	ptrs := make([]*int64, len(vals))
	for i := 0; i < len(vals); i++ {
		ptrs[i] = &vals[i]
	}
	return ptrs
}

// UintPtrs converts a slice of uints to a slice of uint pointers.
func UintPtrs(vals []uint) []*uint {
	ptrs := make([]*uint, len(vals))
	for i := 0; i < len(vals); i++ {
		ptrs[i] = &vals[i]
	}
	return ptrs
}

// Uint64Ptrs converts a slice of uint64s to a slice of uint64 pointers.
func Uint64Ptrs(vals []uint64) []*uint64 {
	ptrs := make([]*uint64, len(vals))
	for i := 0; i < len(vals); i++ {
		ptrs[i] = &vals[i]
	}
	return ptrs
}

// Float64Ptrs converts a slice of float64s to a slice of float64 pointers.
func Float64Ptrs(vals []float64) []*float64 {
	ptrs := make([]*float64, len(vals))
	for i := 0; i < len(vals); i++ {
		ptrs[i] = &vals[i]
	}
	return ptrs
}

// BoolPtrs converts a slice of bools to a slice of bool pointers.
func BoolPtrs(vals []bool) []*bool {
	ptrs := make([]*bool, len(vals))
	for i := 0; i < len(vals); i++ {
		ptrs[i] = &vals[i]
	}
	return ptrs
}

// StringPtrs converts a slice of strings to a slice of string pointers.
func StringPtrs(vals []string) []*string {
	ptrs := make([]*string, len(vals))
	for i := 0; i < len(vals); i++ {
		ptrs[i] = &vals[i]
	}
	return ptrs
}
