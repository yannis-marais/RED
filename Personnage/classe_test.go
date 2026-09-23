package ProjetRED
package ProjetRED

import "testing"

func TestClassesKeysMatchMenuChoices(t *testing.T) {
	if _, ok := Classes["Cuirassé"]; !ok {
		t.Fatal("la classe 'Cuirassé' doit exister dans Classes")
	}

	if _, ok := Classes["Mage spirituel"]; !ok {
		t.Fatal("la classe 'Mage spirituel' doit exister dans Classes")
	}
}
