package main

import "fmt"

type SaveLogInterface interface {
	SaveLog()
}

func SaveLog(st SaveLogInterface) {
	st.SaveLog()
}

type TransferBBL struct {
	name string
}

func (tBBL *TransferBBL) SaveLog() {
	tBBL.name = "test 1"
	fmt.Println("save to database", tBBL.name)
}

type TransferKTB struct {
	name string
}

func (tKTB TransferKTB) SaveLog() {
	tKTB.name = "test 2"
	fmt.Println("save to database", tKTB.name)
}

func main() {
	transA := TransferBBL{name: "BBL"}
	transB := TransferKTB{name: "KTB"}
	SaveLog(&transA)
	SaveLog(transB)

	fmt.Println()
}
