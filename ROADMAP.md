# Roadmap
This is Veil's protocol roadmap

### Disclaimer
Here you will only find *planned* features, not *promised* features. Everything found here can be removed and new things can be added.

### Format
Each task can be marked with a symbol:
- Tasks marked with `?` are just ideas and have a high risk of being removed.
- Tasks marked with `!` will most surely be added.
- Tasks marked with two `!` (`!!`) are max priority.

### Aliases
These are some common aliases:
- **PK(s)**: Public Key(s)
- **WS**: WebSocket
- **mgmt**: Management

---

## Phase 1: Core systems
- [x] Auth:
    - [x] Password hashing
    - [x] Token generation
- [x] Identity:
    - [x] ! PK generation
- [ ] WS:
    - [ ] !! Basic server-side code
    - [ ] !! Base server events
    - [ ] !! Basic client-side code
    - [ ] !! Base client events
- [ ] REST API:
    - [ ] !! Base routes
- [ ] Protocol API:
    - [ ] !! Authentication (tokens, hashing) and generation (public keys) API
    - [ ] ! Server -> client communication helpers
    - [ ] ! Client -> server communication helpers
    - [ ] ! REST API helpers
    - [ ] ! Message creation and formatting helpers
    - [ ] ? JavaScript/TypeScript bindings

## Phase 2: Main protocol
- [ ] Auth:
    - [ ] !! Apply basic authentication
    - [ ] ! Make simple decentralized signup, login and logout system
    - [ ] Use private tokens for internal auth-related things
- [ ] WS:
    - [ ] Chatting:
        - [ ] !! Rooms:
            - [ ] !! Public rooms (link-based)
            - [ ] ! Private rooms (link-based + password)
        - [ ] !! Direct messages (PK-based)
            - [ ] ? Friends (trusted PKs)
        - [ ] !! Storing and compressing messages per-session
        - [ ] !! E2EE (end-to-end encryption) for chats
- [ ] REST API:
    - [ ] Chatting:
        - [ ] !! Rooms:
            - [ ] ! Private room auth (link-based + password)
- [ ] Protocol API:
    - [ ] ! Chatting helpers
    - [ ] ! Room mgmt helpers
