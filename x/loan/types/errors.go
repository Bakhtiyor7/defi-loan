package types

import (
	sdkerrors "cosmossdk.io/errors"
)

var (
	ErrInvalidSigner = sdkerrors.Register(ModuleName, 1, "invalid authority")
    ErrWrongLoanState = sdkerrors.Register(ModuleName, 2, "wrong loan state")
    ErrDeadline       = sdkerrors.Register(ModuleName, 3, "deadline")
)