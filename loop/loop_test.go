package loop

import (
	"testing"

	"github.com/golang/geo/s2"
)

func Test3PLoop(t *testing.T) {
	s2pts := make([]s2.Point, 3)
	s2pts[0] = s2.PointFromLatLng(s2.LatLngFromDegrees(53.51249996891279, -113.53350000023157))
	s2pts[1] = s2.PointFromLatLng(s2.LatLngFromDegrees(53.56133333693556, -113.5860000459106))
	s2pts[2] = s2.PointFromLatLng(s2.LatLngFromDegrees(53.5543333277993, -113.4728333483211))

	l := LoopFenceFromPoints(s2pts)
	t.Log(l)
	rc := &s2.RegionCoverer{MinLevel: 1, MaxLevel: 30, MaxCells: 32}
	if s2.RobustSign(s2pts[0], s2pts[1], s2pts[2]) == s2.Clockwise {
		t.Log("NOT CCW")
		s2pts[0], s2pts[1] = s2pts[1], s2pts[0]
	}
	l = LoopFenceFromPoints(s2pts)
	t.Log(l)
	covering := rc.Covering(l)
	t.Log(covering)
}

func Test4PLoop(t *testing.T) {
	//[47.61649997100236 -122.20166669972009] [47.63100000228975 -122.18449997736279] [47.64100002912518 -122.1180000087143] [47.629666669247634 -122.03883334651307]
	s2pts := make([]s2.Point, 4)
	s2pts[0] = s2.PointFromLatLng(s2.LatLngFromDegrees(47.61649997100236, -122.20166669972009))
	s2pts[1] = s2.PointFromLatLng(s2.LatLngFromDegrees(47.63100000228975, -122.18449997736279))
	s2pts[2] = s2.PointFromLatLng(s2.LatLngFromDegrees(47.64100002912518, -122.1180000087143))
	s2pts[3] = s2.PointFromLatLng(s2.LatLngFromDegrees(47.629666669247634, -122.03883334651307))
	l := LoopFenceFromPoints(s2pts)
	t.Log(l)
	rc := &s2.RegionCoverer{MinLevel: 1, MaxLevel: 30, MaxCells: 32}
	if s2.RobustSign(s2pts[0], s2pts[1], s2pts[2]) == s2.Clockwise {
		t.Log("NOT CCW")
		s2pts[0], s2pts[1] = s2pts[1], s2pts[0]
	}
	l = LoopFenceFromPoints(s2pts)
	t.Log(l)
	covering := rc.Covering(l)
	t.Log(covering)

}

func TestCoverIsNotRectBases(t *testing.T) {
	points := []s2.Point{
		s2.PointFromLatLng(s2.LatLng{Lat: 48.7195648396, Lng: 2.3574256897}),
		s2.PointFromLatLng(s2.LatLng{Lat: 48.7278882281, Lng: 2.40257263184}),
		s2.PointFromLatLng(s2.LatLng{Lat: 48.7273220548, Lng: 2.4028301239}),
		s2.PointFromLatLng(s2.LatLng{Lat: 48.7189985726, Lng: 2.35776901245}),
		s2.PointFromLatLng(s2.LatLng{Lat: 48.7195648396, Lng: 2.3574256897}),
	}

	marker := s2.CellIDFromLatLng(s2.LatLng{Lat: 48.72163165982755, Lng: 2.35806941986084})
	inside := s2.CellIDFromLatLng(s2.LatLng{Lat: 48.71964977907006, Lng: 2.358584403991699})

	rc := &s2.RegionCoverer{MinLevel: 1, MaxLevel: 30, MaxCells: 32}

	if s2.RobustSign(points[0], points[1], points[2]) == s2.Clockwise {
		t.Log("NOT CCW reversing")
		// reverse the slice
		for i := len(points)/2 - 1; i >= 0; i-- {
			opp := len(points) - 1 - i
			points[i], points[opp] = points[opp], points[i]
		}
	}

	// test with s2 rect cover
	lr := s2.LoopFromPoints(points)
	if lr.IsEmpty() || lr.IsFull() {
		t.Fatal("invalid loop")
	}

	t.Log(lr.RectBound())
	if !lr.RectBound().ContainsCell(s2.CellFromCellID(inside)) {
		t.Fatal("loop bound is invalid inside shoud be contained!")
	}
	if !lr.RectBound().ContainsCell(s2.CellFromCellID(marker)) {
		t.Fatal("loop bound is invalid inside shoud not be contained!")
	}

	covering := rc.Covering(lr.RectBound())
	if len(covering) != 32 {
		t.Fatal("covering failed got only", len(covering))
	}

	coverString := ""
	for _, cov := range covering {
		coverString = coverString + cov.ToToken() + " "
	}

	t.Log(coverString)

	// the marker will be inside if it's using a rect bound
	found := false
	for _, c := range covering {
		if c.Contains(marker) {
			found = true
			break
		}
	}

	if !found {
		t.Fatal("our test is not working anymore we should have found Split in the rect based cover")
	}

	l := LoopFenceFromPoints(points)
	if l.IsEmpty() || l.IsFull() {
		t.Fatal("invalid loop")
	}

	if !l.ContainsCell(s2.CellFromCellID(inside)) {
		t.Fatal("loop is invalid inside shoud be contained!")
	}
	if l.ContainsCell(s2.CellFromCellID(marker)) {
		t.Fatal("loop is invalid inside shoud not be contained!")
	}

	// the marker will be outside if it's using a real shape bound
	covering = rc.Covering(l)
	coverString = ""
	for _, cov := range covering {
		coverString = coverString + cov.ToToken() + " "
	}

	t.Log(coverString)

	if len(covering) != 32 {
		t.Fatal("covering failed got only", len(covering))
	}

	for _, c := range covering {
		if c.Contains(marker) {
			t.Fatal("covering should be precise and not rect based, Split should not be contained in the answer")
		}
	}

}
