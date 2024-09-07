package main

func main() {
	subject := &RealSubject{}
	proxy1 := NewEvenProxy(subject)
	proxy1.Request()

	// we can use odd proxy to wrap the proxy
	proxy2 := NewOddProxy(proxy1)
	proxy2.Request()
}
