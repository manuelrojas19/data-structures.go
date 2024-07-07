package list

func test() {
	list := createSingleLinkedList[int]()

	list.AddAtBeg(4)
	list.AddAtBeg(1)
}

func main() {
	test()
}
