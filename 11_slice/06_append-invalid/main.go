package main

func main() {
	greeting := make([]string, 3, 5) // 3 length ["","",""]  5 capacity

	/* "Good morning!", //0
	"Bonjour!",      //1
	"你好",            //2
	"dias!",         // 3
	"Bongiorno!",    //4
	"Ohayo!",        //5
	"Selamat pagi!", // 6
	"Gutten morgen!",
	*/
	greeting[0] = "Good morning"

	greeting[1] = "Bonjour!"
	greeting[2] = "你好"
	greeting[3] = "Ohayo!"
}
