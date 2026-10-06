# Web App Guidelines

These rules apply to all code written in this project. Follow them for every change.

## Principles

- Follow the **SOLID** principles.
- Organize files by the **Single Responsibility Principle**: each file has one clear purpose. A component, hook, util, constant group, or service gets its own file. Do not mix several concerns in one file.

## Folder Structure (Feature-Based)

Code is grouped by feature, not by technical type. Everything that belongs only to one feature (components, hooks, utils, constants, types, etc.) lives inside that feature's folder.

```
src/
├── app/                  # App shell
│   ├── router/           # Router setup and route definitions
│   └── components/       # App-wide components (Navbar, layouts, etc.)
│
├── modules/              # One folder per feature
│   └── moduleA/
│       ├── components/
│       ├── hooks/
│       ├── utils/
│       ├── constants/
│       └── ...
│
└── shared/               # Code used by 2+ modules (same structure as a module)
    ├── components/
    ├── hooks/
    ├── utils/
    ├── constants/
    ├── ...
    └── api/
        ├── axios.ts      # App-wide Axios client config (base URL, interceptors, etc.)
        ├── dtos/         # Request/response DTO types
        └── services/     # Axios functions that make the HTTP requests
```

### `modules/`

- Each feature is its own folder inside `modules/`.
- Feature-specific code stays inside its module. Only create the subfolders a module actually needs.

### `shared/`

- Follows the same structure as a module.
- Only for code used by **two or more** modules. If code is used by a single module, keep it in that module; move it to `shared/` once a second module needs it.

### `shared/api/`

- `axios.ts` — the single app-wide Axios client instance and its config (interceptors, headers, base URL, etc.). All requests must use this client.
- `dtos/` — DTO types for request and response payloads.
- `services/` — all functions that make HTTP requests through the Axios client. Components and hooks must not call Axios directly; they go through these services.

### `app/`

- Holds the router, route definitions, and app-wide components (navbar, layouts, etc.).
