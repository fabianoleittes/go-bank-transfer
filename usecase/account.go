package usecase

import (
	"strings"
	"time"

	"github.com/gsabadini/go-bank-transfer/domain"
)

//transferOutput armazena a estrutura de dados de retorno do caso de uso
type accountOutput struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CPF       string    `json:"cpf"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

//transferOutput armazena a estrutura de dados de retorno do caso de uso
type accountBalanceOutput struct {
	Balance float64 `json:"balance"`
}

//Account armazena as dependências para os casos de uso de Account
type Account struct {
	repository domain.AccountRepository
}

//NewAccount constrói um Account com suas dependências
func NewAccount(repository domain.AccountRepository) Account {
	return Account{repository: repository}
}

//Store cria uma nova Account
func (a Account) Store(name, CPF string, balance float64) (accountOutput, error) {
	var account = domain.NewAccount(name, a.cleanCPF(CPF), balance)

	account, err := a.repository.Store(account)
	if err != nil {
		return accountOutput{}, err
	}

	return accountOutput{
		ID:        account.ID,
		Name:      account.Name,
		CPF:       account.CPF,
		Balance:   account.Balance,
		CreatedAt: account.CreatedAt,
	}, nil
}

//FindAll retorna uma lista de Accounts
func (a Account) FindAll() ([]accountOutput, error) {
	var output = make([]accountOutput, 0)

	accounts, err := a.repository.FindAll()
	if err != nil {
		return output, err
	}

	for _, account := range accounts {
		var account = accountOutput{
			ID:        account.ID,
			Name:      account.Name,
			CPF:       account.CPF,
			Balance:   account.Balance,
			CreatedAt: account.CreatedAt,
		}

		output = append(output, account)
	}

	return output, nil
}

//FindBalance retorna o saldo de uma Account
func (a Account) FindBalance(ID string) (accountBalanceOutput, error) {
	account, err := a.repository.FindBalance(ID)
	if err != nil {
		return accountBalanceOutput{}, err
	}

	return accountBalanceOutput{
		Balance: account.Balance,
	}, nil
}

func (a Account) cleanCPF(cpf string) string {
	return strings.Replace(strings.Replace(cpf, ".", "", -1), "-", "", -1)
}
