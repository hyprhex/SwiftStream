package main

func main() {
	// Start the server
	server := NewAPIServer(":3000")
	server.Run()
}
