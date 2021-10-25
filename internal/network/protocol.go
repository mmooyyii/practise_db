package network

type Status uint8

const (
	Head = Status(0)
	Body = Status(1)
	Ok   = Status(2)

	defaultSize = 128
)

type Protocol struct {
	data   []byte
	index  int
	status Status
}

func NewProtocol() *Protocol {
	return &Protocol{
		data:   make([]byte, defaultSize),
		status: Head,
		index:  0,
	}
}

func (p *Protocol) Write(b []byte) int {
	return 0
}

func (p *Protocol) IsFinish() bool {
	return true
}

func (p *Protocol) ToQuery() string {
	p.status = Head
	p.data = p.data[:defaultSize]
	p.index = 0
	return "select 1;"
}
