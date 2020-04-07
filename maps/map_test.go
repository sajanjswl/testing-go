package  maps

import "testing"

func TestSearch(t *testing.T){

	dictionary := Dictionary {"test":"something to search"}

	t.Run("for know word", func(t *testing.T){
		got,_ := dictionary.Search("test")
		want:= "something to search"


		assertStrings(t,got,want)
	})


	t.Run("for unknow word", func(t *testing.T){
		_,err := dictionary.Search("unkwnow")
		




		assertError(t,err,ErrNotFound)
	})


}


func assertStrings(t *testing.T, got, want string){

	t.Helper()

	if got!=want {
		t.Errorf("got %s want %s",got,want)
	}
}

func assertError(t *testing.T, got, want error) {
    t.Helper()

    if got != want {
        t.Errorf("got error %q want %q", got, want)
    }
}
