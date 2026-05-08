package frontend

// Scanner is a generic scanner for any type of data, for example: runes in the case of lexing or tokens in the case of
// parsing.
type Scanner[T any] struct {
	source []T // Source data
	empty  T   // Empty value if scanner is out of bounds
	cursor int // Cursor position
}

// NewScanner creates a new scanner for the given source data and empty value.
func NewScanner[T any](source []T, empty T) *Scanner[T] {
	return &Scanner[T]{source: source, empty: empty, cursor: 0}
}

// Cursor returns the current cursor position.
func (s *Scanner[T]) Cursor() int { return s.cursor }

// Reset resets the scanner cursor to the beginning of the source data.
func (s *Scanner[T]) Reset() { s.cursor = 0 }

// IsFinished checks if the scanner has reached the end of the source data.
func (s *Scanner[T]) IsFinished() bool {
	return s.cursor >= len(s.source)
}

// Eat returns the current value and advances the cursor by one. If the scanner is out of bounds, it returns the empty
// value.
func (s *Scanner[T]) Eat() T {
	if s.IsFinished() {
		return s.empty
	}
	value := s.source[s.cursor]
	s.cursor++
	return value
}

// EatN returns the next n values and advances the cursor by n. If the scanner is out of bounds, it returns the empty
// value.
func (s *Scanner[T]) EatN(n int) []T {
	values := make([]T, n)
	for i := 0; i < n; i++ {
		values[i] = s.Eat()
	}
	return values
}

// Peek returns the current value without advancing the cursor. If the scanner is out of bounds, it returns the empty
// value.
func (s *Scanner[T]) Peek() T { return s.PeekAt(0) }

// PeekAt returns the value at the given offset without advancing the cursor. If the scanner is out of bounds, it
func (s *Scanner[T]) PeekAt(offset int) T {
	if s.cursor+offset >= len(s.source) {
		return s.empty
	}
	return s.source[s.cursor+offset]
}

// Prev returns the previous value without advancing the cursor. If the scanner is at the beginning, it returns the
// empty value.
func (s *Scanner[T]) Prev() T {
	if s.cursor == 0 {
		return s.empty
	}
	return s.source[s.cursor-1]
}

// PrevAt returns the value at the given offset from the current cursor position without advancing the cursor. If the
// scanner is out of bounds, it returns the empty value.
func (s *Scanner[T]) PrevAt(offset int) T {
	if s.cursor-offset < 0 {
		return s.empty
	}
	return s.source[s.cursor-offset]
}
