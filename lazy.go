package ggk

type Lazier interface{}

//  Efficient way to defer allocating/initializing a class until it is needed
//  (if ever).
type Lazy struct {
	ptr Lazier
}

func NewLazy() *Lazy {
	return &Lazy{}
}

/** Set
Copy src into this, and return a pointer to a copy of it. Note this
will always return the same pointer, so if it is called on a lazy that
has already been initialized, then this will copy over the previous
contents. */
func (lazy *Lazy) Set(src Lazier) Lazier {
	lazy.ptr = src
	return lazy.ptr
}

// Destroy the lazy object (if it was created via init() or set())
func (lazy *Lazy) Reset() {
	toimpl()
}

/** IsValid
Returns true if a valid object has been initialized in the SkTLazy,
false otherwise. */
func (lazy *Lazy) IsValid() bool {
	toimpl()
	return false
}

// Returns the object. This version should only be called when the caller
// knows that the object has been initialized.
func (lazy *Lazy) Get() Lazier {
	toimpl()
	return nil
}

// Like above but doesn't assert if object isn't initialized (in which case
// nullptr is returned).
func (lazy *Lazy) GetMaybeNull() Lazier {
	toimpl()
	return nil
}