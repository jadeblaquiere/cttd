// Copyright (c) 2016 The ciphrtxt developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package ciphrtxt

import (
    "testing"
    //"math/big"
    "net/http"
    "io/ioutil"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "strconv"
)

type THeaderListResponse struct {
    HeaderList []string `json:"header_list"`
}

func TestDeserializeSerialize (t *testing.T) {
    res, err := http.Get("http://ciphrtxt.com:7754/api/header/list/since/0")
    if err != nil {
        fmt.Println("whoops:", err)
        t.Fail()
    }
    
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println("whoops:", err)
        t.Fail()
    }
    
    s := new(THeaderListResponse)
    err = json.Unmarshal(body, &s)
    if(err != nil){
        fmt.Println("whoops:", err)
        t.Fail()
    }
    
    count := int(0)
    for _, hdr := range s.HeaderList {
        h := new(RawMessageHeader)
        h.Deserialize(hdr)
        hdr_out := h.Serialize()
        count += 1
        
        if hdr != hdr_out {
            t.Fail()
        }
    
        dbk, err := h.DBKeys()
        if(err != nil){
            fmt.Println("whoops:", err)
            t.Fail()
        }
        //fmt.Println(hdr)
        //fmt.Println("    " + hex.EncodeToString(dbk.date))
        //fmt.Println("    " + hex.EncodeToString(dbk.expire))
        //fmt.Println("    " + hex.EncodeToString(dbk.I))
        //fmt.Println()
        
        t64, err := strconv.ParseUint(hex.EncodeToString(dbk.date)[1:9], 16, 32)
        time := uint32(t64)
        if (h.time != time) || (h.I != hex.EncodeToString(dbk.date)[9:75]) {
            t.Fail()
        }
        
        t64, err = strconv.ParseUint(hex.EncodeToString(dbk.expire)[1:9], 16, 32)
        expire := uint32(t64)
        if (h.expire != expire) || (h.I != hex.EncodeToString(dbk.expire)[9:75]) {
            t.Fail()
        }
        
        if h.I != hex.EncodeToString(dbk.I) {
            t.Fail()
        }
    }
}

func TestOpenHeaderCache (t *testing.T) {
    hc1, err := OpenHeaderCache("violet.ciphrtxt.com", 7754, "testdb")
    if err != nil {
        t.Fail()
    }
    defer hc1.Close()

    hc2, err := OpenHeaderCache("indigo.ciphrtxt.com", 7754, "")
    if err != nil {
        t.Fail()
    }
    defer hc2.Close()
}

func TestSyncHeaderCache (t *testing.T) {
    hc1, err := OpenHeaderCache("violet.ciphrtxt.com", 7754, "testdb")
    if err != nil {
        t.Fail()
    }
    defer hc1.Close()

    hc2, err := OpenHeaderCache("indigo.ciphrtxt.com", 7754, "")
    if err != nil {
        t.Fail()
    }
    defer hc2.Close()
}

