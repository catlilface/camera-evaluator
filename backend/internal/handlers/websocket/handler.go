package websocket

import (
	"fmt"
	"net/http"
	wsModels "photo-upload-service/internal/pkg/api/ws"
	ws "photo-upload-service/pkg/websocket"

	httpUtils "photo-upload-service/pkg/utils/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type WSHandler struct {
	manager *ws.Manager
}

func NewWSHandler(manager *ws.Manager) *WSHandler {
	return &WSHandler{manager: manager}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *WSHandler) ConnectWebSocket(c *gin.Context, params wsModels.ConnectWebSocketParams) {
	channelID := params.ChannelID
	if channelID == uuid.Nil {
		httpUtils.AbortWithStatus(c, http.StatusBadRequest, fmt.Errorf("channel_id is required"))
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	h.manager.AddClient(channelID.String(), conn)

	go func() {
		defer func() {
			h.manager.RemoveClient(channelID.String(), conn)
			conn.Close()
		}()

		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				break
			}
		}
	}()
}
