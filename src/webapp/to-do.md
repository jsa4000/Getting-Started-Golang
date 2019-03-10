# To-Do

## Sprints

### Current

- [DONE] Make configurable Decode function on net/http/Controller to increase performance avoid unnecessary decodings.
- Create Build for DecodeOptions instead using the default constructor (user-firendly).
- Create second micro-service to interchange messages (http/2, events, etc.)
- GRPC abstraction between code-generated and logic.
- Create generic side-car (envoy, nginx, etc.) for service-discovery, tracing, logging, metrics, circuit-breaker, load-balancing, fault-tolerance, etc.. to allows lightweight clients. Support for cilium and BPF
- [DONE] Add support for distributed cache: Redis, etcd...
- [DONE] Create cache layer (decorator) between repositories and services
- [DONE] Centralize distributed cache into a package store/cache
- [DONE] Add Roles to Users, Model and endpoints.
- [DONE] Add RedirectURL support for OAuth
- Messaging Systems: NATs, Kafka, RabbitMQ, etc..
- [DONE] Support for bootstrapping
- Add etcd support for distributed cache

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

- Continue with the OAut implementation, allowing implicit and Authorization Code Grants. (refresh token)
- Organize the core packages to allow hexagonal architecture. (core, framework, metrics, system, transport, storage, etc..)
- Way to centralize timeouts per context (by configuration per item)
- Add necessary unit testing to check and validate the functionality (uses cases and corner cases). Use Tables cases instead.
- Perform benchmark testing: speed, latency, throughput, memory, etc...
- Improve log quiality for better understanding: log levels, errors accoracy, etc..
- Prepare documentation for the Core framework. Cleanup obsolete and wrong comments in code.
- Thinnk about the usage for Builders, managers, singletone, etc... See other packages and libraries.
- Code generation using Go Templates. This is also to reduce the boiler plate
- Create Frontend using any Framework (React, AngularJS, etc..) to start interacting with the API
- Allow OAuth delegating the Authorization to another system such as Google, Facebook, etc..
- Mode current core into different git repository. Prepare examples and documentation getting started guide.