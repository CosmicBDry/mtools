package response

type Data struct {
	RequestId string `json:"request_id,omitempty"`

	Namespace string `json:"namespace,omitempty"`

	Code *int `json:"code"`

	Message string `json:"message,omitempty"`

	Reason string `json:"reason,omitempty"`

	Meta interface{} `json:"meta,omitempty"`

	Data interface{} `json:"data,omitempty"`
}

type Options interface {
	Apply(*Data)
}

type FuncDataOps struct {
	Fun func(*Data)
}

func (F *FuncDataOps) Apply(d *Data) {

	F.Fun(d)
}

func NewOptions(f func(*Data)) Options {

	return &FuncDataOps{
		Fun: f,
	}
}

func SetRequestId(rid string) Options {

	return NewOptions(func(d *Data) {
		d.RequestId = rid
	})

}
