package checker

import "testing"

func BenchmarkIp(b *testing.B) {
	type Test struct {
		IP string
	}
	test := Test{IP: "127.0.0.1"}

	ipChecker := NewChecker()
	ipChecker.Add(Ip("IP"), "wrong ip")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _, _ = ipChecker.Check(test)
	}
}

func BenchmarkNot(b *testing.B) {
	type Test struct {
		NotIP string
	}
	test := Test{NotIP: "127.0.0.1.1"}

	notIPChecker := NewChecker()
	notIPChecker.Add(Not(Ip("IP")), "wrong ip")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _, _ = notIPChecker.Check(test)
	}
}

func BenchmarkMap(b *testing.B) {

	kvMap := make(map[keyStruct]valueStruct)
	keys := []keyStruct{{1}, {2}, {3}}
	for _, key := range keys {
		kvMap[key] = valueStruct{Value: 9}
	}
	m := mapStruct{
		kvMap,
	}

	mapChecker := NewChecker()
	mapRule := Map("Map",
		RangeInt("Key", 1, 10),
		InInt("Value", 8, 9, 10))
	mapChecker.Add(mapRule, "invalid map")

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _, _ = mapChecker.Check(m)
	}
}
