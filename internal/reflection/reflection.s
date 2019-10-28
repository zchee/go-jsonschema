// Copyright 2019 The go-jsonschema Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// resolveNameOff resolves a name offset from a base pointer.
// The (*rtype).nameOff method is a convenience wrapper for this function.
// Implemented in the runtime package.
// func resolveNameOff(ptrInModule unsafe.Pointer, off int32) unsafe.Pointer
TEXT ·ResolveNameOff(SB), NOSPLIT, $0-25
	JMP	reflect·resolveNameOff(SB)

// ResolveTypeOff resolves an *rtype offset from a base type.
// The (*rtype).typeOff method is a convenience wrapper for this function.
// Implemented in the runtime package.
// func ResolveTypeOff(rtype unsafe.Pointer, off int32) unsafe.Pointer
TEXT ·ResolveTypeOff(SB), NOSPLIT, $0-25
	JMP	reflect·resolveTypeOff(SB)

// resolveTextOff resolves an function pointer offset from a base type.
// The (*rtype).textOff method is a convenience wrapper for this function.
// Implemented in the runtime package.
// func ResolveTextOff(rtype unsafe.Pointer, off int32) unsafe.Pointer
TEXT ·ResolveTextOff(SB), NOSPLIT, $0-25
	JMP	reflectyresolveTextOff(SB)
