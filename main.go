package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"protobuf/pb"

	"github.com/golang/protobuf/jsonpb"

	"google.golang.org/protobuf/proto"
)

func main() {
	emproyee := &pb.Employee{
		Id:          1,
		Name:        "tarochan",
		Email:       "tarotaro@gmail.com",
		Occupation:  pb.Occupation_ENGINNEER, // この定数呼び出しに関して
		PhoneNumber: []string{"090-2222-2222"},
		Project:     map[string]*pb.Company_Project{"project-x": &pb.Company_Project{}}, // キーバリュー
		Profile: &pb.Employee_Text{
			Text: "My name is tarochan",
		},
		Birthday: &pb.Date{
			Year:  2019,
			Month: 3,
			Day:   10,
		},
	}

	// データのシリアライズを行う。
	binData, err := proto.Marshal(emproyee)
	if err != nil {
		log.Fatalln("Can't seriarize", err)
	}

	// binファイルの作成を行う。
	// 戻り値がerrorのみだったら1行で書くことが出来る。
	if err := ioutil.WriteFile("test.bin", binData, 0666); err != nil {
		log.Fatalln("Can't write", err)
	}

	// 試しにデシリアライズをしてみる。
	in, err := ioutil.ReadFile("test.bin")
	if err != nil {
		log.Fatalln("Can't read file", err)
	}

	// proto.Unmarshalで使用するので、空の構造体を用意。
	readEmployee := &pb.Employee{}

	// デシリアライズされた結果が構造体にデシリアライズされる。JSONの扱い方と似ている。
	if err = proto.Unmarshal(in, readEmployee); err != nil {
		log.Fatalln("Can't deseriarize", err)
	}

	fmt.Println(readEmployee)
	// id:1  name:"tarochan"  email:"tarotaro@gmail.com"  occupation:ENGINNEER  phone_number:"090-2222-2222"  project:{key:"project-x"  value:{}}  text:"My name is tarochan"  birthday:{year:2019  month:3  day:10}

	// jsonにしてみる。
	m := jsonpb.Marshaler{}
	out, err := m.MarshalToString(emproyee) // jsonにシリアライズする
	if err != nil {
		log.Fatalln("Can't marshal to json", err)
	}
	fmt.Printf("json -> %v\n", out)

	/* Deprecated Use https://pkg.go.dev/google.golang.org/protobuf@v1.28.0/encoding/protojson#MarshalOptions.Format*/
	// 構造体に戻す
	readEmployee = &pb.Employee{} // 空の構造体を用意
	if err := jsonpb.UnmarshalString(out, readEmployee); err != nil {
		log.Fatalln("Can't unmarshal to json", err)
	}
	fmt.Printf("struct -> %v\n", readEmployee)
}
