module Bouton exposing (..)

import Browser
import Html exposing (Html, button, div, text, input)
import Html.Attributes exposing (..)
import Html.Events exposing (onClick, onInput)


-- MAIN

main =
  Browser.sandbox { init = init, view = view, update = update }


-- MODEL

type alias Model = { userInput : String, answer : String, isChecked : Bool}

init : Model
init = Model "" "hello" False


-- UPDATE

type Msg = Change String | Erase | Check

update : Msg -> Model -> Model
update msg model =
    case msg of
        Change newInput -> { model | userInput = newInput }
        Erase -> { model | userInput = "" }
        Check -> { model | isChecked = not model.isChecked }


-- VIEW

view : Model -> Html Msg
view model = 
    div [style "font-family" "Noto Sans, sans-serif" ]
    [ viewShowAnswer model
    , div [style "font-size" "20px"] [text "Guess the word according to its definition :"]
    , input [ placeholder "Type a word", value model.userInput, onInput Change ] []
    , button [ onClick Erase ] [ text "Erase" ]
    , button [ onClick Check ] [ text "Show the answer" ]
    , viewValidation model
    ]

viewShowAnswer : Model -> Html msg
viewShowAnswer model =
    if model.userInput == model.answer || model.isChecked == True then
        div [style "font-size" "50px", style "font-weight" "bold"] [text model.answer]
    else
        div [style "font-size" "50px", style "font-weight" "bold"] [text "Guess it !"]

viewValidation : Model -> Html msg
viewValidation model =
    if model.userInput == model.answer then
        div [ style "color" "green" ] [ text ("Correct answer ! The word was indeed " ++ model.answer ++ " !") ]
    else
        div [ ] [ text "" ]
        
