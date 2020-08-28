package statsd

type bufferPool struct {
	pool sync.Pool // of *statsdBuffer
}

func newBufferPool(poolSize, bufferMaxSize, bufferMaxElements int) *bufferPool {
	return &bufferPool{
		pool: sync.Pool{
			New: func() interface{} {
				return newStatsdBuffer(bufferMaxSize, bufferMaxElements)
			},
		},
	}
}

func (p *bufferPool) borrowBuffer() *statsdBuffer {
	return p.pool.Get().(*statsdBuffer)
}

func (p *bufferPool) returnBuffer(buffer *statsdBuffer) {
	p.pool.Put(buffer)
}
