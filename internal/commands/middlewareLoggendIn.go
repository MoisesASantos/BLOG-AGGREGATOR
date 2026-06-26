package commands

import (
	"context"
	"github.com/MoisesASantos/BLOG-AGGREGATOR/internal/database"
)


//Aqui o middleware recebe uma função com uma certa assinatura, State, commanda e GetUserRow, essa mesma estrutura é a estrutura das funcoes dos nossos comandos, logo ele vai retornar uma outra funcao com a assinatura State e command, dentro do middleware a gente retorna essa funcao que queremos, a funcao que queremos vai executar e retornar as funcoes dos nossos comandos, mas só quando for executada

func MiddlewareLoggedIn(
	handler func(*State, Command, database.GetUserRow) error,
) func(*State, Command) error {

	return func(s *State, cmd Command) error {

		user, err := s.Db.GetUser(
			context.Background(),
			s.Data.Current_user_name,
		)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
