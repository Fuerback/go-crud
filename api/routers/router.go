package router

// func Router(r *mux.Router, s services.UserAccountService) {
// 	r.HandleFunc("/api/useraccount/{id}", getUser(s)).Methods("GET", "OPTIONS")
// }

// func getUser(s services.UserAccountService) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		//@TODO este código está duplicado em todos os handlers. Pergunta: como podemos melhorar isso?
// 		w.Header().Set("Content-Type", "application/json")

// 		//vamos pegar o ID da URL
// 		//na definição do protocolo http, os parâmetros são enviados no formato de texto
// 		//por isso precisamos converter em int64
// 		vars := mux.Vars(r)
// 		id, err := strconv.ParseInt(vars["id"], 10, 64)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Write(formatJSONError(err.Error()))
// 			return
// 		}
// 		b, err := service.Get(id)
// 		if err != nil {
// 			w.WriteHeader(http.StatusNotFound)
// 			w.Write(formatJSONError(err.Error()))
// 			return
// 		}
// 		//vamos converter o resultado em JSON e gerar a resposta
// 		err = json.NewEncoder(w).Encode(b)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write(formatJSONError("Erro convertendo em JSON"))
// 			return
// 		}
// 	})
// }
