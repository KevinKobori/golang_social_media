package modelos

// AuthenticationData contém o id e o token do usuário autenticado
type AuthenticationData struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
