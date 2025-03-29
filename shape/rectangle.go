package shape

// In Java, a class has to explicitly mention that it is implementing an interface for its functions
// In Go, the struct method does not have to mention that it is implementing an interface.
// It just has to implement the functions/methods mentioned in the interface without explicitly stating

type Rectangle struct {
	Length, Width float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return (r.Length + r.Width) * 2
}
