#include "textflag.h"

TEXT ·IsEnabled(SB), NOSPLIT, $0-1
	CMPL runtime·writeBarrier(SB), $0
	SETNE ret+0(FP)
	RET
