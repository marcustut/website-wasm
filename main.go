package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"syscall"

	"github.com/maxence-charriere/go-app/v8/pkg/app"
	"github.com/maxence-charriere/go-app/v8/pkg/cli"
	"github.com/maxence-charriere/go-app/v8/pkg/errors"
	"github.com/maxence-charriere/go-app/v8/pkg/logs"
)

type options struct {
	Port int `env:"PORT" help:"The port used to run the server."`
}

type githubOptions struct {
  Output string `cli:"o" env:"-" help:"The directory where static resources are saved."`
}

func main() {
	app.Route("/", &App{})
	app.RunWhenOnBrowser()

  ctx, cancel := cli.ContextWithSignals(
    context.Background(),
    os.Interrupt,
    syscall.SIGTERM,
  )
  defer cancel()
  defer exit()

  h := app.Handler{
		Author:          "Marcus Lee",
		BackgroundColor: "#24283b",
		ThemeColor:      "#24283b",
		Name:            "Marcus Lee",
		Title:           "Marcus Lee",
		Description:     "A personal website",
		LoadingLabel:    "Welcome to my site!",
		Styles: []string{
			"/web/tailwind.css",
		},
    Resources: app.GitHubPages("marcustut/website-wasm"),
	}

	opts := options{Port: 8001}
  cli.Register("local").
    Help("Launches a server that serves the app in a local environment.").
    Options(&opts)

  githubOpts := githubOptions{Output: "docs"}
  cli.Register("github").
    Help("Generates the required resources to run app on GitHub Pages.").
    Options(&githubOpts)
  cli.Load()

  switch cli.Load() {
  case "local":
    runLocal(ctx, &h, opts)
  case "github":
    generateGitHubPages(ctx, &h, githubOpts)
  }
}

func runLocal(ctx context.Context, h http.Handler, opts options) {
  app.Logf("%s", logs.New("starting app server").Tag("port", opts.Port))

  s := http.Server{
    Addr: fmt.Sprintf(":%v", opts.Port),
    Handler: h,
  }

  go func() {
    <-ctx.Done()
    s.Shutdown(context.Background())
  }()

  if err := s.ListenAndServe(); err != nil {
    panic(err)
  }
}

func generateGitHubPages(ctx context.Context, h *app.Handler, opts githubOptions) {
  if err := app.GenerateStaticWebsite(opts.Output, h); err != nil {
    panic(err)
  }
}

func exit() {
  err := recover()
  if err != nil {
    app.Logf("command failed: %s", errors.Newf("%v", err))
    os.Exit(-1)
  }
}
