# Security

## Authorization Server

This enables the application to authenticate/authorize and grant clients to web resources.

There are two main standars currently used for web applications:

- **OAuth**: Delegation protocol for authorization decisions (implicit, authentication code, client credentials, Resource Owner password, etc..).
- **OpenID Connect (OIDC)**: Build upon **OAuth2**, it provides authentication information (id_token, userInfo endpoint, default scopes, etc..)

OAuth2 Flows:

- **Authentication Code**:
- **Implicit**:
- **Client Credentials**:
- **Resource Owner Password Crendecials Grant**:

## Resource Server

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


