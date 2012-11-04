#mqtt

An MQTT encoder and decoder,written in Golang.

This library was modified heavily from https://github.com/plucury/mqtt.go and
is API-incompatible with it.

##Functions
* `func Encode(mqtt *Mqtt) (byte[], error)`

	Convert Mqtt struct to bit stream.

* `func EncodeWrite(w io.Writer, mqtt *Mqtt) error`

	Write Mqtt struct to `w`.

* `func Decode(bitstream byte[]) (*Mqtt, error)`

	Convert bit stream to Mqtt struct.

* `func DecodeRead(r io.Reader) (*Mqtt, error)`

	Read Mqtt struct from `r`.
