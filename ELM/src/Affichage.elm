module Affichage exposing (..)

-- Press a button to generate a random number between 1 and 6.
--
-- Read how it works:
--   https://guide.elm-lang.org/effects/random.html
--

import Browser
import Html exposing (..)
import Html.Events exposing (..)
import Html.Attributes exposing (..)
import Http
import Random
import List.Extra exposing (getAt)
import Json.Decode exposing (Decoder, field, int, string, at, decodeString)



-- MAIN


main =
  Browser.element
    { init = init
    , update = update
    , subscriptions = subscriptions
    , view = view
    }



-- MODEL

type alias Model =
  { userInput : String
  , answer : String
  , text : String
  , isChecked : Bool
  }

type alias Def =
  { word : String
  -- , partOfSpeech : String
  -- , def : String
  }


init : () -> ((Model, Cmd Msg))
init _ =
  ( Model "" "" "" False
  , Http.get
      { url = "https://raw.githubusercontent.com/gaiiouch/ELP_gr22/main/ELM/thousand_words_things_explainer.txt"
      , expect = Http.expectString GotText
      }
  )



-- UPDATE


type Msg
  = GotText (Result Http.Error String)
  | GotDef (Result Http.Error String)
  | NewWord Int
  | Change String
  | Erase
  | Check


update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
  case msg of
    GotText result ->
      case result of
        Ok fullText ->
          ({ model | text = fullText }, Random.generate NewWord (Random.int 0 (List.length (String.split " " fullText))))

        Err _ ->
          ({ model | text = "Error" }, Cmd.none)

    GotDef result ->
      case result of
        Ok fullText ->
          ({ model | text = fullText }, Random.generate NewWord (Random.int 0 (List.length (String.split " " fullText))))

        Err _ ->
          ({ model | text = "Error" }, Cmd.none)

    NewWord number ->
        ({ model | answer = (getRandomString (String.split " " model.text) number) }, Cmd.none)

    Change newInput -> ({ model | userInput = newInput }, Cmd.none)
    
    Erase -> 
        ({ model | userInput = "" }, Cmd.none)
    
    Check -> 
        ({ model | isChecked = not model.isChecked }, Cmd.none)


getRandomString : List String -> Int -> String
getRandomString list x =
    case (getAt x list) of
        Just a -> a
        Nothing -> "Valeur inexistante"


-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions model =
  Sub.none



-- VIEW


view : Model -> Html Msg
view model =
    if model.text == "Error" then
        div [style "font-family" "Noto Sans, sans-serif"] [text "I was unable to load the text file."]
    else 
        div [style "font-family" "Noto Sans, sans-serif"]
            [ viewShowAnswer model
            , div [style "font-size" "20px"] [text "Guess the word according to its definition :"]
            , input [ placeholder "Type a word", value model.userInput, onInput Change ] []
            , button [ onClick Erase ] [ text "Erase" ]
            , button [ onClick Check ] [ text "Show the answer" ]
            , viewValidation model
            --h1 [] [ text model.answer ]
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
    else if String.length model.userInput > 0 then
        div [ style "color" "red" ] [ text ("Wrong answer ! The word is not " ++ model.userInput ++ ".") ]
    else
        div [ ] [ text "" ]

-- HTTP

getDefinition : Cmd Msg
getDefinition =
  Http.get
    { url = "https://api.dictionaryapi.dev/api/v2/entries/en/" ++ model.answer
    , expect = Http.expectJson GotDef defDecoder
    }

defDecoder : Decoder Def
defDecoder =
    decodeString Def (at ["meanings","definitions","definition"] string) 
    -- decodeString (at ["meanings","partOfSpeech"] string)
    -- decodeString (at ["word"] string)