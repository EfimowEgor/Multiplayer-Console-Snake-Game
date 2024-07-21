package components

import "net"

type Lobby struct {
	LobbyID string `json:"lobbyID"`
	Cap     int    `json:"cap"`
}

func CreateLobby(conn net.Conn, lobbyID string, cap int) *Lobby {
	return &Lobby{
		LobbyID: lobbyID,
		Cap:     cap,
	}
}
