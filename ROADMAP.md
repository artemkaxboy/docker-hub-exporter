# Roadmap

## Enhancements

- [ ] Place html templates to app/templates folder
- [ ] Try chi.Router
- [ ] Tests
- [ ] Legacy branch for easy migration from the original version
  - [ ] remove healthcheck from dockerfile
- [ ] Cache for values for not to DDOS Docker Hub

## CI

- [x] Add GitHub actions to build and push the image
  - [x] `:snapshot`
  - [x] `:latest`
  - [x] `:X.Y.Z`
  - [x] `:X.Y`
  - [x] `:X`
- [ ] Add test reports for CI
- [ ] Automated CodeReview for PRs
