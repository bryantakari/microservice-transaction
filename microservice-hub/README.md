# Microservice Hub

`microservice-hub` is the frontend dashboard used to illustrate the business flow of the microservice architecture in this repository. It is a Vue 3 + TypeScript + Vite application that shows how orders, items, inventory, and users can be navigated and managed from a single dashboard interface.

## Local setup requirements

- Node.js 18+ (recommended) or the latest LTS version
- npm, pnpm, or yarn package manager
- A terminal or shell environment

## Run locally

1. Open a terminal in `microservice-hub`
2. Install dependencies:
   ```bash
   npm install
   ```
3. Configure the API URLs in `.env.development` (see `.env.example`). The order and inventory URLs assume those services listen on ports `3001` and `3002`; set each Go service's `APP_PORT` to match. The user service URL uses its current default port, `8080`.
4. Start the local development server:
   ```bash
   npm run dev
   ```

The app should be available at `http://localhost:5173` by default.

## Build for production

```bash
npm run build
```

Optionally preview the production build:

```bash
npm run preview
```

## Notes

- This project is primarily used to illustrate the business flow and UI experience for the microservice architecture.
- The menu is currently static, but the structure is easier to extend later if a menu management feature is added.

## Related documentation

See the root project README at `../README.md` for the overall repository structure and backend service details.
