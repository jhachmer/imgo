package transform

import (
	"math"
	"sync"

	"github.com/jhachmer/imgo/internal/types"

	"github.com/jhachmer/imgo/internal/ops"

	"github.com/jhachmer/imgo/pkg/img"
)

type DFT struct {
	Transformed [][]types.Complex
	Magnitude   *DFTMagnitude
	Phase       *DFTPhase
}

func NewDFT(input [][]uint8) *DFT {
	dft := &DFT{
		Transformed: ops.GenerateComplexSlice(input),
	}
	dft.Transformed = DFT2D(dft.Transformed, true)
	dft.Magnitude = NewDFTMagnitude(dft)
	dft.Phase = NewDFTPhase(dft)
	return dft
}

func (dft *DFT) update() {
	dft.Magnitude.Values = dft.DFTMagnitude()
	dft.Phase.Values = dft.DFTPhase()
}

type InverseDFT struct {
	ImageReal [][]float64
}

func NewInverseDFT(dft *DFT) *InverseDFT {
	cols, rows := len(dft.Transformed[0]), len(dft.Transformed)
	var helper = DFT{
		Transformed: DFT2D(dft.Transformed, false),
	}

	idft := ops.GenerateSlice[float64](cols, rows)

	for j := range helper.Transformed {
		for i := range helper.Transformed[j] {
			idft[j][i] = helper.Transformed[j][i].Re
		}
	}

	inverse := &InverseDFT{
		ImageReal: idft,
	}
	return inverse
}

var _ img.Outputable = (*InverseDFT)(nil)

func (iDFT *InverseDFT) Output() [][]uint8 {
	return makeInverseOutput(iDFT.ImageReal)()
}

type DFTMagnitude struct {
	Values [][]float64
}

func NewDFTMagnitude(dft *DFT) *DFTMagnitude {
	return &DFTMagnitude{
		Values: dft.DFTMagnitude(),
	}
}

var _ img.Outputable = (*DFTMagnitude)(nil)

func (dftM *DFTMagnitude) Output() [][]uint8 {
	return makeLogarithmicOutput(dftM.Values)()
}

type DFTPhase struct {
	Values [][]float64
}

func NewDFTPhase(dft *DFT) *DFTPhase {
	return &DFTPhase{
		Values: dft.DFTPhase(),
	}
}

var _ img.Outputable = (*DFTPhase)(nil)

func (dftP *DFTPhase) Output() [][]uint8 {
	return makeLogarithmicOutput(dftP.Values)()
}

// dft1D performs a 1D Discrete Fourier Transform on the input slice of complex numbers.
// forward flag sets whether to use inverse DFT
func dft1D(g []types.Complex, forward bool) []types.Complex {
	M := len(g)
	s := 1 / math.Sqrt(float64(M))

	G := make([]types.Complex, M)

	for m := 0; m < M; m++ {
		sumRe := 0.0
		sumIm := 0.0
		var phim = 2.0 * math.Pi * float64(m) / float64(M)

		for u := 0; u < M; u++ {
			gRe := g[u].Re
			gIm := g[u].Im
			cosw := math.Cos(phim * float64(u))
			sinw := math.Sin(phim * float64(u))
			if !forward {
				sinw = -sinw
			}
			sumRe += gRe*cosw + gIm*sinw
			sumIm += gIm*cosw - gRe*sinw
		}
		G[m] = *types.NewComplex(s*sumRe, s*sumIm)
	}

	return G
}

// DFT2D applies DFT to 2D-slice of complex numbers
// real number image slices can be converted to complex slices using GenerateComplexSlice in util package
// forward flag sets whether to use inverse DFT
func DFT2D(in [][]types.Complex, forward bool) [][]types.Complex {
	rows := len(in)
	cols := len(in[0])
	ret := make([][]types.Complex, rows)
	for i := range ret {
		ret[i] = make([]types.Complex, cols)
	}

	var wg sync.WaitGroup

	for i := 0; i < rows; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ret[i] = dft1D(in[i], forward)
		}()

	}
	wg.Wait()
	ret = ops.TransposeMatrix(ret)

	for i := 0; i < cols; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ret[i] = dft1D(ret[i], forward)
		}()
	}
	wg.Wait()
	return ops.TransposeMatrix(ret)
}

