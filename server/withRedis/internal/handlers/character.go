package handlers

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	"withredis/internal/models"
	"withredis/internal/utils"

	"github.com/redis/go-redis/v9"
)

type Handler struct {
	redis *redis.Client
	url   string
}

func NewCharacterHandler(redis *redis.Client, url string) *Handler {
	return &Handler{redis: redis, url: url}
}

func (h *Handler) GetCharacter(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	values := strings.Split(r.URL.Path, "/")

	id := values[len(values)-1]

	var character models.Character

	val, err := h.GetCharacterFromRedis(ctx, id)

	if err != nil {
		fmt.Println("err in redis req", err)
		_ = val
	} else {
		character = val
		utils.WriteJSON(w, http.StatusOK, character)
		return
	}

	time.Sleep(time.Second * 2)

	res, err := http.Get(h.url + id)

	if err != nil || res.StatusCode != http.StatusOK {
		utils.WriteJSON(w, http.StatusNotFound, map[string]string{"error": "not found"})
		return
	}

	if err := json.NewDecoder(res.Body).Decode(&character); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "bad request"})
		return
	}

	err = h.SetCharacterToRedis(ctx, id, character)

	if err != nil {
		fmt.Println("err during redis set", err)
	}

	utils.WriteJSON(w, http.StatusOK, character)

}

func (h *Handler) GetCharacterFromRedis(ctx context.Context, id string) (models.Character, error) {
	cmd := h.redis.Get(ctx, id)

	cmdb, err := cmd.Bytes()
	if err != nil {
		return models.Character{}, err
	}

	b := bytes.NewReader(cmdb)

	var res models.Character

	if err := gob.NewDecoder(b).Decode(&res); err != nil {
		return models.Character{}, err
	}

	return res, nil
}

func (h *Handler) SetCharacterToRedis(ctx context.Context, id string, char models.Character) error {
	var b bytes.Buffer

	if err := gob.NewEncoder(&b).Encode(char); err != nil {
		return err
	}

	return h.redis.Set(ctx, id, b.Bytes(), 25*time.Second).Err()
}
