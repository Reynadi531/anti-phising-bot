package events

import (
	"fmt"

	"github.com/Reynadi531/anti-phsing-discord/utils"
	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Type != discordgo.MessageType(discordgo.ChannelTypeGuildText) {
		return
	}
	links := utils.ExtractURL(m.Content)

	if len(links) == 0 {
		return
	}

	for _, v := range links {
		phising, err := utils.CheckPhising(string(v))
		if err != nil {
			fmt.Println("Error when checking link: ", err)
			return
		}

		if phising {
			s.ChannelMessageDelete(m.ChannelID, m.ID)
			channel, err := s.UserChannelCreate(m.Author.ID)
			if err != nil {
				fmt.Println("Failed create DM", err)
			}

			guildInfo, err := s.Guild(m.GuildID)
			if err != nil {
				fmt.Println("Failed retrive guild info", err)
			}

			timemessagesended, _ := m.Timestamp.Parse()
			dmMessage := fmt.Sprintf("You're sending phising link at `%s` with message contains `%s`. This message contains phising was sended at <t:%d:R>. Message contains phising will be deleted", guildInfo.Name, v, timemessagesended.Unix())
			_, err = s.ChannelMessageSend(channel.ID, dmMessage)
			if err != nil {
				fmt.Println("Failed send DM", err)
			}

			break
		}
	}
}
