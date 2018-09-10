// Code generated by thriftrw v1.14.0. DO NOT EDIT.
// @generated

package services

import (
	"errors"
	"fmt"
	"go.uber.org/multierr"
	"go.uber.org/thriftrw/gen/internal/tests/exceptions"
	"go.uber.org/thriftrw/gen/internal/tests/unions"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/zap/zapcore"
	"strings"
)

// KeyValue_GetManyValues_Args represents the arguments for the KeyValue.getManyValues function.
//
// The arguments for getManyValues are sent and received over the wire as this struct.
type KeyValue_GetManyValues_Args struct {
	Range []Key `json:"range,omitempty"`
}

type _List_Key_ValueList []Key

func (v _List_Key_ValueList) ForEach(f func(wire.Value) error) error {
	for _, x := range v {
		w, err := x.ToWire()
		if err != nil {
			return err
		}
		err = f(w)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v _List_Key_ValueList) Size() int {
	return len(v)
}

func (_List_Key_ValueList) ValueType() wire.Type {
	return wire.TBinary
}

func (_List_Key_ValueList) Close() {}

// ToWire translates a KeyValue_GetManyValues_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *KeyValue_GetManyValues_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Range != nil {
		w, err = wire.NewValueList(_List_Key_ValueList(v.Range)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _List_Key_Read(l wire.ValueList) ([]Key, error) {
	if l.ValueType() != wire.TBinary {
		return nil, nil
	}

	o := make([]Key, 0, l.Size())
	err := l.ForEach(func(x wire.Value) error {
		i, err := _Key_Read(x)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	l.Close()
	return o, err
}

// FromWire deserializes a KeyValue_GetManyValues_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a KeyValue_GetManyValues_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v KeyValue_GetManyValues_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *KeyValue_GetManyValues_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TList {
				v.Range, err = _List_Key_Read(field.Value.GetList())
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a KeyValue_GetManyValues_Args
// struct.
func (v *KeyValue_GetManyValues_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Range != nil {
		fields[i] = fmt.Sprintf("Range: %v", v.Range)
		i++
	}

	return fmt.Sprintf("KeyValue_GetManyValues_Args{%v}", strings.Join(fields[:i], ", "))
}

func _List_Key_Equals(lhs, rhs []Key) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for i, lv := range lhs {
		rv := rhs[i]
		if !(lv == rv) {
			return false
		}
	}

	return true
}

// Equals returns true if all the fields of this KeyValue_GetManyValues_Args match the
// provided KeyValue_GetManyValues_Args.
//
// This function performs a deep comparison.
func (v *KeyValue_GetManyValues_Args) Equals(rhs *KeyValue_GetManyValues_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Range == nil && rhs.Range == nil) || (v.Range != nil && rhs.Range != nil && _List_Key_Equals(v.Range, rhs.Range))) {
		return false
	}

	return true
}

type _List_Key_Zapper []Key

// MarshalLogArray implements zapcore.ArrayMarshaler, enabling
// fast logging of _List_Key_Zapper.
func (l _List_Key_Zapper) MarshalLogArray(enc zapcore.ArrayEncoder) (err error) {
	for _, v := range l {
		enc.AppendString((string)(v))
	}
	return err
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of KeyValue_GetManyValues_Args.
func (v *KeyValue_GetManyValues_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v.Range != nil {
		err = multierr.Append(err, enc.AddArray("range", (_List_Key_Zapper)(v.Range)))
	}
	return err
}

// GetRange returns the value of Range if it is set or its
// zero value if it is unset.
func (v *KeyValue_GetManyValues_Args) GetRange() (o []Key) {
	if v.Range != nil {
		return v.Range
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "getManyValues" for this struct.
func (v *KeyValue_GetManyValues_Args) MethodName() string {
	return "getManyValues"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *KeyValue_GetManyValues_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// KeyValue_GetManyValues_Helper provides functions that aid in handling the
// parameters and return values of the KeyValue.getManyValues
// function.
var KeyValue_GetManyValues_Helper = struct {
	// Args accepts the parameters of getManyValues in-order and returns
	// the arguments struct for the function.
	Args func(
		range2 []Key,
	) *KeyValue_GetManyValues_Args

	// IsException returns true if the given error can be thrown
	// by getManyValues.
	//
	// An error can be thrown by getManyValues only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for getManyValues
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// getManyValues into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by getManyValues
	//
	//   value, err := getManyValues(args)
	//   result, err := KeyValue_GetManyValues_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from getManyValues: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func([]*unions.ArbitraryValue, error) (*KeyValue_GetManyValues_Result, error)

	// UnwrapResponse takes the result struct for getManyValues
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if getManyValues threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := KeyValue_GetManyValues_Helper.UnwrapResponse(result)
	UnwrapResponse func(*KeyValue_GetManyValues_Result) ([]*unions.ArbitraryValue, error)
}{}

func init() {
	KeyValue_GetManyValues_Helper.Args = func(
		range2 []Key,
	) *KeyValue_GetManyValues_Args {
		return &KeyValue_GetManyValues_Args{
			Range: range2,
		}
	}

	KeyValue_GetManyValues_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *exceptions.DoesNotExistException:
			return true
		default:
			return false
		}
	}

	KeyValue_GetManyValues_Helper.WrapResponse = func(success []*unions.ArbitraryValue, err error) (*KeyValue_GetManyValues_Result, error) {
		if err == nil {
			return &KeyValue_GetManyValues_Result{Success: success}, nil
		}

		switch e := err.(type) {
		case *exceptions.DoesNotExistException:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for KeyValue_GetManyValues_Result.DoesNotExist")
			}
			return &KeyValue_GetManyValues_Result{DoesNotExist: e}, nil
		}

		return nil, err
	}
	KeyValue_GetManyValues_Helper.UnwrapResponse = func(result *KeyValue_GetManyValues_Result) (success []*unions.ArbitraryValue, err error) {
		if result.DoesNotExist != nil {
			err = result.DoesNotExist
			return
		}

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// KeyValue_GetManyValues_Result represents the result of a KeyValue.getManyValues function call.
//
// The result of a getManyValues execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type KeyValue_GetManyValues_Result struct {
	// Value returned by getManyValues after a successful execution.
	Success      []*unions.ArbitraryValue          `json:"success,omitempty"`
	DoesNotExist *exceptions.DoesNotExistException `json:"doesNotExist,omitempty"`
}

type _List_ArbitraryValue_ValueList []*unions.ArbitraryValue

func (v _List_ArbitraryValue_ValueList) ForEach(f func(wire.Value) error) error {
	for i, x := range v {
		if x == nil {
			return fmt.Errorf("invalid [%v]: value is nil", i)
		}
		w, err := x.ToWire()
		if err != nil {
			return err
		}
		err = f(w)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v _List_ArbitraryValue_ValueList) Size() int {
	return len(v)
}

func (_List_ArbitraryValue_ValueList) ValueType() wire.Type {
	return wire.TStruct
}

func (_List_ArbitraryValue_ValueList) Close() {}

// ToWire translates a KeyValue_GetManyValues_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *KeyValue_GetManyValues_Result) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = wire.NewValueList(_List_ArbitraryValue_ValueList(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.DoesNotExist != nil {
		w, err = v.DoesNotExist.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("KeyValue_GetManyValues_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _ArbitraryValue_Read(w wire.Value) (*unions.ArbitraryValue, error) {
	var v unions.ArbitraryValue
	err := v.FromWire(w)
	return &v, err
}

func _List_ArbitraryValue_Read(l wire.ValueList) ([]*unions.ArbitraryValue, error) {
	if l.ValueType() != wire.TStruct {
		return nil, nil
	}

	o := make([]*unions.ArbitraryValue, 0, l.Size())
	err := l.ForEach(func(x wire.Value) error {
		i, err := _ArbitraryValue_Read(x)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	l.Close()
	return o, err
}

// FromWire deserializes a KeyValue_GetManyValues_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a KeyValue_GetManyValues_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v KeyValue_GetManyValues_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *KeyValue_GetManyValues_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TList {
				v.Success, err = _List_ArbitraryValue_Read(field.Value.GetList())
				if err != nil {
					return err
				}

			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.DoesNotExist, err = _DoesNotExistException_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if v.DoesNotExist != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("KeyValue_GetManyValues_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a KeyValue_GetManyValues_Result
// struct.
func (v *KeyValue_GetManyValues_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.DoesNotExist != nil {
		fields[i] = fmt.Sprintf("DoesNotExist: %v", v.DoesNotExist)
		i++
	}

	return fmt.Sprintf("KeyValue_GetManyValues_Result{%v}", strings.Join(fields[:i], ", "))
}

func _List_ArbitraryValue_Equals(lhs, rhs []*unions.ArbitraryValue) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for i, lv := range lhs {
		rv := rhs[i]
		if !lv.Equals(rv) {
			return false
		}
	}

	return true
}

// Equals returns true if all the fields of this KeyValue_GetManyValues_Result match the
// provided KeyValue_GetManyValues_Result.
//
// This function performs a deep comparison.
func (v *KeyValue_GetManyValues_Result) Equals(rhs *KeyValue_GetManyValues_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && _List_ArbitraryValue_Equals(v.Success, rhs.Success))) {
		return false
	}
	if !((v.DoesNotExist == nil && rhs.DoesNotExist == nil) || (v.DoesNotExist != nil && rhs.DoesNotExist != nil && v.DoesNotExist.Equals(rhs.DoesNotExist))) {
		return false
	}

	return true
}

type _List_ArbitraryValue_Zapper []*unions.ArbitraryValue

// MarshalLogArray implements zapcore.ArrayMarshaler, enabling
// fast logging of _List_ArbitraryValue_Zapper.
func (l _List_ArbitraryValue_Zapper) MarshalLogArray(enc zapcore.ArrayEncoder) (err error) {
	for _, v := range l {
		err = multierr.Append(err, enc.AppendObject(v))
	}
	return err
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of KeyValue_GetManyValues_Result.
func (v *KeyValue_GetManyValues_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v.Success != nil {
		err = multierr.Append(err, enc.AddArray("success", (_List_ArbitraryValue_Zapper)(v.Success)))
	}
	if v.DoesNotExist != nil {
		err = multierr.Append(err, enc.AddObject("doesNotExist", v.DoesNotExist))
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *KeyValue_GetManyValues_Result) GetSuccess() (o []*unions.ArbitraryValue) {
	if v.Success != nil {
		return v.Success
	}

	return
}

// GetDoesNotExist returns the value of DoesNotExist if it is set or its
// zero value if it is unset.
func (v *KeyValue_GetManyValues_Result) GetDoesNotExist() (o *exceptions.DoesNotExistException) {
	if v.DoesNotExist != nil {
		return v.DoesNotExist
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "getManyValues" for this struct.
func (v *KeyValue_GetManyValues_Result) MethodName() string {
	return "getManyValues"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *KeyValue_GetManyValues_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
