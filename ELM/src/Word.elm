module Word exposing (..)

-- Press a button to generate a random number between 1 and 6.
--
-- Read how it works:
--   https://guide.elm-lang.org/effects/random.html
--

import Browser
import Html exposing (..)
import Html.Events exposing (..)
import Http
import Random
import List.Extra exposing (getAt)



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
  { word : String
  , text : String
  }


init : () -> ((Model, Cmd Msg))
init _ =
  ( Model "" ""
  , Http.get
      { url = "https://raw.githubusercontent.com/gaiiouch/ELP_gr22/debut_elm/ELM/thousand_words_things_explainer.txt"
      , expect = Http.expectString GotText
      }
  )



-- UPDATE


type Msg
  = GotText (Result Http.Error String)
  | NewWord Int


update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
  case msg of
    GotText result ->
      case result of
        Ok fullText ->
          ( { model | text = fullText }, Random.generate NewWord (Random.int 0 (List.length (String.split " " fullText))))

        Err _ ->
          ({ model | text = "Error" }, Cmd.none)

    NewWord number ->
        ({ model | word = (getRandomString (String.split " " model.text) number) }, Cmd.none)


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
  div []
    [ h1 [] [ text model.word ]
    ]