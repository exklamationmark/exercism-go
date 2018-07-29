package letter

import (
	"bytes"
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

// In the separate file frequency.go, you are given a function, Frequency(),
// to sequentially count letter frequencies in a single text.
//
// Perform this exercise on parallelism using Go concurrency features.
// Make concurrent calls to Frequency and combine results to obtain the answer.

var (
	euro = `Freude schöner Götterfunken
Tochter aus Elysium,
Wir betreten feuertrunken,
Himmlische, dein Heiligtum!
Deine Zauber binden wieder
Was die Mode streng geteilt;
Alle Menschen werden Brüder,
Wo dein sanfter Flügel weilt.`

	dutch = `Wilhelmus van Nassouwe
ben ik, van Duitsen bloed,
den vaderland getrouwe
blijf ik tot in den dood.
Een Prinse van Oranje
ben ik, vrij, onverveerd,
den Koning van Hispanje
heb ik altijd geëerd.`

	us = `O say can you see by the dawn's early light,
What so proudly we hailed at the twilight's last gleaming,
Whose broad stripes and bright stars through the perilous fight,
O'er the ramparts we watched, were so gallantly streaming?
And the rockets' red glare, the bombs bursting in air,
Gave proof through the night that our flag was still there;
O say does that star-spangled banner yet wave,
O'er the land of the free and the home of the brave?`
)

func OriginalFrequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func TestConcurrentFrequency(t *testing.T) {
	seq := OriginalFrequency(euro + dutch + us)
	con := ConcurrentFrequency([]string{euro, dutch, us})
	if !reflect.DeepEqual(con, seq) {
		t.Fatal("ConcurrentFrequency wrong result")
	}
}

func TestSequentialFrequency(t *testing.T) {
	oSeq := OriginalFrequency(euro + dutch + us)
	seq := Frequency(euro + dutch + us)
	if !reflect.DeepEqual(oSeq, seq) {
		t.Fatal("Frequency wrong result")
	}
}

func BenchmarkSequentialFrequency(b *testing.B) {
	str := euro + dutch + us

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Frequency(str)
	}
}

func BenchmarkConcurrentFrequency(b *testing.B) {
	input := []string{euro, dutch, us}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConcurrentFrequency(input)
	}
}

func makeWorkload(count, size int) []string {
	rand.Seed(12345)

	var res []string

	for i := 0; i < count; i++ {
		var buf bytes.Buffer
		for j := 0; j < size; j++ {
			buf.WriteByte(byte(rand.Intn(126-32) + 32))
		}
		res = append(res, buf.String())
	}

	return res
}

func BenchmarkApproaches(b *testing.B) {
	// benchmark different approaches

	var testCases = []struct {
		count int
		size  int
		fn    func([]string) FreqMap
		name  string
	}{
		{10000, 10, ConcurrentFrequencyDevillexio, "devillexio-buffered  "},
		{10000, 10, ConcurrentFrequencyDevillexioUnbuffered, "devillexio-unbuffered"},
		{10000, 10, ConcurrentFrequencyExample, "example"},
		{10000, 10, SequentialFrequency, "sequential"},
		{1000, 100, ConcurrentFrequencyDevillexio, "devillexio-buffered  "},
		{1000, 100, ConcurrentFrequencyDevillexioUnbuffered, "devillexio-unbuffered"},
		{1000, 100, ConcurrentFrequencyExample, "example"},
		{1000, 100, SequentialFrequency, "sequential"},
		{100, 1000, ConcurrentFrequencyDevillexio, "devillexio-buffered  "},
		{100, 1000, ConcurrentFrequencyDevillexioUnbuffered, "devillexio-unbuffered"},
		{100, 1000, ConcurrentFrequencyExample, "example"},
		{100, 1000, SequentialFrequency, "sequential"},
		{10, 10000, ConcurrentFrequencyDevillexio, "devillexio-buffered  "},
		{10, 10000, ConcurrentFrequencyDevillexioUnbuffered, "devillexio-unbuffered"},
		{10, 10000, ConcurrentFrequencyExample, "example"},
		{10, 10000, SequentialFrequency, "sequential"},
	}

	for _, tc := range testCases {
		input := makeWorkload(tc.count, tc.size)
		b.Run(fmt.Sprintf("count= %d, size= %d, name= %s", tc.count, tc.size, tc.name), func(b *testing.B) {

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tc.fn(input)
			}
		})
	}
}
