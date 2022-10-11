// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Tiles Tiles
//
// A list of Tile instances.
//
// swagger:discriminator Tiles A list of Tile instances.
type Tiles interface {
	runtime.Validatable
	runtime.ContextValidatable

	Tiles() []TileInfo
	SetTiles([]TileInfo)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type tiles struct {
	tilesField []TileInfo
}

// Tiles gets the tiles of this polymorphic type
func (m *tiles) Tiles() []TileInfo {
	return m.tilesField
}

// SetTiles sets the tiles of this polymorphic type
func (m *tiles) SetTiles(val []TileInfo) {
	m.tilesField = val
}

// UnmarshalTilesSlice unmarshals polymorphic slices of Tiles
func UnmarshalTilesSlice(reader io.Reader, consumer runtime.Consumer) ([]Tiles, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []Tiles
	for _, element := range elements {
		obj, err := unmarshalTiles(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalTiles unmarshals polymorphic Tiles
func UnmarshalTiles(reader io.Reader, consumer runtime.Consumer) (Tiles, error) {
	// we need to read this twice, so first into a buffer
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalTiles(data, consumer)
}

func unmarshalTiles(data []byte, consumer runtime.Consumer) (Tiles, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the A list of Tile instances. property.
	var getType struct {
		AListOfTileInstances string `json:"A list of Tile instances."`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("A list of Tile instances.", "body", getType.AListOfTileInstances); err != nil {
		return nil, err
	}

	// The value of A list of Tile instances. is used to determine which type to create and unmarshal the data into
	switch getType.AListOfTileInstances {
	case "Tiles":
		var result tiles
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid A list of Tile instances. value: %q", getType.AListOfTileInstances)
}

// Validate validates this tiles
func (m *tiles) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *tiles) validateTiles(formats strfmt.Registry) error {
	if swag.IsZero(m.Tiles()) { // not required
		return nil
	}

	for i := 0; i < len(m.Tiles()); i++ {

		if err := m.tilesField[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tiles" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tiles" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this tiles based on the context it is used
func (m *tiles) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *tiles) contextValidateTiles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tiles()); i++ {

		if err := m.tilesField[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tiles" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tiles" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}
