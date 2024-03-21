package numcase

import (
	"fmt"

	"github.com/ikonglong/domainerr"
)

// NumRange represents a number range which is from start to end, both inclusive.
type NumRange struct {
	start int
	end   int
}

func NewNumRange(start int, end int) (*NumRange, error) {
	err := domainerr.CheckArgument(start >= 0, "start < 0")
	if err != nil {
		return nil, err
	}
	if err = domainerr.CheckArgument(end >= 0, "end < 0"); err != nil {
		return nil, err
	}
	if err = domainerr.CheckArgument(end >= 0, "end < 0"); err != nil {
		return nil, err
	}
	if err = domainerr.CheckArgument(end >= start, "end < start"); err != nil {
		return nil, err
	}
	return &NumRange{
		start: start,
		end:   end,
	}, err
}

func (r *NumRange) Start() int {
	return r.start
}

func (r *NumRange) End() int {
	return r.end
}

func (r *NumRange) include(num int) bool {
	return num >= r.start && num <= r.end
}

func (r *NumRange) includeRange(r1 *NumRange) bool {
	return r.start <= r1.start && r1.end <= r.end
}

func (r *NumRange) String() string {
	return fmt.Sprintf("[%d, %d]", r.start, r.end)
}
