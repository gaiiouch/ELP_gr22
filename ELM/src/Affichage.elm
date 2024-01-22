module Affichage exposing (..)

-- Press a button to generate a random number between 1 and 6.
--
-- Read how it works:
--   https://guide.elm-lang.org/effects/random.html
--

import Browser
import Html exposing (Html, text, pre, div, input, button)
import Html.Events exposing (onInput, onClick)
import Html.Attributes exposing (placeholder, value, style)
import Http
import Random
import List.Extra exposing (getAt)
import Json.Decode exposing (Decoder, map2, field, string, list)



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
  , def : List Def
  }

type alias Def =
    { word : String
    , meanings : List Meaning
    }

type alias Meaning =
    { partOfSpeech : String 
    , definitions : List String
    }


init : () -> ((Model, Cmd Msg))
init _ =
  ( Model "" "" "" False []
  , Http.get
      { url = "https://raw.githubusercontent.com/gaiiouch/ELP_gr22/main/ELM/thousand_words_things_explainer.txt"
      , expect = Http.expectString GotText
      }
  )



-- UPDATE


type Msg
  = GotText (Result Http.Error String)
  | NewWord Int
  | GotDef (Result Http.Error (List Def))
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
        Ok definition ->
          ({model | def = definition }, Cmd.none)

        Err _ ->
          ({model | def = []}, Cmd.none)
      
  --  NewWord number ->
    --    ({ model | answer = (getRandomString (String.split " " model.text) number) }, getWord model.answer)
    
    NewWord number ->
      let
        newWord = getRandomString (String.split " " model.text) number
        newModel = { model | answer = newWord, userInput = "", isChecked = False }
      in
      (newModel, getWord newWord)

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

getWord : String -> Cmd Msg
getWord word =
  Http.get
    { url = "https://api.dictionaryapi.dev/api/v2/entries/en/" ++ word 
    , expect = Http.expectJson GotDef  defDecoder
    }

defDecoder : Decoder (List Def)
defDecoder =
    list listDecodage


listDecodage : Decoder Def
listDecodage =
    map2 Def
        (field "word" string)
        (field "meanings" (list meaningDecodage))


meaningDecodage : Decoder Meaning
meaningDecodage =
    map2 Meaning
        (field "partOfSpeech" string)
        (field "definitions" (list (field "definition" string)))


-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions model =
  Sub.none


-- VIEW


view : Model -> Html Msg
view model =
    if model.text == "Error" then
        div [style "font-family" "Noto Sans, sans-serif"] [text "I was unable to load the text file."]
    else if model.def == [] then
      div [style "font-family" "Noto Sans, sans-serif"] [text ("I was unable to load the definition of the word " ++ model.answer)]
    else 
        div [style "font-family" "Noto Sans, sans-serif"]
            [ viewShowAnswer model
            , div [style "font-size" "20px"] [text "Guess the word according to its definition : \n"]
            , pre [style "font-family" "Noto Sans, sans-serif"] (recur1 model.def)
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
    else if String.length model.userInput > 0 then
        div [ style "color" "red" ] [ text ("Wrong answer ! The word is not " ++ model.userInput ++ ".") ]
    else
        div [] [ text "" ]


recur1 : List Def -> List (Html Msg)
recur1 list =
    case list of 
        [] -> []
        (x :: xs) -> [ text ("Meaning :\n") ] ++ recur2(x.meanings) ++ recur1(xs)


recur2 : List Meaning -> List (Html Msg)
recur2  list =
    case list of 
        [] -> []
        (x :: xs) -> [ text ("  " ++ x.partOfSpeech ++ "\n") ] ++ recur3(x.definitions) ++ recur2(xs)


recur3 : List String -> List (Html Msg)
recur3  list =
    case list of 
        [] -> []
        (x :: xs) -> [text ("       - " ++ x ++ "\n")] ++ recur3(xs)