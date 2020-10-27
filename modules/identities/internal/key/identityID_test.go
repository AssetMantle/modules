package key

import (
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"reflect"
	"testing"
)

func Test_IdentityID_Help(t *testing.T) {

	classificationID := base.NewID("classificationID")
	hashID := base.NewID("hashID")
	testIdentityID := NewIdentityID(classificationID, hashID)
	testIdentityID2 := NewIdentityID(classificationID, base.NewID(""))

	t.Run("PositiveCase - is Partial check false", func(t *testing.T) {
		want := false
		if got := testIdentityID.(identityID).IsPartial(); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("PositiveCase - is Partial check true", func(t *testing.T) {
		want := true
		if got := testIdentityID2.(identityID).IsPartial(); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("PositiveCase - Equal Check", func(t *testing.T) {
		want := false
		if got := testIdentityID2.(identityID).Equals(testIdentityID); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

}
