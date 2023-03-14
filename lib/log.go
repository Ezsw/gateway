package lib

type TraceContext struct {
	Trace
	CSpanId string
}

type Trace struct {
	TraceId     string
	SpanId      string
	Caller      string
	SrcMethod   string
	HintCode    int64
	HintContent string
}
