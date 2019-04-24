package phase1

import (
	"errors"

	"github.com/uzimaru0000/ie03project-gnocchi/back/model"
	"github.com/uzimaru0000/ie03project-gnocchi/back/utils"
)

func CheckCrossPoint(road1, road2 *model.Road) (*model.Point, error) {
	mat := &mat2x2{
		m00: road1.To.Coord.X - road1.From.Coord.X,
		m10: road2.From.Coord.X - road2.To.Coord.X,
		m01: road1.To.Coord.Y - road1.From.Coord.Y,
		m11: road2.From.Coord.Y - road2.To.Coord.Y,
	}
	det := calcDet(mat)
	if utils.NearEqual(det, 0) {
		return nil, errors.New("NA")
	}

	s, t := calsParams(det, road1, road2)

	if !utils.NearEqual(s, 0) && !utils.NearEqual(s, 1) &&
		!utils.NearEqual(t, 0) && !utils.NearEqual(t, 1) &&
		(0 < s && s < 1) && (0 < t && t < 1) {
		x := road1.From.Coord.X + (road1.To.Coord.X-road1.From.Coord.X)*s
		y := road1.From.Coord.Y + (road1.To.Coord.Y-road1.From.Coord.Y)*s
		return &model.Point{X: utils.Round(x, 6), Y: utils.Round(y, 6)}, nil
	}

	return nil, errors.New("NA")
}

func calsParams(det float64, road1, road2 *model.Road) (float64, float64) {
	mat := &mat2x2{
		m00: road2.From.Coord.Y - road2.To.Coord.Y,
		m10: road2.To.Coord.X - road2.From.Coord.X,
		m01: road1.From.Coord.Y - road1.To.Coord.Y,
		m11: road1.To.Coord.X - road1.From.Coord.X,
	}
	vec := &model.Point{
		X: road2.From.Coord.X - road1.From.Coord.X,
		Y: road2.From.Coord.Y - road1.From.Coord.Y,
	}

	dot := calcDot(mat, vec)

	return dot.X / det, dot.Y / det
}

type mat2x2 struct {
	m00 float64
	m10 float64
	m01 float64
	m11 float64
}

func calcDot(mat *mat2x2, vec *model.Point) *model.Point {
	return &model.Point{
		X: mat.m00*vec.X + mat.m10*vec.Y,
		Y: mat.m01*vec.X + mat.m11*vec.Y,
	}
}

func calcDet(mat *mat2x2) float64 {
	return mat.m00*mat.m11 - mat.m10*mat.m01
}
