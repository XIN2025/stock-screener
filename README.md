```markdown
# Stock Screener

This project provides a responsive Next.js 13 frontend for filtering and sorting stock data. It communicates with a separate Go backend via a POST endpoint.

## Environment Variables

1. **NEXT_PUBLIC_BACKEND**  
   Placeholder for the backend URL (e.g., `https://my-backend.com/api/screen`).  
   This is used in the frontend to make all requests.

2. **ALLOWED_ORIGIN**  
   In the Go backend, specify which origin is allowed for CORS (e.g., `http://localhost:3000`).

## Setup and Usage

1. **Install dependencies**
   ```bash
   npm install
   ```
2. **Create `.env.local`** in the project root and set:
   ```bash
   NEXT_PUBLIC_BACKEND=<your_backend_url>
   ```
3. **Run the dev server**
   ```bash
   npm run dev
   ```
   Access the app at [http://localhost:3000](http://localhost:3000).

4. **Build for production**
   ```bash
   npm run build
   npm run start
   ```

## Backend Notes

- The Go backend should enable CORS with the allowed origin.
- Example environment variable:
  ```bash
  ALLOWED_ORIGIN=<your_frontend_url>
  ```
- The backend URL specified in **NEXT_PUBLIC_BACKEND** is where the frontend sends requests.

## Features

- Enter a query with multiple conditions (AND-based).
- Sort by any column (e.g., Market Cap, P/E Ratio).
- Pagination with custom limits (10, 25, 50, 100).
- Download all results as CSV.
- Responsive design for both desktop and mobile.

```
Market Capitalization (B) > 100 AND
ROE (%) > 15 AND
P/E Ratio < 20
```
