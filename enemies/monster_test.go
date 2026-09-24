package enemies

import "testing"

func TestNewMonsterByName_SelectsExactMonster(t *testing.T) {
	monster, ok := NewMonsterByName("squelette")
	if !ok {
		t.Fatalf("expected skeleton to be found")
	}
	if monster.NOM != "squelette" {
		t.Fatalf("expected skeleton, got %q", monster.NOM)
	}

	monster, ok = NewMonsterByName("orc")
	if !ok {
		t.Fatalf("expected orc to be found")
	}
	if monster.NOM != "orc" {
		t.Fatalf("expected orc, got %q", monster.NOM)
	}
}
