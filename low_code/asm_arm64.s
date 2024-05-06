// +build arm64

#include "textflag.h"

// func Add(a, b int) int
TEXT ·Add(SB), NOSPLIT, $0
  MOVD	a+0(FP), R0
  MOVD	b+8(FP), R1
  ADD	R0, R1
  MOVD	R1, ret+16(FP)
  RET
