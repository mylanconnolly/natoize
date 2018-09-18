package main

import "testing"

func BenchmarkNatoize(b *testing.B) {
	tests := []struct {
		Name string
		Func func(string) string
	}{
		{"Phonetic", Phonetic},
		{"PhoneticFast", PhoneticFast},
	}
	testString := "this is a test"
	for _, tt := range tests {
		b.Run(tt.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tt.Func(testString)
			}
		})
	}
}
