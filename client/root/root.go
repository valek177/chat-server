package root

import (
	"context"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/valek177/chat-server/client"
	"github.com/valek177/chat-server/grpc/pkg/chat_v1"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "chat-client",
	Short: "Chat client app",
}

var createChatCmd = &cobra.Command{
	Use:   "create",
	Short: "Create chat",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.NewChatV1Client()
		if err != nil {
			log.Fatalf("unable to create client")
		}
		defer c.Close()

		chatID, err := createChat(cmd.Context(), c.C)
		if err != nil {
			log.Fatalf("failed to connect chat: %v", err)
		}

		log.Printf("was created chat with id %d", chatID)
	},
}

var deleteChatCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete chat",
}

var connectChatCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to chat",
	Run: func(cmd *cobra.Command, args []string) {
		username, err := cmd.Flags().GetString("username")
		if err != nil {
			log.Fatalf("failed to get usernames: %s\n", err.Error())
		}

		chatID, err := cmd.Flags().GetInt64("chat-id")
		if err != nil {
			log.Fatalf("failed to get chat-id: %s\n", err.Error())
		}

		c, err := client.NewChatV1Client()
		if err != nil {
			log.Fatalf("unable to create client")
		}
		defer c.Close()

		wg := sync.WaitGroup{}
		wg.Add(1)

		go func() {
			defer wg.Done()

			err = connectChat(cmd.Context(), c.C, chatID, username, 5*time.Second)
			if err != nil {
				log.Fatalf("failed to connect chat: %v", err)
			}
		}()

		wg.Wait()
	},
}

var disconnectChatCmd = &cobra.Command{
	Use:   "disconnect",
	Short: "Disconnect from chat",
}

var createUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Create new chat user",
	Run: func(cmd *cobra.Command, args []string) {
		usernamesStr, err := cmd.Flags().GetString("username")
		if err != nil {
			log.Fatalf("failed to get usernames: %s\n", err.Error())
		}

		log.Printf("user %s created\n", usernamesStr)
	},
}

var deleteUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Delete chat user",
	Run: func(cmd *cobra.Command, args []string) {
		usernamesStr, err := cmd.Flags().GetString("username")
		if err != nil {
			log.Fatalf("failed to get usernames: %s\n", err.Error())
		}

		log.Printf("user %s deleted\n", usernamesStr)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(createChatCmd, deleteChatCmd, createUserCmd, deleteUserCmd)

	deleteChatCmd.Flags().Int64("chat-id", 0, "Chat ID")

	createUserCmd.Flags().StringP("username", "u", "", "User name")
	err := createUserCmd.MarkFlagRequired("username")
	if err != nil {
		log.Fatalf("failed to mark username flag as required: %s\n", err.Error())
	}

	deleteUserCmd.Flags().StringP("username", "u", "", "User name")
	err = deleteUserCmd.MarkFlagRequired("username")
	if err != nil {
		log.Fatalf("failed to mark username flag as required: %s\n", err.Error())
	}
}

func connectChat(ctx context.Context, client chat_v1.ChatV1Client, chatID int64, username string, period time.Duration) error {
	stream, err := client.ConnectChat(ctx, &chat_v1.ConnectChatRequest{
		ChatId:   chatID,
		Username: username,
	})
	if err != nil {
		return err
	}

	go func() {
		for {
			message, errRecv := stream.Recv()
			if errRecv == io.EOF {
				return
			}
			if errRecv != nil {
				log.Println("failed to receive message from stream: ", errRecv)
				return
			}

			log.Printf("[%v] - [from: %s]: %s\n",
				color.YellowString(message.GetCreatedAt().AsTime().Format(time.RFC3339)),
				color.BlueString(message.GetFrom()),
				message.GetText(),
			)
		}
	}()

	for {
		// Ниже пример того, как можно считывать сообщения из консоли
		// в демонстрационных целях будем засылать в чат рандомный текст раз в 5 секунд
		//scanner := bufio.NewScanner(os.Stdin)
		//var lines strings.Builder
		//
		//for {
		//	scanner.Scan()
		//	line := scanner.Text()
		//	if len(line) == 0 {
		//		break
		//	}
		//
		//	lines.WriteString(line)
		//	lines.WriteString("\n")
		//}
		//
		//err = scanner.Err()
		//if err != nil {
		//	log.Println("failed to scan message: ", err)
		//}

		time.Sleep(period)

		text := gofakeit.Word()

		_, err = client.SendMessage(ctx, &chat_v1.SendMessageRequest{
			ChatId: chatID,
			Message: &chat_v1.Message{
				From:      username,
				Text:      text,
				CreatedAt: timestamppb.Now(),
			},
		})
		if err != nil {
			log.Println("failed to send message: ", err)
			return err
		}
	}
}

func createChat(ctx context.Context, client chat_v1.ChatV1Client) (int64, error) {
	res, err := client.CreateChat(ctx, &chat_v1.CreateChatRequest{
		Name: "newchat", UserIds: []int64{1, 2},
	})
	if err != nil {
		return 0, err
	}

	return res.GetId(), nil
}
