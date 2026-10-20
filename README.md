# SpectraRoleplayDataService

`SpectraRoleplayDataService` is a Go HTTP service that provides read access to
Spectra roleplay guild data stored in MongoDB. At present, it exposes a health
endpoint and a lookup endpoint for a guild document in the `bot_data.guild`
MongoDB collection.

## Requirements

- Go **1.25+** (the module declares Go 1.25.0)
- Access to a MongoDB instance containing the `bot_data` database and `guild`
  collection

## How to run locally

1. From the repository root, create a local configuration file from the
   provided template:

   ```powershell
   Copy-Item .env.example .env
   ```

   On macOS or Linux, use:

   ```sh
   cp .env.example .env
   ```

2. Update `.env` with the port to listen on and a valid MongoDB connection
   string:

   ```dotenv
   PORT="8080"
   MONGODB_CONNECTION_STRING="mongodb://<user>:<password>@<host>:<port>"
   ```

   The local configuration file is intentionally ignored by Git. Do not commit
   credentials.

3. Download dependencies and start the API:

   ```sh
   go mod download
   go run ./cmd/api
   ```

   By default, the service listens at `http://localhost:8080` when `PORT` is
   set to `8080`.

4. Verify that the service is running:

   ```sh
   curl http://localhost:8080/health
   ```

### Configuration modes

`go run ./cmd/api` uses the `local` configuration mode by default and loads
environment values from the repository-root `.env` file. A non-local mode can
be selected with the `-env` flag; for example:

```sh
go run ./cmd/api -env production
```

Non-local modes load `configs/configuration-<environment>.json`. Each runtime
configuration needs both `PORT` and `MONGODB_CONNECTION_STRING` for the API to
start and serve guild data successfully.

## API contract (current)

The API currently has no authentication, version prefix, or request body for
its available endpoints. Routes are implemented for `GET` requests only.

### `GET /health`

Returns a lightweight liveness response. This endpoint does not query MongoDB.

**Example request**

```sh
curl http://localhost:8080/health
```

**Successful response — `200 OK`**

```json
{
  "message": "API Live",
  "currentTime": "Mon, 19 Oct 2026 12:00:00 UTC"
}
```

`currentTime` is generated when the request is handled and formatted as an
RFC 1123 timestamp; the value above is illustrative.

### `GET /guild-data/{guild_id}`

Looks up a single document in MongoDB's `bot_data.guild` collection where the
stored `guildId` field exactly matches the `{guild_id}` path parameter.

**Path parameter**

| Name | Type | Required | Description |
| --- | --- | --- | --- |
| `guild_id` | string | Yes | Guild identifier used to match the MongoDB `guildId` field. |

**Example request**

```sh
curl http://localhost:8080/guild-data/1234567890
```

**Successful response — `200 OK`**

```json
{
  "guildId": "1234567890",
  "guildData": {
    "guildId": "1234567890"
  }
}
```

`guildData` contains the complete MongoDB document, so its fields depend on
the stored guild record. The shown document is a minimal example, not a fixed
schema.

### Current error behavior

There is not yet a structured API error contract. If a guild document cannot
be read (including when it is not found), the service logs the error and
currently returns an empty response with the default `200 OK` status. Consumers
should treat an empty guild-data response as an unsuccessful lookup until an
explicit error response contract is introduced.
