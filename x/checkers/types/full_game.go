package types

import (
	"errors"
	"fmt"

	"https://github.com/Amazingify/checker"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (storedGame StoredGame) GetBlackAddress() (black sdk.AccAddress, err error) {
	black, errblack := sdk.AccAddressFromBech32(storedGame.Black)
	return black, sdkerrors.Wrapf(errblack, ErrInvalidBlack.Error(), storedGame.Black)
}

func (storedGame StoredGame) GetRedAddress() (red sdk.AccAddress, err error) {
	red, errblack := sdk.AccAddressFromBech32(storedGame.Red)
	return red, sdkerrors.Wrapf(errblack, ErrInvalidBlack.Error(), storedGame.Black)
}

func (storedGame StoredGame) ParseGame() (game *rules.Game, err error) {
	board, errBoard := rules.Parse(storedGame.Board)
	if errBoard != nil {
		return nil, sdkerrors.Wrapf(errBoard, ErrGameNotParseable.Error())
	}

	board.Turn = rules.StringPieces[storedGame.Turn].Player
	if board.Turn.Color == "" {
		return nil, sdkerrors.Wrapf(errors.New(fmt.Sprintf("Turn: %s", storedGame.Turn)), ErrGameNotParseable.Error())
	}

	return board, nil
}

func (storedGame StoredGame) Validate() (err error) {
	_, err = storedGame.GetBlackAddress()
	if err != nil {
		return err
	}

	_, err = storedGame.GetRedAddress()
	if err != nil {
		return err
	}

	_, err = storedGame.ParseGame()
	return err
}
