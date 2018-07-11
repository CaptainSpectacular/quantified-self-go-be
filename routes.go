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
        "/api/v1/foods",
        FoodIndex,
    },

    Route{
        "FoodsShow",
        "GET",
        "/api/v1/foods/{id}",
        FoodShow,
    },

    Route{
        "FoodsCreate",
        "POST",
        "/api/v1/foods",
        FoodCreate,
    },

    Route{
        "FoodsDelete",
        "DELETE",
        "/api/v1/foods/{id}",
        FoodDelete,
    },

    Route{
        "FoodsUpdate",
        "PUT",
        "/api/v1/foods/{id}",
        FoodUpdate,
    },

    Route{
        "FoodsUpdate",
        "PATCH",
        "/api/v1/foods/{id}",
        FoodUpdate,
    },

    Route{
        "MealsIndex",
        "GET",
        "/api/v1/meals",
        MealIndex,
    },

    Route{
        "MealsShow",
        "GET",
        "/api/v1/meals/{id}/foods",
        MealShow,
    },

    Route{
        "MealsUpdate",
        "POST",
        "/api/v1/meals/{mid}/foods/{fid}",
        MealUpdate,
    },

    Route{
        "MealsDelete",
        "DELETE",
        "/api/v1/meals/{mid}/foods/{fid}",
        MealDelete,
    },
}
