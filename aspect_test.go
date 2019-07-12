package aspect

import "testing"

func TestHelloWorld(t *testing.T) {
  want := "hello"
  got := hello();
  if got != want { t.Errorf("hello() = %q, want %q", got, want) }
}
