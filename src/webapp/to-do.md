# To-Do

## Sprints

### Current

- Add support for distrbuted cache: Redis, etcd...
- Messaging Systems: NATs, Kafka, RabbitMQ, etc..
- [DONE] Add Roles to Users, Model and endpoints.
- Support for bootstraping

### 2019/03/07

- [DONE] Order filters used in security middleware by priority (basic, jwt, no auth, ..). Currently this is orderer by order during insertion. <sub>[*EPIC Security*]<sub>
- [DONE] Add authorities validation inside Basic-Auth, oauth, etc.. <sub>[*EPIC Security*]<sub> 
- [DONE] Add Oauth grant_types abstraction <sub>[*EPIC Security*]<sub>
- [DONE] Add controller selection into security manager: oauth, oauth2, next implementations...  <sub>[*EPIC Security*]<sub>
- [DONE] Move CORS, Content Type, etc.. to another middleware, so it can be enabled or disabled by the manager.
- [DONE] Refactor Target to it allows to add more fields than url and authorities. i.e. origins, enabled, methods, scopes, etc...
- [DONE] Add support for roles (scopes) in HTTP Routes. Roles in authentication must match with the roles in the endoint.
- [DONE] Add Metrics and health-checks endpoints

## Backlog

- Organize the core packages to allow hexagonal architecture
- Way to centralize timeouts per context (by configuration per item)
- GRPC abstraction
- Create generic side-car, using envoy for example, to create and ensure lightweight clients.
- Add necessary unit testing to check and validate the functionality (uses cases and corner cases)
- Perform benchmark testing.