-- Press a button to send a GET request for random quotes.
--
-- Read how it works:
--   https://guide.elm-lang.org/effects/json.html
--
module Definitions exposing (..)
import Browser
import Html exposing (..)
import Html.Attributes exposing (style)
import Html.Events exposing (..)
import Http
import Json.Decode exposing (Decoder, map2, field, int, string, list)



-- MAIN


main =
  Browser.element
    { init = init
    , update = update
    , subscriptions = subscriptions
    , view = view
    }



-- MODEL


type Model
  = Failure
  | Loading
  | Success (List Def)


type alias Def =
    { word : String
    , meanings : List Meaning
    }

type alias Meaning =
    { partOfSpeech : String 
    , definitions : List String
    }


init : () -> (Model, Cmd Msg)
init _ =
  (Loading, getWord)



-- UPDATE


type Msg
  = MorePlease
  | GotQuote (Result Http.Error (List Def))


update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
  case msg of
    MorePlease ->
      (Loading, getWord)

    GotQuote result ->
      case result of
        Ok word ->
          (Success word, Cmd.none)

        Err _ ->
          (Failure, Cmd.none)



-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions model =
  Sub.none



-- VIEW


view : Model -> Html Msg
view model =
  div []
    [ h2 [] [ text "Random Quotes" ]
    , viewQuote model
    ]


viewQuote : Model -> Html Msg
viewQuote model =
  case model of
    Failure ->
      div []
        [ text "I could not load a random quote for some reason. "
        , button [ onClick MorePlease ] [ text "Try Again!" ]
        ]

    Loading ->
      div []
        [ text "Loading..."    ]

    Success lst ->
        case lst of 
            [] -> div []
                [text "vide..."] 
            (x :: xs) ->
                case x.meanings of
                    [] -> div [] [text "vide..."] 
                    (y::ys) ->  
                        case y.definitions of
                            [] -> div [] [text "vide..."] 
                            (z::zs) -> div [] [ blockquote [] [ text x.word ]
                                , blockquote [] [ text y.partOfSpeech ]
                                , blockquote [] [ text z ]
                                ]



-- HTTP

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


getWord : Cmd Msg
getWord =
  Http.get
    { url = "https://api.dictionaryapi.dev/api/v2/entries/en/hello"
    , expect = Http.expectJson GotQuote defDecoder
    }


{-- quoteDecoder : Decoder Quote
quoteDecoder =
  Json.Decode.map Quote 
    (field "word" string) --}