// DFTMagnitude calculates the magnitude of every complex number in DFT result
func (dft *DFT) DFTMagnitude() [][]float64 {
	rows := len(dft.Transformed)
	cols := len(dft.Transformed[0])
	magnitude := make([][]float64, rows)

	for j := 0; j < rows; j++ {
		magnitude[j] = make([]float64, cols)
		for i := 0; i < cols; i++ {
			magnitude[j][i] = dft.Transformed[j][i].Abs()
		}
	}
	return magnitude
}

// DFTPhase calculates the phase of every complex number in DFT result
func (dft *DFT) DFTPhase() [][]float64 {
	rows := len(dft.Transformed)
	cols := len(dft.Transformed[0])
	phase := make([][]float64, rows)

	for j := 0; j < rows; j++ {
		phase[j] = make([]float64, cols)
		for i := 0; i < cols; i++ {
			phase[j][i] = dft.Transformed[j][i].Phase()
		}
	}
	return phase
}

// makeLogarithmicOutput adjusts numbers in given 2D-slice to uint8-range
// number range are in logarithmic scale
// lower frequencies (DC-Values) are shifted to the middle
func makeLogarithmicOutput(values [][]float64) img.OutputFunc {
	cols, rows := len(values[0]), len(values)
	var c float64

	maxMagnitude := ops.FindMaxIn2DSlice(values)

	c = 255 / math.Log(1+math.Abs(maxMagnitude))

	normalized := make([][]uint8, rows)
	for j := 0; j < rows; j++ {
		normalized[j] = make([]uint8, cols)
		for i := 0; i < cols; i++ {
			v := int(c * math.Log(1+math.Abs(values[j][i])))
			normalized[j][i] = ops.ClampPixel(v)
		}
	}
	return func() [][]uint8 {
		return dftShift(normalized)
	}
}

// makeInverseOutput returns displayable 2D-slice
func makeInverseOutput(values [][]float64) img.OutputFunc {
	return func() [][]uint8 {
		cols, rows := len(values[0]), len(values)
		ret := ops.GenerateSlice[uint8](cols, rows)
		// curMax := ops.FindMaxIn2DSlice(values)

		//factor := 255.0 / curMax
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				// val := values[i][j] * factor
				val := values[i][j]
				ret[i][j] = ops.ClampPixel(val)
			}
		}
		return ret
	}
}

// dftShift uses symmetry of DFT to align low-frequency parts of signal (DC) to the center of the image
func dftShift[T any](matrix [][]T) [][]T {
	cols, rows := len(matrix[0]), len(matrix)
	shifted := ops.GenerateSlice[T](cols, rows)
	for j := 0; j < rows; j++ {
		for i := 0; i < cols; i++ {
			newI := (i + cols/2) % cols
			newJ := (j + rows/2) % rows
			shifted[newJ][newI] = matrix[j][i]
		}
	}
	return shifted
}

// ApplyLowPassFilter applies a low-pass filter to the DFT result.
// The cutoff frequency is specified as a fraction of the maximum frequency.
func (dft *DFT) ApplyLowPassFilter(cutoff float64) {
	applyPassFilter(dft, cutoff, false)
}

// ApplyHighPassFilter applies a high-pass filter on the DFT of the image
func (dft *DFT) ApplyHighPassFilter(cutoff float64) {
	applyPassFilter(dft, cutoff, true)
}

func applyPassFilter(dft *DFT, cutoff float64, high bool) {
	rows := len(dft.Transformed)
	cols := len(dft.Transformed[0])
	cutoff = cutoff * float64(rows) // Adjust cutoff to the size of the DFT

	dft.Transformed = dftShift(dft.Transformed)

	for j := 0; j < rows; j++ {
		for i := 0; i < cols; i++ {
			distance := math.Sqrt(math.Pow(float64(j-rows/2), 2) + math.Pow(float64(i-cols/2), 2))
			if high {
				if distance < cutoff {
					dft.Transformed[j][i] = types.Complex{Re: 0, Im: 0}
				}
			} else {
				if distance > cutoff {
					dft.Transformed[j][i] = types.Complex{Re: 0, Im: 0}
				}
			}

		}
	}
	dft.Transformed = dftShift(dft.Transformed)
	dft.update()
}
