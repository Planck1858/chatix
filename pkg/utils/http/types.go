package utils

import (
	"encoding/json"
	"time"
)

type JsonNullTimestamp struct {
	time.Time
}

type JsonNullString struct {
	String string
}

type JsonNullTime struct {
	time.Time
}

func (v *JsonNullTimestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String)
}

func (v *JsonNullTimestamp) UnmarshalJSON(data []byte) error {
	var x *string
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	return nil
}

func (v *JsonNullString) MarshalJSON() ([]byte, error) {
	if v.String != "" {
		return json.Marshal(v.String)
	} else {
		return json.Marshal(nil)
	}
}

func (v *JsonNullString) UnmarshalJSON(data []byte) error {
	var x *string

	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.String = *x
	} else {
		v.String = ""
	}

	return nil
}

func (v *JsonNullTime) MarshalJSON() ([]byte, error) {
	if !v.IsZero() {
		return json.Marshal(v.Time)
	} else {
		return json.Marshal(nil)
	}
}

func (v *JsonNullTime) UnmarshalJSON(data []byte) error {
	var x *time.Time

	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}

	if x == nil || x.IsZero() {
		v.Time = time.Time{}
	} else {
		v.Time = *x
	}

	return nil
}
