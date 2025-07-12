package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)
 

func printCommandEvents(analyticsChannel <- chan * slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main()  {
	 os.Setenv("SLACK_BOT_TOKEN","xoxb-9210208325680-9189408971156-m1hfqG0T0RWBftd7ZdrMkyVz")
	 os.Setenv("SLACK_APP_TOKEN","xapp-1-A095F8U9N7Q-9201466128177-5a403bc31e1ec690058f19340af8e5879a00e938c21c33e23da0b45aa9bbbac6")

	 bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"),os.Getenv("SLACK_APP_TOKEN"))

	 go printCommandEvents(bot.CommandEvents())

	 bot.Command("my yob is <year>",&slacker.CommandDefinition{
		Description: "yob calculator",
		Examples:[]string{ "my yob is 2020"},
		Handler: func(botCtx slacker.BotContext,request slacker.Request,resopnse slacker.ResponseWriter){
			year := request.Param("year")
			yob,err :=strconv.Atoi(year)
			if err != nil {
				fmt.Println(err)
			}
			age := 2025 - yob
			r := fmt.Sprintf("age is %d",age)
			resopnse.Reply(r)
		},
	 })

	 ctx,cancel := context.WithCancel(context.Background())
	 defer cancel()

	 err := bot.Listen(ctx)
	 if err != nil {
		log.Fatal(err)
	 }
}