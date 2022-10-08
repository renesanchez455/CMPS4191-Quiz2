-- CMPS4191 - Quiz #2
-- Rene Sanchez - 2018118383

module Star exposing (main)

import Html exposing (..)
import Html.Attributes exposing (class, src)
import Html.Events exposing (onClick)
import Browser

type Msg = 
    Like | Unlike

initialModel : { liked : Bool }
initialModel = 
    {
        liked = False
    }
viewStar : { liked : Bool } -> Html Msg
viewStar model = 
    let
        buttonType = 
            if model.liked then "star" else "star_border"
        msg =
                if model.liked then Unlike else Like
    in
    
    div [ class "star-button" ]
        [ span
            [ class "material-icons md-48", onClick msg ]
            [ text buttonType ]
        ]


view : {liked : Bool } -> Html Msg
view model =
    div []
    [
        div [ class "header" ]
        [ h1 [] [text "Star" ] ]
        ,div [ class "content-flow" ]
        [
            viewStar model
        ]
    ]

update : Msg -> {liked : Bool } -> {liked : Bool }
update msg model = 
    case msg of
        Like ->
            { model | liked = True }
        Unlike ->
            { model | liked = False }



main : Program () {liked : Bool } Msg
main = 
    Browser.sandbox
    {
        init = initialModel
        ,view = view
        ,update = update
    }