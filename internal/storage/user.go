package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"goTgExample/internal/models"
	"log"
)

type UserStorage struct {
	pool *pgxpool.Pool
}

func NewUserStorage(pool *pgxpool.Pool) *UserStorage {
	return &UserStorage{
		pool: pool,
	}
}

func (us *UserStorage) ResetContext(ctx context.Context, userId int) error {
	query := `UPDATE "user" SET chat=$1 WHERE user_id=$2`

	_, err := us.pool.Exec(ctx, query, "[]", userId)
	if err != nil {
		return fmt.Errorf("error uery execution: %s", err)
	}

	return nil
}

func (us *UserStorage) AddUser(ctx context.Context, userId int) error {
	query := `INSERT INTO "user"(user_id, chat) VALUES ($1, $2)`

	_, err := us.pool.Exec(ctx, query, userId, "[]")
	if err != nil {
		return fmt.Errorf("error qery execution: %s", err)
	}

	return nil
}

func (us *UserStorage) GetChatHistory(ctx context.Context, userId int) (*models.Chat, error) {
	query := `SELECT chat FROM "user" WHERE user_id=$1`
	log.Printf("Executing query: %s with userId=%d, chatId=%d\n", query, userId)
	row := us.pool.QueryRow(ctx, query, userId)

	var messagesJSON string
	err := row.Scan(&messagesJSON)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, fmt.Errorf("no chat found for userId=%d and chatId=%d", userId)
		}
		return nil, fmt.Errorf("unable to get chat: %w", err)
	}

	var message []models.Message
	err = json.Unmarshal([]byte(messagesJSON), &message)
	if err != nil {
		return nil, fmt.Errorf("unable to decode messages: %w", err)
	}

	chat := models.Chat{
		Messages: message,
	}

	return &chat, nil
}

func (us *UserStorage) NewMessage(ctx context.Context, userId int, message models.Message) error {
	query := `SELECT chat FROM "user" WHERE user_id=$1`
	row := us.pool.QueryRow(ctx, query, userId)

	var messagesJSON string
	if err := row.Scan(&messagesJSON); err != nil {
		return fmt.Errorf("unable to get chat: %w", err)
	}

	var messages []models.Message
	if err := json.Unmarshal([]byte(messagesJSON), &messages); err != nil {
		return fmt.Errorf("unable to decode messages: %w", err)
	}
	fmt.Println(messages)
	// Добавляем новое сообщение к списку сообщений
	messages = append(messages, message)

	// Обновляем JSON в базе данных
	newMessagesJSON, err := json.Marshal(messages)
	if err != nil {
		return fmt.Errorf("unable to encode messages: %w", err)
	}

	updateQuery := `UPDATE "user" SET chat=$1 WHERE user_id=$2`
	_, err = us.pool.Exec(ctx, updateQuery, newMessagesJSON, userId)
	if err != nil {
		return fmt.Errorf("unable to update chat: %w", err)
	}

	return nil
}
