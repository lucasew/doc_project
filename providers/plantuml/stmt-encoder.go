package plantuml

import (
	"bytes"
	"compress/zlib"
	"fmt"
)

func compress(stmt string) ([]byte, error) {
    var err error
    var compressedBuf bytes.Buffer
    compressorSink, err := zlib.NewWriterLevel(&compressedBuf, zlib.BestCompression)
    if err != nil {
        return nil, err
    }
    _, err = compressorSink.Write([]byte(stmt))
    if err != nil {
        return nil, err
    }
    err = compressorSink.Close()
    if err != nil {
        return nil, err
    }
    return compressedBuf.Bytes(), nil
}

func EncodeStatement(stmt string) (string, error) {
    compressed, err := compress(stmt)
    if err != nil {
        return "", err
    }
    return fmt.Sprintf("~1%s", base64_encode(compressed)), nil
}


// based on: https://github.com/yogendra/plantuml-go/blob/1e0758c537a343cb48d3de014ef23b61109527ef/plantuml-go.go#L89
const mapper = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"
func base64_encode(input []byte) string {
    var buffer bytes.Buffer
    inputLength := len(input)
    for i := 0; i < 3 - inputLength % 3; i++ {
        input = append(input, byte(0))
    }

    for i := 0; i < inputLength; i += 3 {
        b1, b2, b3, b4 := input[i], input[i+1], input[i+2], byte(0)

        b4 = b3 & 0x3f
        b3 = ((b2 & 0xf) << 2) | (b3 >> 6)
        b2 = ((b1 & 0x3) << 4) | (b2 >> 4)
        b1 = b1 >> 2

        for _,b := range []byte{b1,b2,b3,b4} {
            buffer.WriteByte(byte(mapper[b]))
        }
    }
    return string(buffer.Bytes())
}

// func encode64(data []byte) string {
//     ret := []byte{}
//     datalen := len(data)
//     for i := 0; i < datalen; i += 3 {
//         if (i + 2) == datalen {
//             ret = append(ret, append3bytes(data[i], data[i + 1], 0)...)
//         } else if (i + 1 == datalen) {
//             ret = append(ret, append3bytes(data[i], 0, 0)...)
//         } else {
//             ret = append(ret, append3bytes(data[i], data[i + 1], data[i + 2])...)
//         }
//     }
//     return string(ret)
// }

// func append3bytes(b1, b2, b3 byte) []byte {
//     c1 := b1 >> 2
//     c2 := ((b1 & 0x3) << 4) | (b2 >> 4)
//     c3 := ((b2 & 0xF) << 2) | (b3 >> 6)
//     c4 := b3 & 0x3F
//     return []byte{
//         encode6bit(c1 & 0x3F),
//         encode6bit(c2 & 0x3F),
//         encode6bit(c3 & 0x3F),
//         encode6bit(c4 & 0x3F),
//     }
// }

// func encode6bit(b byte) byte {
//     if b < 10 {
//         return b + 48
//     }
//     b -= 10
//     if b < 26 {
//         return b + 65
//     }
//     b -= 26
//     if b < 26 {
//         return b + 97
//     }
//     b -= 26
//     if b == 0 {
//         return '-'
//     }
//     if b == 1 {
//         return '_'
//     }
//     return '?'
// }
