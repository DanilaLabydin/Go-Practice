package main

func main() {
	var arr [3]int

}

func fillArray(arr [3]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
}
