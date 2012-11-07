# mqtt

An MQTT encoder and decoder,written in Golang.

This library was modified heavily from https://github.com/plucury/mqtt.go and
is API-incompatible with it.

Currently the library's API is unstable.

## Usage

### Decoding Messages

Use the `DecodeOneMessage` function to read a Message from an `io.Reader`, it
will return a Message value. The function can be implemented if more control is
required.

### Encoding Messages

Use the `Encode` method on a Message value to write to an `io.Writer`.
