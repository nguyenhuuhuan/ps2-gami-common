package queue

import (
	"math"
	"strconv"

	"github.com/Shopify/sarama"
	"github.com/spaolacci/murmur3"
)

type ModuloPartitioner struct {
	dividend int32 // q will be used in case the dividend was not found
}

func (p *ModuloPartitioner) Partition(message *sarama.ProducerMessage, numPartitions int32) (int32, error) {
	var (
		dividend int32 = p.dividend
	)

	bytes, _ := message.Key.Encode()
	if d, err := strconv.Atoi(string(bytes)); err == nil {
		dividend = int32(d)
	}

	p.dividend = (dividend + 1) % numPartitions

	return dividend % numPartitions, nil
}

func (p *ModuloPartitioner) RequiresConsistency() bool {
	return false
}

func NewModuloPartitioner(topic string) sarama.Partitioner {
	return &ModuloPartitioner{}
}

type HashPartitioner struct {
	d    int32
	hash murmur3.Hash128
}

func (p *HashPartitioner) Partition(message *sarama.ProducerMessage, numPartitions int32) (int32, error) {
	var (
		d = p.d
	)

	key, _ := message.Key.Encode()

	if len(key) > 0 {
		p.hash.Reset()
		_, _ = p.hash.Write(key)
		h := murmur2(key)
		d = int32(h % math.MaxInt32)
		if d < 0 {
			d = -d
		}
	}

	p.d = (d + 1) % numPartitions
	return d % numPartitions, nil
}

func (p *HashPartitioner) RequiresConsistency() bool {
	return false
}

func NewHashPartitioner(topic string) sarama.Partitioner {
	return &HashPartitioner{
		hash: murmur3.New128(),
	}
}

func murmur2(data []byte) uint32 {
	var (
		k      uint32
		length = len(data)
		seed   = 0x9747b28c
		m      = uint32(0x5bd1e995)
		r      = uint32(24)
		// Initialize the hash to a random value
		h       = uint32(seed ^ length)
		length4 = length / 4
	)

	for index := 0; index < length4; index++ {
		i4 := index * 4
		k += uint32((data[i4] & 0xff)) + uint32(uint64(data[i4+1]&0xff))<<8 +
			uint32(uint64(data[i4+2]&0xff)<<16) + uint32(uint64(data[i4+3]&0xff)<<24)
		k *= m
		k ^= k >> r
		k *= m
		k *= m
		h *= m
		h ^= k
	}

	// Handle the last few bytes of the input array
	m4 := length % 4
	switch m4 {
	case 3:
		h ^= uint32(uint64(data[(length-m4)+2]&0xff) << 16)
		fallthrough
	case 2:
		h ^= uint32(uint64(data[(length-m4)+1]&0xff) << 8)
		fallthrough
	case 1:
		h ^= uint32(data[length-m4] & 0xff)
		h *= m
	}

	h ^= h >> 13
	h *= m
	h ^= h >> 15

	return h
}
