package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to Leonardo", func(t *testing.T) {
		got := Hello("Leonardo Da Vinci", "")
		want := "Hello, Leonardo Da Vinci!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to the world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Juan", "Spanish")
		want := "Hola, Juan!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Bonnie", "French")
		want := "Bonjour, Bonnie!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in German", func(t *testing.T) {
		got := Hello("Hans", "German")
		want := "Hallo, Hans!"

		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
