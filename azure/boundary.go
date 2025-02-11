package main

//
//import (
//	"fmt"
//	"os"
//	"time"
//
//	"github.com/Microsoft/cognitive-services-speech-sdk-go/audio"
//	"github.com/Microsoft/cognitive-services-speech-sdk-go/common"
//	"github.com/Microsoft/cognitive-services-speech-sdk-go/speech"
//)
//
//func bookmarkReachedHandler(event speech.SpeechSynthesisBookmarkEventArgs) {
//	defer event.Close()
//	fmt.Println("BookmarkReached event")
//}
//
//func synthesisCanceledHandler(event speech.SpeechSynthesisEventArgs) {
//	defer event.Close()
//	fmt.Println("SynthesisCanceled event")
//}
//
//func synthesisCompletedHandler(event speech.SpeechSynthesisEventArgs) {
//	defer event.Close()
//	fmt.Println("SynthesisCompleted event")
//	fmt.Printf("\tAudioData: %d bytes\n", len(event.Result.AudioData))
//	fmt.Printf("\tAudioDuration: %d\n", event.Result.AudioDuration)
//}
//
//func synthesisStartedHandler(event speech.SpeechSynthesisEventArgs) {
//	defer event.Close()
//	fmt.Println("SynthesisStarted event")
//}
//
//func synthesizingHandler(event speech.SpeechSynthesisEventArgs) {
//	defer event.Close()
//	fmt.Println("Synthesizing event")
//	fmt.Printf("\tAudioData %d bytes\n", len(event.Result.AudioData))
//}
//
//func visemeReceivedHandler(event speech.SpeechSynthesisVisemeEventArgs) {
//	defer event.Close()
//	fmt.Println("VisemeReceived event")
//	fmt.Printf("\tAudioOffset: %dms\n", (event.AudioOffset+5000)/10000)
//	fmt.Printf("\tVisemeID %d\n", event.VisemeID)
//}
//
//func wordBoundaryHandler(event speech.SpeechSynthesisWordBoundaryEventArgs) {
//	defer event.Close()
//	boundaryType := ""
//	switch event.BoundaryType {
//	case 0:
//		boundaryType = "Word"
//	case 1:
//		boundaryType = "Punctuation"
//	case 2:
//		boundaryType = "Sentence"
//	}
//	fmt.Println("WordBoundary event")
//	fmt.Printf("\tBoundaryType %v\n", boundaryType)
//	fmt.Printf("\tAudioOffset: %dms\n", (event.AudioOffset+5000)/10000)
//	fmt.Printf("\tDuration %d\n", event.Duration)
//	fmt.Printf("\tText %s\n", event.Text)
//	fmt.Printf("\tTextOffset %d\n", event.TextOffset)
//	fmt.Printf("\tWordLength %d\n", event.WordLength)
//}
//
//func main1() {
//	// This example requires environment variables named "SPEECH_KEY" and "SPEECH_REGION"
//	speechKey := os.Getenv("SPEECH_KEY")
//	speechRegion := os.Getenv("SPEECH_REGION")
//
//	audioConfig, err := audio.NewAudioConfigFromDefaultSpeakerOutput()
//	if err != nil {
//		fmt.Println("Got an error: ", err)
//		return
//	}
//	defer audioConfig.Close()
//	speechConfig, err := speech.NewSpeechConfigFromSubscription(speechKey, speechRegion)
//	if err != nil {
//		fmt.Println("Got an error: ", err)
//		return
//	}
//	defer speechConfig.Close()
//
//	// Required for WordBoundary event sentences.
//	speechConfig.SetProperty(common.SpeechServiceResponseRequestSentenceBoundary, "true")
//
//	speechSynthesizer, err := speech.NewSpeechSynthesizerFromConfig(speechConfig, audioConfig)
//	if err != nil {
//		fmt.Println("Got an error: ", err)
//		return
//	}
//	defer speechSynthesizer.Close()
//
//	speechSynthesizer.BookmarkReached(bookmarkReachedHandler)
//	speechSynthesizer.SynthesisCanceled(synthesisCanceledHandler)
//	speechSynthesizer.SynthesisCompleted(synthesisCompletedHandler)
//	speechSynthesizer.SynthesisStarted(synthesisStartedHandler)
//	speechSynthesizer.Synthesizing(synthesizingHandler)
//	speechSynthesizer.VisemeReceived(visemeReceivedHandler)
//	speechSynthesizer.WordBoundary(wordBoundaryHandler)
//
//	speechSynthesisVoiceName := "en-US-AvaMultilingualNeural"
//
//	ssml := fmt.Sprintf(`<speak version='1.0' xml:lang='en-US' xmlns='http://www.w3.org/2001/10/synthesis' xmlns:mstts='http://www.w3.org/2001/mstts'>
//            <voice name='%s'>
//                <mstts:viseme type='redlips_front'/>
//                The rainbow has seven colors: <bookmark mark='colors_list_begin'/>Red, orange, yellow, green, blue, indigo, and violet.<bookmark mark='colors_list_end'/>.
//            </voice>
//        </speak>`, speechSynthesisVoiceName)
//
//	// Synthesize the SSML
//	fmt.Printf("SSML to synthesize: \n\t%s\n", ssml)
//	task := speechSynthesizer.SpeakSsmlAsync(ssml)
//
//	var outcome speech.SpeechSynthesisOutcome
//	select {
//	case outcome = <-task:
//	case <-time.After(60 * time.Second):
//		fmt.Println("Timed out")
//		return
//	}
//	defer outcome.Close()
//	if outcome.Error != nil {
//		fmt.Println("Got an error: ", outcome.Error)
//		return
//	}
//
//	if outcome.Result.Reason == common.SynthesizingAudioCompleted {
//		fmt.Println("SynthesizingAudioCompleted result")
//	} else {
//		cancellation, _ := speech.NewCancellationDetailsFromSpeechSynthesisResult(outcome.Result)
//		fmt.Printf("CANCELED: Reason=%d.\n", cancellation.Reason)
//
//		if cancellation.Reason == common.Error {
//			fmt.Printf("CANCELED: ErrorCode=%d\nCANCELED: ErrorDetails=[%s]\nCANCELED: Did you set the speech resource key and region values?\n",
//				cancellation.ErrorCode,
//				cancellation.ErrorDetails)
//		}
//	}
//}
