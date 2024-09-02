package utils

import (
    "encoding/json"
    "encoding/gob"
    "bytes"
)

// SerializeJSON serializa una estructura a JSON.
func SerializeJSON(data interface{}) ([]byte, error) {
    return json.Marshal(data)
}

// DeserializeJSON deserializa datos JSON a una estructura.
func DeserializeJSON(data []byte, v interface{}) error {
    return json.Unmarshal(data, v)
}

// SerializeBinary serializa una estructura a un formato binario.
func SerializeBinary(data interface{}) ([]byte, error) {
    var buf bytes.Buffer
    encoder := gob.NewEncoder(&buf)
    if err := encoder.Encode(data); err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}

// DeserializeBinary deserializa datos binarios a una estructura.
func DeserializeBinary(data []byte, v interface{}) error {
    buf := bytes.NewBuffer(data)
    decoder := gob.NewDecoder(buf)
    return decoder.Decode(v)
}

