// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package inference

import (
	"encoding/json"
	"errors"
)

// MarshalJSON marshals as JSON.
func (a AvailableWork) MarshalJSON() ([]byte, error) {
	type AvailableWork struct {
		Type         InferType `json:"type" gencodec:"required"`
		InfoHash     string    `json:"infohash" gencodec:"required"`
		RawSize      int64     `json:"rawSize" gencodec:"required"`
		CvmNetworkId int64     `json:"cvm_networkid"`
	}
	var enc AvailableWork
	enc.Type = a.Type
	enc.InfoHash = a.InfoHash
	enc.RawSize = a.RawSize
	enc.CvmNetworkId = a.CvmNetworkId
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (a *AvailableWork) UnmarshalJSON(input []byte) error {
	type AvailableWork struct {
		Type         *InferType `json:"type" gencodec:"required"`
		InfoHash     *string    `json:"infohash" gencodec:"required"`
		RawSize      *int64     `json:"rawSize" gencodec:"required"`
		CvmNetworkId *int64     `json:"cvm_networkid"`
	}
	var dec AvailableWork
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Type == nil {
		return errors.New("missing required field 'type' for AvailableWork")
	}
	a.Type = *dec.Type
	if dec.InfoHash == nil {
		return errors.New("missing required field 'infohash' for AvailableWork")
	}
	a.InfoHash = *dec.InfoHash
	if dec.RawSize == nil {
		return errors.New("missing required field 'rawSize' for AvailableWork")
	}
	a.RawSize = *dec.RawSize
	if dec.CvmNetworkId != nil {
		a.CvmNetworkId = *dec.CvmNetworkId
	}
	return nil
}