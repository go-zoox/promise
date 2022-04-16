package promise

// func TestRace(t *testing.T) {
// 	p := Race(
// 		New(func(resolve Resolver, reject Rejecter) {
// 			time.Sleep(100 * time.Millisecond)
// 			resolve(1)
// 		}),
// 		New(func(resolve Resolver, reject Rejecter) {
// 			time.Sleep(10 * time.Microsecond)
// 			resolve(2)
// 		}),
// 	)

// 	resultsX, err := p.Wait()
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	results := resultsX.(int)
// 	if results != 2 {
// 		t.Errorf("expected 2, got %v", results)
// 	}
// }
