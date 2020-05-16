package message

import (
	"log"
	"testing"

	"github.com/goui2/ui/base"
)

func TestMessageProcessor(t *testing.T) {
	mp := MD_MessageProcessor.GetClass().New("mp").(MessageProcessor)
	mp.FireMessageChange(nil)
	md2 := base.Extend(MD_MessageProcessor, "mpC1")
	if _, ok := md2.(base.EventProviderMetadata); !ok {
		t.Error("MessageProcessor has invalid type")
	}
	log.Println(MD_MessageProcessor.GetName())
}
