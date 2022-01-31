func personFirst(p person) {
	got := p.First
	if got != "James" {
		t.Errorf(`First is %v; want "James"`, got)
	}
}