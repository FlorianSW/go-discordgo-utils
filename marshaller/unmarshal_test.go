package marshaller_test

import (
	"github.com/bwmarrin/discordgo"
	"github.com/floriansw/go-discordgo-utils/marshaller"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Marshaller", func() {
	Describe("Unmarshal []*discordgo.ApplicationCommandInteractionDataOption", func() {
		It("unmarshals data into struct", func() {
			var v struct {
				String    string `discordgo:"someString"`
				Int       int    `discordgo:"someInt"`
				SecondInt int    `discordgo:"anotherInt"`
				Boolean   bool   `discordgo:"someBoolean"`
			}

			err := marshaller.Unmarshal([]*discordgo.ApplicationCommandInteractionDataOption{
				{Name: "someString", Value: "A_STRING", Type: discordgo.ApplicationCommandOptionString},
				{Name: "someInt", Value: 150, Type: discordgo.ApplicationCommandOptionInteger},
				{Name: "anotherInt", Value: float64(20), Type: discordgo.ApplicationCommandOptionNumber},
				{Name: "someBoolean", Value: true, Type: discordgo.ApplicationCommandOptionBoolean},
			}, &v)
			Expect(err).ToNot(HaveOccurred())

			Expect(v.String).To(Equal("A_STRING"))
			Expect(v.Int).To(Equal(150))
			Expect(v.SecondInt).To(Equal(20))
			Expect(v.Boolean).To(BeTrue())
		})
	})
	Describe("Unmarshal []discordgo.MessageComponent", func() {
		It("unmarshals data into struct", func() {
			var v struct {
				Input        string `discordgo:"my-input"`
				AnotherInput string `discordgo:"another-input"`
			}

			err := marshaller.Unmarshal([]discordgo.MessageComponent{
				&discordgo.TextInput{
					Label:    "AN_INPUT",
					CustomID: "my-input",
					Value:    "SOME_VALUE",
				},
				&discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						&discordgo.TextInput{
							Label:    "ANOTHER_INPUT",
							CustomID: "another-input",
							Value:    "ANOTHER_VALUE",
						},
					},
				},
			}, &v)
			Expect(err).ToNot(HaveOccurred())

			Expect(v.Input).To(Equal("SOME_VALUE"))
			Expect(v.AnotherInput).To(Equal("ANOTHER_VALUE"))
		})
	})
})
