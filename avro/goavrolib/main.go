package goavrolib

import (
	"fmt"
	"time"

	"github.com/linkedin/goavro"
)

func TestGoavro(size int) {
	avroSchema := `
	{
	  "type": "record",
	  "name": "test_schema",
	  "fields": [
		{
		  "name": "time",
		  "type": "long"
		},
		{
		  "name": "customer",
		  "type": "string"
		}
	  ]
	}`

	// Writing OCF data
	// var ocfFileContents bytes.Buffer
	// writer, err := goavro.NewOCFWriter(goavro.OCFConfig{
	// 	W:      &ocfFileContents,
	// 	Schema: avroSchema,
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = writer.Append([]map[string]interface{}{
	// 	{
	// 		"time":     1617104831727,
	// 		"customer": "customer1",
	// 	},
	// 	{
	// 		"time":     1717104831727,
	// 		"customer": "customer2",
	// 	},
	// })
	// fmt.Println("ocfFileContents", ocfFileContents.String())
	// ocfFileContents.Reset()
	// err = writer.Append([]map[string]interface{}{
	// 	{
	// 		"time":     1617104831727,
	// 		"customer": "customer1",
	// 	},
	// })
	// fmt.Println("ocfFileContents", ocfFileContents.String())
	// ocfFileContents.Reset()
	// err = writer.Append([]map[string]interface{}{
	// 	{
	// 		"time":     1717104831727,
	// 		"customer": "customer2",
	// 	},
	// })
	// fmt.Println("ocfFileContents", ocfFileContents.String())
	// ocfFileContents.Reset()

	// Reading OCF data
	// ocfReader, err := goavro.NewOCFReader(strings.NewReader(ocfFileContents.String()))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("Records in OCF File")
	// for ocfReader.Scan() {
	// 	record, err := ocfReader.Read()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println("record", record)
	// }

	c, err := goavro.NewCodec(avroSchema)
	if err != nil {
		return
	}
	// writer, err := goavro.NewOCFWriter(goavro.OCFConfig{
	// 	W:     &ocfFileContents,
	// 	Codec: c,
	// })
	// err = writer.Append([]map[string]interface{}{
	// 	{
	// 		"time":     1617104831727,
	// 		"customer": "customer1",
	// 	},
	// })
	// fmt.Println("ocfFileContents", ocfFileContents.String())
	// ocfFileContents.Reset()
	// err = writer.Append([]map[string]interface{}{
	// 	{
	// 		"time":     1717104831727,
	// 		"customer": "customer2",
	// 	},
	// })
	// fmt.Println("ocfFileContents", ocfFileContents.String())
	// ocfFileContents.Reset()

	var b []byte
	// var s interface{}
	data := make([]interface{}, size)

	for i := range size {
		data[i] = map[string]interface{}{
			"time":     1717104831727 + i,
			"customer": fmt.Sprintf("customer%v", i),
		}
	}

	start := time.Now()
	for _, datum := range data {
		b, err = c.BinaryFromNative(b, datum)
		// fmt.Println(string(b))
	}
	fmt.Println(time.Since(start))

	// b, err = c.BinaryFromNative(b, map[string]interface{}{
	// 	"time":     1617104831727,
	// 	"customer": "customer1",
	// })
	// fmt.Println(string(b))
	// s, b, err = c.NativeFromBinary(b)
	// fmt.Println(s)

	// b, err = c.BinaryFromNative(b, map[string]interface{}{
	// 	"time":     1717104831727,
	// 	"customer": "customer2",
	// })
	// fmt.Println(string(b))
	// s, b, err = c.NativeFromBinary(b)
	// fmt.Println(s)
}
