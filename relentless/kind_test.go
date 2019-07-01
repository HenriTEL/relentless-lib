package relentless

import "testing"

func TestToKind(t *testing.T) {
	if toKind("SPELL") != SPELL {
		t.Fail()
	}
}

func TestKindString(t *testing.T) {
	if CREATURE.String() != "CREATURE" {
		t.Fail()
	}
}

func TestKindUnmarshalJSON(t *testing.T) {
	var k Kind
	jsonType := []byte(`"spell"`)

	if k.UnmarshalJSON(jsonType); k != SPELL {
		t.Fail()
	}
}
