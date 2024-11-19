package main

//
//import (
//	"bytes"
//	"compress/gzip"
//	"context"
//	"encoding/binary"
//	"fmt"
//	"io"
//	"net/http"
//	"net/http/httputil"
//	"sync"
//	"time"
//
//	"github.com/gorilla/websocket"
//)
//
//type volcRepo struct {
//	log  Logger
//	conf Config
//	pool sync.Pool
//}
//
//type Logger interface {
//	WithContext(ctx context.Context) Logger
//	Infof(format string, args ...interface{})
//	Errorf(format string, args ...interface{})
//}
//
//type Config struct {
//	VolcTts VolcTtsConfig
//}
//
//type VolcTtsConfig struct {
//	BearerToken string
//	Url         string
//}
//
//type dto struct{}
//
//func (d *dto) TtsRequest() {}
//
//func (d *dto) TtsResponse() {}
//
//func gzipCompress(input []byte) []byte {
//	var b bytes.Buffer
//	w := gzip.NewWriter(&b)
//	_, _ = w.Write(input)
//	_ = w.Close()
//	return b.Bytes()
//}
//
//func (v *volcRepo) StreamSynth(ctx context.Context, req *dto.TtsRequest) (chan []byte, chan dto.TtsResponse, error) {
//	fmt.Printf("StreamSynth, req: %+v\n", req)
//	v.log.WithContext(ctx).Infof("StreamSynth req: %+v", req)
//
//	fileChan := make(chan []byte, 10)
//	sendChan := make(chan dto.TtsResponse, 10)
//
//	go func() {
//		defer close(fileChan)
//		defer close(sendChan)
//
//		startTime := ctx.Value("StartTime")
//
//		var err error
//		var c *websocket.Conn
//		header := http.Header{"Authorization": []string{fmt.Sprintf("Bearer;%s", v.conf.VolcTts.BearerToken)}}
//
//		for attempt := 0; attempt <= 1; attempt++ {
//			conn, cRes, err := websocket.DefaultDialer.Dial(v.conf.VolcTts.Url, header)
//			if err != nil {
//				fmt.Printf("volc websocket dial err:%s\n", err)
//				var wsResBody string
//				if cRes != nil {
//					cResB, _ := httputil.DumpResponse(cRes, true)
//					wsResBody = string(cResB)
//				}
//				v.log.WithContext(ctx).Errorf("volc_websocket_dial_retry, err: %+v, res: %s", err, wsResBody)
//				time.Sleep(time.Millisecond * 100)
//				continue
//			}
//			c = conn
//			break
//		}
//
//		if c == nil {
//			v.log.WithContext(ctx).Errorf("volc_websocket_dial_fail, err:%s", err.Error())
//			return
//		}
//
//		defer c.Close()
//
//		connCost := time.Since(startTime.(time.Time))
//		fmt.Printf("volc_tts, conn_cost: %v\n", connCost)
//		v.log.WithContext(ctx).Infof("volc_tts, conn_cost: %v", connCost)
//
//		input := v.formatReqParams(ctx, req)
//		input = gzipCompress(input)
//		payloadSize := len(input)
//		payloadArr := make([]byte, 4)
//		binary.BigEndian.PutUint32(payloadArr, uint32(payloadSize))
//
//		clientRequest := make([]byte, len(defaultHeader))
//		copy(clientRequest, defaultHeader)
//		clientRequest = append(clientRequest, payloadArr...)
//		clientRequest = append(clientRequest, input...)
//
//		err = c.WriteMessage(websocket.BinaryMessage, clientRequest)
//
//		if err != nil {
//			fmt.Println("write message fail, err:", err.Error())
//			v.log.WithContext(ctx).Errorf("volc_write_message_fail, err:%s", err.Error())
//			return
//		}
//
//		write_cost := time.Since(startTime.(time.Time))
//		fmt.Printf("volc_tts, write_cost: %v\n", write_cost)
//		v.log.WithContext(ctx).Infof("volc_tts, write_cost: %v", write_cost)
//
//		firstA := true
//		first := true
//
//		chunkSize := entity.DefaultChunkSize
//
//		if req.ChunkSize > 0 {
//			chunkSize = req.ChunkSize
//		}
//
//		for {
//			var message []byte
//
//			_, message, err := c.ReadMessage()
//
//			if err != nil {
//				if err == io.EOF {
//					break
//				}
//				fmt.Println("read message fail, err:", err.Error())
//				v.log.WithContext(ctx).Errorf("volc_read_message_fail, err:%s", err.Error())
//				break
//			}
//
//			if firstA {
//				firstA_cost := time.Since(startTime.(time.Time))
//				fmt.Printf("volc_tts, firstA_cost: %v\n", firstA_cost)
//				v.log.WithContext(ctx).Infof("volc_tts, firstA_cost: %v", firstA_cost)
//				firstA = false
//			}
//
//			resp, err := v.parseResponse(ctx, message)
//
//			fmt.Printf("audio: %d,isLast: %v\n", len(resp.Audio), resp.IsLast)
//
//			v.log.WithContext(ctx).Infof("volc_message,audio: %d,isLast:%v", len(resp.Audio), resp.IsLast)
//
//			if err != nil {
//				fmt.Println("parse response fail ,err:", err.Error())
//
//				v.log.WithContext(ctx).Errorf(
//					"volc_parse_response_fail ,err:%s",
//					err.Error(),
//				)
//
//				break
//
//			}
//
//			if len(resp.Audio) == 0 && !resp.IsLast {
//
//				continue
//
//			}
//
//			reader := bytes.NewReader(resp.Audio)
//
//			for {
//
//				chunk := make([]byte, chunkSize)
//
//				n, err := reader.Read(chunk)
//
//				if err != nil {
//
//					if err == io.EOF {
//
//						break
//
//					}
//
//					fmt.Println(
//						"read chunk fail ,err:",
//						err.Error(),
//					)
//
//					v.log.WithContext(
//						ctx,
//					).Errorf(
//						"volc_read_chunk_fail ,err:%s",
//						err.Error(),
//					)
//
//					break
//
//				}
//
//				if first {
//
//					first_cost := time.Since(startTime.(time.Time))
//
//					fmt.Printf(
//						"volc_tts ,first_cost:%v\n",
//						first_cost,
//					)
//
//					v.log.WithContext(
//						ctx,
//					).Infof(
//						"volc_tts ,first_cost:%v",
//						first_cost,
//					)
//
//					first = false
//
//				}
//
//				fileChan <- chunk[:n]
//
//				sendChan <- dto.TtsResponse{
//
//					Result: chunk[:n],
//				}
//
//			}
//
//			if resp.IsLast {
//
//				break
//
//			}
//
//		}
//
//		return
//
//	}()
//
//	return fileChan, sendChan, nil
//
//}
