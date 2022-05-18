package password

type Encoder interface {
	Encode(raw string) (string, error)
	Compare(raw string, encoded string) bool
}
