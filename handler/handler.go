package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	control "secondAssignment/controller"
	"strconv"
)

type CollectionHandler struct {
	Collection *control.Collection
}

func NewCollectionHandler(collection *control.Collection) *CollectionHandler {
	return &CollectionHandler{Collection: collection}
}

func (h *CollectionHandler) ItemPushHandler(w http.ResponseWriter, r *http.Request) {
	if !control.IsAuthorized {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	var item control.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Endpoint hit: itemPushHandler")
		return
	}
	h.Collection.ItemPush(item.ItemName, item.Price, item.Rating)
	w.WriteHeader(http.StatusCreated)
	fmt.Println("Endpoint hit: itemPushHandler")
}

func (h *CollectionHandler) GetItemByIdHandler(w http.ResponseWriter, r *http.Request) {
	if !control.IsAuthorized {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["itemId"])
	//fmt.Println(key)
	for _, item := range h.Collection.Items {
		if item.Id == key {
			//fmt.Println(item)
			json.NewEncoder(w).Encode(item)
			w.WriteHeader(http.StatusOK)
			fmt.Println("Endpoint hit: getItemByIdHandler")
			return
		}
	}
	w.WriteHeader(http.StatusNoContent)
	fmt.Println("Endpoint hit: getItemByIdHandler")
}

func (h *CollectionHandler) GetItemsHandler(w http.ResponseWriter, r *http.Request) {
	if !control.IsAuthorized {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println("Endpoint hit: getItemsHandler")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(h.Collection.Items)
	fmt.Println("Endpoint hit: getItemsHandler")
}

func (h *CollectionHandler) SearchItemsByNameHandler(w http.ResponseWriter, r *http.Request) {
	if !control.IsAuthorized {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println("Endpoint hit: searchItemsByNameHandler")
		return
	}
	vars := mux.Vars(r)
	name := vars["itemName"]
	items := h.Collection.SearchItemsByName(name)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
	fmt.Println("Endpoint hit: searchItemsByNameHandler")
}

func (h *CollectionHandler) FilterItemsByPriceHandler(w http.ResponseWriter, r *http.Request) {
	if !control.IsAuthorized {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println("Endpoint hit: filterItemsByPriceHandler")
		return
	}
	var item control.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Endpoint hit: filterItemsByPriceHandler")
		return
	}
	items := h.Collection.FilterItemsByPrice(item.Price)
	json.NewEncoder(w).Encode(items)
	fmt.Println("Endpoint hit: filterItemsByPriceHandler")
}

func (h *CollectionHandler) FilterItemsByRatingHandler(w http.ResponseWriter, r *http.Request) {
	if !control.IsAuthorized {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println("Endpoint hit: filterItemsByRatingHandler")
		return
	}
	var item control.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Endpoint hit: filterItemsByRatingHandler")
		return
	}
	items := h.Collection.FilterItemsByRating(item.Rating)
	json.NewEncoder(w).Encode(items)
	fmt.Println("Endpoint hit: filterItemsByRatingHandler")
}

func (h *CollectionHandler) SetRatingHandler(w http.ResponseWriter, r *http.Request) {
	if !control.IsAuthorized {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println("Endpoint hit: setRatingHandler")
		return
	}
	var item control.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Endpoint hit: setRatingHandler")
		return
	}
	h.Collection.SetRating(item.Rating, item.Id)
	//json.NewEncoder(w).Encode(h.Collection.items)
	w.WriteHeader(http.StatusNoContent)
	fmt.Println("Endpoint hit: setRatingHandler")
}

func (h *CollectionHandler) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var user control.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Endpoint hit: signUpHandler")
		return
	}
	h.Collection.SignUp(user.Username, user.Password)
	//h.Collection.UserSaveData()
	w.WriteHeader(http.StatusCreated)
	fmt.Println("Endpoint hit: signUpHandler")
}

func (h *CollectionHandler) SignInHandler(w http.ResponseWriter, r *http.Request) {
	var user control.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Endpoint hit: signInHandler")
		return
	}
	if h.Collection.SignIn(user.Username, user.Password) {
		control.IsAuthorized = true
		w.WriteHeader(http.StatusOK)
		fmt.Println("Endpoint hit: signInHandler")
	} else {
		http.Error(w, "invalid username or password", http.StatusUnauthorized)
		fmt.Println("Endpoint hit: signInHandler")
	}
}

func (h *CollectionHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if !control.IsAuthorized {
		fmt.Fprintf(w, "You were not authorized yet.")
		fmt.Println("Endpoint hit: logoutHandler")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	control.IsAuthorized = false
	fmt.Println("Endpoint hit: logoutHandler")
	w.WriteHeader(http.StatusOK)
}

func (h *CollectionHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	if !control.IsAuthorized {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println("Endpoint hit: getUserHandler")
		return
	}
	var user control.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Endpoint hit: getUserHandler")
		return
	}
	users := h.Collection.GetUser(user.Username)
	json.NewEncoder(w).Encode(users)
	fmt.Println("Endpoint hit: getUserHandler")
}
