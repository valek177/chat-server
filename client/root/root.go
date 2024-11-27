package root

import (
	"context"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/valek177/auth/grpc/pkg/auth_v1"
	"github.com/valek177/chat-server/client"
	"github.com/valek177/chat-server/grpc/pkg/chat_v1"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "chat-client",
	Short: "Chat client app",
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create object",
}

var createChatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Create chat",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.NewChatV1Client()
		if err != nil {
			log.Fatalf("unable to create client for connect")
		}
		defer c.Close()

		chatID, err := createChat(cmd.Context(), c.C)
		if err != nil {
			log.Fatalf("failed to connect chat: %v", err)
		}

		log.Printf("was created chat with id %d", chatID)
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete object",
}

var deleteChatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Delete chat",

	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.NewChatV1Client()
		if err != nil {
			log.Fatalf("unable to create client for connect")
		}
		defer c.Close()

		// chatID, err := createChat(cmd.Context(), c.C)
		// if err != nil {
		// 	log.Fatalf("failed to connect chat: %v", err)
		// }

		// log.Printf("chat with id %d was deleted", chatID)
	},
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
	Run: func(cmd *cobra.Command, args []string) {
		// usernamesStr, err := cmd.Flags().GetString("username")
		// if err != nil {
		// 	log.Fatalf("failed to get usernames: %s\n", err.Error())
		// }

		// log.Printf("user %s created\n", usernamesStr)
	},
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

var sendMessageCmd = &cobra.Command{
	Use:   "send",
	Short: "Send message",
	Run: func(cmd *cobra.Command, args []string) {
		chatID, err := cmd.Flags().GetInt64("chat-id")
		if err != nil {
			log.Fatalf("failed to get flag chat-id: %s\n", err.Error())
		}

		message, err := cmd.Flags().GetString("message")
		if err != nil {
			log.Fatalf("failed to get flag message: %s\n", err.Error())
		}

		from, err := cmd.Flags().GetString("from")
		if err != nil {
			log.Fatalf("failed to get flag from: %s\n", err.Error())
		}

		c, err := client.NewChatV1Client()
		if err != nil {
			log.Fatalf("unable to create client")
		}
		defer c.Close()

		// login to auth service; get token and use it for send message

		err = sendMessage(cmd.Context(), c.C, from, chatID, message)
		if err != nil {
			log.Fatalf("failed to send message: %v", err)
		}
		log.Printf("was sended message")
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
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(connectChatCmd)
	rootCmd.AddCommand(disconnectChatCmd)
	rootCmd.AddCommand(sendMessageCmd)

	createCmd.AddCommand(createChatCmd)
	createCmd.AddCommand(createUserCmd)

	deleteCmd.AddCommand(deleteChatCmd)
	deleteCmd.AddCommand(deleteUserCmd)

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

	connectChatCmd.Flags().StringP("username", "u", "", "User name")
	err = connectChatCmd.MarkFlagRequired("username")
	if err != nil {
		log.Fatalf("failed to mark username flag as required: %s\n", err.Error())
	}

	connectChatCmd.Flags().Int64("chat-id", 0, "Chat ID")
	err = connectChatCmd.MarkFlagRequired("chat-id")
	if err != nil {
		log.Fatalf("failed to mark chat-id flag as required: %s\n", err.Error())
	}

	sendMessageCmd.Flags().Int64P("chat-id", "i", 0, "Chat ID")
	err = sendMessageCmd.MarkFlagRequired("chat-id")
	if err != nil {
		log.Fatalf("failed to mark chat-id flag as required: %s\n", err.Error())
	}
	sendMessageCmd.Flags().StringP("message", "m", "", "Message text")
	err = sendMessageCmd.MarkFlagRequired("message")
	if err != nil {
		log.Fatalf("failed to mark message flag as required: %s\n", err.Error())
	}
	sendMessageCmd.Flags().StringP("from", "f", "", "From")
	err = sendMessageCmd.MarkFlagRequired("from")
	if err != nil {
		log.Fatalf("failed to mark from flag as required: %s\n", err.Error())
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
	return nil
}

func createChat(ctx context.Context, client chat_v1.ChatV1Client) (int64, error) {
	res, err := client.CreateChat(ctx, &chat_v1.CreateChatRequest{
		Name: "newchat", UserIds: []int64{1, 2, 3, 4},
	})
	if err != nil {
		return 0, err
	}

	return res.GetId(), nil
}

func sendMessage(ctx context.Context, client chat_v1.ChatV1Client, from string, chatID int64, message string) error {
	// for {
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
	// }
	// text := gofakeit.Word()
	authClient, err := auth_v1.AuthV1Client
	if err != nil {
		log.Fatalf("unable to create auth client for connect")
	}
	defer authClient.Close()
	token, err := authClient.Login()

	_, err = client.SendMessage(ctx, &chat_v1.SendMessageRequest{
		ChatId: chatID,
		Message: &chat_v1.Message{
			From:      from,
			Text:      message,
			CreatedAt: timestamppb.Now(),
		},
	})
	if err != nil {
		log.Println("failed to send message: ", err)
		return err
	}

	return nil
}
