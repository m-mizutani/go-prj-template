package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/m-mizutani/go-prj-template/pkg/model"
	"github.com/m-mizutani/go-prj-template/pkg/usecase"
	"github.com/m-mizutani/goerr"
)

func Listen(uc *usecase.Usecase) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var req model.CreateUserRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			ctx := model.NewContext(model.WithBaseCtx(r.Context()))
			resp, err := uc.CreateUser(ctx, &req)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(resp.ID))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not found"))
		}
	})

	fmt.Println("Starting server...")
	if err := http.ListenAndServe("127.0.0.1:8080", mux); err != nil {
		return goerr.Wrap(err)
	}

	return nil
}
