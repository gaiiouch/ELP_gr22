module Main exposing (..)

--
-- From a list a word, pick a random word and display its definition
-- The user have to guess which word is behind that definition
-- 


import Browser
import Html exposing (Html, text, blockquote, pre, div, input, button)
import Html.Events exposing (onInput, onClick)
import Html.Attributes exposing (placeholder, value, style)
import Http
import Random
import List.Extra exposing (getAt)
import Json.Decode exposing (Decoder, map2, field, string, list)


-- MAIN


main : Program () Model Msg
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
  , state : State
  }


type State =
   FailureText
  | FailureDef
  | Loading
  | Success


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
  ( Model "" "" "" False [] Loading
  , getText
  )


getText : Cmd Msg
getText =
  Http.get
      { url = "https://raw.githubusercontent.com/gaiiouch/ELP_gr22/main/ELM/thousand_words_things_explainer.txt"
      , expect = Http.expectString GotText
      }


-- UPDATE


type Msg
  = GotText (Result Http.Error String)
  | NewWord Int
  | GotDef (Result Http.Error (List Def))
  | Change String
  | Erase
  | Check
  | New 


update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
  case msg of
    GotText result ->
      case result of
        Ok fullText ->
          ({ model | text = fullText }, Random.generate NewWord (Random.int 0 (List.length (String.split " " fullText))))

        Err _ ->
          ({ model | state = FailureText }, Cmd.none)
    
    NewWord number ->
      let
        newWord = getRandomString (String.split " " model.text) number
        newModel = { model | answer = newWord }
      in
      (newModel, getWord newWord)
   
    GotDef result ->
      case result of
        Ok definition ->
          ({model | def = definition, state = Success }, Cmd.none)

        Err _ ->
          ({model | state = FailureDef }, Cmd.none)

    Change newInput -> ({ model | userInput = newInput }, Cmd.none)
    
    Erase -> 
        ({ model | userInput = "" }, Cmd.none)
    
    Check -> 
        ({ model | isChecked = not model.isChecked }, Cmd.none)

    New -> init ()
    


getRandomString : List String -> Int -> String
getRandomString list x =
    case (getAt x list) of
        Just a -> a
        Nothing -> "Valeur inexistante"


getWord : String -> Cmd Msg
getWord word =
  Http.get
    { url = "https://api.dictionaryapi.dev/api/v2/entries/en/" ++ word 
    , expect = Http.expectJson GotDef defDecoder
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
  case model.state of
    Loading ->
      blockquote [style "font-family" "Noto Sans, sans-serif"] [text "Loading..."]
    FailureText ->
      blockquote [style "font-family" "Noto Sans, sans-serif"] [ div [] [text "I was unable to load the text file."]]
    FailureDef ->
      blockquote [style "font-family" "Noto Sans, sans-serif"] [ div [] [text "I was unable to load the definition of the word."]]
    Success -> 
      blockquote [style "font-family" "Noto Sans, sans-serif"]
            [ viewShowAnswer model
            , div [style "font-size" "20px"] [text "Guess the word according to its definition : \n"]
            , pre [style "font-family" "Noto Sans, sans-serif", style "font-size" "15px"] (recur1 model.def)
            , input [ placeholder "Type a word", value model.userInput, onInput Change ] []
            , button [ onClick Erase ] [ text "Erase" ]
            , button [ onClick Check ] [ text "Show the answer" ]
            , button [ onClick New ] [ text "New Definition !" ]
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
    else if String.length model.userInput > 0 && model.isChecked == False then
        div [ style "color" "red" ] [ text ("Wrong answer ! The word is not " ++ model.userInput ++ ".") ]
    else if model.isChecked == True then
        div [ style "color" "black" ] [ text ("The right answer was " ++ model.answer ++ ".") ]
    else
        div [] [ text "" ]


recur1 : List Def -> List (Html Msg)
recur1 list =
    case list of 
        [] -> []
        (x :: xs) -> [ div [style "font-weight" "bold"] [text ("Meaning :\n")] ] ++ recur2(x.meanings) ++ recur1(xs)


recur2 : List Meaning -> List (Html Msg)
recur2  list =
    case list of 
        [] -> []
        (x :: xs) -> [ div [style "font-style" "italic"] [text ("  " ++ x.partOfSpeech ++ "\n") ]] ++ recur3(x.definitions) ++ recur2(xs)


recur3 : List String -> List (Html Msg)
recur3  list =
    case list of 
        [] -> [text ("\n")]
        (x :: xs) -> [text ("       - " ++ x ++ "\n")] ++ recur3(xs)