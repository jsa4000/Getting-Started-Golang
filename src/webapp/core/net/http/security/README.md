# Security

## Introduction

Security is a quality attribute which interacts heavily with other such attributes, including availability, safety, and robustness. It is the sum of all of the attributes of an information system or product which contributes towards ensuring that processing, storing, and communicating of information sufficiently protects confidentiality, integrity, and authenticity.

Security can be delegated to another systems, for example for authorization (oauth2). Security is also delegated onto upper layers in the infrastructure (TLS termination, API Gateway, Firewalls, etc..).

## Authentication, Authorization, and Accounting (AAA)

Authentication, authorization, and accounting (AAA) is a term for a framework for intelligently controlling access to computer **resources**, enforcing policies, auditing usage, and providing the information necessary to bill for services. These combined processes are considered important for effective network management and security.

- **Authentication**: provides a way of identifying a user, typically by having the user enter a valid user name and valid password before access is granted. The process of authentication is based on each user having a unique set of criteria for gaining access. The AAA server compares a user's authentication credentials with other user credentials stored in a database. If the credentials match, the user is granted access to the network. If the credentials are at variance, authentication fails and network access is denied.
- **Authentication**: a user must gain authorization for doing certain tasks (scopes, roles, etc..). After logging into a system, for instance, the user may try to issue commands. The authorization process determines whether the user has the authority to issue such commands. Simply put, authorization is the process of enforcing policies: determining what types or qualities of activities, resources, or services a user is permitted. Usually, authorization occurs within the context of authentication. Once you have authenticated a user, they may be authorized for different types of access or activity.
- **Accounting**: which measures the resources a user consumes during access. This can include the amount of system time or the amount of data a user has sent and/or received during a session. Accounting is carried out by logging of session statistics and usage information and is used for authorization control, billing, trend analysis, resource utilization, and capacity planning activities.

Authentication, authorization, and accounting services are often provided by a dedicated AAA server, a program that performs these functions. A current standard by which network access servers interface with the AAA server is the *Remote Authentication Dial-In User Service* (*RADIUS*).

### Authentication/Authorization

Authentication provides a way to identify **clients** and **users**. **Clients** are the applications that consume the resources and the **users** are the *resource owner*.

There are different ways to authenticate/authorize. Here are some of them:

- OAuth2: is an open standard for access **delegation**, commonly used as a way for Internet users to grant websites or applications access to their information on other websites but without giving them the passwords (directly). This mechanism is used by companies such as Amazon, Google, Facebook, Microsoft and Twitter to permit the users to share information about their accounts with third party applications or websites.
- OpenId-Connect (OIDC): is a simple identity layer on top of the OAuth 2.0 protocol, which allows computing clients to verify the identity of an end-user based on the authentication performed by an authorization server, as well as to obtain basic profile information about the end-user in an interoperable and REST-like manner. In technical terms, OpenID Connect specifies a RESTful HTTP API, using JSON as a data format.
- Basic-Auth: HTTP Basic authentication (BA) implementation is the simplest technique for enforcing access controls to web resources because it does not require cookies, session identifiers, or login pages; rather, HTTP Basic authentication uses standard fields in the HTTP header, removing the need for handshakes.
- No-Auth: this means the resources don not require any kind of authentication. In this case, resources are opened to the outside without the need of any credentials or grants.

Basically, authenticate/authorize methods enable and grant applications and users to access to web resources (Resources Owner). The goal are workflows (protocols) for previous methods differs to each other:

- OAuth: generates an `access_token`. This token can be *state-less* (JWT) or *state-ful* (session). This token is used by applications to access to resources, this way, credentials are no longer needed. Depending on the client and its purposes, it can be used different flows to grant access to resources.
- Basic-Auth: it must be used for all request performed by the client (application). This does not generates any token, since the credentials are passed using `Basic {Base64("username:password")}`. This method can be used for not highly secured or private resources, such as client information, confidential data, etc.. For example, it can be used for *health-check* or *metrics* end-points. Sometimes, *basic-auth* is replaced by *OAuth*, using the  *client-crendencials* flow, that it is a similar auth process but it needs the **client-secret** to be secured by the client. Using OAuth allows users to use  a standard way to access resources but using scopes, authorities, etc..
- No-Auth: this method must be restricted to static resources, so any user can access to them (via GET method) without any kind of auth. ie. *.html, *.images, *.css, *.js, etc..

> Independently of the auth method used the Accounting is important to ensoure attacks limiting the rate, requests, etc..

#### OAuth2

Allowed grant types or flows are:

- **Authentication Code**:
- **Implicit**:
- **Client Credentials**:
- **Resource Owner Password Crendecials Grant**:

### Server Resources

Resource Servers verify if incomming requests are **allowed** to **access** to a particular resource (based on scopes, userinfo, roles, etc..).

There are some ways to check if the current request is `authorized`:

- **No-Auth**: Open to the outside world; *if no other rule or firewall inbetween*. All requests are allowed.
- **Token-Based**: The requests must have the `Authorization` header. It starts with `Bearer {token}`. There are different token specifications:
  - [Self-contained token (JWT)](https://jwt.io/):it contains the enough information so Resource Servers can check if the requests have enough priviledges to access to a particular resource.
    > i.e `Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`
  - *Reference Token*: These token are just pointers to metadata in the authorization server. So it is needed to call to auth server to gather information so resurce servers can decide if the request (token based) is allows to access to that particular resource.
    > i.e `Authorization: Bearer RGFuJ3MgVG9vbHMgYXJlIGNvb2wh`
- **Basic-Auth**: The requests must have the `Authorization` header. It starts with `Basic {Base64("username:password")}`.
    > i.e `Authorization: Basic YWhhbWlsdG9uQGFwaWdlZS5jb206bXlwYXNzdzByZAo`

Resource Servers basically performs the following information per request:

- Check request authorization type for that particular resource/endpoint (no-auth, basic-auth, etc..)
- Depending on the previous step, it looks for the proper header. i.e  `Authorization: Basic ..`
- Verify if current request has access to that particular resource. Checks for scopes, roles, userinfo etc.. Depending on the `authorization` method it would be needed to access to the authorization server in order to validate the token, expiration, user, etc.. It depends also, if the token is `self-contained`, `id_token`,  (OpenID Connect), no secret-key (HMAC with SHA-256) or public-key (RSA Signature with SHA-256)) in resource server to verify the token, audience, client_id, etc...
- Finally, resource server must `allows` or `forbid` current request.
