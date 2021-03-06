// Copyright The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

// Package rand implements pseudo-random number generators.
//
// Random numbers are generated by a Source. Top-level functions, such as Float64
// and Int, use a default shared Source that produces a deterministic sequence of
// values each time a program is run. Use the Seed function to initialize the
// default Source if different behavior is required for each run. The default
// Source is safe for concurrent use by multiple goroutines.

// rand 包实现了伪随机数生成器.
//
// 随机数由一个 Source 生成。像 Float64 和 Int
// 这样的顶级函数使用默认共享的 Source，
// 它会在每次程序运行时产生一系列确定的值。若每次运行都需要不同的行为，需使用 Seed 函数来初始化默认的
// Source。对于多Go程并发来说，默认的 Source 是安全的。
package rand

// ExpFloat64 returns an exponentially distributed float64 in the range (0,
// +math.MaxFloat64] with an exponential distribution whose rate parameter (lambda)
// is 1 and whose mean is 1/lambda (1) from the default Source. To produce a
// distribution with a different rate parameter, callers can adjust the output
// using:
//
//	sample = ExpFloat64() / desiredRateParameter
func ExpFloat64() float64

// Float32 returns, as a float32, a pseudo-random number in [0.0,1.0) from the
// default Source.
func Float32() float32

// Float64 returns, as a float64, a pseudo-random number in [0.0,1.0) from the
// default Source.
func Float64() float64

// Int returns a non-negative pseudo-random int from the default Source.
func Int() int

// Int31 returns a non-negative pseudo-random 31-bit integer as an int32 from the
// default Source.
func Int31() int32

// Int31n returns, as an int32, a non-negative pseudo-random number in [0,n) from
// the default Source. It panics if n <= 0.
func Int31n(n int32) int32

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64 from the
// default Source.
func Int63() int64

// Int63n returns, as an int64, a non-negative pseudo-random number in [0,n) from
// the default Source. It panics if n <= 0.
func Int63n(n int64) int64

// Intn returns, as an int, a non-negative pseudo-random number in [0,n) from the
// default Source. It panics if n <= 0.
func Intn(n int) int

// NormFloat64 returns a normally distributed float64 in the range
// [-math.MaxFloat64, +math.MaxFloat64] with standard normal distribution (mean =
// 0, stddev = 1) from the default Source. To produce a different normal
// distribution, callers can adjust the output using:
//
//	sample = NormFloat64() * desiredStdDev + desiredMean
func NormFloat64() float64

// Perm returns, as a slice of n ints, a pseudo-random permutation of the integers
// [0,n) from the default Source.
func Perm(n int) []int

// Seed uses the provided seed value to initialize the default Source to a
// deterministic state. If Seed is not called, the generator behaves as if seeded
// by Seed(1).
func Seed(seed int64)

// Uint32 returns a pseudo-random 32-bit value as a uint32 from the default Source.
func Uint32() uint32

// A Rand is a source of random numbers.
type Rand struct {
	// contains filtered or unexported fields
}

// New returns a new Rand that uses random values from src to generate other random
// values.
func New(src Source) *Rand

// ExpFloat64 returns an exponentially distributed float64 in the range (0,
// +math.MaxFloat64] with an exponential distribution whose rate parameter (lambda)
// is 1 and whose mean is 1/lambda (1). To produce a distribution with a different
// rate parameter, callers can adjust the output using:
//
//	sample = ExpFloat64() / desiredRateParameter

// ExpFloat64 按照率参数（lambda）为 1，均值为 1/lambda (1) 来返回一个在区间 (0, +math.MaxFloat64]
// 内程指数分布的 float64。要以不同的率参数产生一个分布， 调用者只需通过：
//
//	范例 = ExpFloat64() / 所需的率参数
//
// 来调整输出即可。
func (r *Rand) ExpFloat64() float64

// Float32 returns, as a float32, a pseudo-random number in [0.0,1.0).
func (r *Rand) Float32() float32

// Float64 returns, as a float64, a pseudo-random number in [0.0,1.0).
func (r *Rand) Float64() float64

// Int returns a non-negative pseudo-random int.
func (r *Rand) Int() int

// Int31 returns a non-negative pseudo-random 31-bit integer as an int32.
func (r *Rand) Int31() int32

// Int31n returns, as an int32, a non-negative pseudo-random number in [0,n). It
// panics if n <= 0.
func (r *Rand) Int31n(n int32) int32

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (r *Rand) Int63() int64

// Int63n returns, as an int64, a non-negative pseudo-random number in [0,n). It
// panics if n <= 0.
func (r *Rand) Int63n(n int64) int64

// Intn returns, as an int, a non-negative pseudo-random number in [0,n). It panics
// if n <= 0.
func (r *Rand) Intn(n int) int

// NormFloat64 returns a normally distributed float64 in the range
// [-math.MaxFloat64, +math.MaxFloat64] with standard normal distribution (mean =
// 0, stddev = 1). To produce a different normal distribution, callers can adjust
// the output using:
//
//	sample = NormFloat64() * desiredStdDev + desiredMean

// NormFloat64 按照标准正态分布（均值 = 0，标准差 = 1）来返回一个在区间 (0, +math.MaxFloat64]
// 内程正态分布的 float64。要产生一个不同的正态分布， 调用者只需通过：
//
//	范例 = NormFloat64() * 所需的标准差 + 所需的均值
//
// 来调整输出即可。
func (r *Rand) NormFloat64() float64

// Perm returns, as a slice of n ints, a pseudo-random permutation of the integers
// [0,n).
func (r *Rand) Perm(n int) []int

// Seed uses the provided seed value to initialize the generator to a deterministic
// state.
func (r *Rand) Seed(seed int64)

// Uint32 returns a pseudo-random 32-bit value as a uint32.
func (r *Rand) Uint32() uint32

// A Source represents a source of uniformly-distributed pseudo-random int64 values
// in the range [0, 1<<63).
type Source interface {
	Int63() int64
	Seed(seed int64)
}

// NewSource returns a new pseudo-random Source seeded with the given value.
func NewSource(seed int64) Source

// A Zipf generates Zipf distributed variates.

// Zipf 生成齐夫分布变量。
type Zipf struct {
	// contains filtered or unexported fields
}

// NewZipf returns a Zipf generating variates p(k) on [0, imax] proportional to
// (v+k)**(-s) where s>1 and k>=0, and v>=1.
func NewZipf(r *Rand, s float64, v float64, imax uint64) *Zipf

// Uint64 returns a value drawn from the Zipf distribution described by the Zipf
// object.
func (z *Zipf) Uint64() uint64
