# To-Do

## Sprints

### Current

- Order filters used in security middleware by priority (basic, jwt, no auth, ..). Currently this is orderer by order during insertion. <sub>[*EPIC Security*]<sub>
- [DONE] Add authorities validation inside Basic-Auth, oauth, etc.. <sub>[*EPIC Security*]<sub> 
- Add Oauth grant_types abstraction <sub>[*EPIC Security*]<sub>
- [DONE] Add controller selection into security manager: oauth, oauth2, next implementations...  <sub>[*EPIC Security*]<sub>
- [DONE] Move CORS, Content Type, etc.. to another middleware, so it can be enabled or disabled by the manager.
- [DONE] Refactor Target to it allows to add more fields than url and authorities. i.e. origins, enabled, methods, scopes, etc...

## Backlog

- Messaging Systems: NATs, Kafka, RabbitMQ, etc..
- GRPC abstraction
- Metrics and health-checks endpoints
- Create generic side-car, using envoy for example, to create and ensure lightweight clients.