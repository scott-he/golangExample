var c chan string;
c = make(chan string); // HL

c <- "Hello"; // infix send // HL
// in a different goroutine
greeting := <-c; // prefix receive // HL

cc := new(chan chan string);
cc <- c; // handing off a capability
