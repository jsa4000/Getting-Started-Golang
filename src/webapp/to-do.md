# To-Do

## Sprints

### Current

- Order filters used in security middleware by priority (basic, jwt, no auth, ..). Currently this is orderer by order during insertion. <sub>[*EPIC Security*]<sub>
- Add authorities validation inside Basic-Auth, oauth, etc.. This must can be automatic, by using context to add authorities to next middleware.  <sub>[*EPIC Security*]<sub>
- Add Oauth grant_types abstraction  <sub>[*EPIC Security*]<sub>
- Add controller selection into security manager: oauth, oauth2, next implementations...  <sub>[*EPIC Security*]<sub>

## Backlog

- Messaging Systems: NATs, Kafka, RabbitMQ, etc..
- GRPC abstraction
- Metrics and health-checks endpoints
- Create generic sire-car, using envoy for example, to create and ensure lightweight clients.