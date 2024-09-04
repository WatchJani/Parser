package main

import "testing"

func BenchmarkSpeedParser(b *testing.B) {
	b.StopTimer()
	var (
		name     = "Janko"
		lastName = "Kondic"
		year     = "21"
		tel      = "+386 66 311 063"
		sex      = "mail"
	)

	user := New(name, lastName, tel, year, sex)
	build := user.Build()

	bufIndex := make([]byte, 4096)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		CreateIndex(build, bufIndex, []int{3, 2, 1})
	}
}
