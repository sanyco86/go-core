package point

import "testing"

func TestNew(t *testing.T) {
	p1, err1 := New(1., -1.)
	p2, err2 := New(0., 0.)
	p3, err3 := New(2., 5.)
	t.Logf("result1: %+v error1: %+v", p1, err1)
	t.Logf("result2: %+v error2: %+v", p2, err2)
	t.Logf("result3: %+v error3: %+v", p3, err3)
}
