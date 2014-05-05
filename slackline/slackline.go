package main

import (
  "appengine"
  "appengine/urlfetch"
  "bytes"
  "encoding/json"
  "errors"
  "github.com/codegangsta/martini"
  "io"
  "io/ioutil"
  "net/http"
  "os"
  "strings"
)

const postMessageURL = "/services/hooks/incoming-webhook?token="

type slackMessage struct {
  Channel  string `json:"channel"`
  Username string `json:"username"`
  Text     string `json:"text"`
  Mentions int    `json:"link_names"`
}

func AppEngine(c martini.Context, r *http.Request) {
  c.Map(appengine.NewContext(r))
}

func (s slackMessage) payload(c appengine.Context) io.Reader {
  content := []byte("payload=")
  json, _ := json.Marshal(s)
  content = append(content, json...)

  c.Debugf("payload: %v\n", string(content[:]))

  return bytes.NewReader(content)
}

func (s slackMessage) sendTo(domain, token string, c appengine.Context) (err error) {
  payload := s.payload(c)

  client := urlfetch.Client(c)

  res, err := client.Post(
    "https://"+domain+postMessageURL+token,
    "application/x-www-form-urlencoded",
    payload,
  )

  if res.StatusCode != 200 {
    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)
    return errors.New(res.Status + " - " + string(body))
  }

  return
}

func init() {
  m := martini.Classic()

  m.Use(AppEngine)

  m.Get("/", func(c appengine.Context) (int, string) {
    return 200, "OK"
  })

  m.Get("/test", func(c appengine.Context) (int, string) {
    return 200, "OK"
  })

  m.Post("/bridge", func(c appengine.Context, res http.ResponseWriter, req *http.Request) {

    // Hack to fix issue with mentions sending user id not user actual name
    // Add any userid / username to be replaced
    r := strings.NewReplacer(
      "U0297UAP2", "csullivan",
      "U0298L117", "drebelo",
    )

    username := req.PostFormValue("user_name")
    text := req.PostFormValue("text")

    text = r.Replace(text)

    if username == "slackbot" {
      // Avoid infinite loop
      return
    }

    msg := slackMessage{
      Username: username,
      Text:     text,
      Mentions: 1,
    }

    domain := req.URL.Query().Get("domain")
    token := req.URL.Query().Get("token")

    if os.Getenv("DEBUG_BRIDGE") == domain {
      c.Infof("Request: %v\n", req.PostForm)
      c.Infof("Message: %v\n", msg)
    }

    err := msg.sendTo(domain, token, c)

    if err != nil {
      c.Infof("Error: %s\n", err.Error())
      res.WriteHeader(500)
    }
  })

  http.Handle("/", m)
}
