package main

type uint16SliceSortAsc []uint16

func (u uint16SliceSortAsc) Len() int {
	return len(u)
}

func (u uint16SliceSortAsc) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}
func (u uint16SliceSortAsc) Less(i, j int) bool {
	return u[i] < u[j]
}
