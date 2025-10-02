package model

import (
	"fmt"
	"math"
)

type Stats interface {
	Update(value float64)
	Average() float64
	ToString(name string) string
}

type stats struct {
	min   float64
	max   float64
	sum   float64
	count int
}

func NewStats() Stats {
	return &stats{
		min: math.Inf(1),
		max: math.Inf(-1),
	}
}

func (s *stats) Update(value float64) {
	if value < s.min {
		s.min = value
	}
	if value > s.max {
		s.max = value
	}
	s.sum += value
	s.count++
}

func (s *stats) Average() float64 {
	if s.count == 0 {
		return 0
	}
	return s.sum / float64(s.count)
}

func (s *stats) ToString(name string) string {
	if s.count == 0 {
		return fmt.Sprintf("%s: no data", name)
	}
	return fmt.Sprintf("%s - Max: %.2f, Min: %.2f, Avg: %.2f", name, s.max, s.min, s.Average())
}
