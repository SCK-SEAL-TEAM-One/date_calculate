#!/bin/bash

# ตั้งค่า Environment เริ่มต้นของโปรเจค
export PATH=$PATH:./node_modules/.bin
export GOPATH=$(pwd)

# สั่งจัด Format ของโค้ด
go fmt dateCalculate

# สั่งติดตั้ง package dateCalculate
go install dateCalculate

# สั่งรัน Unit Test ของ package dateCalculate
go test dateCalculate

# สั่ง build โปรเจค ให้ไฟล์ Binary อยู่ที่ directory bin
go build -o ./bin/mainDateCalculate main

# สั่งรัน Binary File ของโปรเจค
./bin/mainDateCalculate &

# สั่งรัน Acceptance Test โดยใช้ newman (Postman CLI)
newman run duration.postman_collection.json

# ปิด Process API
killall -9 mainDateCalculate