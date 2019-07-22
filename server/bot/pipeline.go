package bot

import (
	"github.com/rumblefrog/source-chat-relay/server/entity"
	"github.com/rumblefrog/source-chat-relay/server/relay"
)

func Listen() {
	for {
		select {
		case message := <-relay.Instance.Bot:
			for _, guild := range RelayBot.State.Guilds {
				for _, channel := range guild.Channels {
					tEntity, err := entity.GetEntity(channel.ID)

					if err != nil {
						continue
					}

					if channel.ID != message.Author() &&
						tEntity.ReceiveIntersectsWith(entity.DeliverableSendChannels(message)) {
						// log.Println("Gottttt it")
					}
				}
			}
		}
	}
}
