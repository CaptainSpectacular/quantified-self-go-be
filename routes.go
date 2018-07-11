package main

import ("net/http"
        "github.com/gorilla/mux")

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route 

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }

    return router
}

var routes = Routes{
    Route{
        "FoodsIndex",
        "GET",
        "/foods",
        FoodIndex,
    },

    Route{
        "FoodsShow",
        "GET",
        "/foods/{id}",
        FoodShow,
    },

    Route{
        "FoodsCreate",
        "POST",
        "/foods",
        FoodCreate,
    },

    Route{
        "FoodsDelete",
        "DELETE",
        "/foods/{id}",
        FoodDelete,
    },

    Route{
        "FoodsUpdate",
        "PUT",
        "/foods/{id}",
        FoodUpdate,
    },

    Route{
        "FoodsUpdate",
        "PATCH",
        "/foods/{id}",
        FoodUpdate,
    },

    Route{
        "MealsIndex",
        "GET",
        "/meals",
        MealIndex,
    },

    Route{
        "MealsShow",
        "GET",
        "/meals/{id}/foods",
        MealShow,
    },

    Route{
        "MealsUpdate",
        "POST",
        "/meals/{mid}/foods/{fid}",
        MealUpdate,
    },

    Route{
        "MealsDelete",
        "DELETE",
        "/meals/{mid}/foods/{fid}",
        MealDelete,
    },
}
