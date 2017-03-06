package tweetstorm_test

import (
	"testing"

	tweetstorm "github.com/andermalk16/tweetstorm/core"
	"github.com/stretchr/testify/mock"
)

//TestGenerateTweetsText
func TestGenerateTweetsText(t *testing.T) {
	text := "Lorem Ipsum é simplesmente uma simulação de texto da indústria tipográfica e de impressos, e vem sendo utilizado desde o século XVI, quando um impressor desconhecido pegou uma bandeja de tipos e os embaralhou para fazer um livro de modelos de tipos. Lorem Ipsum sobreviveu não só a cinco séculos, como também ao salto para a editoração eletrônica, permanecendo essencialmente inalterado. Se popularizou na década de 60, quando a Letraset lançou decalques contendo passagens de Lorem Ipsum, e mais recentemente quando passou a ser integrado a softwares de editoração eletrônica como Aldus PageMaker."
	tweets := tweetstorm.GenerateTweetsText(text)
	if len(tweets) <= 0 {
		t.Error("expected tweets > 0")
	}
}

//TestGenerateTweetsText
func TestStorm(t *testing.T) {
	text := "Lorem Ipsum é simplesmente uma simulação de texto da indústria tipográfica e de impressos, e vem sendo utilizado desde o século XVI, quando um impressor desconhecido pegou uma bandeja de tipos e os embaralhou para fazer um livro de modelos de tipos. Lorem Ipsum sobreviveu não só a cinco séculos, como também ao salto para a editoração eletrônica, permanecendo essencialmente inalterado. Se popularizou na década de 60, quando a Letraset lançou decalques contendo passagens de Lorem Ipsum, e mais recentemente quando passou a ser integrado a softwares de editoração eletrônica como Aldus PageMaker."
	api := new(TwiterApiSuccessMock)
	tweetstorm.Storm(text, api)
}

//TestGenerateTweetsText
func TestTweetSuccess(t *testing.T) {
	text := []string{"Test.", "Teste2."}
	api := new(TwiterApiSuccessMock)
	tweetstorm.DoTweets(text, api)
}

//TestGenerateTweetsText
func TestTweetFail(t *testing.T) {
	text := []string{"Test.", "Teste2."}
	api := new(TwiterApiFailMock)
	tweetstorm.DoTweets(text, api)

}

type TwiterApiSuccessMock struct {
	mock.Mock
}

func (api TwiterApiSuccessMock) Tweet(text string) bool {
	return true
}

func (api TwiterApiSuccessMock) GetTwitterCredentials() tweetstorm.TwitterCredentials {
	return tweetstorm.TwitterCredentials{}
}

type TwiterApiFailMock struct {
	mock.Mock
}

func (api TwiterApiFailMock) Tweet(text string) bool {
	return false
}

func (api TwiterApiFailMock) GetTwitterCredentials() tweetstorm.TwitterCredentials {
	return tweetstorm.TwitterCredentials{}
}
