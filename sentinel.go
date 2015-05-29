package main

type Sentinel struct {
}

func NewSentinel() *Sentinel {
	this := new(Sentinel)

	return this
}

func (this *Sentinel) RunForever() {

}
