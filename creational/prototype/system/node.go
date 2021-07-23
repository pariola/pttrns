package system

type node interface {
	print(string)
	clone() node
}
