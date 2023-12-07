package encoder

import "sync"

var _consolePool = sync.Pool{New: func() interface{} {
	return &consoleEncoder{}
}}

func getConsoleEncoder() *consoleEncoder {
	return _consolePool.Get().(*consoleEncoder)
}

func putConsoleEncoder(enc *consoleEncoder) {
	if enc.reflectBuf != nil {
		enc.reflectBuf.Free()
	}
	enc.Config = nil
	enc.buf = nil
	enc.openNamespaces = 0
	enc.reflectBuf = nil
	enc.reflectEnc = nil
	_consolePool.Put(enc)
}

var _sliceEncoderPool = sync.Pool{
	New: func() interface{} {
		return &sliceArrayEncoder{elems: make([]interface{}, 0, 2)}
	},
}

func getSliceEncoder() *sliceArrayEncoder {
	return _sliceEncoderPool.Get().(*sliceArrayEncoder)
}

func putSliceEncoder(e *sliceArrayEncoder) {
	e.elems = e.elems[:0]
	_sliceEncoderPool.Put(e)
}
