func b() bool { printInt(4); return true }
func main() { if true && b() { printInt(3) } }