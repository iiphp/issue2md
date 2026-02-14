docker run -it --rm `
  --name go-claude-env `
  -v ${PWD}:/app `
  -v claude-config:/root/.anthropic `
  -p 8080:8080 `
  godev-claude